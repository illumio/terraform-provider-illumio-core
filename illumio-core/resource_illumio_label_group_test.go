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

var providerLG *schema.Provider

func TestAccIllumioLabelGroup_CreateUpdate(t *testing.T) {
	lgAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerLG),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerLG, "illumio-core_label_group", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelGroupConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioLabelGroupExists("illumio-core_label_group.test", lgAttr),
					testAccCheckIllumioLabelGroupAttributes("creation from terraform", lgAttr),
				),
			},
			{
				Config: testAccCheckIllumioLabelGroupConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioLabelGroupExists("illumio-core_label_group.test", lgAttr),
					testAccCheckIllumioLabelGroupAttributes("updation from terraform", lgAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioLabelGroupConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_label_group" "test" {
		key           = "role"
		name          = "acc. test label group"
		description   = "%s"
		external_data_set = "illumio-core_label_group_external_data_set_1"
		external_data_reference = "illumio-core_label_group_external_data_reference_1"
		labels {
			href = "/orgs/1/labels/7147"
		}
		sub_groups {
			href = "/orgs/1/sec_policy/draft/label_groups/4dcfedff-6236-454b-acb7-92827c9a7a2f"
		}
	}
	`, val)
}

func testAccCheckIllumioLabelGroupExists(name string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Label Group %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerLG).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"key",
			"name",
			"description",
			"labels.0.href",
			"sub_groups.0.href",
			"external_data_set",
			"external_data_reference",
		} {
			lgAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioLabelGroupAttributes(val string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"key":                     "role",
			"description":             val,
			"labels.0.href":           "/orgs/1/labels/7147",
			"sub_groups.0.href":       "/orgs/1/sec_policy/draft/label_groups/4dcfedff-6236-454b-acb7-92827c9a7a2f",
			"external_data_set":       "illumio-core_label_group_external_data_set_1",
			"external_data_reference": "illumio-core_label_group_external_data_reference_1",
		}
		for k, v := range expectation {
			if lgAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, lgAttr[k], v)
			}
		}

		return nil
	}
}
