// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLTL string = "TF-ACC-LTL"

func init() {
	resource.AddTestSweepers("label_types", &resource.Sweeper{
		Name:         "label_types",
		F:            sweep("label type", "display_name", prefixLTL, "/orgs/%d/label_dimensions"),
		Dependencies: []string{"labels"},
	})
}

func TestAccIllumioLTL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_label_types.ltl_test"
	labelTypeName := acctest.RandomWithPrefix(prefixLTL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLTLDataSourceConfig_basic(labelTypeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioLTLDataSourceConfig_exactMatch(labelTypeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.display_name", labelTypeName),
				),
			},
		},
	})
}

func ltlConfig(labelTypeName string) string {
	rName := acctest.RandomWithPrefix(prefixLTL)

	return fmt.Sprintf(`
resource "illumio-core_label_type" "ltl_test1" {
	key          = %[1]q
	display_name = %[1]q
}

resource "illumio-core_label_type" "ltl_test2" {
	key          = %[2]q
	display_name = %[2]q
}
`, rName, labelTypeName)
}

func testAccCheckIllumioLTLDataSourceConfig_basic(labelTypeName string) string {
	return ltlConfig(labelTypeName) + fmt.Sprintf(`
data "illumio-core_label_types" "ltl_test" {
	# lookup based on partial match
	display_name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_label_type.ltl_test1,
		illumio-core_label_type.ltl_test2,
	]
}
`, prefixLTL)
}

func testAccCheckIllumioLTLDataSourceConfig_exactMatch(labelTypeName string) string {
	return ltlConfig(labelTypeName) + fmt.Sprintf(`
data "illumio-core_label_types" "ltl_test" {
	# lookup using exact match
	display_name = %[1]q
	match_type   = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_label_type.ltl_test1,
		illumio-core_label_type.ltl_test2,
	]
}
`, labelTypeName)
}
