// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSSyslogDestL *schema.Provider

func TestAccIllumioSyslogDestL_Read(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSSyslogDestL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSyslogDestLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceSyslogDestLExists("data.illumio-core_syslog_destinations.test"),
				),
			},
		},
	})
}

func testAccCheckIllumioSyslogDestLDataSourceConfig_basic() string {
	return `
	data "illumio-core_syslog_destinations" "test" {}
	`
}

func testAccCheckIllumioDataSourceSyslogDestLExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Syslog Destinations %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		return nil
	}
}
