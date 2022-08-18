// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixCCWP string = "TF-ACC-CCWP"

func TestAccIllumioCCWP_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_container_cluster_workload_profile.ccwp_test"
	resourceName := "illumio-core_container_cluster_workload_profile.ccwp_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCWPDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "namespace", resourceName, "namespace"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enforcement_mode", resourceName, "enforcement_mode"),
					resource.TestCheckResourceAttrPair(dataSourceName, "visibility_level", resourceName, "visibility_level"),
					resource.TestCheckResourceAttrPair(dataSourceName, "managed", resourceName, "managed"),
					resource.TestCheckResourceAttrPair(dataSourceName, "linked", resourceName, "linked"),
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

func testAccCheckIllumioCCWPDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixCCWP)
	rName2 := acctest.RandomWithPrefix(prefixCCWP)
	rName3 := acctest.RandomWithPrefix(prefixCCWP)

	return fmt.Sprintf(`
resource "illumio-core_label" "ccwp_test" {
	key   = "role"
	value = %[1]q
}

resource "illumio-core_container_cluster" "ccwp_test" {
	name = %[2]q
	description = "Terraform Container Cluster test"
}

resource "illumio-core_container_cluster_workload_profile" "ccwp_test" {
	container_cluster_href = illumio-core_container_cluster.ccwp_test.href
	name = %[3]q
	description = "Terraform Container Cluster Workload Profile test"
	managed = true
	enforcement_mode = "visibility_only"

	assign_labels {
		href = illumio-core_label.ccwp_test.href
	}
}

data "illumio-core_container_cluster_workload_profile" "ccwp_test" {
	href = illumio-core_container_cluster_workload_profile.ccwp_test.href
}
`, rName1, rName2, rName3)
}
