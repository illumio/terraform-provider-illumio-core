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

var providerVService *schema.Provider

func TestAccIllumioVirtualService_CreateUpdate(t *testing.T) {
	vsAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerVService),
		CheckDestroy:      testAccCheckIllumioVirtualServiceDestroy(providerVService),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVirtualServiceConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioVirtualServiceExists("illumio-core_virtual_service.test", vsAttr),
					testAccCheckIllumioVirtualServiceAttributes("creation from terraform", vsAttr),
				),
			},
			{
				Config: testAccCheckIllumioVirtualServiceConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioVirtualServiceExists("illumio-core_virtual_service.test", vsAttr),
					testAccCheckIllumioVirtualServiceAttributes("updation from terraform", vsAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioVirtualServiceConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_virtual_service" "test" {
		name = "Acc. Test Virtual Service"
		description = "%s"
		apply_to = "host_only"

		service_ports {
		  proto = 6
		}
		service_addresses {
		  fqdn = "*.illumio.com"
		}
		ip_overrides = [ "1.2.3.4" ]
	  }
	`, val)
}

func testAccCheckIllumioVirtualServiceDestroy(providerInstance *schema.Provider) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		pConfig := (*providerInstance).Meta().(Config)
		for _, rs := range s.RootModule().Resources {
			if rs.Type == "illumio-core_virtual_service" {
				// Provisioning Deleteion of Virtual Service resource
				e := pConfig.ProvisionAResource(pConfig.OrgID, "virtual_services", rs.Primary.ID)
				if e != nil {
					return fmt.Errorf("illumio-core_virtual_service resource deletion not successfully provisioned")
				}

				_, _, err := pConfig.IllumioClient.Get(rs.Primary.ID, nil)

				if err == nil { // got successful response
					return fmt.Errorf("illumio-core_virtual_service still exists, ResourceID: %v", rs.Primary.ID)
				}
				if !strings.Contains(err.Error(), "not-found") {
					return fmt.Errorf("illumio-core_virtual_service still exists, ResourceID: %v", rs.Primary.ID)
				}

			}
		}

		return nil
	}
}

func testAccCheckIllumioVirtualServiceExists(name string, vsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Virtual Service %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerVService).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"apply_to",
			"service_addresses.0.fqdn",
			"ip_overrides.0",
		} {
			vsAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		vsAttr["service_ports.0.proto"] = int(cont.S("service_ports", "0", "proto").Data().(float64))

		return nil
	}
}

func testAccCheckIllumioVirtualServiceAttributes(val string, vsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		expectation := map[string]interface{}{
			"name":                     "Acc. Test Virtual Service",
			"description":              val,
			"apply_to":                 "host_only",
			"service_ports.0.proto":    6,
			"service_addresses.0.fqdn": "*.illumio.com",
			"ip_overrides.0":           "1.2.3.4",
		}
		for k, v := range expectation {
			if vsAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, vsAttr[k], v)
			}
		}

		return nil
	}
}
