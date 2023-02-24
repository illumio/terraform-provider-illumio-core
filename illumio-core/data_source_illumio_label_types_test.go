// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLTL string = "TF-ACC-LTL"

func TestAccIllumioLTL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_label_types.ltl_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLTLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioLTLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixLTL)
	rName2 := acctest.RandomWithPrefix(prefixLTL)

	return fmt.Sprintf(`
resource "illumio-core_label_type" "ltl_test1" {
	key          = %[1]q
	display_name = %[1]q
}

resource "illumio-core_label_type" "ltl_test2" {
	key          = %[2]q
	display_name = %[2]q
}

data "illumio-core_label_types" "ltl_test" {
	# lookup based on partial match
	display_name = %[3]q

	# enforce dependencies
	depends_on = [
		illumio-core_label_type.ltl_test1,
		illumio-core_label_type.ltl_test2,
	]
}
`, rName1, rName2, prefixLTL)
}
