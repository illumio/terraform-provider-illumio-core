package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSVirtualServiceL *schema.Provider

func TestAccIllumioVirtualServiceL_Read(t *testing.T) {
	listAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSVirtualServiceL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVirtualServiceLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceVirtualServiceLExists("data.illumio-core_virtual_services.test", listAttr),
					testAccCheckIllumioListDataSourceSize(listAttr, "5"),
				),
			},
		},
	})
}

func testAccCheckIllumioVirtualServiceLDataSourceConfig_basic() string {
	return `
	data "illumio-core_virtual_services" "test" {
		max_results = 5
	}
	`
}

func testAccCheckIllumioDataSourceVirtualServiceLExists(name string, listAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Virtual Services %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		listAttr["length"] = rs.Primary.Attributes["items.#"]

		return nil
	}
}
