package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSLG *schema.Provider

func TestAccIllumioLG_Read(t *testing.T) {
	lgAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSLG),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLGDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceLGExists("data.illumio-core_label_group.test", lgAttr),
					testAccCheckIllumioLGDataSourceAttributes(lgAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioLGDataSourceConfig_basic() string {
	return `
	data "illumio-core_label_group" "test" {
		href = "/orgs/1/sec_policy/draft/label_groups/b347346a-7aff-4334-ac90-6f64a7a98f05"
	}
	`
}

func testAccCheckIllumioDataSourceLGExists(name string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Label Group %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSLG).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"key",
			"labels.0.href",
			"labels.0.key",
			"labels.0.value",
		} {
			lgAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioLGDataSourceAttributes(lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":           "Acc. test name",
			"description":    "Acc. test description",
			"key":            "role",
			"labels.0.href":  "/orgs/1/labels/2",
			"labels.0.key":   "role",
			"labels.0.value": "Database",
		}
		for k, v := range expectation {
			if lgAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, lgAttr[k], v)
			}
		}

		return nil
	}
}
