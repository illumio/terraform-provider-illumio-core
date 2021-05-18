// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSTCSL *schema.Provider

func TestAccIllumioTCSL_Read(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSTCSL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioTCSLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceTCSLExists("data.illumio-core_traffic_collector_settings_list.test"),
				),
			},
		},
	})
}

func testAccCheckIllumioTCSLDataSourceConfig_basic() string {
	return `
	data "illumio-core_traffic_collector_settings_list" "test" {}
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
