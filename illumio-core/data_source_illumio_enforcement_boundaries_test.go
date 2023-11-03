// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixEBL string = "TF-ACC-EBL"

func init() {
	resource.AddTestSweepers("enforcement_boundaries", &resource.Sweeper{
		Name: "enforcement_boundaries",
		F:    sweep("enforcement boundary", "name", prefixEBL, "/orgs/%d/sec_policy/draft/enforcement_boundaries"),
	})
}

func TestAccIllumioEBL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_enforcement_boundaries.ebl_test"
	boundaryName := acctest.RandomWithPrefix(prefixEBL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioEBLDataSourceConfig_basic(boundaryName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioEBLDataSourceConfig_exactMatch(boundaryName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", boundaryName),
				),
			},
		},
	})
}

func eblConfig(boundaryName string) string {
	serviceName := acctest.RandomWithPrefix(prefixSL)
	ipListName := acctest.RandomWithPrefix(prefixIPL)
	labelName := acctest.RandomWithPrefix(prefixLL)
	rName := acctest.RandomWithPrefix(prefixEBL)

	return fmt.Sprintf(`
resource "illumio-core_service" "ebl_test" {
	name          = %[1]q
	description   = "Terraform Enforcement Boundaries test"
	service_ports {
		proto = 6
		port = 80
	}

	service_ports {
		proto = 6
		port = 443
	}
}

resource "illumio-core_ip_list" "ebl_test" {
	name        = %[2]q
	description = "Terraform Enforcement Boundaries test"
	ip_ranges {
		from_ip = "10.0.0.0"
		to_ip = "10.255.255.255"
		description = "Terraform Enforcement Boundaries test"
		exclusion = false
	}
	fqdns {
		fqdn = "*.example.com"
	}
}

resource "illumio-core_label" "ebl_test" {
	key   = "role"
	value = %[3]q
}

resource "illumio-core_enforcement_boundary" "ebl_test1" {
	name = %[4]q
	ingress_services {
		href = illumio-core_service.ebl_test.href
	}
	consumers {
		ip_list {
			href = illumio-core_ip_list.ebl_test.href
		}
	}
	providers {
		label {
			href = illumio-core_label.ebl_test.href
		}
	}
}

resource "illumio-core_enforcement_boundary" "ebl_test2" {
	name = %[5]q
	ingress_services {
		href = illumio-core_service.ebl_test.href
	}
	consumers {
		ip_list {
			href = illumio-core_ip_list.ebl_test.href
		}
	}
	providers {
		label {
			href = illumio-core_label.ebl_test.href
		}
	}
}
`, serviceName, ipListName, labelName, rName, boundaryName)
}

func testAccCheckIllumioEBLDataSourceConfig_basic(boundaryName string) string {
	return eblConfig(boundaryName) + fmt.Sprintf(`
data "illumio-core_enforcement_boundaries" "ebl_test" {
	# lookup based on partial match
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_enforcement_boundary.ebl_test1,
		illumio-core_enforcement_boundary.ebl_test2,
	]
}
`, prefixEBL)
}

func testAccCheckIllumioEBLDataSourceConfig_exactMatch(boundaryName string) string {
	return eblConfig(boundaryName) + fmt.Sprintf(`
data "illumio-core_enforcement_boundaries" "ebl_test" {
	# lookup using exact match
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_enforcement_boundary.ebl_test1,
		illumio-core_enforcement_boundary.ebl_test2,
	]
}
`, boundaryName)
}
