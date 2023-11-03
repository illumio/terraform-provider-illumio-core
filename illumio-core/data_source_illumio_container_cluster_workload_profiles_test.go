// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixCCWPL string = "TF-ACC-CCWPL"

// no sweeper needed - the container cluster sweeper will remove all associated profiles

func TestAccIllumioCCWPL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_container_cluster_workload_profiles.ccwpl_test"
	ccwpName := acctest.RandomWithPrefix(prefixCCWPL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCWPLConfig_basic(ccwpName),
				Check: resource.ComposeTestCheckFunc(
					// Container clusters have a Default Workload Profile, included
					// with the two created below there should be 3 total
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "3"),
				),
			},
			{
				Config: testAccCheckIllumioCCWPLConfig_partialMatchName(ccwpName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioCCWPLConfig_exactMatchName(ccwpName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", ccwpName),
				),
			},
		},
	})
}

func ccwplConfig(ccwpName string) string {
	labelName := acctest.RandomWithPrefix(prefixLL)
	clusterName := acctest.RandomWithPrefix(prefixCCL)
	rName := acctest.RandomWithPrefix(prefixCCWPL)

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
	name                   = %[3]q
	description            = "Terraform Container Cluster Workload Profile test"
	managed                = true
	enforcement_mode       = "visibility_only"

	assign_labels {
		href = illumio-core_label.ccwpl_test.href
	}
}

resource "illumio-core_container_cluster_workload_profile" "ccwpl_test2" {
	container_cluster_href = illumio-core_container_cluster.ccwpl_test.href
	name                   = %[4]q
	description            = "Terraform Container Cluster Workload Profile test"
	managed                = true
	enforcement_mode       = "visibility_only"

	assign_labels {
		href = illumio-core_label.ccwpl_test.href
	}
}
`, labelName, clusterName, rName, ccwpName)
}

func testAccCheckIllumioCCWPLConfig_basic(ccwpName string) string {
	return ccwplConfig(ccwpName) + fmt.Sprintf(`
data "illumio-core_container_cluster_workload_profiles" "ccwpl_test" {
	# lookup using just cluster HREF
	container_cluster_href = illumio-core_container_cluster.ccwpl_test.href

	# enforce dependencies
	depends_on = [
		illumio-core_container_cluster_workload_profile.ccwpl_test1,
		illumio-core_container_cluster_workload_profile.ccwpl_test2,
	]
}
`)
}

func testAccCheckIllumioCCWPLConfig_partialMatchName(ccwpName string) string {
	return ccwplConfig(ccwpName) + fmt.Sprintf(`
data "illumio-core_container_cluster_workload_profiles" "ccwpl_test" {
	# lookup based on partial name match
	container_cluster_href = illumio-core_container_cluster.ccwpl_test.href
	name                   = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_container_cluster_workload_profile.ccwpl_test1,
		illumio-core_container_cluster_workload_profile.ccwpl_test2,
	]
}
`, prefixCCWPL)
}

func testAccCheckIllumioCCWPLConfig_exactMatchName(ccwpName string) string {
	return ccwplConfig(ccwpName) + fmt.Sprintf(`
data "illumio-core_container_cluster_workload_profiles" "ccwpl_test" {
	# lookup name using exact match
	container_cluster_href = illumio-core_container_cluster.ccwpl_test.href
	name                   = %[1]q
	match_type             = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_container_cluster_workload_profile.ccwpl_test1,
		illumio-core_container_cluster_workload_profile.ccwpl_test2,
	]
}
`, ccwpName)
}
