// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerLabel *schema.Provider

func TestAccIllumioLabel_CreateUpdate(t *testing.T) {
	labelAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerLabel),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerLabel, "illumio-core_label", true),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioLabelExists("illumio-core_label.test", labelAttr),
					testAccCheckIllumioLabelAttributes("creation from terraform", labelAttr),
				),
			},
			{
				Config: testAccCheckIllumioLabelConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioLabelExists("illumio-core_label.test", labelAttr),
					testAccCheckIllumioLabelAttributes("updation from terraform", labelAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioLabelConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_label" "test" {
		key = "role"
		value = "%s"
		external_data_set = "illumio-core_label_external_data_set_1"
		external_data_reference = "illumio-core_label_external_data_reference_1"
	}
	`, val)
}

func testAccCheckIllumioLabelExists(name string, labelAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Label %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerLabel).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"key",
			"value",
			"external_data_set",
			"external_data_reference",
		} {
			labelAttr[k] = cont.S(k).Data()
		}

		return nil
	}
}

func testAccCheckIllumioLabelAttributes(val string, labelAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"key":                     "role",
			"value":                   val,
			"external_data_set":       "illumio-core_label_external_data_set_1",
			"external_data_reference": "illumio-core_label_external_data_reference_1",
		}
		for k, v := range expectation {
			if labelAttr[k] != v {
				return fmt.Errorf("Bad %s %v", k, v)
			}
		}

		return nil
	}
}
