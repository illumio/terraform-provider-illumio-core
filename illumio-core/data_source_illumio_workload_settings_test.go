// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
					testAccCheckIllumioWSResource_valueSet(resourceName, "server", "workload_disconnected_timeout_seconds", discTimeoutSecs),
					testAccCheckIllumioWSResource_valueSet(resourceName, "server", "workload_goodbye_timeout_seconds", goodbyeTimeoutSecs),
				),
			},
			{
				SkipFunc: skipIfPCEVersionBelow("23.1.0"),
				Config:   testAccCheckIllumioWSResource_updateWithVENTypes(discTimeoutSecs, goodbyeTimeoutSecs),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWSResource_valueSet(resourceName, "server", "workload_disconnected_timeout_seconds", discTimeoutSecs),
					testAccCheckIllumioWSResource_valueSet(resourceName, "server", "workload_goodbye_timeout_seconds", goodbyeTimeoutSecs),
					testAccCheckIllumioWSResource_valueSet(resourceName, "endpoint", "workload_disconnected_timeout_seconds", discTimeoutSecs),
					testAccCheckIllumioWSResource_valueSet(resourceName, "endpoint", "workload_goodbye_timeout_seconds", goodbyeTimeoutSecs),
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

func testAccCheckIllumioWSResource_valueSet(resourceName, venType, setting, expectedValue string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		settingCount, err := strconv.Atoi(rs.Primary.Attributes[setting+".#"])
		if err != nil {
			return err
		}

		for i := 0; i < settingCount; i++ {
			key := fmt.Sprintf("%s.%d", setting, i)
			if rs.Primary.Attributes[key+".ven_type"] == venType {
				value := rs.Primary.Attributes[key+".value"]
				if value != expectedValue {
					return fmt.Errorf(`Attribute '%s' expected "%s" got "%s"`, key+".value", expectedValue, value)
				}

				return nil
			}
		}

		return fmt.Errorf(`Couldn't find value for setting '%s' with ven_type '%s'`, setting, venType)
	}
}
