// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixVSL string = "TF-ACC-VSL"

func init() {
	resource.AddTestSweepers("virtual_services", &resource.Sweeper{
		Name: "virtual_services",
		F:    sweep("virtual service", "name", prefixVSL, "/orgs/%d/sec_policy/draft/virtual_services"),
		Dependencies: []string{
			"enforcement_boundaries",
			"rule_sets",
		},
	})
}

func TestAccIllumioVSL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_virtual_services.vsl_test"
	virtualServiceName := acctest.RandomWithPrefix(prefixVSL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVSLDataSourceConfig_basic(virtualServiceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioVSLDataSourceConfig_exactMatch(virtualServiceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", virtualServiceName),
				),
			},
		},
	})
}

func vslConfig(virtualServiceName string) string {
	rName := acctest.RandomWithPrefix(prefixVSL)

	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "vsl_test1" {
	name = %[1]q
	description = "Terraform Virtual Services test"
	apply_to = "host_only"

	service_ports {
		proto = 6
		port = 8080
	}
}

resource "illumio-core_virtual_service" "vsl_test2" {
	name = %[2]q
	description = "Terraform Virtual Services test"
	apply_to = "host_only"

	service_ports {
		proto = 6
		port = 8443
	}
}
`, rName, virtualServiceName)
}

func testAccCheckIllumioVSLDataSourceConfig_basic(virtualServiceName string) string {
	return vslConfig(virtualServiceName) + fmt.Sprintf(`
data "illumio-core_virtual_services" "vsl_test" {
	# lookup based on partial match
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_virtual_service.vsl_test1,
		illumio-core_virtual_service.vsl_test2,
	]
}
`, prefixVSL)
}

func testAccCheckIllumioVSLDataSourceConfig_exactMatch(virtualServiceName string) string {
	return vslConfig(virtualServiceName) + fmt.Sprintf(`
data "illumio-core_virtual_services" "vsl_test" {
	# lookup using exact match
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_virtual_service.vsl_test1,
		illumio-core_virtual_service.vsl_test2,
	]
}
`, virtualServiceName)
}
