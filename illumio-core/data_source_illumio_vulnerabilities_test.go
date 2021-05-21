// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSVulnerabilitiesL *schema.Provider

func TestAccIllumioVulnerabilitiesL_Read(t *testing.T) {
	listAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSVulnerabilitiesL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVulnerabilitiesLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceVulnerabilitiesLExists("data.illumio-core_vulnerabilities.test", listAttr),
					testAccCheckIllumioListDataSourceSize(listAttr, "5"),
				),
			},
		},
	})
}

func testAccCheckIllumioVulnerabilitiesLDataSourceConfig_basic() string {
	return `
	data "illumio-core_vulnerabilities" "test" {
		max_results = 5
	}
	`
}

func testAccCheckIllumioDataSourceVulnerabilitiesLExists(name string, listAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Vulnerabilities %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		listAttr["length"] = rs.Primary.Attributes["items.#"]

		return nil
	}
}
