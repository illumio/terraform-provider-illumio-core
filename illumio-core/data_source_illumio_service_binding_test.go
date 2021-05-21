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

var providerDSServiceBinding *schema.Provider

func TestAccIllumioServiceBinding_Read(t *testing.T) {
	serviceBindingAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSServiceBinding),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioServiceBindingDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceServiceBindingExists("data.illumio-core_service_binding.test", serviceBindingAttr),
					testAccCheckIllumioServiceBindingDataSourceAttributes(serviceBindingAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioServiceBindingDataSourceConfig_basic() string {
	return `
	data "illumio-core_service_binding" "test" {
		href = "/orgs/1/service_bindings/5087544a-84b7-47c4-9d82-3ca1732d3242"
	}
	`
}

func testAccCheckIllumioDataSourceServiceBindingExists(name string, serviceBindingAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Service Binding %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSServiceBinding).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"workload.href",
			"virtual_service.href",
		} {
			serviceBindingAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioServiceBindingDataSourceAttributes(serviceBindingAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"workload.href":        "/orgs/1/workloads/c754a713-2bde-4427-af1f-bff145be509b",
			"virtual_service.href": "/orgs/1/sec_policy/active/virtual_services/9fa52991-3c2e-4be1-9e66-189e0bb724e2",
		}
		for k, v := range expectation {
			if serviceBindingAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, serviceBindingAttr[k], v)
			}
		}

		return nil
	}
}
