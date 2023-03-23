// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixVS string = "TF-ACC-VS"

func TestAccIllumioVS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_virtual_service.vs_test"
	resourceName := "illumio-core_virtual_service.vs_test"

	vsName := acctest.RandomWithPrefix(prefixVS)
	serviceName := acctest.RandomWithPrefix(prefixVS)
	updatedName := acctest.RandomWithPrefix(prefixVS)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVSDataSourceConfig_basic(vsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "apply_to", resourceName, "apply_to"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_ports", resourceName, "service_ports"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_addresses", resourceName, "service_addresses"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ip_overrides", resourceName, "ip_overrides"),
				),
			},
			{
				Config: testAccCheckIllumioVSResource_updateRemoveAddresses(vsName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ip_overrides.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "service_addresses.#", "0"),
				),
			},
			{
				Config: testAccCheckIllumioVSResource_updateToService(vsName, serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "apply_to", "internal_bridge_network"),
					resource.TestCheckResourceAttr(resourceName, "service_ports.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "service.#", "1"),
				),
			},
			{
				Config: testAccCheckIllumioVSResource_updateNameAndDesc(updatedName, serviceName),
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

func vsServiceConfig(serviceName string) string {
	return fmt.Sprintf(`
resource "illumio-core_service" "vs_https" {
	name = %[1]q

	service_ports {
		proto = 6
		port = 443
	}
}
`, serviceName)
}

func testAccCheckIllumioVSDataSourceConfig_basic(vsName string) string {
	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "vs_test" {
	name = %[1]q
	description = "Terraform Virtual Service test"
	apply_to = "host_only"

	service_ports {
		proto = 6
		port = 1234
	}

	service_addresses {
		fqdn = "*.illumio.com"
	}

	ip_overrides = [
		"1.2.3.4"
	]
}

data "illumio-core_virtual_service" "vs_test" {
	href = illumio-core_virtual_service.vs_test.href
}
`, vsName)
}

func testAccCheckIllumioVSResource_updateRemoveAddresses(vsName string) string {
	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "vs_test" {
	name = %[1]q
	description = "Terraform Virtual Service test"
	apply_to = "host_only"

	service_ports {
		proto = 6
		port = 1234
	}
}
`, vsName)
}

func testAccCheckIllumioVSResource_updateToService(vsName, serviceName string) string {
	return vsServiceConfig(serviceName) + fmt.Sprintf(`
resource "illumio-core_virtual_service" "vs_test" {
	name = %[1]q
	description = "Terraform Virtual Service test"
	apply_to = "internal_bridge_network"

	service {
		href = illumio-core_service.vs_https.href
	}
}
`, vsName)
}

func testAccCheckIllumioVSResource_updateNameAndDesc(updatedName, serviceName string) string {
	return vsServiceConfig(serviceName) + fmt.Sprintf(`
resource "illumio-core_virtual_service" "vs_test" {
	name = %[1]q
	description = ""
	apply_to = "internal_bridge_network"

	service {
		href = illumio-core_service.vs_https.href
	}
}
`, updatedName)
}
