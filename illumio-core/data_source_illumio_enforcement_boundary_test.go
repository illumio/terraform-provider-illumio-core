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

var providerDSEB *schema.Provider

func TestAccIllumioEB_Read(t *testing.T) {
	ebAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSEB),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioEBDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceEBExists("data.illumio-core_enforcement_boundary.test", ebAttr),
					testAccCheckIllumioEBDataSourceAttributes(ebAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioEBDataSourceConfig_basic() string {
	return `
	data "illumio-core_enforcement_boundary" "test" {
		href = "/orgs/1/sec_policy/draft/enforcement_boundaries/37"
	}
	`
}

func testAccCheckIllumioDataSourceEBExists(name string, ebAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Enforcement Boundary %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSEB).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"providers.0.label.href",
			"consumers.0.ip_list.href",
			"ingress_services.0.href",
		} {
			ebAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioEBDataSourceAttributes(ebAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                     "Acc. test name",
			"providers.0.label.href":   "/orgs/1/labels/1",
			"consumers.0.ip_list.href": "/orgs/1/sec_policy/draft/ip_lists/1",
			"ingress_services.0.href":  "/orgs/1/sec_policy/draft/services/3",
		}
		for k, v := range expectation {
			if ebAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ebAttr[k], v)
			}
		}

		return nil
	}
}
