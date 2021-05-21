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

var providerDSOS *schema.Provider

func TestAccIllumioOS_Read(t *testing.T) {
	osAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSOS),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioOSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceOSExists("data.illumio-core_organization_settings.test", osAttr),
					testAccCheckIllumioOSDataSourceAttributes(osAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioOSDataSourceConfig_basic() string {
	return `
	data "illumio-core_organization_settings" "test" {
	}
	`
}

func testAccCheckIllumioDataSourceOSExists(name string, osAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Organization Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSOS).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"audit_event_retention_seconds",
			"audit_event_min_severity",
			"format",
		} {
			osAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioOSDataSourceAttributes(osAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"audit_event_retention_seconds": float64(7776000),
			"audit_event_min_severity":      "informational",
			"format":                        "JSON",
		}
		for k, v := range expectation {
			if osAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, osAttr[k], v)
			}
		}

		return nil
	}
}
