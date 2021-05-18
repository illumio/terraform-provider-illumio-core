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

var providerEnforcementBoundary *schema.Provider

func TestAccIllumioEnforcementBoundary_CreateUpdate(t *testing.T) {
	ebAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerEnforcementBoundary),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerEnforcementBoundary, "illumio-core_enforcement_boundary", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioEnforcementBoundaryConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioEnforcementBoundaryExists("illumio-core_enforcement_boundary.test", ebAttr),
					testAccCheckIllumioEnforcementBoundaryAttributes("creation from terraform", ebAttr),
				),
			},
			{
				Config: testAccCheckIllumioEnforcementBoundaryConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioEnforcementBoundaryExists("illumio-core_enforcement_boundary.test", ebAttr),
					testAccCheckIllumioEnforcementBoundaryAttributes("updation from terraform", ebAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioEnforcementBoundaryConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_enforcement_boundary" "test" {
		name = "acc. test Enforcement Boundary %s"
		ingress_service {
		  href = "/orgs/1/sec_policy/draft/services/3"
		}
		consumer {
		  ip_list {
			href = "/orgs/1/sec_policy/draft/ip_lists/1"
		  }
		}
		illumio_provider {
		  label {
			href = "/orgs/1/labels/1"
		  }
		}
	}
	`, val)
}

func testAccCheckIllumioEnforcementBoundaryExists(name string, ebAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Enforcement Boundary %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerEnforcementBoundary).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"ingress_services.0.href",
			"consumers.0.ip_list.href",
			"providers.0.label.href",
		} {
			ebAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioEnforcementBoundaryAttributes(val string, ebAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                     fmt.Sprintf("acc. test Enforcement Boundary %s", val),
			"ingress_services.0.href":  "/orgs/1/sec_policy/draft/services/3",
			"consumers.0.ip_list.href": "/orgs/1/sec_policy/draft/ip_lists/1",
			"providers.0.label.href":   "/orgs/1/labels/1",
		}

		for k, v := range expectation {
			if ebAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ebAttr[k], v)
			}
		}

		return nil
	}
}
