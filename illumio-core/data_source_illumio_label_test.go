// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLabel string = "TF-ACC-Label"

func TestAccIllumioLabel_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_label.label_test"
	resourceName := "illumio-core_label.label_test"
	updatedValue := acctest.RandomWithPrefix(prefixLabel)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "key", resourceName, "key"),
					resource.TestCheckResourceAttrPair(dataSourceName, "value", resourceName, "value"),
				),
			},
			{
				Config: testAccCheckIllumioLabeResource_updateValue(updatedValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "value", updatedValue),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIllumioLabelDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixLabel)

	return fmt.Sprintf(`
resource "illumio-core_label" "label_test" {
	key   = "app"
	value = %[1]q
}

data "illumio-core_label" "label_test" {
	href = illumio-core_label.label_test.href
}
`, rName1)
}

func testAccCheckIllumioLabeResource_updateValue(updatedValue string) string {
	return fmt.Sprintf(`
resource "illumio-core_label" "label_test" {
	key   = "app"
	value = %[1]q
}
`, updatedValue)
}
