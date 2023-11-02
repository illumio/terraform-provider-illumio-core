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
	ruleSetName := acctest.RandomWithPrefix(prefixRSL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioRSLDataSourceConfig_basic(ruleSetName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioRSLDataSourceConfig_exactMatch(ruleSetName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
				),
			},
		},
	})
}

func rslConfig(ruleSetName string) string {
	rName1 := acctest.RandomWithPrefix(prefixRSL)
	rName2 := acctest.RandomWithPrefix(prefixRSL)

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
`, rName1, rName2, ruleSetName)
}

func testAccCheckIllumioRSLDataSourceConfig_basic(ruleSetName string) string {
	return rslConfig(ruleSetName) + fmt.Sprintf(`
data "illumio-core_rule_sets" "rsl_test" {
	# lookup based on partial match
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_rule_set.rsl_test1,
		illumio-core_rule_set.rsl_test2,
	]
}
`, prefixRSL)
}

func testAccCheckIllumioRSLDataSourceConfig_exactMatch(ruleSetName string) string {
	return rslConfig(ruleSetName) + fmt.Sprintf(`
data "illumio-core_rule_sets" "rsl_test" {
	# lookup using exact match
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_rule_set.rsl_test1,
		illumio-core_rule_set.rsl_test2,
	]
}
`, ruleSetName)
}
