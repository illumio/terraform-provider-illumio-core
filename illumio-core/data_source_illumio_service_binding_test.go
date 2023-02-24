// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var prefixSB string = "TF-ACC-SB"

func TestAccIllumioSB_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_service_binding.sb_test"
	resourceName := "illumio-core_service_binding.sb_test"

	virtualServiceName := acctest.RandomWithPrefix(prefixSB)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		CheckDestroy:      testAccCheckIllumioSBDataSource_VSDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSBDataSourceConfig_VSSetup(virtualServiceName),
				Check:  testAccCheckIllumioSBDataSource_VSExists("illumio-core_virtual_service.test"),
			},
			{
				Config: testAccCheckIllumioSBDataSourceConfig_basic(virtualServiceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "virtual_service", resourceName, "virtual_service"),
					resource.TestCheckResourceAttrPair(dataSourceName, "workload", resourceName, "workload"),
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

func testAccCheckIllumioSBDataSourceConfig_VSSetup(name string) string {
	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "test" {
	name = %[1]q
	description = "Terraform Service Binding test"
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
`, name)
}

func testAccCheckIllumioSBDataSourceConfig_basic(virtualServiceName string) string {
	rName := acctest.RandomWithPrefix(prefixSB)

	var sb strings.Builder

	sb.WriteString(testAccCheckIllumioSBDataSourceConfig_VSSetup(virtualServiceName))
	sb.WriteString(fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" "sb_test" {
	name               = %[1]q
	description        = "Terraform Service Binding test"
	hostname           = "example.workload"
}

resource "illumio-core_service_binding" "sb_test" {
	virtual_service {
		href = replace(illumio-core_virtual_service.test.href, "draft", "active")
	}
	workload {
		href = illumio-core_unmanaged_workload.sb_test.href
	}
}

data "illumio-core_service_binding" "sb_test" {
	href = illumio-core_service_binding.sb_test.href
}
`, rName))
	return sb.String()
}

func testAccCheckIllumioSBDataSource_VSExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		href := rs.Primary.ID

		if href == "" {
			return fmt.Errorf("Virtual Service HREF is not set")
		}

		conf := TestAccProvider.Meta().(Config)

		// try to provision the virtual service
		err := conf.ProvisionAResource("virtual_services", href)

		if err != nil {
			return fmt.Errorf("Provisioning failed for Virtual Service")
		}

		return nil
	}
}

func testAccCheckIllumioSBDataSource_VSDestroyed(s *terraform.State) error {
	conf := TestAccProvider.Meta().(Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "illumio-core_virtual_service" {
			continue
		}

		// try to provision the virtual service deletion
		err := conf.ProvisionAResource("virtual_services", rs.Primary.ID)

		if err != nil {
			return fmt.Errorf("Teardown failed for Virtual Service")
		}
	}

	return nil
}
