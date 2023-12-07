// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixIP string = "TF-ACC-IP"

func TestAccIllumioIP_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_ip_list.ip_test"
	resourceName := "illumio-core_ip_list.ip_test"
	ipListName := acctest.RandomWithPrefix(prefixIP)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioIPDataSourceConfig_basic(ipListName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_ranges", resourceName, "ip_ranges"),
					resource.TestCheckResourceAttrPair(dataSourceName, "fqdns", resourceName, "fqdns"),
				),
			},
			{
				Config: testAccCheckIllumioIPResource_updateIPRange(ipListName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ip_ranges.0.from_ip", "10.1.0.0/16"),
					resource.TestCheckResourceAttr(resourceName, "ip_ranges.0.to_ip", ""),
					resource.TestCheckResourceAttr(resourceName, "ip_ranges.0.description", ""),
				),
			},
			{
				Config: testAccCheckIllumioIPResource_IPRangeConsolidation(ipListName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ip_ranges.#", "5"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ip_ranges.*", map[string]string{"from_ip": "10.1.0.0"}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ip_ranges.*", map[string]string{"from_ip": "2001:4860:4860::8844"}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ip_ranges.*", map[string]string{"from_ip": "10.0.0.0/14", "description": "", "exclusion": "false"}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ip_ranges.*", map[string]string{"from_ip": "10.0.0.0/14", "description": "test desc", "exclusion": "false"}),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "ip_ranges.*", map[string]string{"from_ip": "10.0.0.0/14", "description": "", "exclusion": "true"}),
				),
			},
			{
				Config: testAccCheckIllumioIPResource_removeIPRanges(ipListName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ip_ranges.#", "0"),
				),
			},
			{
				Config: testAccCheckIllumioIPResource_emptyDescription(ipListName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIllumioIPDataSourceConfig_basic(ipListName string) string {
	return fmt.Sprintf(`
resource "illumio-core_ip_list" "ip_test" {
	name        = %[1]q
	description = "Terraform IP List test"
	ip_ranges {
		from_ip = "10.1.0.0"
		to_ip = "10.10.0.0"
		description = "Terraform IP List test"
		exclusion = false
	}
	fqdns {
		fqdn = "app.example.com"
	}
}

data "illumio-core_ip_list" "ip_test" {
	href = illumio-core_ip_list.ip_test.href
}
`, ipListName)
}

func testAccCheckIllumioIPResource_updateIPRange(ipListName string) string {
	return fmt.Sprintf(`
resource "illumio-core_ip_list" "ip_test" {
	name        = %[1]q
	description = "Terraform IP List test"
	ip_ranges {
		from_ip = "10.1.0.0/16"
		description = ""
		exclusion = false
	}
	fqdns {
		fqdn = "app.example.com"
	}
}
`, ipListName)
}

func testAccCheckIllumioIPResource_IPRangeConsolidation(ipListName string) string {
	return fmt.Sprintf(`
resource "illumio-core_ip_list" "ip_test" {
	name        = %[1]q
	description = "Terraform IP List test"
	# test CIDR notation strip for individual address ranges
	ip_ranges {
		from_ip = "10.1.0.0/32"
		description = ""
		exclusion = false
	}
	ip_ranges {
		from_ip = "2001:4860:4860::8844/128"
		description = ""
		exclusion = false
	}
	# test consolidation of identical subnets
	# and change to network address
	ip_ranges {
		from_ip = "10.1.0.10/14"
		description = ""
		exclusion = false
	}
	ip_ranges {
		from_ip = "10.0.0.10/14"
		description = ""
		exclusion = false
	}
	ip_ranges {
		from_ip = "10.0.0.10/14"
		description = "test desc"
		exclusion = false
	}
	ip_ranges {
		from_ip = "10.0.0.10/14"
		description = ""
		exclusion = true
	}
	fqdns {
		fqdn = "app.example.com"
	}
}
`, ipListName)
}

func testAccCheckIllumioIPResource_removeIPRanges(ipListName string) string {
	return fmt.Sprintf(`
resource "illumio-core_ip_list" "ip_test" {
	name        = %[1]q
	description = "Terraform IP List test"
	fqdns {
		fqdn = "app.example.com"
	}
}
`, ipListName)
}

func testAccCheckIllumioIPResource_emptyDescription(ipListName string) string {
	return fmt.Sprintf(`
resource "illumio-core_ip_list" "ip_test" {
	name        = %[1]q
	description = ""
	fqdns {
		fqdn = "app.example.com"
	}
}
`, ipListName)
}
