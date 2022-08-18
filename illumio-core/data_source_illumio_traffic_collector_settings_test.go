// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIllumioTCS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_traffic_collector_settings.tcs_test"
	resourceName := "illumio-core_traffic_collector_settings.tcs_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioTCSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "transmission", resourceName, "transmission"),
					resource.TestCheckResourceAttrPair(dataSourceName, "action", resourceName, "action"),
					resource.TestCheckResourceAttrPair(dataSourceName, "target", resourceName, "target"),
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

func testAccCheckIllumioTCSDataSourceConfig_basic() string {
	return `
resource "illumio-core_traffic_collector_settings" "tcs_test" {
	transmission = "broadcast"
	action       = "drop"

	target {
		dst_ip   = "127.0.0.1"
		dst_port = 65535
		proto    = 6
	}
}

data "illumio-core_traffic_collector_settings" "tcs_test" {
	href = illumio-core_traffic_collector_settings.tcs_test.href
}
`
}
