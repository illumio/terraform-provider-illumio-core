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

var providerSR *schema.Provider

func TestAccIllumioSecurityRule_CreateUpdate(t *testing.T) {
	srAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerSR),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerSR, "illumio-core_security_rule", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSecurityRuleConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioSecurityRuleExists("illumio-core_security_rule.test", srAttr),
					testAccCheckIllumioSecurityRuleAttributes("creation from terraform", srAttr),
				),
			},
			{
				Config: testAccCheckIllumioSecurityRuleConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioSecurityRuleExists("illumio-core_security_rule.test", srAttr),
					testAccCheckIllumioSecurityRuleAttributes("updation from terraform", srAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioSecurityRuleConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_security_rule" "test" {
		rule_set_href = "/orgs/1/sec_policy/draft/rule_sets/6"
		enabled = true
		description = "%s"

		resolve_labels_as {
		  consumers = ["workloads"]
		  providers = ["workloads"]
		}
	  
		consumers {
		  actors = "ams"
		}
	  
		providers {
		  label {
			href = "/orgs/1/labels/715"
		  }
		}
	  
		ingress_services {
		  proto = 6
		  port  = 12
		}
	  }
	`, val)
}

func testAccCheckIllumioSecurityRuleExists(name string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Security Rule %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerSR).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"enabled",
			"description",
			"resolve_labels_as.providers.0",
			"resolve_labels_as.consumers.0",
			"consumers.0.actors",
			"providers.0.label.href",
		} {
			lgAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		for _, k := range []string{"ingress_services.0.proto", "ingress_services.0.port"} {
			lgAttr[k] = int(cont.S(strings.Split(k, ".")...).Data().(float64))
		}

		return nil
	}
}

func testAccCheckIllumioSecurityRuleAttributes(val string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"enabled":                       true,
			"description":                   val,
			"resolve_labels_as.providers.0": "workloads",
			"resolve_labels_as.consumers.0": "workloads",
			"consumers.0.actors":            "ams",
			"providers.0.label.href":        "/orgs/1/labels/715",
			"ingress_services.0.proto":      6,
			"ingress_services.0.port":       12,
		}
		for k, v := range expectation {
			if lgAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, lgAttr[k], v)
			}
		}

		return nil
	}
}
