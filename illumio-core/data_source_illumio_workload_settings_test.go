// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIllumioWS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_workload_settings.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "workload_disconnected_timeout_seconds.0.value"),
					resource.TestCheckResourceAttrSet(dataSourceName, "workload_goodbye_timeout_seconds.0.value"),
				),
			},
		},
	})
}

func testAccCheckIllumioWSDataSourceConfig_basic() string {
	return `
data "illumio-core_workload_settings" "test" {}
`
}
