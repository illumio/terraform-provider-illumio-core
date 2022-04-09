// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIllumioTCSL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_traffic_collector_settings_list.tcsl_test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioTCSLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceTCSLExists(dataSourceName),
					testAccCheckIllumioDataSourceTCSLNonEmpty(dataSourceName),
				),
			},
		},
	})
}

func testAccCheckIllumioTCSLDataSourceConfig_basic() string {
	return `
resource "illumio-core_traffic_collector_settings" "tcsl_test" {
	transmission = "broadcast"
	action       = "drop"

	target {
		dst_ip   = "192.168.0.1"
		dst_port = 22
		proto    = 6
	}
}

data "illumio-core_traffic_collector_settings_list" "tcsl_test" {
	# enforce dependency
	depends_on = [
		illumio-core_traffic_collector_settings.tcsl_test,
	]
}
`
}

func testAccCheckIllumioDataSourceTCSLExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Traffic Collector Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		return nil
	}
}

func testAccCheckIllumioDataSourceTCSLNonEmpty(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Traffic Collector Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		if v, ok := rs.Primary.Attributes["items.#"]; ok {
			count, _ := strconv.Atoi(v)
			if count <= 0 {
				return fmt.Errorf("Empty list of Traffic Collector Settings")
			}
			return nil
		}

		return fmt.Errorf("Error reading item count in Traffic Collector Settings list data source")
	}
}
