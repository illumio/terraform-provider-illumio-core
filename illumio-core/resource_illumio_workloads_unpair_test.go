// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerWUnpair *schema.Provider

func TestAccIllumioWorkloadsUnpair_CreateUpdate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerWUnpair),
		// CheckDestroy is ignored as illumio-core_workloads_unpair does not support delete operation
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadsUnpairConfig_basic("63bf19d1-1efa-49ec-b712-c51d5c0aa552"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWorkloadsUnpairExists("illumio-illumio-core_workloads_unpair.test"),
				),
			},
			{
				Config: testAccCheckIllumioWorkloadsUnpairConfig_basic("e683b686-8afe-4675-88a1-4463395f0482"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWorkloadsUnpairExists("illumio-illumio-core_workloads_unpair.test"),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadsUnpairConfig_basic(id string) string {
	return fmt.Sprintf(`
	resource "illumio-core_workloads_unpair" "test" {
		workloads {
		  href = "/orgs/1/workloads/%s"
		}
	  }
	`, id)
}

func testAccCheckIllumioWorkloadsUnpairExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Workloads Unpair %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		return nil
	}
}
