package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerPairingProfile *schema.Provider

func TestAccIllumioPairingProfile_CreateUpdate(t *testing.T) {
	ppAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerPairingProfile),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerPairingProfile, "illumio-core_pairing_profile", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioPairingProfileConfig_basic("full"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioPairingProfileExists("illumio-core_pairing_profile.test", ppAttr),
					testAccCheckIllumioPairingProfileAttributes("full", ppAttr),
				),
			},
			{
				Config: testAccCheckIllumioPairingProfileConfig_basic("idle"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioPairingProfileExists("illumio-core_pairing_profile.test", ppAttr),
					testAccCheckIllumioPairingProfileAttributes("idle", ppAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioPairingProfileConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_pairing_profile" "test" {
		name    = "test_code_example"
		enabled = false
		label {
		  href = "/orgs/1/labels/1"
		}
		label {
		  href = "/orgs/1/labels/7"
		}
		allowed_uses_per_key  = "unlimited"
		env_label_lock        = false
		loc_label_lock        = true
		role_label_lock       = true
		app_label_lock        = true
		log_traffic           = false
		log_traffic_lock      = true
		visibility_level      = "flow_off"
		visibility_level_lock = false
		enforcement_mode      = "%s"
	  }
	`, val)
}

func testAccCheckIllumioPairingProfileExists(name string, ppAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Pairing Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerPairingProfile).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"enabled",
			"label.0.href",
			"label.1.href",
			"allowed_uses_per_key",
			"env_label_lock",
			"loc_label_lock",
			"role_label_lock",
			"app_label_lock",
			"log_traffic",
			"log_traffic_lock",
			"visibility_level",
			"visibility_level_lock",
			"enforcement_mode",
			"labels.0.href",
			"labels.1.href",
		} {
			ppAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioPairingProfileAttributes(val string, ppAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                  "test_code_example",
			"enabled":               false,
			"allowed_uses_per_key":  "unlimited",
			"env_label_lock":        false,
			"loc_label_lock":        true,
			"role_label_lock":       true,
			"app_label_lock":        true,
			"log_traffic":           false,
			"log_traffic_lock":      true,
			"visibility_level":      "flow_off",
			"visibility_level_lock": false,
			"enforcement_mode":      val,
			"labels.0.href":         "/orgs/1/labels/1",
			"labels.1.href":         "/orgs/1/labels/7",
		}

		for k, v := range expectation {
			if ppAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ppAttr[k], v)
			}
		}

		return nil
	}
}
