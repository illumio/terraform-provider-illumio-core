package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerOrgSettings *schema.Provider

func TestAccIllumioOrgSettings_CreateUpdate(t *testing.T) {
	wsAttr := map[string]interface{}{}

	var err error

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerOrgSettings),
		Steps: []resource.TestStep{
			{
				Config:             testAccCheckIllumioOrgSettingsConfig_basic(),
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      "illumio-core_organization_settings.test",
				ImportStateId:     "/orgs/1/settings/events",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"test_href",
				},
			},
			{
				Config: testAccCheckIllumioOrgSettingsConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioOrgSettingsExists("illumio-core_organization_settings.test", wsAttr),
					testAccCheckIllumioOrgSettingsAttributes(wsAttr, &err),
				),
			},
			{
				Config: testAccCheckIllumioOrgSettingsConfig_revert(),
				Check: resource.ComposeTestCheckFunc(
					throwErrorOccuredInValidationStep(err),
				),
			},
		},
	})
}

// Used for updating the settings
func testAccCheckIllumioOrgSettingsConfig_basic() string {
	return `
	resource "illumio-core_organization_settings" "test" {
		test_href = "/orgs/1/settings/events"
		audit_event_retention_seconds = 7775000
		audit_event_min_severity = "informational"
		format = "JSON"
	  }
	  `
}

// Used for reseting the settings
func testAccCheckIllumioOrgSettingsConfig_revert() string {
	return `
	resource "illumio-core_organization_settings" "test" {
		test_href = "/orgs/1/settings/events"
		audit_event_retention_seconds = 7776000
		audit_event_min_severity = "informational"
		format = "JSON"
	  }
	`
}

func testAccCheckIllumioOrgSettingsExists(name string, wsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Organization Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerOrgSettings).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"audit_event_retention_seconds",
			"audit_event_min_severity",
			"format",
		} {
			wsAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioOrgSettingsAttributes(wsAttr map[string]interface{}, err *error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"audit_event_retention_seconds": 7775000,
			"audit_event_min_severity":      "informational",
			"format":                        "JSON",
		}
		for k, v := range expectation {
			if wsAttr[k] != v {
				*err = fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, wsAttr[k], v)
				return nil
			}
		}

		return nil
	}
}
