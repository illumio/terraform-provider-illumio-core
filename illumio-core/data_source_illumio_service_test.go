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

var providerDSS *schema.Provider

func TestAccIllumioS_Read(t *testing.T) {
	sAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSS),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceSExists("data.illumio-core_service.test", sAttr),
					testAccCheckIllumioSDataSourceAttributes(sAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioSDataSourceConfig_basic() string {
	return `
	data "illumio-core_service" "test" {
		href = "/orgs/1/sec_policy/draft/services/17"
	}
	`
}

func testAccCheckIllumioDataSourceSExists(name string, sAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Service %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSS).Meta().(Config)
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
			"service_ports.0.port",
			"service_ports.0.to_port",
			"service_ports.0.proto",
		} {
			sAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioSDataSourceAttributes(sAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                    "Acc. test name",
			"description":             "Acc. test description",
			"external_data_set":       "illumio-core_service_external_data_set_1",
			"external_data_reference": "illumio-core_service_external_data_reference_1",
			"service_ports.0.port":    float64(10),
			"service_ports.0.to_port": float64(100),
			"service_ports.0.proto":   float64(6),
		}
		for k, v := range expectation {
			if sAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, sAttr[k], v)
			}
		}

		return nil
	}
}
