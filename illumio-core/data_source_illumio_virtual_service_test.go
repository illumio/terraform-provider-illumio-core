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

var providerDSVirtualService *schema.Provider

func TestAccIllumioVirtualService_Read(t *testing.T) {
	virtualserAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSVirtualService),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVirtualServiceDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceVirtualServiceExists("data.illumio-core_virtual_service.test", virtualserAttr),
					testAccCheckIllumioVirtualServiceDataSourceAttributes(virtualserAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioVirtualServiceDataSourceConfig_basic() string {
	return `
	data "illumio-core_virtual_service" "test" {
		href = "/orgs/1/sec_policy/draft/virtual_services/91a20432-bec8-4b2b-9a2d-9185c9dd75e4"
	}
	`
}

func testAccCheckIllumioDataSourceVirtualServiceExists(name string, virtualserAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Virtual Service %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSVirtualService).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"description",
			"name",
			"service_ports.0.proto",
			"labels.0.href",
			"service_addresses.0.ip",
		} {
			virtualserAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioVirtualServiceDataSourceAttributes(virtualserAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"description":            "acc-test description",
			"name":                   "acc-test-VS",
			"service_ports.0.proto":  float64(6),
			"labels.0.href":          "/orgs/1/labels/1",
			"service_addresses.0.ip": "1.1.1.1",
		}
		for k, v := range expectation {
			if virtualserAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, virtualserAttr[k], v)
			}
		}

		return nil
	}
}
