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

var providerRS *schema.Provider

func TestAccIllumioRuleSet_CreateUpdate(t *testing.T) {
	srAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerRS),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerRS, "illumio-core_rule_set", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioRuleSetConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioRuleSetExists("illumio-core_rule_set.test", srAttr),
					testAccCheckIllumioRuleSetAttributes("creation from terraform", srAttr),
				),
			},
			{
				Config: testAccCheckIllumioRuleSetConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioRuleSetExists("illumio-core_rule_set.test", srAttr),
					testAccCheckIllumioRuleSetAttributes("updation from terraform", srAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioRuleSetConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_rule_set" "test" {
		name = "terraform-test-1"
		description = "%s"
		
		ip_tables_rules {
		  actors {
			actors = "ams"
		  }

		  actors {
			label {
			  href = "/orgs/1/labels/69"
			}
		  }
	  
		  enabled = false
	  
		  ip_version = 6
		  statements {
			table_name = "nat"
			chain_name = "PREROUTING"
			parameters = "value"
		  }
		}
	  
		scopes {
		  label {
			href = "/orgs/1/labels/1"
		  }
		  label {
			href = "/orgs/1/labels/7"
		  }
		}
	  
		rules {
		  enabled = false
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
	  
		  providers {
			label {
			  href = "/orgs/1/labels/294"
			}
		  }
	  
		  ingress_services {
			proto = 6
			port  = 4
		  }
		}
	  }
	`, val)
}

func testAccCheckIllumioRuleSetExists(name string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Rule set %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerRS).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"ip_tables_rules.0.enabled",
			"ip_tables_rules.0.statements.0.table_name",
			"rules.0.enabled",
		} {
			lgAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioRuleSetAttributes(val string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                      "terraform-test-1",
			"description":               val,
			"ip_tables_rules.0.enabled": false,
			"ip_tables_rules.0.statements.0.table_name": "nat",
			"rules.0.enabled": false,
		}
		for k, v := range expectation {
			if lgAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, lgAttr[k], v)
			}
		}

		return nil
	}
}
