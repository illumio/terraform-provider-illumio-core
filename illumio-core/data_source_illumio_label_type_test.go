// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLabelType string = "TF-ACC-LabelType"

func TestAccIllumioLabelType_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_label_type.label_type_test"
	resourceName := "illumio-core_label_type.label_type_test"
	labelTypeKey := acctest.RandomWithPrefix(prefixLabelType)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelTypeDataSourceConfig_basic(labelTypeKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "key", resourceName, "key"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_info", resourceName, "display_info"),
				),
			},
			{
				Config: testAccCheckIllumioLabelTypeResource_updateInitial(labelTypeKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "display_info.0.initial", "UP"),
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

func TestAccIllumioLabelType_Delete(t *testing.T) {
	labelTypeHref := new(string)
	newLabelTypeHref := new(string)
	resourceName := "illumio-core_label_type.label_type_test"
	labelTypeKey := acctest.RandomWithPrefix(prefixLabelType)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelTypeResourceConfig_basic(labelTypeKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(resourceName, labelTypeHref),
					resource.TestCheckResourceAttr(resourceName, "key", labelTypeKey),
					resource.TestCheckResourceAttr(resourceName, "display_name", labelTypeKey),
				),
			},
			{
				// check that an apply called after a label type has been deleted
				// correctly destroys and recreates the resource
				PreConfig: deleteFromPCE(labelTypeHref, t),
				Config:    testAccCheckIllumioLabelTypeResourceConfig_basic(labelTypeKey),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(resourceName, newLabelTypeHref),
					testAccCheckCompareRefs(labelTypeHref, newLabelTypeHref, false),
					resource.TestCheckResourceAttr(resourceName, "key", labelTypeKey),
					resource.TestCheckResourceAttr(resourceName, "display_name", labelTypeKey),
				),
			},
			{
				// check that a destroy called after a label type has been deleted
				// doesn't throw an error
				PreConfig: deleteFromPCE(labelTypeHref, t),
				Destroy:   true,
				Config:    testAccCheckIllumioLabelTypeResourceConfig_basic(labelTypeKey),
			},
		},
	})
}

func testAccCheckIllumioLabelTypeResourceConfig_basic(labelTypeKey string) string {
	return fmt.Sprintf(`
resource "illumio-core_label_type" "label_type_test" {
	key          = %[1]q
	display_name = %[1]q

	display_info {
		initial = "TS"
	}
}
`, labelTypeKey)
}

func testAccCheckIllumioLabelTypeDataSourceConfig_basic(labelTypeKey string) string {
	return testAccCheckIllumioLabelTypeResourceConfig_basic(labelTypeKey) + `
data "illumio-core_label_type" "label_type_test" {
	href = illumio-core_label_type.label_type_test.href
}`
}

func testAccCheckIllumioLabelTypeResource_updateInitial(labelTypeKey string) string {
	return fmt.Sprintf(`
resource "illumio-core_label_type" "label_type_test" {
	key          = %[1]q
	display_name = %[1]q

	display_info {
		initial = "UP"
	}
}
`, labelTypeKey)
}
