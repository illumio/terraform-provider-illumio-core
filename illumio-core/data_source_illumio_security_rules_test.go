// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSSRL *schema.Provider

func TestAccIllumioSRL_Read(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSSRL),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSRLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceSRLExists("data.illumio-core_security_rules.test"),
				),
			},
		},
	})
}

func testAccCheckIllumioSRLDataSourceConfig_basic() string {
	return `
	data "illumio-core_security_rules" "test" {
		rule_set_href = "/orgs/1/sec_policy/draft/rule_sets/6"
	}
	`
}

func testAccCheckIllumioDataSourceSRLExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Security Rules %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		// listAttr["length"] = rs.Primary.Attributes["items.#"]

		return nil
	}
}
