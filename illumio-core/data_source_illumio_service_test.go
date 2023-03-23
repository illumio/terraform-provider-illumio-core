// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixService string = "TF-ACC-Service"

func TestAccIllumioService_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_service.test"
	resourceName := "illumio-core_service.test"

	svcName := acctest.RandomWithPrefix(prefixService)
	updatedName := acctest.RandomWithPrefix(prefixService)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioServiceDataSourceConfig_basic(svcName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_ports", resourceName, "service_ports"),
				),
			},
			{
				Config: testAccCheckIllumioServiceResource_updatenameAndDesc(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
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

func testAccCheckIllumioServiceDataSourceConfig_basic(svcName string) string {
	return fmt.Sprintf(`
resource "illumio-core_service" "test" {
	name          = %[1]q
	description   = "Terraform Service test"
	service_ports {
		proto = 6
		port = 1
		to_port = 1023
	}
}

data "illumio-core_service" "test" {
	href = illumio-core_service.test.href
}
`, svcName)
}

func testAccCheckIllumioServiceResource_updatenameAndDesc(updatedName string) string {
	return fmt.Sprintf(`
resource "illumio-core_service" "test" {
	name          = %[1]q
	description   = ""
	service_ports {
		proto = 6
		port = 1
		to_port = 1023
	}
}
`, updatedName)
}
