// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSCCWPL *schema.Provider

func TestAccIllumioCCWPL_Read(t *testing.T) {
	listAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSCCWPL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCWPLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceCCWPLExists("data.illumio-core_container_cluster_workload_profiles.test", listAttr),
					testAccCheckIllumioListDataSourceSize(listAttr, "5"),
				),
			},
		},
	})
}

func testAccCheckIllumioCCWPLDataSourceConfig_basic() string {
	return `
	data "illumio-core_container_cluster_workload_profiles" "test" {
		max_results = "5"
  		container_cluster_href = "/orgs/1/container_clusters/f959d2d0-fe56-4bd9-8132-b7a31d1cbdde"
	}
	`
}

func testAccCheckIllumioDataSourceCCWPLExists(name string, listAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Container Clusters Workload Profiles %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		listAttr["length"] = rs.Primary.Attributes["items.#"]

		return nil
	}
}
