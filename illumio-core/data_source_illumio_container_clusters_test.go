// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixCCL string = "TF-ACC-CCL"

func TestAccIllumioCCL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_container_clusters.ccl_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioCCLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixCCL)
	rName2 := acctest.RandomWithPrefix(prefixCCL)

	return fmt.Sprintf(`
resource "illumio-core_container_cluster" "ccl_test1" {
	name = %[1]q
	description = "Terraform Container Cluster test"
}

resource "illumio-core_container_cluster" "ccl_test2" {
	name = %[2]q
	description = "Terraform Container Cluster test"
}

data "illumio-core_container_clusters" "ccl_test" {
	# lookup based on partial match
	name = %[3]q

	# enforce dependencies
	depends_on = [
		illumio-core_container_cluster.ccl_test1,
		illumio-core_container_cluster.ccl_test2,
	]
}
`, rName1, rName2, prefixCCL)
}
