// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixIPL string = "TF-ACC-IPL"

func init() {
	resource.AddTestSweepers("ip_lists", &resource.Sweeper{
		Name: "ip_lists",
		F:    sweep("IP list", "name", prefixIPL, "/orgs/%d/sec_policy/draft/ip_lists"),
		Dependencies: []string{
			"enforcement_boundaries",
			"rule_sets",
			"virtual_services",
		},
	})
}

func TestAccIllumioIPL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_ip_lists.ipl_test"
	ipListName := acctest.RandomWithPrefix(prefixIPL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioIPLDataSourceConfig_basic(ipListName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "3"),
				),
			},
			{
				Config: testAccCheckIllumioIPLDataSourceConfig_exactMatch(ipListName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", ipListName),
				),
			},
		},
	})
}

func iplConfig(ipListName string) string {
	rName1 := acctest.RandomWithPrefix(prefixIPL)
	rName2 := acctest.RandomWithPrefix(prefixIPL)

	return fmt.Sprintf(`
resource "illumio-core_ip_list" "ipl_test1" {
	name        = %[1]q
	description = "Terraform IP Lists test"
	ip_ranges {
		from_ip = "10.1.0.0"
		to_ip = "10.10.0.0"
		description = "Terraform IP Lists test"
		exclusion = false
	}
	fqdns {
		fqdn = "app.example.com"
	}
}

resource "illumio-core_ip_list" "ipl_test2" {
	name        = %[2]q
	description = "Terraform IP Lists test"
	ip_ranges {
		from_ip = "172.168.0.0"
		to_ip = "172.168.0.255"
		description = "Terraform IP Lists test"
		exclusion = false
	}
	fqdns {
		fqdn = "*.illum.io"
	}
}

resource "illumio-core_ip_list" "ipl_test3" {
	name        = %[3]q
	description = "Terraform IP Lists test"
	fqdns {
		fqdn = "*.example.com"
	}
}
`, rName1, rName2, ipListName)
}

func testAccCheckIllumioIPLDataSourceConfig_basic(ipListName string) string {
	return iplConfig(ipListName) + fmt.Sprintf(`
data "illumio-core_ip_lists" "ipl_test" {
	# lookup based on partial match
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_ip_list.ipl_test1,
		illumio-core_ip_list.ipl_test2,
		illumio-core_ip_list.ipl_test3,
	]
}
`, prefixIPL)
}

func testAccCheckIllumioIPLDataSourceConfig_exactMatch(ipListName string) string {
	return iplConfig(ipListName) + fmt.Sprintf(`
data "illumio-core_ip_lists" "ipl_test" {
	# lookup using exact match
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_ip_list.ipl_test1,
		illumio-core_ip_list.ipl_test2,
		illumio-core_ip_list.ipl_test3,
	]
}
`, ipListName)
}
