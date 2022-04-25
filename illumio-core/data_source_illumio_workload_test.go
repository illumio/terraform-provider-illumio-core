// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixWorkload string = "TF-ACC-WL"

func TestAccIllumioWorkload_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_workload.wl_test"
	resourceName := "illumio-core_unmanaged_workload.wl_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "hostname", resourceName, "hostname"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_principal_name", resourceName, "service_principal_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enforcement_mode", resourceName, "enforcement_mode"),
					resource.TestCheckResourceAttrPair(dataSourceName, "visibility_level", resourceName, "visibility_level"),
					resource.TestCheckResourceAttrPair(dataSourceName, "distinguished_name", resourceName, "distinguished_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "public_ip", resourceName, "public_ip"),
					resource.TestCheckResourceAttrPair(dataSourceName, "data_center", resourceName, "data_center"),
					resource.TestCheckResourceAttrPair(dataSourceName, "data_center_zone", resourceName, "data_center_zone"),
					resource.TestCheckResourceAttrPair(dataSourceName, "os_id", resourceName, "os_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "os_detail", resourceName, "os_detail"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_provider", resourceName, "service_provider"),
					resource.TestCheckResourceAttrPair(dataSourceName, "online", resourceName, "online"),
					resource.TestCheckResourceAttrPair(dataSourceName, "labels", resourceName, "labels"),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixWorkload)
	rName2 := acctest.RandomWithPrefix(prefixWorkload)

	return fmt.Sprintf(`
resource "illumio-core_label" "wl_test" {
	key   = "role"
	value = %[1]q
}

resource "illumio-core_unmanaged_workload" "wl_test" {
	name               = %[2]q
	description        = "Terraform Workload test"
	hostname           = "example.workload"
	distinguished_name = ""
	service_provider   = "SPN"

	labels {
		href = illumio-core_label.wl_test.href
	}
}

data "illumio-core_workload" "wl_test" {
	href = illumio-core_unmanaged_workload.wl_test.href
}
`, rName1, rName2)
}
