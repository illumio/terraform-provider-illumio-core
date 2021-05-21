// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSCCSBL *schema.Provider

func TestAccIllumioCCSBL_Read(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSCCSBL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCSBLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceCCSBLExists("data.illumio-core_container_cluster_service_backends.test"),
				),
			},
		},
	})
}

func testAccCheckIllumioCCSBLDataSourceConfig_basic() string {
	return `
	data "illumio-core_container_cluster_service_backends" "test" {
		href = "/orgs/1/container_clusters/f959d2d0-fe56-4bd9-8132-b7a31d1cbdde/service_backends"
	  }
	`
}

func testAccCheckIllumioDataSourceCCSBLExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Container Cluster Service Backends %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		// listAttr["length"] = rs.Primary.Attributes["items.#"]

		return nil
	}
}
