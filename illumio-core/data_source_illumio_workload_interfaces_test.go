// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixWLIL string = "TF-ACC-WLIL"

func TestAccIllumioWLIL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_workload_interfaces.wlil_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWLILDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioWLILDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixWLIL)

	return fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" "wlil_test" {
	name               = %[1]q
	description        = "Terraform Workload Interfaces test"
	hostname           = "example.workload"

	interfaces {
		name          = "eth0"
		address       = "172.22.1.14"
		friendly_name = "Wired Network (Ethernet)"
		link_state    = "up"
	}

	interfaces {
		name          = "lo0"
		address       = "127.0.0.1"
		friendly_name = "Loopback Interface"
		link_state    = "up"
	}
}

data "illumio-core_workload_interfaces" "wlil_test" {
	workload_href = illumio-core_unmanaged_workload.wlil_test.href
}
`, rName1)
}
