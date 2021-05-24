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

var providerSer *schema.Provider

func TestAccIllumioService_CreateUpdate(t *testing.T) {
	serAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerSer),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerSer, "illumio-core_service", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioServiceConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioServiceExists("illumio-core_service.test", serAttr),
					testAccCheckIllumioServiceAttributes("creation from terraform", serAttr),
				),
			},
			{
				Config: testAccCheckIllumioServiceConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioServiceExists("illumio-core_service.test", serAttr),
					testAccCheckIllumioServiceAttributes("updation from terraform", serAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioServiceConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_service" "test" {
		name          = "acc test service"
		description   = "%s"
		service_ports {
			proto = 6
			port = 10
			to_port = 100
		}
	}
	`, val)
}

func testAccCheckIllumioServiceExists(name string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Service %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerSer).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"service_ports.0.proto",
			"service_ports.0.port",
			"service_ports.0.to_port",
		} {
			serAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioServiceAttributes(val string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                    "acc test service",
			"description":             val,
			"service_ports.0.proto":   float64(6),
			"service_ports.0.port":    float64(10),
			"service_ports.0.to_port": float64(100),
		}

		for k, v := range expectation {
			if serAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, serAttr[k], v)
			}
		}

		return nil
	}
}
