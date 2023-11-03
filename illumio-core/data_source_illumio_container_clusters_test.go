// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixCCL string = "TF-ACC-CCL"

func init() {
	resource.AddTestSweepers("container_clusters", &resource.Sweeper{
		Name: "container_clusters",
		F:    sweep("contain cluster", "name", prefixCCL, "/orgs/%d/container_clusters"),
	})
}

func TestAccIllumioCCL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_container_clusters.ccl_test"
	ccName := acctest.RandomWithPrefix(prefixCCL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCLDataSourceConfig_basic(ccName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioCCLDataSourceConfig_exactMatch(ccName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", ccName),
				),
			},
		},
	})
}

func cclConfig(ccName string) string {
	rName := acctest.RandomWithPrefix(prefixCCL)

	return fmt.Sprintf(`
resource "illumio-core_container_cluster" "ccl_test1" {
	name = %[1]q
	description = "Terraform Container Cluster test"
}

resource "illumio-core_container_cluster" "ccl_test2" {
	name = %[2]q
	description = "Terraform Container Cluster test"
}
`, rName, ccName)
}

func testAccCheckIllumioCCLDataSourceConfig_basic(ccName string) string {
	return cclConfig(ccName) + fmt.Sprintf(`
data "illumio-core_container_clusters" "ccl_test" {
	# lookup based on partial match
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_container_cluster.ccl_test1,
		illumio-core_container_cluster.ccl_test2,
	]
}
`, prefixCCL)
}

func testAccCheckIllumioCCLDataSourceConfig_exactMatch(ccName string) string {
	return cclConfig(ccName) + fmt.Sprintf(`
data "illumio-core_container_clusters" "ccl_test" {
	# lookup using exact match
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_container_cluster.ccl_test1,
		illumio-core_container_cluster.ccl_test2,
	]
}
`, ccName)
}
