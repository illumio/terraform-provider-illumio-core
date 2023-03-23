// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIllumioOS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_organization_settings.os_test"
	resourceName := "illumio-core_organization_settings.os_test"

	updatedFormat := "JSON"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t); testAccPreCheckSaaSPCE(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioOSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceName, "audit_event_retention_seconds"),
					resource.TestCheckResourceAttrSet(dataSourceName, "audit_event_min_severity"),
					resource.TestCheckResourceAttrSet(dataSourceName, "format"),
				),
			},
			{
				Config: testAccCheckIllumioOSResource_updateFormat(updatedFormat),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "format", updatedFormat),
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

func testAccCheckIllumioOSDataSourceConfig_basic() string {
	return `
resource "illumio-core_organization_settings" "os_test" {}

data "illumio-core_organization_settings" "os_test" {}
`
}

func testAccCheckIllumioOSResource_updateFormat(updatedFormat string) string {
	return fmt.Sprintf(`
resource "illumio-core_organization_settings" "os_test" {
	format = %[1]q
}
`, updatedFormat)
}
