// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixCCWPL string = "TF-ACC-CCWPL"

func TestAccIllumioCCWPL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_container_cluster_workload_profiles.ccwpl_test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCWPLConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					// Container clusters have a Default Workload Profile, included
					// with the two created below there should be 3 total
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIllumioCCWPLConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixCCWPL)
	rName2 := acctest.RandomWithPrefix(prefixCCWPL)
	rName3 := acctest.RandomWithPrefix(prefixCCWPL)
	rName4 := acctest.RandomWithPrefix(prefixCCWPL)

	return fmt.Sprintf(`
resource "illumio-core_label" "ccwpl_test" {
	key   = "app"
	value = %[1]q
}

resource "illumio-core_container_cluster" "ccwpl_test" {
	name = %[2]q
	description = "Terraform Container Cluster Workload Profile test"
}

resource "illumio-core_container_cluster_workload_profile" "ccwpl_test1" {
	container_cluster_href = illumio-core_container_cluster.ccwpl_test.href
	name = %[3]q
	description = "Terraform Container Cluster Workload Profile test"
	managed = true
	enforcement_mode = "visibility_only"

	assign_labels {
		href = illumio-core_label.ccwpl_test.href
	}
}

resource "illumio-core_container_cluster_workload_profile" "ccwpl_test2" {
	container_cluster_href = illumio-core_container_cluster.ccwpl_test.href
	name = %[4]q
	description = "Terraform Container Cluster Workload Profile test"
	managed = true
	enforcement_mode = "visibility_only"

	assign_labels {
		href = illumio-core_label.ccwpl_test.href
	}
}

data "illumio-core_container_cluster_workload_profiles" "ccwpl_test" {
	# lookup based on partial match
	container_cluster_href = illumio-core_container_cluster.ccwpl_test.href

	# enforce dependencies
	depends_on = [
		illumio-core_container_cluster_workload_profile.ccwpl_test1,
		illumio-core_container_cluster_workload_profile.ccwpl_test2,
	]
}
`, rName1, rName2, rName3, rName4)
}
