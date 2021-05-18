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

var providerDSIP *schema.Provider

func TestAccIllumioIP_Read(t *testing.T) {
	ipAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSIP),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioIPDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceIPExists("data.illumio-core_ip_list.test", ipAttr),
					testAccCheckIllumioIPDataSourceAttributes(ipAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioIPDataSourceConfig_basic() string {
	return `
	data "illumio-core_ip_list" "test" {
		href = "/orgs/1/sec_policy/draft/ip_lists/27"
	}
	`
}

func testAccCheckIllumioDataSourceIPExists(name string, ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("IP List %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSIP).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"fqdns.0.fqdn",
			"ip_ranges.0.from_ip",
		} {
			ipAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioIPDataSourceAttributes(ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                "Acc. test name",
			"description":         "Acc. test description",
			"fqdns.0.fqdn":        "app.example.com",
			"ip_ranges.0.from_ip": "1.1.0.0/24",
		}
		for k, v := range expectation {
			if ipAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ipAttr[k], v)
			}
		}

		return nil
	}
}
