// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixEB string = "TF-ACC-EB"

func TestAccIllumioEB_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_enforcement_boundary.eb_test"
	resourceName := "illumio-core_enforcement_boundary.eb_test"

	svcName := acctest.RandomWithPrefix(prefixEB)
	ipListName := acctest.RandomWithPrefix(prefixEB)
	labelName := acctest.RandomWithPrefix(prefixEB)

	updatedName := acctest.RandomWithPrefix(prefixEB)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioEBDataSourceConfig_basic(svcName, ipListName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
					resource.TestCheckResourceAttrPair(dataSourceName, "providers", resourceName, "providers"),
					resource.TestCheckResourceAttrPair(dataSourceName, "consumers", resourceName, "consumers"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ingress_services", resourceName, "ingress_services"),
				),
			},
			{
				Config: testAccCheckIllumioEBResource_updateName(updatedName, svcName, ipListName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
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

func enforcementBoundaryReqs(svcName, ipListName, labelName string) string {
	return fmt.Sprintf(`
resource "illumio-core_service" "eb_test" {
	name          = %[1]q
	description   = "Terraform Enforcement Boundary test"
	service_ports {
		proto = 6
		port = 1
		to_port = 1023
	}
}

resource "illumio-core_ip_list" "eb_test" {
	name        = %[2]q
	description = "Terraform Enforcement Boundary test"
	ip_ranges {
		from_ip = "10.1.0.0"
		to_ip = "10.10.0.0"
		description = "Terraform Enforcement Boundary test"
		exclusion = false
	}
	fqdns {
		fqdn = "app.example.com"
	}
}

resource "illumio-core_label" "eb_test" {
	key   = "role"
	value = %[3]q
}
`, svcName, ipListName, labelName)
}

func testAccCheckIllumioEBDataSourceConfig_basic(svcName, ipListName, labelName string) string {
	rName := acctest.RandomWithPrefix(prefixEB)

	return enforcementBoundaryReqs(svcName, ipListName, labelName) + fmt.Sprintf(`
resource "illumio-core_enforcement_boundary" "eb_test" {
	name = %[1]q
	ingress_services {
		href = illumio-core_service.eb_test.href
	}
	consumers {
		ip_list {
			href = illumio-core_ip_list.eb_test.href
		}
	}
	providers {
		label {
			href = illumio-core_label.eb_test.href
		}
	}
}

data "illumio-core_enforcement_boundary" "eb_test" {
	href = illumio-core_enforcement_boundary.eb_test.href
}
`, rName)
}

func testAccCheckIllumioEBResource_updateName(updatedName, svcName, ipListName, labelName string) string {
	return enforcementBoundaryReqs(svcName, ipListName, labelName) + fmt.Sprintf(`
resource "illumio-core_enforcement_boundary" "eb_test" {
	name    = %[1]q
	enabled = false

	ingress_services {
		href = illumio-core_service.eb_test.href
	}
	consumers {
		ip_list {
			href = illumio-core_ip_list.eb_test.href
		}
	}
	providers {
		label {
			href = illumio-core_label.eb_test.href
		}
	}
}
`, updatedName)
}
