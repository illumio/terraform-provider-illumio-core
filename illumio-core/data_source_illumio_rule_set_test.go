// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSRS *schema.Provider

func TestAccIllumioRS_Read(t *testing.T) {
	ppAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSRS),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioRSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceRSExists("data.illumio-core_rule_set.test", ppAttr),
					testAccCheckIllumioRSDataSourceAttributes(ppAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioRSDataSourceConfig_basic() string {
	return `
	data "illumio-core_rule_set" "test" {
		href = "/orgs/1/sec_policy/draft/rule_sets/54"
	}
	`
}

func testAccCheckIllumioDataSourceRSExists(name string, ppAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Ruleset %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSRS).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"enabled",
			"scopes.0.0.label.href",
			"rules.0.enabled",
		} {
			ppAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioRSDataSourceAttributes(ppAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                  "Acc. test name",
			"description":           "Acc. test description",
			"enabled":               true,
			"scopes.0.0.label.href": "/orgs/1/labels/69",
			"rules.0.enabled":       true,
		}
		for k, v := range expectation {
			if ppAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ppAttr[k], v)
			}
		}

		return nil
	}
}
