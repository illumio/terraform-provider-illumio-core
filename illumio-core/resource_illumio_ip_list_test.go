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

var providerIPList *schema.Provider

func TestAccIllumioIPList_CreateUpdate(t *testing.T) {
	ipAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerIPList),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerIPList, "illumio-core_ip_list", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioIPListConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioIPListExists("illumio-core_ip_list.test", ipAttr),
					testAccCheckIllumioIPListAttributes("creation from terraform", ipAttr),
				),
			},
			{
				Config: testAccCheckIllumioIPListConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioIPListExists("illumio-core_ip_list.test", ipAttr),
					testAccCheckIllumioIPListAttributes("updation from terraform", ipAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioIPListConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_ip_list" "test" {
		name        = "acc. test iplist"
		description = "%s"
		external_data_set = "illumio-core_ip_list_external_data_set_1"
		external_data_reference = "illumio-core_ip_list_external_data_reference_1"
		ip_ranges {
		  from_ip = "1.1.0.0"
		  to_ip = "1.10.0.0"
		  description = "test ip_ranges description"
		  exclusion = false
		}
		fqdns {
		  fqdn = "app.example.com"
		}
	  }
	`, val)
}

func testAccCheckIllumioIPListExists(name string, ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("IP List %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerIPList).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"external_data_set",
			"external_data_reference",
			"ip_ranges.0.from_ip",
			"ip_ranges.0.to_ip",
			"ip_ranges.0.description",
			"ip_ranges.0.exclusion",
			"fqdns.0.fqdn",
		} {
			ipAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioIPListAttributes(val string, ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                    "acc. test iplist",
			"description":             val,
			"external_data_set":       "illumio-core_ip_list_external_data_set_1",
			"external_data_reference": "illumio-core_ip_list_external_data_reference_1",
			"ip_ranges.0.from_ip":     "1.1.0.0",
			"ip_ranges.0.to_ip":       "1.10.0.0",
			"ip_ranges.0.description": "test ip_ranges description",
			"ip_ranges.0.exclusion":   false,
			"fqdns.0.fqdn":            "app.example.com",
		}
		for k, v := range expectation {
			if ipAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ipAttr[k], v)
			}
		}

		return nil
	}
}
