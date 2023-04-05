// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixRS string = "TF-ACC-RS"

func TestAccIllumioRS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_rule_set.rs_test"
	resourceName := "illumio-core_rule_set.rs_test"

	rsName := acctest.RandomWithPrefix(prefixRS)
	labelName := acctest.RandomWithPrefix(prefixRS)

	updatedName := acctest.RandomWithPrefix(prefixRS)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioRSDataSourceConfig_basic(rsName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_tables_rules", resourceName, "ip_tables_rules"),
					resource.TestCheckResourceAttrPair(dataSourceName, "scopes", resourceName, "scopes"),
				),
			},
			{
				Config: testAccCheckIllumioRSResource_updateRemoveIPTablesRule(rsName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ip_tables_rules.#", "0"),
				),
			},
			{
				Config: testAccCheckIllumioRSResource_updateRemoveScopes(rsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scopes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scopes.0.#", "0"),
				),
			},
			{
				Config: testAccCheckIllumioRSResource_updateNameAndDesc(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// ignore the field count (%) and rules param in the import
				// state as rules use the security_rule resource rather
				// than being added directly to the rule set
				ImportStateVerifyIgnore: []string{"%", "rules"},
			},
		},
	})
}

func TestAccIllumioRS_MT4L(t *testing.T) {
	resourceName := "illumio-core_rule_set.rs_mt4l_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioRSResource_moreThanFourLabels(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "scopes.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "scopes.0.label.#", "4"),
				),
			},
		},
	})
}

func rsLabel(labelKey string, labelName string) string {
	return fmt.Sprintf(`
resource "illumio-core_label" "rs_%[1]s" {
	key   = %[1]q
	value = %[2]q
}
`, labelKey, labelName)
}

func testAccCheckIllumioRSDataSourceConfig_basic(rsName, labelName string) string {
	return rsLabel("env", labelName) + fmt.Sprintf(`
resource "illumio-core_rule_set" "rs_test" {
	name = %[1]q
	description = "Terraform Rule Set test"

	ip_tables_rules {
		actors {
			actors = "ams"
		}

		actors {
			label {
				href = illumio-core_label.rs_env.href
			}
		}

		enabled = false

		ip_version = 6
		statements {
			table_name = "nat"
			chain_name = "PREROUTING"
			parameters = "value"
		}
	}

	scopes {
		label {
			href = illumio-core_label.rs_env.href
		}
	}

	lifecycle {
		create_before_destroy = true
	}
}

data "illumio-core_rule_set" "rs_test" {
	href = illumio-core_rule_set.rs_test.href
}
`, rsName)
}

func testAccCheckIllumioRSResource_updateRemoveIPTablesRule(rsName, labelName string) string {
	return rsLabel("env", labelName) + fmt.Sprintf(`
resource "illumio-core_rule_set" "rs_test" {
	name = %[1]q
	description = "Terraform Rule Set test"

	scopes {
		label {
			href = illumio-core_label.rs_env.href
		}
	}

	lifecycle {
		create_before_destroy = true
	}
}
`, rsName)
}

func testAccCheckIllumioRSResource_updateRemoveScopes(rsName string) string {
	return fmt.Sprintf(`
resource "illumio-core_rule_set" "rs_test" {
	name = %[1]q
	description = "Terraform Rule Set test"

	scopes {}
}
`, rsName)
}

func testAccCheckIllumioRSResource_updateNameAndDesc(updatedName string) string {
	return fmt.Sprintf(`
resource "illumio-core_rule_set" "rs_test" {
	name = %[1]q
	description = ""

	scopes {}
}
`, updatedName)
}

func testAccCheckIllumioRSResource_moreThanFourLabels() string {
	rsName := acctest.RandomWithPrefix(prefixRS)
	lName1 := acctest.RandomWithPrefix(prefixRS)
	lName2 := acctest.RandomWithPrefix(prefixRS)
	lName3 := acctest.RandomWithPrefix(prefixRS)
	lName4 := acctest.RandomWithPrefix(prefixRS)
	labelTypeKey := acctest.RandomWithPrefix(prefixRS)

	return rsLabel("app", lName1) +
		rsLabel("env", lName2) + rsLabel("loc", lName3) + fmt.Sprintf(`
resource "illumio-core_label_type" "rs_label_type_test" {
	key          = %[1]q
	display_name = %[1]q

	display_info {
		initial = "TS"
	}
}

resource "illumio-core_label" "rs_custom_lt" {
	key   = illumio-core_label_type.rs_label_type_test.key
	value = %[2]q
}

resource "illumio-core_rule_set" "rs_mt4l_test" {
	name = %[3]q
	description = "Terraform Rule Set test"

	ip_tables_rules {
		actors {
			actors = "ams"
		}

		enabled = false

		ip_version = 6
		statements {
			table_name = "nat"
			chain_name = "PREROUTING"
			parameters = "value"
		}
	}

	scopes {
		label {
			href = illumio-core_label.rs_app.href
		}

		label {
			href = illumio-core_label.rs_env.href
		}

		label {
			href = illumio-core_label.rs_loc.href
		}

		label {
			href = illumio-core_label.rs_custom_lt.href
		}
	}
}
`, labelTypeKey, lName4, rsName)
}
