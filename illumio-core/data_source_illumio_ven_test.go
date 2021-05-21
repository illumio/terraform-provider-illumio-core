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

var providerDSVEN *schema.Provider

func TestAccIllumioVEN_Read(t *testing.T) {
	venAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSVEN),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVENDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceVENExists("data.illumio-core_ven.test", venAttr),
					testAccCheckIllumioVENDataSourceAttributes(venAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioVENDataSourceConfig_basic() string {
	return `
	data "illumio-core_ven" "test" {
		href = "/orgs/1/vens/95291fef-d71e-448e-9b7c-a70aa15ab402"
	}
	`
}

func testAccCheckIllumioDataSourceVENExists(name string, venAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VEN %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSVEN).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"hostname",
			"description",
			"name",
			"interfaces.0.name",
			"conditions.0.first_reported_timestamp",
		} {
			venAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioVENDataSourceAttributes(venAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"hostname":                              "perf-workload-47045",
			"description":                           "acc-test description",
			"name":                                  "acc-test-VEN",
			"interfaces.0.name":                     "eth0",
			"conditions.0.first_reported_timestamp": "2020-10-22T00:10:50.556Z",
		}
		for k, v := range expectation {
			if venAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, venAttr[k], v)
			}
		}

		return nil
	}
}
