// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLGL string = "TF-ACC-LGL"

func init() {
	resource.AddTestSweepers("label_groups", &resource.Sweeper{
		Name: "label_groups",
		F:    sweep("label group", "name", prefixLGL, "/orgs/%d/sec_policy/draft/label_groups"),
		Dependencies: []string{
			"container_clusters",
			"enforcement_boundaries",
			"pairing_profiles",
			"rule_sets",
			"virtual_services",
			"workloads",
		},
	})
}

func TestAccIllumioLGL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_label_groups.lgl_test"
	labelGroupName := acctest.RandomWithPrefix(prefixLGL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLGLDataSourceConfig_basic(labelGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioLGLDataSourceConfig_exactMatch(labelGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", labelGroupName),
				),
			},
		},
	})
}

func lglConfig(labelGroupName string) string {
	labelName := acctest.RandomWithPrefix(prefixLL)
	rName := acctest.RandomWithPrefix(prefixLGL)

	return fmt.Sprintf(`
resource "illumio-core_label" "lgl_test" {
	key   = "app"
	value = %[1]q
}

resource "illumio-core_label_group" "lgl_test1" {
	key           = "app"
	name          = %[2]q
	description   = "Terraform Label Group subgroup"
}

resource "illumio-core_label_group" "lgl_test2" {
	key           = "app"
	name          = %[3]q
	description   = "Terraform Label Group test"
	labels {
		href = illumio-core_label.lgl_test.href
	}
	sub_groups {
		href = illumio-core_label_group.lgl_test1.href
	}
}
`, labelName, rName, labelGroupName)
}

func testAccCheckIllumioLGLDataSourceConfig_basic(labelGroupName string) string {
	return lglConfig(labelGroupName) + fmt.Sprintf(`
data "illumio-core_label_groups" "lgl_test" {
	# lookup based on partial match
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_label_group.lgl_test1,
		illumio-core_label_group.lgl_test2,
	]
}
`, prefixLGL)
}

func testAccCheckIllumioLGLDataSourceConfig_exactMatch(labelGroupName string) string {
	return lglConfig(labelGroupName) + fmt.Sprintf(`
data "illumio-core_label_groups" "lgl_test" {
	# lookup using exact match
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_label_group.lgl_test1,
		illumio-core_label_group.lgl_test2,
	]
}
`, labelGroupName)
}
