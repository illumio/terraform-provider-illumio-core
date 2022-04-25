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
}

resource "illumio-core_workload_interface" "wlil_test1" {
	workload_href = illumio-core_unmanaged_workload.wlil_test.href
	name = "eth0"
	friendly_name = "Terraform Workload Interface 1"
	link_state = "up"
}

resource "illumio-core_workload_interface" "wlil_test2" {
	workload_href = illumio-core_unmanaged_workload.wlil_test.href
	name = "eth1"
	friendly_name = "Terraform Workload Interface 2"
	link_state = "up"
}

data "illumio-core_workload_interfaces" "wlil_test" {
	workload_href = illumio-core_unmanaged_workload.wlil_test.href

	# enforce dependencies
	depends_on = [
		illumio-core_workload_interface.wlil_test1,
		illumio-core_workload_interface.wlil_test2,
	]
}
`, rName1)
}
