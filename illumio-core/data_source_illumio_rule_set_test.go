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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioRSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_tables_rules", resourceName, "ip_tables_rules"),
					resource.TestCheckResourceAttrPair(dataSourceName, "scopes", resourceName, "scopes"),
				),
			},
		},
	})
}

func testAccCheckIllumioRSDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixRS)
	rName2 := acctest.RandomWithPrefix(prefixRS)

	return fmt.Sprintf(`
resource "illumio-core_label" "rs_test" {
	key   = "env"
	value = %[1]q
}

resource "illumio-core_rule_set" "rs_test" {
	name = %[2]q
	description = "Terraform Rule Set test"

	ip_tables_rules {
		actors {
			actors = "ams"
		}

		actors {
			label {
				href = illumio-core_label.rs_test.href
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
			href = illumio-core_label.rs_test.href
		}
	}
}

data "illumio-core_rule_set" "rs_test" {
	href = illumio-core_rule_set.rs_test.href
}
`, rName1, rName2)
}
