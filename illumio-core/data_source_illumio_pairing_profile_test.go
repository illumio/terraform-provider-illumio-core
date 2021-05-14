package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSPP *schema.Provider

func TestAccIllumioPP_Read(t *testing.T) {
	ppAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSPP),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioPPDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourcePPExists("data.illumio-core_pairing_profile.test", ppAttr),
					testAccCheckIllumioPPDataSourceAttributes(ppAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioPPDataSourceConfig_basic() string {
	return `
	data "illumio-core_pairing_profile" "test" {
		href = "/orgs/1/pairing_profiles/32"
	}
	`
}

func testAccCheckIllumioDataSourcePPExists(name string, ppAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Pairing Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSPP).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"total_use_count",
			"enabled",
			"is_default",
		} {
			ppAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioPPDataSourceAttributes(ppAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":            "Acc. test pp name",
			"description":     "Acc. test pp description",
			"total_use_count": float64(0),
			"enabled":         true,
			"is_default":      false,
		}
		for k, v := range expectation {
			if ppAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ppAttr[k], v)
			}
		}

		return nil
	}
}
