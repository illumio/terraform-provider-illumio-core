// The import binary scrapes the target PCE for all object types
// and creates HCL and .tfstate files for all remote objects
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"

	"github.com/brian1917/illumioapi"
)

const illumioProviderName = "registry.terraform.io/illumio/illumio-core"

var pce *illumioapi.PCE
var currentDirectory string
var ctx context.Context

// HCL normalization regex - names can only use alphanum, dashes and underscores
var hclNormRe = regexp.MustCompile(`[^a-zA-Z0-9_-]`)
var containerClusterRe = regexp.MustCompile(`\/orgs\/\d+\/container_clusters\/(?P<ClusterID>[\w-]+)`)

var tfStateMap = map[string]struct{}{}

// map of HCL addresses to object IDs (HREFs)
var tfImportMap = map[string]string{}

// TODO: build another index of HREFs to HCL addresses and
// do another pass after we write the files to replace refs.
// this will provide dependency resolution between resources

func handleError(err error, exitOnErr bool) {
	if err != nil {
		log.Printf("%v", err)
		if exitOnErr {
			os.Exit(1)
		}
	}
}

func main() {
	var err error
	ctx = context.Background()

	log.Print("Connecting to the PCE")
	connectToPCE()

	currentDirectory, err = os.Getwd()
	handleError(err, true)

	log.Print("Initializing Terraform")
	tf := initializeTerraform()
	log.Print("Terraform initialized")

	log.Print("Loading PCE objects")
	// load all objects from the PCE
	pce.Load(illumioapi.LoadInput{
		Labels:                      true,
		LabelGroups:                 true,
		IPLists:                     true,
		Workloads:                   true,
		VirtualServices:             true,
		VirtualServers:              true,
		Services:                    true,
		ConsumingSecurityPrincipals: true,
		RuleSets:                    true,
		VENs:                        true,
		ContainerClusters:           true,
		ContainerWorkloads:          true,
	})

	// fetch all remaining objects from the PCE:
	// - enforcement boundaries
	// - pairing profiles
	// - vulnerabilities
	// - vulnerability reports
	// - rules (can maybe pull these out from the rule sets?)
	// - service bindings (similar, maybe can get from virtual services)
	// - workload/container cluster subobjects
	// - firewall settings
	// - org settings
	// - traffic collector settings
	// - syslog destinations
	emptyParams := map[string]string{}
	vulnerabilities, _, err := pce.GetVulns(emptyParams)
	handleError(err, false)

	vulnerabilityReports, _, err := pce.GetVulnReports(emptyParams)
	handleError(err, false)

	enforcementBoundaries := []illumioapi.EnforcementBoundary{}
	err = getAllObjects(pce, "/sec_policy/draft/enforcement_boundaries", &enforcementBoundaries)
	handleError(err, false)

	pairingProfiles := []illumioapi.PairingProfile{}
	err = getAllObjects(pce, "/pairing_profiles", &pairingProfiles)
	handleError(err, false)
	log.Print("Loaded all PCE objects")

	providerSchemas, err := tf.ProvidersSchema(ctx)
	handleError(err, true)

	// if the Illumio provider isn't defined, create a main.tf
	// with the provider definition
	if _, ok := providerSchemas.Schemas[illumioProviderName]; !ok {
		mainTfFilePath := path.Join(currentDirectory, "main.tf")
		if _, err := os.Stat(mainTfFilePath); err == nil {
			log.Fatal("main.tf exists but illumio-core provider has not been defined. Exiting.")
		}
		log.Print("Writing main.tf")
		// Create main.tf which includes the provider configuration
		writeMainTf()
	}

	pceObjectMap := map[string]string{
		"container_clusters":     buildHCL(pce.ContainerClustersSlice),
		"enforcement_boundaries": buildHCL(enforcementBoundaries),
		"ip_lists":               buildHCL(pce.IPListsSlice),
		"labels":                 buildHCL(pce.LabelsSlice),
		"label_groups":           buildHCL(pce.LabelGroupsSlice),
		"pairing_profiles":       buildHCL(pairingProfiles),
		"rule_sets":              buildHCL(pce.RuleSets),
		"security_principals":    buildHCL(pce.ConsumingSecurityPrincipals),
		"services":               buildHCL(pce.ServicesSlice),
		"vens":                   buildHCL(pce.VENsSlice),
		"virtual_servers":        buildHCL(pce.VirtualServers),
		"virtual_services":       buildHCL(pce.VirtualServices),
		"vulnerabilities":        buildHCL(vulnerabilities),
		"vulnerability_reports":  buildHCL(vulnerabilityReports),
		"workloads":              buildHCL(pce.WorkloadsSlice),
		// TODO: rules
		// TODO: container_workloads
	}

	// Create HCL entries for each object type
	for objectType, hcl := range pceObjectMap {
		writeTfFile(objectType, hcl)
	}

	// import objects into tfstate
	log.Print("Importing PCE objects into tfstate")
	for address, id := range tfImportMap {
		err = tf.Import(ctx, address, id)
		handleError(err, false)
	}

	log.Print("Formatting TF files")
	err = tf.FormatWrite(ctx)
	handleError(err, false)

	// run TF plan to check that we imported everything correctly
	ok, err := tf.Plan(ctx)
	handleError(err, true)
	if !ok {
		log.Print("Plan succeeded, no changes required")
	} else {
		log.Print("Plan succeeded, changes need to be applied")
	}
}

// Runs terraform init and loads existing state. If one exists, reads the
// state into memory so we don't try to import the same object twice.
// This assumes that the object hasn't been updated on the remote, but any
// change will be resolved in a refresh or apply
func initializeTerraform() *tfexec.Terraform {
	installer := &releases.LatestVersion{
		Product: product.Terraform,
	}

	execPath, err := installer.Install(ctx)
	handleError(err, true)

	tf, err := tfexec.NewTerraform(currentDirectory, execPath)
	handleError(err, true)

	err = tf.Init(ctx, tfexec.Upgrade(true))
	handleError(err, true)

	// load state from default state path
	// TODO: custom state file definition
	state, err := tf.Show(ctx)
	handleError(err, false)

	if state != nil {
		stateResources := state.Values.RootModule.Resources
		for _, res := range stateResources {
			tfStateMap[res.Address] = struct{}{}
		}
	}

	return tf
}

func connectToPCE() {
	pce_host := os.Getenv("ILLUMIO_PCE_HOSTNAME")
	pce_port, err := strconv.Atoi(os.Getenv("ILLUMIO_PCE_PORT"))
	handleError(err, true)

	pce_org_id, err := strconv.Atoi(os.Getenv("ILLUMIO_PCE_ORG_ID"))
	handleError(err, true)

	pce_api_key := os.Getenv("ILLUMIO_API_KEY_USERNAME")
	pce_api_secret := os.Getenv("ILLUMIO_API_KEY_SECRET")

	pce = &illumioapi.PCE{
		FQDN: pce_host,
		Port: pce_port,
		Org:  pce_org_id,
		User: pce_api_key,
		Key:  pce_api_secret,
	}
}

// Fetches all objects from the given endpoint by first getting the
// X-Total-Count header and then setting the max_results value so we
// pull everything at once
func getAllObjects(pce *illumioapi.PCE, endpoint string, target interface{}) error {
	response, err := pce.GetCollection(endpoint, false, map[string]string{"max_results": "0"}, target)
	if err != nil {
		return err
	}

	objectCount := response.Header.Get("x-total-count")
	response, err = pce.GetCollection(endpoint, false, map[string]string{"max_results": objectCount}, target)
	return err
}

// Creates the main.tf HCL file that defines the illumio provider
func writeMainTf() {
	writeTfFile("main", `terraform {
  required_providers {
    illumio-core = {
      source = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  request_timeout = 30
}
`)
}

// Creates a .tf HCL file with the given name and contents
func writeTfFile(fileName, hcl string) {
	log.Printf("Writing %s.tf", fileName)
	err := os.WriteFile(path.Join(currentDirectory, fmt.Sprintf("%s.tf", fileName)), []byte(hcl), 0644)
	handleError(err, true)
}

// iterates over a slice or map loaded into the PCE to write HCL resource blocks
func buildHCL(objects interface{}) string {
	v := reflect.ValueOf(objects)
	var hcl strings.Builder
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			hcl.WriteString(hclFromObject(v.Index(i).Interface()))
		}
	} else if v.Kind() == reflect.Map {
		iter := v.MapRange()
		for iter.Next() {
			// each map may have multiple keys pointing to the same object, so use
			// the HREF prefix to dedup
			if strings.HasPrefix(iter.Key().String(), "/orgs/") {
				hcl.WriteString(hclFromObject(iter.Value().Interface()))
			}
		}
		v.MapKeys()
	} else {
		handleError(fmt.Errorf("passed type %v to buildHCL, expected slice or map", v.Kind()), true)
	}
	return hcl.String()
}

func hclFromObject(obj interface{}) string {
	switch o := obj.(type) {
	case illumioapi.ConsumingSecurityPrincipals:
	case illumioapi.ContainerCluster:
		return buildContainerClusterHCL(o)
	case illumioapi.EnforcementBoundary:
		return buildEnforcementBoundaryHCL(o)
	case illumioapi.IPList:
		return buildIPListHCL(o)
	case illumioapi.Label:
		return buildLabelHCL(o)
	case illumioapi.LabelGroup:
		return buildLabelGroupHCL(o)
	case illumioapi.PairingProfile:
		return buildPairingProfileHCL(o)
	case illumioapi.RuleSet:
		return buildRuleSetHCL(o)
	case illumioapi.Service:
		return buildServiceHCL(o)
	case illumioapi.VEN:
	case illumioapi.VirtualServer:
	case illumioapi.VirtualService:
		return buildVirtualServiceHCL(o)
	case illumioapi.Workload:
		return buildWorkloadHCL(o)
	case illumioapi.Vulnerability:
	case illumioapi.VulnerabilityReport:
	default:
		handleError(fmt.Errorf("invalid type: %v", o), true)
	}
	return ""
}

// normalizes a given string to fit HCL name constraints
func hclNormalize(s string) string {
	s = strings.ToLower(s)
	// replace all spaces first
	s = strings.ReplaceAll(s, " ", "-")
	// strip any remaining special characters
	s = hclNormRe.ReplaceAllString(s, "")
	return s
}

func buildLabelHCL(label illumioapi.Label) string {
	var hcl strings.Builder
	hclName := hclNormalize(label.Value)
	address := fmt.Sprintf("illumio-core_label.%s_%s", label.Key, hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = label.Href
	}
	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_label" "%[1]s_%[2]s" {
  key   = %[1]q
  value = %[3]q`, label.Key, hclName, label.Value, label.ExternalDataSet, label.ExternalDataReference))

	if label.ExternalDataSet != "" && label.ExternalDataReference != "" {
		hcl.WriteString(fmt.Sprintf(`

  external_data_set       = %q
  external_data_reference = %q`, label.ExternalDataSet, label.ExternalDataReference))
	}

	hcl.WriteString(`
}
`)
	return hcl.String()
}

func buildLabelGroupHCL(labelGroup illumioapi.LabelGroup) string {
	var hcl strings.Builder
	hclName := hclNormalize(labelGroup.Name)
	address := fmt.Sprintf("illumio-core_label_group.%s_%s", labelGroup.Key, hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = labelGroup.Href
	}
	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_label_group" "%[1]s_%[2]s" {
  key         = %[1]q
  name        = %[3]q
  description = %[4]q
`, labelGroup.Key, hclName, labelGroup.Name, labelGroup.Description))

	// labels blocks
	for _, label := range labelGroup.Labels {
		hcl.WriteString(fmt.Sprintf(`
  labels {
    href = illumio-core_label.%s_%s.href
  }`, label.Key, hclNormalize(label.Value)))
	}

	if labelGroup.ExternalDataSet != "" && labelGroup.ExternalDataReference != "" {
		hcl.WriteString(fmt.Sprintf(`

  external_data_set       = %q
  external_data_reference = %q`, labelGroup.ExternalDataSet, labelGroup.ExternalDataReference))
	}

	hcl.WriteString(`
}
`)
	return hcl.String()
}

func buildIPListHCL(ipList illumioapi.IPList) string {
	var hcl strings.Builder
	hclName := hclNormalize(ipList.Name)
	address := fmt.Sprintf("illumio-core_ip_list.%s", hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = ipList.Href
	}
	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_ip_list" %q {
  name        = %q
  description = %q
`, hclName, ipList.Name, ipList.Description))

	// ip_ranges blocks
	for _, ipRange := range *ipList.IPRanges {
		hcl.WriteString(fmt.Sprintf(`
  ip_ranges {
    exclusion   = %v
    description = %q
    from_ip     = %q`, ipRange.Exclusion, ipRange.Description, ipRange.FromIP))
		if ipRange.ToIP != "" {
			hcl.WriteString(fmt.Sprintf(`
    to_ip       = %q`, ipRange.ToIP))
		}
		hcl.WriteString(`
  }`)
	}

	// fqdns blocks
	for _, fqdn := range *ipList.FQDNs {
		// illumioapi FQDN object doesn't include a description field
		hcl.WriteString(fmt.Sprintf(`
  fqdns {
    fqdn = %q
  }`, fqdn.FQDN))
	}

	if ipList.ExternalDataSet != "" && ipList.ExternalDataReference != "" {
		hcl.WriteString(fmt.Sprintf(`

  external_data_set       = %q
  external_data_reference = %q`, ipList.ExternalDataSet, ipList.ExternalDataReference))
	}

	hcl.WriteString(`
}
`)
	return hcl.String()
}

func buildServiceHCL(service illumioapi.Service) string {
	var hcl strings.Builder
	hclName := hclNormalize(service.Name)
	address := fmt.Sprintf("illumio-core_service.%s", hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = service.Href
	}
	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_service" %q {
  name         = %q
  description  = %q
`, hclName, service.Name, service.Description))
	if service.ProcessName != "" {
		hcl.WriteString(fmt.Sprintf(`
  process_name = %q
`, service.ProcessName))
	}

	for _, svcPort := range service.ServicePorts {
		svcValueMap := map[string]int{
			"port": svcPort.Port, "to_port": svcPort.ToPort, "proto": svcPort.Protocol,
			"icmp_type": svcPort.IcmpType, "icmp_code": svcPort.IcmpCode,
		}
		hcl.WriteString(`
  service_ports {`)
		for field, val := range svcValueMap {
			if val != 0 {
				hcl.WriteString(fmt.Sprintf(`
    %-9s = "%d"`, field, val))
			}
		}
		hcl.WriteString(`
  }`)
	}

	for _, winSvc := range service.WindowsServices {
		svcValueMap := map[string]int{
			"port": winSvc.Port, "to_port": winSvc.ToPort, "proto": winSvc.Protocol,
			"icmp_type": winSvc.IcmpType, "icmp_code": winSvc.IcmpCode,
		}
		hcl.WriteString(`
  windows_services {`)
		if winSvc.ServiceName != "" {
			hcl.WriteString(fmt.Sprintf(`
    service_name = %q`, winSvc.ServiceName))
		}
		if winSvc.ProcessName != "" {
			hcl.WriteString(fmt.Sprintf(`
    process_name = %q`, winSvc.ProcessName))
		}
		for field, val := range svcValueMap {
			if val != 0 {
				hcl.WriteString(fmt.Sprintf(`
    %-12s = "%d"`, field, val))
			}
		}
		hcl.WriteString(`
  }`)
	}

	if service.ExternalDataSet != "" && service.ExternalDataReference != "" {
		hcl.WriteString(fmt.Sprintf(`

  external_data_set       = %q
  external_data_reference = %q`, service.ExternalDataSet, service.ExternalDataReference))
	}

	hcl.WriteString(`
}
`)
	return hcl.String()
}

func buildEnforcementBoundaryHCL(enforcementBoundary illumioapi.EnforcementBoundary) string {
	var hcl strings.Builder
	hclName := hclNormalize(enforcementBoundary.Name)
	address := fmt.Sprintf("illumio-core_enforcement_boundary.%s", hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = enforcementBoundary.Href
	}

	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_enforcement_boundary" %q {
  name = %q`, hclName, enforcementBoundary.Name))

	for _, service := range enforcementBoundary.IngressServices {
		hcl.WriteString(`
  ingress_services {`)
		if service.Href != nil {
			hcl.WriteString(fmt.Sprintf(`
    href = %q`, *service.Href))
		}
		if service.Port != nil {
			hcl.WriteString(fmt.Sprintf(`
    port = "%d"`, *service.Port))
		}
		if service.ToPort != nil {
			hcl.WriteString(fmt.Sprintf(`
    to_port = "%d"`, *service.ToPort))
		}
		if service.Protocol != nil {
			hcl.WriteString(fmt.Sprintf(`
    proto = "%d"`, *service.Protocol))
		}
		hcl.WriteString(`
  }`)
	}

	for _, consumer := range enforcementBoundary.Consumers {
		hcl.WriteString(`
  consumers {`)
		if consumer.Actors == "ams" {
			hcl.WriteString(`
    actors = "ams"`)
		}
		if consumer.IPList != nil {
			hcl.WriteString(fmt.Sprintf(`
    ip_list {
      href = %q
    }`, consumer.IPList.Href))
		}
		if consumer.Label != nil {
			hcl.WriteString(fmt.Sprintf(`
    label {
      href = %q
    }`, consumer.Label.Href))
		}
		if consumer.LabelGroup != nil {
			hcl.WriteString(fmt.Sprintf(`
    label_group {
      href = %q
    }`, consumer.LabelGroup.Href))
		}
		hcl.WriteString(`
  }`)
	}

	// this duplication is necessary because of Go's lack of inheritance
	// we could maybe use generics, but it's more trouble than it's worth
	for _, provider := range enforcementBoundary.Providers {
		hcl.WriteString(`
  providers {`)
		if provider.Actors == "ams" {
			hcl.WriteString(`
    actors = "ams"`)
		}
		if provider.IPList != nil {
			hcl.WriteString(fmt.Sprintf(`
    ip_list {
      href = %q
    }`, provider.IPList.Href))
		}
		if provider.Label != nil {
			hcl.WriteString(fmt.Sprintf(`
    label {
      href = %q
    }`, provider.Label.Href))
		}
		if provider.LabelGroup != nil {
			hcl.WriteString(fmt.Sprintf(`
    label_group {
      href = %q
    }`, provider.LabelGroup.Href))
		}
		hcl.WriteString(`
  }`)
	}

	hcl.WriteString(`
}
`)
	return hcl.String()
}

func buildRuleSetHCL(ruleSet illumioapi.RuleSet) string {
	var hcl strings.Builder
	hclName := hclNormalize(ruleSet.Name)
	address := fmt.Sprintf("illumio-core_rule_set.%s", hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = ruleSet.Href
	}

	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_rule_set" %q {
  name        = %q
  description = %q
  enabled     = %v
`, hclName, ruleSet.Name, ruleSet.Description, *ruleSet.Enabled))
	for _, scope := range ruleSet.Scopes {
		hcl.WriteString(`
  scopes {`)
		for _, label := range scope {
			if label.Label != nil {
				hcl.WriteString(fmt.Sprintf(`
    label {
      href = %q
    }`, label.Label.Href))
			} else if label.LabelGroup != nil {
				hcl.WriteString(fmt.Sprintf(`
    label_group {
      href = %q
    }`, label.Label.Href))
			}
		}
		hcl.WriteString(`
  }`)
	}

	for _, ipTablesRule := range ruleSet.IPTablesRules {
		hcl.WriteString(fmt.Sprintf(`
  ip_tables_rules {
    enabled     = %v
    description = %q
    ip_version  = %q
`, ipTablesRule.Enabled, ipTablesRule.Description, ipTablesRule.IPVersion))
		for _, statement := range ipTablesRule.Statements {
			hcl.WriteString(fmt.Sprintf(`
    statements {
      table_name = %q
      chain_name = %q
      parameters = %q
    }
`, statement.TableName, statement.ChainName, statement.Parameters))
		}
		for _, actor := range ipTablesRule.Actors {
			hcl.WriteString(`
    actors {`)
			if actor.Actors == "ams" {
				hcl.WriteString(`
      actors = "ams"`)
			}
			if actor.Workload != nil {
				hcl.WriteString(fmt.Sprintf(`
      workload {
        href = %q
      }`, actor.Workload.Href))
			}
			if actor.Label != nil {
				hcl.WriteString(fmt.Sprintf(`
      label {
        href = %q
      }`, actor.Label.Href))
			}
			if actor.LabelGroup != nil {
				hcl.WriteString(fmt.Sprintf(`
      label_group {
        href = %q
      }`, actor.LabelGroup.Href))
			}
			hcl.WriteString(`
    }`)
		}
	}

	if ruleSet.ExternalDataSet != "" && ruleSet.ExternalDataReference != "" {
		hcl.WriteString(fmt.Sprintf(`

  external_data_set       = %q
  external_data_reference = %q`, ruleSet.ExternalDataSet, ruleSet.ExternalDataReference))
	}
	hcl.WriteString(`
}
`)
	return hcl.String()
}

// XXX: illumioapi still uses the deprecated mode parameter rather
// than enforcement_mode and doesn't implement agent_software_release
// so go get them separately
type PairingProfile struct {
	EnforcementMode     string `json:"enforcement_mode"`
	EnforcementModeLock bool   `json:"enforcement_mode_lock"`
	VENVersion          string `json:"agent_software_release"`
}

func buildPairingProfileHCL(pairingProfile illumioapi.PairingProfile) string {
	var hcl strings.Builder
	hclName := hclNormalize(pairingProfile.Name)
	address := fmt.Sprintf("illumio-core_pairing_profile.%s", hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = pairingProfile.Href
	}

	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_pairing_profile" %q {
  name                   = %q
  description            = %q
  enabled                = %v
  allowed_uses_per_key   = %q
  key_lifespan           = %q
  visibility_level       = %q
  visibility_level_lock  = %v
  role_label_lock        = %v
  app_label_lock         = %v
  env_label_lock         = %v
  loc_label_lock         = %v`, hclName, pairingProfile.Name, pairingProfile.Description, pairingProfile.Enabled,
		pairingProfile.AllowedUsesPerKey, pairingProfile.KeyLifespan, pairingProfile.VisibilityLevel,
		pairingProfile.VisibilityLevelLock, pairingProfile.RoleLabelLock, pairingProfile.AppLabelLock,
		pairingProfile.EnvLabelLock, pairingProfile.LocLabelLock))
	pp := PairingProfile{}
	_, err := pce.GetHref(pairingProfile.Href, &pp)
	handleError(err, false)

	hcl.WriteString(fmt.Sprintf(`
  enforcement_mode       = %q
  enforcement_mode_lock  = %v
  agent_software_release = %q
`, pp.EnforcementMode, pp.EnforcementModeLock, pp.VENVersion))

	for _, label := range pairingProfile.Labels {
		hcl.WriteString(fmt.Sprintf(`
  labels {
    href = %q
  }`, label.Href))
	}

	if pairingProfile.ExternalDataSet != "" && pairingProfile.ExternalDataReference != "" {
		hcl.WriteString(fmt.Sprintf(`

  external_data_set       = %q
  external_data_reference = %q`, pairingProfile.ExternalDataSet, pairingProfile.ExternalDataReference))
	}
	hcl.WriteString(`
}
`)
	return hcl.String()
}

func buildWorkloadHCL(workload illumioapi.Workload) string {
	var hcl strings.Builder
	var hclName string
	var address string

	unmanagedWorkloadProperties := ""

	if workload.VEN != nil {
		hclName = hclNormalize(workload.Hostname)
		address = fmt.Sprintf("illumio-core_managed_workload.%s", hclName)
		hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_managed_workload" %q {
  enforcement_mode        = %q`, hclName, workload.EnforcementMode))
	} else {
		hclName = hclNormalize(workload.Name)
		address = fmt.Sprintf("illumio-core_unmanaged_workload.%s", hclName)
		hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" %q {`, hclName))

		if workload.Hostname != "" {
			unmanagedWorkloadProperties += (fmt.Sprintf(`
  hostname                = %q`, workload.Hostname))
		}

		unmanagedWorkloadProperties += (fmt.Sprintf(`
  online                  = %v`, workload.Online))

		if workload.PublicIP != "" {
			unmanagedWorkloadProperties += (fmt.Sprintf(`
  public_ip               = %q`, workload.PublicIP))
		}

		if workload.OsDetail != "" {
			unmanagedWorkloadProperties += (fmt.Sprintf(`
  os_detail               = %q`, workload.OsDetail))
		}

		if workload.OsID != "" {
			unmanagedWorkloadProperties += (fmt.Sprintf(`
  os_id                   = %q`, workload.OsID))
		}

		if workload.DistinguishedName != "" {
			unmanagedWorkloadProperties += (fmt.Sprintf(`
  distinguished_name      = %q`, workload.DistinguishedName))
		}
	}

	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = workload.Href
	}

	hcl.WriteString(fmt.Sprintf(`
  name                    = %q
  description             = %q`, workload.Name, workload.Description))

	hcl.WriteString(unmanagedWorkloadProperties)

	if workload.DataCenter != "" {
		hcl.WriteString(fmt.Sprintf(`
  data_center             = %q`, workload.DataCenter))
	}

	if workload.DataCenterZone != "" {
		hcl.WriteString(fmt.Sprintf(`
  data_center_zone        = %q`, workload.DataCenterZone))
	}

	if len(*workload.IgnoredInterfaceNames) > 0 {
		hcl.WriteString(fmt.Sprintf(`
  ignored_interface_names = ["%v"]`, strings.Join(*workload.IgnoredInterfaceNames, `, "`)))
	}

	if workload.ServicePrincipalName != "" {
		hcl.WriteString(fmt.Sprintf(`
  service_principal_name  = %q`, workload.ServicePrincipalName))
	}

	if workload.ServiceProvider != "" {
		hcl.WriteString(fmt.Sprintf(`
  service_provider        = %q`, workload.ServiceProvider))
	}

	for _, label := range *workload.Labels {
		hcl.WriteString(fmt.Sprintf(`
  labels {
    href = %q
  }`, label.Href))
	}

	if workload.ExternalDataSet != "" && workload.ExternalDataReference != "" {
		hcl.WriteString(fmt.Sprintf(`

  external_data_set       = %q
  external_data_reference = %q`, workload.ExternalDataSet, workload.ExternalDataReference))
	}
	hcl.WriteString(`
}
`)
	return hcl.String()
}

func buildVirtualServiceHCL(virtualService illumioapi.VirtualService) string {
	var hcl strings.Builder
	hclName := hclNormalize(virtualService.Name)
	address := fmt.Sprintf("illumio-core_virtual_service.%s", hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = virtualService.Href
	}

	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_virtual_service" %q {
  name         = %q
  description  = %q
  apply_to     = %q
`, hclName, virtualService.Name, virtualService.Description, virtualService.ApplyTo))

	if len(virtualService.IPOverrides) > 0 {
		hcl.WriteString(fmt.Sprintf(`
  ip_overrides = ["%v"]
`, strings.Join(virtualService.IPOverrides, `, "`)))
	}

	if virtualService.Service != nil {
		hcl.WriteString(fmt.Sprintf(`
  service {
    href = %q
  }`, virtualService.Service.Href))
	}

	for _, svcPort := range virtualService.ServicePorts {
		svcValueMap := map[string]int{
			"port": svcPort.Port, "to_port": svcPort.ToPort, "proto": svcPort.Protocol,
		}
		hcl.WriteString(`
  service_ports {`)
		for field, val := range svcValueMap {
			if val != 0 {
				hcl.WriteString(fmt.Sprintf(`
    %-7s = "%d"`, field, val))
			}
		}
		hcl.WriteString(`
  }`)
	}

	for _, svcAddress := range virtualService.ServiceAddresses {
		hcl.WriteString(`
  service_addresses {`)
		if svcAddress.Fqdn != "" {
			hcl.WriteString(fmt.Sprintf(`
    fqdn         = %q`, svcAddress.Fqdn))
		}
		if svcAddress.Description != "" {
			hcl.WriteString(fmt.Sprintf(`
    description  = %q`, svcAddress.Description))
		}
		if svcAddress.IP != "" {
			hcl.WriteString(fmt.Sprintf(`
    ip           = %q`, svcAddress.IP))
		}
		if svcAddress.Network != nil {
			hcl.WriteString(fmt.Sprintf(`
    network_href = %q`, svcAddress.Network.Href))
		}
		hcl.WriteString(`
  }`)
	}

	for _, label := range virtualService.Labels {
		hcl.WriteString(fmt.Sprintf(`
  labels {
    href = %q
  }`, label.Href))
	}

	if virtualService.ExternalDataSet != "" && virtualService.ExternalDataReference != "" {
		hcl.WriteString(fmt.Sprintf(`

  external_data_set       = %q
  external_data_reference = %q`, virtualService.ExternalDataSet, virtualService.ExternalDataReference))
	}
	hcl.WriteString(`
}
`)
	return hcl.String()
}

func buildContainerClusterHCL(containerCluster illumioapi.ContainerCluster) string {
	var hcl strings.Builder
	hclName := hclNormalize(containerCluster.Name)
	address := fmt.Sprintf("illumio-core_container_cluster.%s", hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = containerCluster.Href
	}

	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_container_cluster" %q {
  name         = %q
  description  = %q
}
`, hclName, containerCluster.Name, containerCluster.Description))

	// append workload profile subobjects
	matches := containerClusterRe.FindStringSubmatch(containerCluster.Href)
	if len(matches) > 1 {
		i := containerClusterRe.SubexpIndex("ClusterID")
		workloadProfiles, _, err := pce.GetContainerWkldProfiles(map[string]string{}, matches[i])
		handleError(err, false)
		for _, workloadProfile := range workloadProfiles {
			hcl.WriteString(buildContainerClusterWorkloadProfileHCL(workloadProfile, containerCluster.Href, containerCluster.Name))
		}
	}
	return hcl.String()
}

func buildContainerClusterWorkloadProfileHCL(clusterWorkloadProfile illumioapi.ContainerWorkloadProfile, clusterHref, clusterName string) string {
	var hcl strings.Builder
	hclName := hclNormalize(fmt.Sprintf("%s-%s", clusterName, clusterWorkloadProfile.Name))
	address := fmt.Sprintf("illumio-core_container_cluster_workload_profile.%s", hclName)
	if _, ok := tfStateMap[address]; !ok {
		tfImportMap[address] = clusterWorkloadProfile.Href
	}

	hcl.WriteString(fmt.Sprintf(`
resource "illumio-core_container_cluster_workload_profile" %q {
  container_cluster_href = %q
  name                   = %q
  description            = %q
  enforcement_mode       = %q
  managed                = %v`, hclName, clusterHref, clusterWorkloadProfile.Name,
		clusterWorkloadProfile.Description, clusterWorkloadProfile.EnforcementMode,
		*clusterWorkloadProfile.Managed))

	for _, label := range clusterWorkloadProfile.AssignLabels {
		hcl.WriteString(fmt.Sprintf(`
  assign_labels {
    href = %q
  }`, label.Href))
	}

	for _, label := range clusterWorkloadProfile.Labels {
		hcl.WriteString(fmt.Sprintf(`
  labels {
    key = %q`, label.Key))
		if label.Assignment.Href != "" {
			hcl.WriteString(fmt.Sprintf(`
    assignment {
      href  = %q
    }`, label.Assignment.Href))
		}
		hcl.WriteString(`
  }`)
	}

	hcl.WriteString(`
}
`)
	return hcl.String()
}
