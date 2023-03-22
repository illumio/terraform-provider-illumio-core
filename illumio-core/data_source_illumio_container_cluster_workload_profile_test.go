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

	ccName := acctest.RandomWithPrefix(prefixCCWP)
	ccwpName := acctest.RandomWithPrefix(prefixCCWP)
	labelName := acctest.RandomWithPrefix(prefixCCWP)
	updatedName := acctest.RandomWithPrefix(prefixCCWP)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCWPDataSourceConfig_basic(ccwpName, ccName, labelName),
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
				Config: testAccCheckIllumioCCWPResource_updateToLabels(ccwpName, ccName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "labels.0.key", "role"),
					resource.TestCheckResourceAttr(resourceName, "labels.0.assignment.0.value", labelName),
				),
			},
			{
				Config: testAccCheckIllumioCCWPResource_updateNameAndDesc(updatedName, ccName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				Config: testAccCheckIllumioCCWPResource_updateManageStateAndEnforcement(updatedName, ccName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "managed", "false"),
					resource.TestCheckResourceAttr(resourceName, "enforcement_mode", "idle"),
					resource.TestCheckResourceAttr(resourceName, "assign_labels.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "0"),
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

func ccwpContainerClusterSchema(ccName string) string {
	return fmt.Sprintf(`
resource "illumio-core_container_cluster" "ccwp_test" {
	name = %[1]q
	description = "Terraform Container Cluster test"
}`, ccName)
}

func ccwpRoleLabelSchema(labelName string) string {
	return fmt.Sprintf(`
resource "illumio-core_label" "ccwp_role" {
	key   = "role"
	value = %[1]q
}`, labelName)
}

func testAccCheckIllumioCCWPDataSourceConfig_basic(ccwpName, ccName, labelName string) string {
	schema := ccwpContainerClusterSchema(ccName) + ccwpRoleLabelSchema(labelName)
	return schema + fmt.Sprintf(`
resource "illumio-core_container_cluster_workload_profile" "ccwp_test" {
	container_cluster_href = illumio-core_container_cluster.ccwp_test.href
	name = %[1]q
	description = "Terraform Container Cluster Workload Profile test"
	managed = true
	enforcement_mode = "visibility_only"

	assign_labels {
		href = illumio-core_label.ccwp_role.href
	}

	lifecycle {
		create_before_destroy = true
	}
}

data "illumio-core_container_cluster_workload_profile" "ccwp_test" {
	href = illumio-core_container_cluster_workload_profile.ccwp_test.href
}
`, ccwpName)
}

func testAccCheckIllumioCCWPResource_updateToLabels(ccwpName, ccName, labelName string) string {
	appLabelName := acctest.RandomWithPrefix(prefixCCWP)

	schema := ccwpContainerClusterSchema(ccName) + ccwpRoleLabelSchema(labelName)
	return schema + fmt.Sprintf(`
resource "illumio-core_label" "ccwp_app" {
	key   = "app"
	value = %[1]q
}

resource "illumio-core_container_cluster_workload_profile" "ccwp_test" {
	container_cluster_href = illumio-core_container_cluster.ccwp_test.href
	name = %[2]q
	description = "Terraform Container Cluster Workload Profile test"
	managed = true
	enforcement_mode = "visibility_only"

	labels {
		key = "role"

		assignment {
			href = illumio-core_label.ccwp_role.href
		}
	}

	labels {
		key = "app"

		assignment {
			href = illumio-core_label.ccwp_app.href
		}
	}

	lifecycle {
		create_before_destroy = true
	}
}
`, appLabelName, ccwpName)
}

func testAccCheckIllumioCCWPResource_updateNameAndDesc(updatedName, ccName, labelName string) string {
	schema := ccwpContainerClusterSchema(ccName) + ccwpRoleLabelSchema(labelName)
	return schema + fmt.Sprintf(`
resource "illumio-core_container_cluster_workload_profile" "ccwp_test" {
	container_cluster_href = illumio-core_container_cluster.ccwp_test.href
	name = %[1]q
	description = ""
	managed = true
	enforcement_mode = "visibility_only"

	labels {
		key = "role"

		assignment {
			href = illumio-core_label.ccwp_role.href
		}
	}

	lifecycle {
		create_before_destroy = true
	}
}
`, updatedName)
}

func testAccCheckIllumioCCWPResource_updateManageStateAndEnforcement(updatedName, ccName string) string {
	return ccwpContainerClusterSchema(ccName) + fmt.Sprintf(`
resource "illumio-core_container_cluster_workload_profile" "ccwp_test" {
	container_cluster_href = illumio-core_container_cluster.ccwp_test.href
	name = %[1]q
	description = ""
	managed = false
	enforcement_mode = "idle"
}
`, updatedName)
}
