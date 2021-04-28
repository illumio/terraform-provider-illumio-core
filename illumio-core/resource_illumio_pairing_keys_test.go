package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerPKs *schema.Provider

func TestAccIllumioPairingKeys_CreateUpdate(t *testing.T) {
	pksAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerPKs),
		// CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerPKs, "illumio_pairing_keys", true),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioPairingKeysConfig_basic(2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioPairingKeysExists("illumio_pairing_keys.test", pksAttr),
					testAccCheckIllumioPairingKeysAttributes(2, pksAttr),
				),
			},
			{
				Config: testAccCheckIllumioPairingKeysConfig_basic(3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioPairingKeysExists("illumio_pairing_keys.test", pksAttr),
					testAccCheckIllumioPairingKeysAttributes(3, pksAttr),
				),
			},
			{
				Config: testAccCheckIllumioPairingKeysConfig_basic(1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioPairingKeysExists("illumio_pairing_keys.test", pksAttr),
					testAccCheckIllumioPairingKeysAttributes(1, pksAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioPairingKeysConfig_basic(val int) string {
	return fmt.Sprintf(`
	resource "illumio_pairing_keys" "test" {
		pairing_profile_href = "/orgs/1/pairing_profiles/1"
		token_count = %d
	}
	`, val)
}

func testAccCheckIllumioPairingKeysExists(name string, attr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Pairing Keys %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}
		attr["activation_tokens.#"] = rs.Primary.Attributes["activation_tokens.#"]
		return nil
	}
}

func testAccCheckIllumioPairingKeysAttributes(val int, attr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"activation_tokens.#": fmt.Sprint(val),
		}
		for k, v := range expectation {
			if attr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, attr[k], v)
			}
		}

		return nil
	}
}
