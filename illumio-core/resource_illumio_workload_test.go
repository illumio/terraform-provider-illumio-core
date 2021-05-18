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

var providerW *schema.Provider

func TestAccIllumioWorkload_CreateUpdate(t *testing.T) {
	lgAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerW),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerW, "illumio-core_workload", true),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWorkloadExists("illumio-core_workload.test", lgAttr),
					testAccCheckIllumioWorkloadAttributes("creation from terraform", lgAttr),
				),
			},
			{
				Config: testAccCheckIllumioWorkloadConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWorkloadExists("illumio-core_workload.test", lgAttr),
					testAccCheckIllumioWorkloadAttributes("updation from terraform", lgAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_workload" "test" {
		name                   = "acc. test workload"
		description            = "%s"
		hostname               = "acc. test workload hostname"
		distinguished_name     = "acc. test distinguished name"
		interfaces {
			name       = "acc. test workload interface"
			link_state = "up"
			address    = "10.10.3.10"
		}
		labels {
			href = "/orgs/1/labels/1"
		}
		service_provider = "acc. test service provider"
	}
	`, val)
}

func testAccCheckIllumioWorkloadExists(name string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Workload %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerW).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"hostname",
			"description",
			"labels.0.href",
			"interfaces.0.name",
			"interfaces.0.link_state",
			"interfaces.0.address",
			"service_provider",
			"distinguished_name",
		} {
			lgAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioWorkloadAttributes(val string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                    "acc. test workload",
			"hostname":                "acc. test workload hostname",
			"description":             val,
			"labels.0.href":           "/orgs/1/labels/1",
			"interfaces.0.name":       "acc. test workload interface",
			"interfaces.0.link_state": "up",
			"interfaces.0.address":    "10.10.3.10",
			"service_provider":        "acc. test service provider",
			"distinguished_name":      "acc. test distinguished name",
		}
		for k, v := range expectation {
			if lgAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, lgAttr[k], v)
			}
		}

		return nil
	}
}
