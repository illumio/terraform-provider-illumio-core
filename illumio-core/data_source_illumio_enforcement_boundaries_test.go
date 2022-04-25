// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixEBL string = "TF-ACC-EBL"

func TestAccIllumioEBL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_enforcement_boundaries.ebl_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioEBLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioEBLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixEBL)
	rName2 := acctest.RandomWithPrefix(prefixEBL)
	rName3 := acctest.RandomWithPrefix(prefixEBL)
	rName4 := acctest.RandomWithPrefix(prefixEBL)
	rName5 := acctest.RandomWithPrefix(prefixEBL)

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

data "illumio-core_enforcement_boundaries" "ebl_test" {
	# lookup based on partial match
	name = %[6]q

	# enforce dependencies
	depends_on = [
		illumio-core_enforcement_boundary.ebl_test1,
		illumio-core_enforcement_boundary.ebl_test2,
	]
}
`, rName1, rName2, rName3, rName4, rName5, prefixEBL)
}
