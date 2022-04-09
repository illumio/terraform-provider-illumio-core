// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixRSL string = "TF-ACC-RSL"

func TestAccIllumioRSL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_rule_sets.rsl_test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioRSLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioRSLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixRSL)
	rName2 := acctest.RandomWithPrefix(prefixRSL)
	rName3 := acctest.RandomWithPrefix(prefixRSL)

	return fmt.Sprintf(`
resource "illumio-core_label" "rsl_test" {
	key   = "env"
	value = %[1]q
}

resource "illumio-core_rule_set" "rsl_test1" {
	name = %[2]q
	description = "Terraform Rule Sets test"

	scopes {
		label {
			href = illumio-core_label.rsl_test.href
		}
	}
}

resource "illumio-core_rule_set" "rsl_test2" {
	name = %[3]q
	description = "Terraform Rule Sets test"

	scopes {
		label {
			href = illumio-core_label.rsl_test.href
		}
	}
}

data "illumio-core_rule_sets" "rsl_test" {
	# lookup based on partial match
	name = %[4]q

	# enforce dependencies
	depends_on = [
		illumio-core_rule_set.rsl_test1,
		illumio-core_rule_set.rsl_test2,
	]
}
`, rName1, rName2, rName3, prefixRSL)
}
