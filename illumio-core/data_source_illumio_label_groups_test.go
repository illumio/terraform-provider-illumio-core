// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLGL string = "TF-ACC-LGL"

func TestAccIllumioLGL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_label_groups.lgl_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLGLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioLGLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixLGL)
	rName2 := acctest.RandomWithPrefix(prefixLGL)

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

data "illumio-core_label_groups" "lgl_test" {
	# lookup based on partial match
	name = %[3]q

	# enforce dependencies
	depends_on = [
		illumio-core_label_group.lgl_test1,
		illumio-core_label_group.lgl_test2,
	]
}
`, rName1, rName2, prefixLGL)
}
