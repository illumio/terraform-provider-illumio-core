// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixWLL string = "TF-ACC-WLL"

func init() {
	resource.AddTestSweepers("workloads", &resource.Sweeper{
		Name: "workloads",
		F:    sweep("workload", "name", prefixWLL, "/orgs/%d/workloads"),
	})
}

func TestAccIllumioWLL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_workloads.wll"
	workloadName := acctest.RandomWithPrefix(prefixWLL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWLLDataSourceConfig_basic(workloadName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioWLLDataSourceConfig_exactMatch(workloadName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", workloadName),
				),
			},
		},
	})
}

func wllConfig(workloadName string) string {
	rName := acctest.RandomWithPrefix(prefixWLL)

	return fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" "wll1" {
	name               = %[1]q
	description        = "Terraform Workloads test 1"
	hostname           = "jumpbox1"
}

resource "illumio-core_unmanaged_workload" "wll2" {
	name               = %[2]q
	description        = "Terraform Workloads test 2"
	hostname           = "jumpbox2"
}
`, rName, workloadName)
}

func testAccCheckIllumioWLLDataSourceConfig_basic(workloadName string) string {
	return wllConfig(workloadName) + fmt.Sprintf(`
data "illumio-core_workloads" "wll" {
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_unmanaged_workload.wll1,
		illumio-core_unmanaged_workload.wll2,
	]
}
`, prefixWLL)
}

func testAccCheckIllumioWLLDataSourceConfig_exactMatch(workloadName string) string {
	return wllConfig(workloadName) + fmt.Sprintf(`
data "illumio-core_workloads" "wll" {
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_unmanaged_workload.wll1,
		illumio-core_unmanaged_workload.wll2,
	]
}
`, workloadName)
}
