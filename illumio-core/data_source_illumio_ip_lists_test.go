// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixIPL string = "TF-ACC-IPL"

func TestAccIllumioIPL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_ip_lists.ipl_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioIPLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIllumioIPLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixIPL)
	rName2 := acctest.RandomWithPrefix(prefixIPL)
	rName3 := acctest.RandomWithPrefix(prefixIPL)

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

data "illumio-core_ip_lists" "ipl_test" {
	# lookup based on partial match
	name = %[4]q

	# enforce dependencies
	depends_on = [
		illumio-core_ip_list.ipl_test1,
		illumio-core_ip_list.ipl_test2,
		illumio-core_ip_list.ipl_test3,
	]
}
`, rName1, rName2, rName3, prefixIPL)
}
