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

var providerServiceBinding *schema.Provider

func TestAccIllumioServiceBinding_CreateUpdate(t *testing.T) {
	ipAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerServiceBinding),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerServiceBinding, "illumio-core_service_binding", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioServiceBindingConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioServiceBindingExists("illumio-core_service_binding.test", ipAttr),
					testAccCheckIllumioServiceBindingAttributes(ipAttr),
				),
			},
			{
				Config: testAccCheckIllumioServiceBindingConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioServiceBindingExists("illumio-core_service_binding.test", ipAttr),
					testAccCheckIllumioServiceBindingAttributes(ipAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioServiceBindingConfig_basic() string {
	return `
	resource "illumio-core_service_binding" "test" {
		virtual_service {
		  href = "/orgs/1/sec_policy/active/virtual_services/91a20432-bec8-4b2b-9a2d-9185c9dd75e4"
		}
		workload {
		  href = "/orgs/1/workloads/d42a430e-b20b-4b2d-853f-2d39fa4cea22"
		}
	  }			
	`
}

func testAccCheckIllumioServiceBindingExists(name string, ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Service Binding %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerServiceBinding).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"workload.href",
			"virtual_service.href",
		} {
			ipAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioServiceBindingAttributes(ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"virtual_service.href": "/orgs/1/sec_policy/active/virtual_services/91a20432-bec8-4b2b-9a2d-9185c9dd75e4",
			"workload.href":        "/orgs/1/workloads/d42a430e-b20b-4b2d-853f-2d39fa4cea22",
		}
		for k, v := range expectation {
			if ipAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ipAttr[k], v)
			}
		}

		return nil
	}
}
