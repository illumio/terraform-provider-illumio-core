// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSLabel *schema.Provider

func TestAccIllumioLabel_Read(t *testing.T) {
	labelAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSLabel),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceLabelExists("data.illumio-core_label.test", labelAttr),
					testAccCheckIllumioLabelDataSourceAttributes(labelAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioLabelDataSourceConfig_basic() string {
	return `
	data "illumio-core_label" "test" {
		href = "/orgs/1/labels/828"
	}
	`
}

func testAccCheckIllumioDataSourceLabelExists(name string, labelAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Label %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSLabel).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"key",
			"value",
		} {
			labelAttr[k] = cont.S(k).Data()
		}

		return nil
	}
}

func testAccCheckIllumioLabelDataSourceAttributes(labelAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"key":   "role",
			"value": "Acc. test label value",
		}
		for k, v := range expectation {
			if labelAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, labelAttr[k], v)
			}
		}

		return nil
	}
}
