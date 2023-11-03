// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLL string = "TF-ACC-LL"

func init() {
	resource.AddTestSweepers("labels", &resource.Sweeper{
		Name: "labels",
		F:    sweep("label", "value", prefixLL, "/orgs/%d/labels"),
		Dependencies: []string{
			"container_clusters",
			"enforcement_boundaries",
			"label_groups",
			"pairing_profiles",
			"rule_sets",
			"virtual_services",
			"workloads",
		},
	})
}

func TestAccIllumioLL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_labels.ll_test"
	labelValue := acctest.RandomWithPrefix(prefixLL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLLDataSourceConfig_basic(labelValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioLLDataSourceConfig_exactMatch(labelValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.value", labelValue),
				),
			},
		},
	})
}

func llConfig(labelValue string) string {
	rVal := acctest.RandomWithPrefix(prefixLL)

	return fmt.Sprintf(`
resource "illumio-core_label" "ll_test1" {
	key   = "app"
	value = %[1]q
}

resource "illumio-core_label" "ll_test2" {
	key   = "app"
	value = %[2]q
}
`, rVal, labelValue)
}

func testAccCheckIllumioLLDataSourceConfig_basic(labelValue string) string {
	return llConfig(labelValue) + fmt.Sprintf(`
data "illumio-core_labels" "ll_test" {
	# lookup based on partial match
	value = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_label.ll_test1,
		illumio-core_label.ll_test2,
	]
}
`, prefixLL)
}

func testAccCheckIllumioLLDataSourceConfig_exactMatch(labelValue string) string {
	return llConfig(labelValue) + fmt.Sprintf(`
data "illumio-core_labels" "ll_test" {
	# lookup using exact match
	value      = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_label.ll_test1,
		illumio-core_label.ll_test2,
	]
}
`, labelValue)
}
