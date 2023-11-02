// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixRSL string = "TF-ACC-RSL"

func init() {
	resource.AddTestSweepers("rule_sets", &resource.Sweeper{
		Name: "rule_sets",
		F: func(region string) error {
			conf := TestAccProvider.Meta().(Config)
			illumioClient := conf.IllumioClient

			endpoint := fmt.Sprintf("/orgs/%d/sec_policy/draft/rule_sets", illumioClient.OrgID)
			_, data, err := illumioClient.Get(endpoint, &map[string]string{
				"name": prefixWLL,
			})

			if err != nil {
				return fmt.Errorf("Error getting rule sets: %s", err)
			}

			for _, ruleSet := range data.Children() {
				href := ruleSet.S("href").Data().(string)
				_, err := illumioClient.Delete(href)
				if err != nil {
					fmt.Printf("Failed to sweep rule set with HREF: %s", href)
				}
			}

			return nil
		},
	})
}

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
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", ruleSetName),
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
