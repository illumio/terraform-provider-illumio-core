// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixWLI string = "TF-ACC-WLI"

func TestAccIllumioWLI_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_workload_interface.wli_test"
	resourceName := "illumio-core_workload_interface.wli_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWLIDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "friendly_name", resourceName, "friendly_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "link_state", resourceName, "link_state"),
				),
			},
		},
	})
}

func testAccCheckIllumioWLIDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixWLI)

	return fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" "wli_test" {
	name               = %[1]q
	description        = "Terraform Workload Interface test"
	hostname           = "example.workload"
}

resource "illumio-core_workload_interface" "wli_test" {
	workload_href = illumio-core_unmanaged_workload.wli_test.href
	name = "eth0"
	friendly_name = "Terraform Workload Interface"
	link_state = "up"
}

data "illumio-core_workload_interface" "wli_test" {
	href = illumio-core_workload_interface.wli_test.href
}
`, rName1)
}
