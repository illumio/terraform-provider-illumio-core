package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerFirewallSettings *schema.Provider

func TestAccIllumioFirewallSettings_CreateUpdate(t *testing.T) {
	attrs := map[string]interface{}{}

	var err error

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerFirewallSettings),
		Steps: []resource.TestStep{
			{
				Config:             testAccCheckIllumioFirewallSettingsConfig_basic(),
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      "illumio-core_firewall_settings.test",
				ImportStateId:     "/orgs/1/sec_policy/draft/firewall_settings",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"test_href",
				},
			},
			{
				Config: testAccCheckIllumioFirewallSettingsConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioFirewallSettingsExists("illumio-core_firewall_settings.test", attrs),
					testAccCheckIllumioFirewallSettingsAttributes(attrs, &err),
				),
			},
			{
				Config: testAccCheckIllumioFirewallSettingsConfig_revert(),
				Check: resource.ComposeTestCheckFunc(
					throwErrorOccuredInValidationStep(err),
				),
			},
		},
	})
}

// Used for updating the settings
func testAccCheckIllumioFirewallSettingsConfig_basic() string {
	return `
	resource "illumio-core_firewall_settings" "test" {
		test_href = "/orgs/1/sec_policy/draft/firewall_settings"
		ike_authentication_type = "certificate"
	  }
	  `
}

// Used for reseting the settings
func testAccCheckIllumioFirewallSettingsConfig_revert() string {
	return `
	resource "illumio-core_firewall_settings" "test" {
		test_href = "/orgs/1/sec_policy/draft/firewall_settings"
		ike_authentication_type = "psk"
	  }
	`
}

func testAccCheckIllumioFirewallSettingsExists(name string, attrs map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Firewall Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerFirewallSettings).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"ike_authentication_type",
		} {
			attrs[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioFirewallSettingsAttributes(attrs map[string]interface{}, err *error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"ike_authentication_type": "certificate",
		}
		for k, v := range expectation {
			if attrs[k] != v {
				*err = fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, attrs[k], v)
				return nil
			}
		}

		return nil
	}
}
