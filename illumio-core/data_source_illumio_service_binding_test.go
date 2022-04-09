// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixSB string = "TF-ACC-SB"

func TestAccIllumioSB_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_service_binding.sb_test"
	resourceName := "illumio-core_service_binding.sb_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSBDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "virtual_service", resourceName, "virtual_service"),
					resource.TestCheckResourceAttrPair(dataSourceName, "workload", resourceName, "workload"),
				),
			},
		},
	})
}

func testAccCheckIllumioSBDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixSB)
	rName2 := acctest.RandomWithPrefix(prefixSB)

	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "sb_test" {
	name = %[1]q
	description = "Terraform Service Binding test"
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

resource "illumio-core_workload" "sb_test" {
	name               = %[2]q
	description        = "Terraform Service Binding test"
	hostname           = "example.workload"
}

resource "illumio-core_service_binding" "sb_test" {
	virtual_service {
		href = illumio-core_virtual_service.sb_test.href
	}
	workload {
		href = illumio-core_workload.sb_test.href
	}
}

data "illumio-core_service_binding" "sb_test" {
	href = illumio-core_service_binding.sb_test.href
}
`, rName1, rName2)
}
