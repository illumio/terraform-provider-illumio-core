// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixVS string = "TF-ACC-VS"

func TestAccIllumioVS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_virtual_service.vs_test"
	resourceName := "illumio-core_virtual_service.vs_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "apply_to", resourceName, "apply_to"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_ports", resourceName, "service_ports"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_addresses", resourceName, "service_addresses"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_overrides", resourceName, "ip_overrides"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIllumioVSDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixVS)

	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "vs_test" {
	name = %[1]q
	description = "Terraform Virtual Service test"
	apply_to = "host_only"

	service_ports {
		proto = 6
		port = 1234
	}

	service_addresses {
		fqdn = "*.illumio.com"
	}

	ip_overrides = [
		"1.2.3.4"
	]
}

data "illumio-core_virtual_service" "vs_test" {
	href = illumio-core_virtual_service.vs_test.href
}
`, rName1)
}
