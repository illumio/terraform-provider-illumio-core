// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLL string = "TF-ACC-LL"

func TestAccIllumioLL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_labels.ll_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioLLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixLL)
	rName2 := acctest.RandomWithPrefix(prefixLL)

	return fmt.Sprintf(`
resource "illumio-core_label" "ll_test1" {
	key   = "app"
	value = %[1]q
}

resource "illumio-core_label" "ll_test2" {
	key   = "app"
	value = %[2]q
}

data "illumio-core_labels" "ll_test" {
	# lookup based on partial match
	value = %[3]q

	# enforce dependencies
	depends_on = [
		illumio-core_label.ll_test1,
		illumio-core_label.ll_test2,
	]
}
`, rName1, rName2, prefixLL)
}
