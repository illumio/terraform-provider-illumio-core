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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioIPDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_ranges", resourceName, "ip_ranges"),
					resource.TestCheckResourceAttrPair(dataSourceName, "fqdns", resourceName, "fqdns"),
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

func testAccCheckIllumioIPDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixIP)

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
`, rName1)
}
