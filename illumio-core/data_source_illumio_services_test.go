// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixSL string = "TF-ACC-SL"

func init() {
	resource.AddTestSweepers("services", &resource.Sweeper{
		Name: "services",
		F:    sweep("service", "name", prefixSL, "/orgs/%d/sec_policy/draft/services"),
	})
}

func TestAccIllumioSL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_services.sl_test"
	serviceName := acctest.RandomWithPrefix(prefixSL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSLDataSourceConfig_basic(serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioSLDataSourceConfig_exactMatch(serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", serviceName),
				),
			},
		},
	})
}

func slConfig(serviceName string) string {
	rName := acctest.RandomWithPrefix(prefixSL)

	return fmt.Sprintf(`
resource "illumio-core_service" "sl_test1" {
	name          = %[1]q
	description   = "Terraform Services test"

	service_ports {
		proto = 6
		port = 137
	}

	service_ports {
		proto = 6
		port = 138
	}
}

resource "illumio-core_service" "sl_test2" {
	name          = %[2]q
	description   = "Terraform Services test"

	service_ports {
		proto = 17
		port = 137
	}

	service_ports {
		proto = 17
		port = 138
	}
}
`, rName, serviceName)
}

func testAccCheckIllumioSLDataSourceConfig_basic(serviceName string) string {
	return slConfig(serviceName) + fmt.Sprintf(`
data "illumio-core_services" "sl_test" {
	# lookup based on partial match
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_service.sl_test1,
		illumio-core_service.sl_test2,
	]
}
`, prefixSL)
}

func testAccCheckIllumioSLDataSourceConfig_exactMatch(serviceName string) string {
	return slConfig(serviceName) + fmt.Sprintf(`
data "illumio-core_services" "sl_test" {
	# lookup using exact match
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_service.sl_test1,
		illumio-core_service.sl_test2,
	]
}
`, serviceName)
}
