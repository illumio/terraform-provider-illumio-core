// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIllumioWS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_workload_settings.ws_test"
	resourceName := "illumio-core_workload_settings.ws_test"

	discTimeoutSecs := "3600"
	goodbyeTimeoutSecs := "900"

	resource.ParallelTest(t, resource.TestCase{
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
			{
				Config: testAccCheckIllumioWSResource_updateSettings(discTimeoutSecs, goodbyeTimeoutSecs),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workload_disconnected_timeout_seconds.0.value", discTimeoutSecs),
					resource.TestCheckResourceAttr(resourceName, "workload_goodbye_timeout_seconds.0.value", goodbyeTimeoutSecs),
				),
			},
			{
				SkipFunc: skipIfPCEVersionBelow("23.1.0"),
				Config:   testAccCheckIllumioWSResource_updateWithVENTypes(discTimeoutSecs, goodbyeTimeoutSecs),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "workload_disconnected_timeout_seconds.0.value", discTimeoutSecs),
					resource.TestCheckResourceAttr(resourceName, "workload_goodbye_timeout_seconds.0.value", goodbyeTimeoutSecs),
				),
			},
		},
	})
}

func testAccCheckIllumioWSDataSourceConfig_basic() string {
	return `
resource "illumio-core_workload_settings" "ws_test" {}

data "illumio-core_workload_settings" "ws_test" {}
`
}

func testAccCheckIllumioWSResource_updateSettings(discTimeoutSecs, goodbyeTimeoutSecs string) string {
	return fmt.Sprintf(`
resource "illumio-core_workload_settings" "ws_test" {
	workload_disconnected_timeout_seconds {
		value = %[1]s
	}

	workload_goodbye_timeout_seconds {
		value = %[2]s
	}
}
`, discTimeoutSecs, goodbyeTimeoutSecs)
}

func testAccCheckIllumioWSResource_updateWithVENTypes(discTimeoutSecs, goodbyeTimeoutSecs string) string {
	return fmt.Sprintf(`
resource "illumio-core_workload_settings" "ws_test" {
	workload_disconnected_timeout_seconds {
		ven_type = "server"
		value    = %[1]s
	}

	workload_disconnected_timeout_seconds {
		ven_type = "endpoint"
		value    = %[1]s
	}

	workload_goodbye_timeout_seconds {
		ven_type = "server"
		value    = %[2]s
	}

	workload_goodbye_timeout_seconds {
		ven_type = "endpoint"
		value    = %[2]s
	}
}
`, discTimeoutSecs, goodbyeTimeoutSecs)
}
