package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerSER *schema.Provider

func TestAccIllumioSelectiveEnforcementRule_CreateUpdate(t *testing.T) {
	serAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerSER),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerSER, "illumio_selective_enforcement_rule", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSelectiveEnforcementRuleConfig_basic("/orgs/1/labels/69"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioSelectiveEnforcementRuleExists("illumio_selective_enforcement_rule.test", serAttr),
					testAccCheckIllumioSelectiveEnforcementRuleAttributes("/orgs/1/labels/69", serAttr),
				),
			},
			{
				Config: testAccCheckIllumioSelectiveEnforcementRuleConfig_basic("/orgs/1/labels/715"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioSelectiveEnforcementRuleExists("illumio_selective_enforcement_rule.test", serAttr),
					testAccCheckIllumioSelectiveEnforcementRuleAttributes("/orgs/1/labels/715", serAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioSelectiveEnforcementRuleConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio_selective_enforcement_rule" "test" {
		name = "SER test 1"
		scope {
		  label {
			href = "%s"
		  }
		}
	  
		enforced_services {
		  href = "/orgs/1/sec_policy/draft/services/3"
		}
	  }
	`, val)
}

func testAccCheckIllumioSelectiveEnforcementRuleExists(name string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Selective Enforcement Rule %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerSER).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"scope.0.label.href",
			"enforced_services.0.href",
		} {
			serAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioSelectiveEnforcementRuleAttributes(val string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                     "SER test 1",
			"scope.0.label.href":       val,
			"enforced_services.0.href": "/orgs/1/sec_policy/draft/services/3",
		}
		for k, v := range expectation {
			if serAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, serAttr[k], v)
			}
		}

		return nil
	}
}
