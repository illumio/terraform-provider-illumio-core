// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/client"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

// convenience method to convert int literals into int pointer
func intPointer(number int) *int { return &number }

// encrypt using AES-GCM algorithm
func aesGcmEncrypt(key, plaintext string) (string, string, error) {
	k, err := hex.DecodeString(key)
	if err != nil {
		return "", "", errors.New("could not decode AES GCM key")
	}
	text := []byte(plaintext)

	block, err := aes.NewCipher(k)
	if err != nil {
		return "", "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}
	nonce := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", errors.New("could not generate nonce")
	}

	// nonce is added as prefix so encrypted string stores the nonce along with ciphertext
	ciphertext := aesgcm.Seal(nil, nonce, text, nil)
	return fmt.Sprintf("%x", ciphertext), fmt.Sprintf("%x", nonce), nil
}

func hashcode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

// for all resource, name attribute has character limit from 1 to 255
var nameValidation = validation.ToDiagFunc(validation.StringLenBetween(1, 255))

// for all resource, name attribute has character limit from 0 to 255
var checkStringZerotoTwoHundredAndFiftyFive = validation.ToDiagFunc(validation.StringLenBetween(0, 255))

var uuidV4RegEx = "[a-f0-9]{8}-?[a-f0-9]{4}-?4[a-f0-9]{3}-?[89ab][a-f0-9]{3}-?[a-f0-9]{12}"
var orgsPrefix = "^/orgs/[1-9][0-9]*/"
var secPMefix = "sec_policy/(draft|active|[0-9]*)/"

var isLabelHref = generateHrefValidationFunction("labels/[1-9][0-9]*", "Label")
var isLabelGroupHref = generateHrefValidationFunction(secPMefix+"label_groups/"+uuidV4RegEx, "Label Group")
var isIPListHref = generateHrefValidationFunction(secPMefix+"ip_lists/[1-9][0-9]*", "IP List")
var isServiceHref = generateHrefValidationFunction(secPMefix+"services/[1-9][0-9]*", "Service")
var isVirtualServiceHref = generateHrefValidationFunction(secPMefix+"virtual_services/"+uuidV4RegEx, "Virtual Service")
var isWorkloadHref = generateHrefValidationFunction("workloads/"+uuidV4RegEx, "Workload")
var isPairingProfileHref = generateHrefValidationFunction("pairing_profiles/[1-9][0-9]*", "Pairing Profile")
var isVulnerabilityHref = generateHrefValidationFunction("vulnerabilities/.*", "Vulnerability")
var isVENHref = generateHrefValidationFunction("vens/"+uuidV4RegEx, "VEN")
var isContainerClusterHref = generateHrefValidationFunction("container_clusters/"+uuidV4RegEx, "Container Cluster")
var isContainerClusterWorkloadProfileHref = generateHrefValidationFunction("container_clusters/"+uuidV4RegEx+"/container_workload_profiles/"+uuidV4RegEx, "Container Cluster Workload Profile")
var isEnforcementBoundaryHref = generateHrefValidationFunction(secPMefix+"enforcement_boundaries/[1-9][0-9]*", "Enforcement Boundary")
var isRuleSetHref = generateHrefValidationFunction(secPMefix+"rule_sets/[1-9][0-9]*", "Rule Set")
var isSecurityRuleHref = generateHrefValidationFunction(secPMefix+"rule_sets/[1-9][0-9]*/sec_rules/[1-9][0-9]*", "Security Rule")
var isFirewallSettingsHref = generateHrefValidationFunction(secPMefix+"firewall_settings", "Firewall Settings")
var isWorkloadInterfaceHref = generateHrefValidationFunction("workloads/"+uuidV4RegEx+"/interfaces/.*", "Workload Interface")
var isVulnerabilityReportHref = generateHrefValidationFunction("vulnerability_reports/.*", "Vulnerability Report")
var isTrafficCollectorSettingsHref = generateHrefValidationFunction("settings/traffic_collector/"+uuidV4RegEx, "Traffic Collector Settings")
var isSyslogDestinationHref = generateHrefValidationFunction("settings/syslog/destinations/"+uuidV4RegEx, "Syslog Destination")
var isServiceBindingHref = generateHrefValidationFunction("service_bindings/"+uuidV4RegEx, "Service Binding")

func generateHrefValidationFunction(regex string, msg string) schema.SchemaValidateDiagFunc {
	return validation.ToDiagFunc(
		validation.StringMatch(
			regexp.MustCompile(
				orgsPrefix+
					regex+
					"$",
			),
			fmt.Sprintf("%v href is not in the correct format", msg),
		),
	)
}

// hrefSchemaRequired returns Href resource as required
func hrefSchemaRequired(rName string, diagValid schema.SchemaValidateDiagFunc) *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: diagValid,
				Description:      fmt.Sprintf("URI of %v", rName),
			},
		},
	}
}

// hrefSchemaComputed returns Href resource as computed
func hrefSchemaComputed(rName string, diagValid schema.SchemaValidateDiagFunc) *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: fmt.Sprintf("URI of %v", rName),
			},
		},
	}
}

// Constructs Href object from href resource schema, supports both TypeList(Maxitem = 1) and regulare type
func getHrefObj(obj interface{}) *models.Href {
	switch obj.(type) {
	case []interface{}: // TypeList
		l := obj.([]interface{})
		if len(l) > 0 {
			m := l[0].(map[string]interface{})
			return &models.Href{Href: m["href"].(string)}
		} else {
			return &models.Href{}
		}

	case map[string]interface{}:
		m := obj.(map[string]interface{})
		return &models.Href{Href: m["href"].(string)}
	default:
		return &models.Href{}
	}
}

// Returns string list from interface type
func getStringList(o interface{}) []string {
	i := o.([]interface{})
	list := []string{}
	for _, v := range i {
		list = append(list, v.(string))
	}
	return list
}

// Checks if all elements of list is in oneOf or not
func validateList(list []string, oneOf []string) bool {
	for _, v := range list {
		if !contains(oneOf, v) {
			return false
		}
	}
	return true
}

// contains checks if element is present in given array
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func isStringABoolean() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics

		if v.(string) != "true" && v.(string) != "false" {
			diags = append(diags, diag.Errorf("expected boolean values (true or false), got %v", v)...)
		}

		return diags
	}
}

func isStringANumber() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics

		if _, k := getInt(v); !k {
			diags = append(diags, diag.Errorf("expected integer, got %v", v)...)
		}

		return diags
	}
}

func isStringInRange(min, max int) schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		k, err := strconv.Atoi(v.(string))

		if err != nil {
			diags = append(diags, diag.Errorf("expected integer value, got: %v", v)...)
			return diags
		}

		if min > k || k > max {
			diags = append(diags, diag.Errorf("expected to be in range %v-%v, got: %v", min, max, v)...)
		}

		return diags
	}
}

func isStringAPortNumber() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		k, err := strconv.Atoi(v.(string))

		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "Not an integer",
				Detail:        fmt.Sprintf("expected an integer, got: %s", v.(string)),
				AttributePath: path,
			})
			return diags
		}

		if 0 > k || k > 65535 {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "Invalid Port Number",
				Detail:        fmt.Sprintf("Port Number should be in the range (0 - 65535): %s", v.(string)),
				AttributePath: path,
			})
		}

		return diags
	}
}

// getInt Returns int from different interface{}
func getInt(v interface{}) (int, bool) {
	switch v := v.(type) {
	case float64:
		in := int(v)
		return in, true

	case string:
		if v == "" {
			return 0, false
		}
		in, err := strconv.Atoi(v)
		return in, err == nil

	case int:
		return v, true

	default:
		return 0, false
	}

}

// validation function for checking "unlimited" or range
func isUnlimitedOrValidRange(min, max int) schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		if v.(string) == "unlimited" {
			return diags

		} else if k, err := strconv.Atoi(v.(string)); err == nil {
			if min > k || k > max {
				diags = append(diags, diag.Errorf("expected to be in range %v-%v, got: %v", min, max, v)...)
			}
		} else {
			diags = append(diags, diag.Errorf(`expected to be one of "unlimited" or range %v-%v, got: %v`, min, max, v)...)
		}

		return diags
	}
}

func isUpdatedToEmptyString(oldVal, newVal interface{}) bool {
	if oldVal.(string) != "" && (newVal == nil || newVal.(string) == "") {
		return true
	}
	return false
}

func extractMapArray(data *gabs.Container, keys []string) []map[string]interface{} {
	ms := []map[string]interface{}{}

	for _, child := range data.Children() {
		ms = append(ms, extractMap(child, keys))
	}

	return ms
}

func extractMap(data *gabs.Container, keys []string) map[string]interface{} {
	m := map[string]interface{}{}

	for _, k := range keys {
		if data.Exists(k) && data.S(k).Data() != nil {
			m[k] = data.S(k).Data()
		} else {
			m[k] = nil
		}
	}

	return m
}

func extractResourceScopes(data *gabs.Container) []map[string]interface{} {

	ms := []map[string]interface{}{}

	for _, data := range data.Children() {
		labels := []map[string]interface{}{}
		labelGroups := []map[string]interface{}{}

		for _, data := range data.Children() {
			for k, v := range data.ChildrenMap() {
				if k == "label" {
					labels = append(labels, v.Data().(map[string]interface{}))
				} else if k == "label_group" {
					labelGroups = append(labelGroups, v.Data().(map[string]interface{}))
				}
			}
		}

		m := map[string]interface{}{}
		m["label"] = labels
		m["label_group"] = labelGroups

		ms = append(ms, m)
	}
	return ms
}

func paramsString(p map[string]string) string {
	var b strings.Builder
	b.Grow(32)

	for _, v := range p {
		fmt.Fprintf(&b, "%v", v)
	}

	return b.String()
}

func resourceDataToMap(d *schema.ResourceData, keys []string) map[string]string {
	m := map[string]string{}
	for _, k := range keys {
		if v, ok := d.GetOk(k); ok {
			m[k] = fmt.Sprintf("%v", v)
		}
	}

	return m
}

func handleUnpairAndUpgradeOperationErrors(e error, res *http.Response, op, r string) diag.Diagnostics {
	var diags diag.Diagnostics

	if e != nil {
		diags = append(diags, diag.Errorf(fmt.Sprintf("[illumio-core_%v_%v] %v", r, op, e.Error()))...)
	} else {
		container, err := client.GetContainer(res)
		if err == nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("[illumio-core_%v_%v] Got possible failure/s in responce", r, op),
				Detail:   container.String(),
			})
		}
	}

	return diags
}

func extractResourceRuleActors(data *gabs.Container) []map[string]interface{} {
	actors := []map[string]interface{}{}

	validRuleActors := []string{
		"label",
		"label_group",
		"workload",
		"virtual_service",
		"virtual_server",
		"ip_list",
	}

	for _, actorArray := range data.Children() {

		actor := map[string]interface{}{}
		for k, v := range actorArray.ChildrenMap() {
			if k == "actors" {
				actor[k] = v.Data().(string)
			} else if contains(validRuleActors, k) {
				vM := v.Data().(map[string]interface{})

				hrefs := map[string]string{}
				hrefs["href"] = vM["href"].(string)

				actor[k] = []map[string]string{hrefs}
			}
		}
		actors = append(actors, actor)
	}

	return actors
}

// extractDataSourceAttrs - extracts list of objects from container using element keys
func extractDataSourceAttrs(data *gabs.Container, key string, elementKeys []string) []map[string]interface{} {
	if data.Exists(key) {
		elements := []map[string]interface{}{}

		for _, elem := range data.S(key).Children() {
			elements = append(elements, extractMap(elem, elementKeys))
		}

		return elements
	}
	return nil
}

func isValidPversion() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics

		sv := v.(string)
		iv, k := getInt(v)
		if !k || (k && iv < 1) {
			if sv == "active" || sv == "draft" {
				return diags
			} else {
				diags = append(diags, diag.Errorf(`expected an integer greater than 0 or one of ["active", "draft"], got %v`, v)...)
			}
		}

		return diags
	}
}

func isStringGreaterThanZero() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics

		if i, k := getInt(v); !k || i < 1 {
			diags = append(diags, diag.Errorf("expected non-zero positive integer, got %v", v)...)
		}

		return diags
	}
}
