// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIllumioOS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_organization_settings.test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
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
		},
	})
}

func testAccCheckIllumioOSDataSourceConfig_basic() string {
	return `
data "illumio-core_organization_settings" "test" {}
`
}
