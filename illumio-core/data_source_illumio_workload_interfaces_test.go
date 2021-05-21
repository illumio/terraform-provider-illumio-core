// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSWorkloadInterfacesL *schema.Provider

func TestAccIllumioWorkloadInterfacesL_Read(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSWorkloadInterfacesL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadInterfacesLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceWorkloadInterfacesLExists("data.illumio-core_workload_interfaces.test"),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadInterfacesLDataSourceConfig_basic() string {
	return `
	data "illumio-core_workload_interfaces" "test" {
		workload_href = "/orgs/1/workloads/d42a430e-b20b-4b2d-853f-2d39fa4cea22"
	}
	`
}

func testAccCheckIllumioDataSourceWorkloadInterfacesLExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Workload Interfaces %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		// listAttr["length"] = rs.Primary.Attributes["items.#"]

		return nil
	}
}
