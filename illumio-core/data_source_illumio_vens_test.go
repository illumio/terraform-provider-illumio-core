package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSVENL *schema.Provider

func TestAccIllumioVENL_Read(t *testing.T) {
	listAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSVENL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVENLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceVENLExists("data.illumio-core_vens.test", listAttr),
					testAccCheckIllumioListDataSourceSize(listAttr, "5"),
				),
			},
		},
	})
}

func testAccCheckIllumioVENLDataSourceConfig_basic() string {
	return `
	data "illumio-core_vens" "test" {
		max_results = 5
	}
	`
}

func testAccCheckIllumioDataSourceVENLExists(name string, listAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of VENs %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		listAttr["length"] = rs.Primary.Attributes["items.#"]

		return nil
	}
}
