// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/illumio/terraform-provider-illumio-core/client"
	"github.com/illumio/terraform-provider-illumio-core/models"
	"golang.org/x/time/rate"
)

const (
	pceHostKey     = "pce_host"
	apiUsernameKey = "api_username"
	apiSecretKey   = "api_secret"
	timeoutKey     = "request_timeout"
	backoffTimeKey = "backoff_time"
	maxRetriesKey  = "max_retries"
	proxyURLKey    = "proxy_url"
	proxyCredsKey  = "proxy_creds"
	orgIDKey       = "org_id"
	insecureKey    = "insecure"
	caFileKey      = "ca_file"

	version = 1

	hrefFilename = "hrefs.csv"

	defaultorgID = 1
)

var (
	// lock to handle write to href
	fileMutex sync.Mutex
)

// Provider - Illumio Core Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			pceHostKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ILLUMIO_PCE_HOST", nil),
				Description: "Host URL of Illumio PCE. This can also be set by environment variable `ILLUMIO_PCE_HOST`",
			},
			apiUsernameKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ILLUMIO_API_KEY_USERNAME", nil),
				Description: "Username of API Key. This can also be set by environment variable `ILLUMIO_API_KEY_USERNAME`",
			},
			apiSecretKey: {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("ILLUMIO_API_KEY_SECRET", nil),
				Description: "Secret of API Key. This can also be set by environment variable `ILLUMIO_API_KEY_SECRET`",
			},
			insecureKey: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ILLUMIO_ALLOW_INSECURE_TLS", nil),
				Description: "Allow insecure TLS. Only `yes` will mark it insecure. This can also be set by environment variable `ILLUMIO_ALLOW_INSECURE_TLS`",
			},
			caFileKey: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ILLUMIO_CA_FILE", nil),
				Description: "The path to CA certificate file (PEM). In case, certificate is based on legacy CN instead of ASN, set env. variable `GODEBUG=x509ignoreCN=0`. This can also be set by environment variable `ILLUMIO_CA_FILE`",
			},
			orgIDKey: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     defaultorgID,
				Description: "ID of the Organization. Default value: 1",
			},
			timeoutKey: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     30,
				Description: "Timeout for HTTP requests. Default value: 30",
			},
			backoffTimeKey: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
				Description: "Backoff Time (in seconds) on getting 429 (Too Many Requests). " +
					"Default value: 10. Note: A default rate limit of 125 requests/min is already in place. " +
					"A jitter of 1-5 seconds will be added to backoff time to randomize backoff.",
			},
			maxRetriesKey: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     3,
				Description: "Maximum retries for an API request. Default value: 3",
			},
			proxyURLKey: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ILLUMIO_PROXY_URL", nil),
				Description: "Proxy Server URL with port number. This can also be set by environment variable `ILLUMIO_PROXY_URL`",
			},
			proxyCredsKey: {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ILLUMIO_PROXY_CREDENTIALS", nil),
				Description: "Proxy credential in format `username:password`. This can also be set by environment variable `ILLUMIO_PROXY_CREDENTIALS`",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"illumio-core_firewall_settings":                  resourceIllumioFirewallSettings(),
			"illumio-core_ip_list":                            resourceIllumioIPList(),
			"illumio-core_label_group":                        resourceIllumioLabelGroup(),
			"illumio-core_label":                              resourceIllumioLabel(),
			"illumio-core_pairing_keys":                       resourceIllumioPairingKeys(),
			"illumio-core_pairing_profile":                    resourceIllumioPairingProfile(),
			"illumio-core_security_rule":                      resourceIllumioSecurityRule(),
			"illumio-core_rule_set":                           resourceIllumioRuleSet(),
			"illumio-core_workload":                           resourceIllumioWorkload(),
			"illumio-core_managed_workload":                   resourceIllumioManagedWorkload(),
			"illumio-core_workloads_unpair":                   resourceIllumioWorkloadsUnpair(),
			"illumio-core_service":                            resourceIllumioService(),
			"illumio-core_syslog_destination":                 resourceIllumioSyslogDestination(),
			"illumio-core_virtual_service":                    resourceIllumioVirtualService(),
			"illumio-core_container_cluster":                  resourceIllumioContainerCluster(),
			"illumio-core_container_cluster_workload_profile": resourceIllumioContainerClusterWorkloadProfileWorkloadProfile(),
			"illumio-core_workload_interface":                 resourceIllumioWorkloadInterface(),
			"illumio-core_ven":                                resourceIllumioVEN(),
			"illumio-core_vens_unpair":                        resourceIllumioVENsUnpair(),
			"illumio-core_vens_upgrade":                       resourceIllumioVENsUpgrade(),
			"illumio-core_vulnerability_report":               resourceIllumioVulnerabilityReport(),
			"illumio-core_vulnerabilities":                    resourceIllumioVulnerabilities(),
			"illumio-core_traffic_collector_settings":         resourceIllumioTrafficCollectorSettings(),
			"illumio-core_workload_settings":                  resourceIllumioWorkloadSettings(),
			"illumio-core_organization_settings":              resourceIllumioOrganizationSettings(),
			"illumio-core_service_binding":                    resourceIllumioServiceBinding(),
			"illumio-core_enforcement_boundary":               resourceIllumioEnforcementBoundary(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"illumio-core_firewall_settings":                   datasourceIllumioFirewallSettings(),
			"illumio-core_ip_list":                             datasourceIllumioIPList(),
			"illumio-core_ip_lists":                            datasourceIllumioIPLists(),
			"illumio-core_label_group":                         datasourceIllumioLabelGroup(),
			"illumio-core_label_groups":                        datasourceIllumioLabelGroups(),
			"illumio-core_label":                               datasourceIllumioLabel(),
			"illumio-core_labels":                              datasourceIllumioLabels(),
			"illumio-core_pairing_profile":                     datasourceIllumioPairingProfile(),
			"illumio-core_pairing_profiles":                    datasourceIllumioPairingProfiles(),
			"illumio-core_security_rule":                       datasourceIllumioSecurityRule(),
			"illumio-core_security_rules":                      datasourceIllumioSecurityRules(),
			"illumio-core_rule_set":                            datasourceIllumioRuleSet(),
			"illumio-core_rule_sets":                           datasourceIllumioRuleSets(),
			"illumio-core_workload":                            datasourceIllumioWorkload(),
			"illumio-core_workloads":                           datasourceIllumioWorkloads(),
			"illumio-core_service":                             datasourceIllumioService(),
			"illumio-core_services":                            datasourceIllumioServices(),
			"illumio-core_syslog_destination":                  datasourceIllumioSyslogDestination(),
			"illumio-core_syslog_destinations":                 datasourceIllumioSyslogDestinations(),
			"illumio-core_virtual_service":                     datasourceIllumioVirtualService(),
			"illumio-core_virtual_services":                    datasourceIllumioVirtualServices(),
			"illumio-core_container_cluster":                   datasourceIllumioContainerCluster(),
			"illumio-core_container_clusters":                  datasourceIllumioContainerClusters(),
			"illumio-core_container_cluster_workload_profile":  datasourceIllumioContainerClusterWorkloadProfile(),
			"illumio-core_container_cluster_workload_profiles": datasourceIllumioContainerClusterWorkloadProfiles(),
			"illumio-core_workload_interface":                  datasourceIllumioWorkloadInterface(),
			"illumio-core_workload_interfaces":                 datasourceIllumioWorkloadInterfaces(),
			"illumio-core_ven":                                 datasourceIllumioVEN(),
			"illumio-core_vens":                                datasourceIllumioVENs(),
			"illumio-core_vulnerability_report":                datasourceIllumioVulnerabilityReport(),
			"illumio-core_vulnerability_reports":               datasourceIllumioVulnerabilityReports(),
			"illumio-core_vulnerability":                       datasourceIllumioVulnerability(),
			"illumio-core_vulnerabilities":                     datasourceIllumioVulnerabilities(),
			"illumio-core_traffic_collector_settings":          datasourceIllumioTrafficCollectorSettings(),
			"illumio-core_traffic_collector_settings_list":     datasourceIllumioTrafficCollectorSettingsList(),
			"illumio-core_workload_settings":                   datasourceIllumioWorkloadSettings(),
			"illumio-core_organization_settings":               datasourceIllumioOrganizationSettings(),
			"illumio-core_service_binding":                     datasourceIllumioServiceBinding(),
			"illumio-core_service_bindings":                    datasourceIllumioServiceBindings(),
			"illumio-core_enforcement_boundary":                datasourceIllumioEnforcementBoundary(),
			"illumio-core_enforcement_boundaries":              datasourceIllumioEnforcementBoundaries(),
			"illumio-core_container_cluster_service_backends":  datasourceIllumioContainerClusterServiceBackends(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diagnostics diag.Diagnostics
	if diagnostics = validateInput(d); diagnostics.HasError() {
		return nil, diagnostics
	}
	insecure := false
	if d.Get(insecureKey).(string) == "yes" {
		insecure = true
	}
	illumioV2Client, err := client.NewV2(
		d.Get(pceHostKey).(string),
		d.Get(apiUsernameKey).(string),
		d.Get(apiSecretKey).(string),
		d.Get(timeoutKey).(int),
		rate.NewLimiter(rate.Limit(float64(125)/float64(60)), 1), // limits API calls 125/min
		d.Get(backoffTimeKey).(int),
		d.Get(maxRetriesKey).(int),
		insecure,
		d.Get(caFileKey).(string),
		d.Get(proxyURLKey).(string),
		d.Get(proxyCredsKey).(string),
	)
	if err != nil {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create illumio client",
			Detail:   fmt.Sprintf("Unable to create illumio client. Error - %v", err),
		})
		return nil, diagnostics
	}
	providerConfig := Config{
		IllumioClient: illumioV2Client,
		OrgID:         d.Get(orgIDKey).(int),
		HrefFilename:  hrefFilename,
	}
	return providerConfig, diagnostics
}

func validateInput(d *schema.ResourceData) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	hostURL := d.Get(pceHostKey).(string)
	if hostURL == "" {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "PCE Host URL is required",
			Detail:   "PCE Host URL must be set for illumio provider",
		})
	}
	apiUsername := d.Get(apiUsernameKey).(string)
	if apiUsername == "" {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API Key Username is required",
			Detail:   "API Key Username must be set for illumio provider",
		})
	}
	apiKeySecret := d.Get(apiSecretKey).(string)
	if apiKeySecret == "" {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API Key Secret is required",
			Detail:   "API Key Secret must be set for illumio provider",
		})
	}
	orgID := d.Get(orgIDKey).(int)
	if orgID == 0 {
		diagnostics = append(diagnostics, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Org ID is required",
			Detail:   "Org ID must be set for illumio provider",
		})
	}
	return diagnostics
}

// Config Configuration for provider
type Config struct {
	IllumioClient *client.V2
	OrgID         int
	HrefFilename  string
}

// StoreHref - Writes href to hrefs.csv file for provisioning of resource
func (c Config) StoreHref(orgID int, resourceType, href string) {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	fp, err := os.OpenFile(c.HrefFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err == nil {
		_, err = fp.WriteString(fmt.Sprintf("%d,%s,%s\n", orgID, resourceType, href))
		if err != nil {
			panic(errors.New("couldn't write to file"))
		}
	} else {
		panic(errors.New("couldn't create file"))
	}
}

// ProvisionAResource - Provision a single resource
func (c Config) ProvisionAResource(orgID int, resourceType, href string) error {
	log.Printf("ProvisionAResource - %s", href)
	cs := models.SecurityPolicyChangeSubset{}
	cs.AppendHref(resourceType, href)
	secPolicy := &models.SecurityPolicy{
		UpdateDesc:   "Provisioned by Terraform",
		ChangeSubset: cs,
	}
	_, _, err := c.IllumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy", orgID), secPolicy)
	return err
}
