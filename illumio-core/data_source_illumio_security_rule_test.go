package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSSR *schema.Provider

func TestAccIllumioSR_Read(t *testing.T) {
	ppAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSSR),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSRDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceSRExists("data.illumio-core_security_rule.test", ppAttr),
					testAccCheckIllumioSRDataSourceAttributes(ppAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioSRDataSourceConfig_basic() string {
	return `
	data "illumio-core_security_rule" "test" {
		href = "/orgs/1/sec_policy/draft/rule_sets/54/sec_rules/56"
	}
	`
}

func testAccCheckIllumioDataSourceSRExists(name string, ppAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Security Rule %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSSR).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"description",
			"enabled",
			"providers.0.label.href",
			"consumers.0.actors",
			"network_type",
		} {
			ppAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioSRDataSourceAttributes(ppAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"description":            "Acc. test description",
			"enabled":                true,
			"providers.0.label.href": "/orgs/1/labels/715",
			"consumers.0.actors":     "ams",
			"network_type":           "brn",
		}
		for k, v := range expectation {
			if ppAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ppAttr[k], v)
			}
		}

		return nil
	}
}
