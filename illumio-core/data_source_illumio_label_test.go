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

func TestAccIllumioLabelResource_Delete(t *testing.T) {
	labelHref := new(string)
	newLabelHref := new(string)
	resourceName := "illumio-core_label.label_test"
	labelValue := acctest.RandomWithPrefix(prefixLabel)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelResourceConfig_basic(labelValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(resourceName, labelHref),
					resource.TestCheckResourceAttr(resourceName, "key", "app"),
					resource.TestCheckResourceAttr(resourceName, "value", labelValue),
				),
			},
			{
				// check that an apply called after a label has been deleted
				// correctly destroys and recreates the resource
				PreConfig: deleteLabelFromPCE("app", labelValue, t),
				Config:    testAccCheckIllumioLabelResourceConfig_basic(labelValue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(resourceName, newLabelHref),
					testAccCheckCompareRefs(labelHref, newLabelHref, false),
					resource.TestCheckResourceAttr(resourceName, "key", "app"),
					resource.TestCheckResourceAttr(resourceName, "value", labelValue),
				),
			},
			{
				// check that a destroy called after a label has been deleted
				// doesn't throw an error
				PreConfig: deleteLabelFromPCE("app", labelValue, t),
				Destroy:   true,
				Config:    testAccCheckIllumioLabelResourceConfig_basic(labelValue),
			},
		},
	})
}

func testAccCheckIllumioLabelResourceConfig_basic(value string) string {
	return fmt.Sprintf(`
resource "illumio-core_label" "label_test" {
	key   = "app"
	value = %[1]q
}
`, value)
}

func testAccCheckIllumioLabelDataSourceConfig_basic() string {
	labelValue := acctest.RandomWithPrefix(prefixLabel)

	return testAccCheckIllumioLabelResourceConfig_basic(labelValue) + `
data "illumio-core_label" "label_test" {
	href = illumio-core_label.label_test.href
}`
}

func testAccCheckIllumioLabeResource_updateValue(updatedValue string) string {
	return fmt.Sprintf(`
resource "illumio-core_label" "label_test" {
	key   = "app"
	value = %[1]q
}
`, updatedValue)
}

// deleteLabelFromPCE removes the label with the given key/value
// via the API to test Terraform drift behaviour
func deleteLabelFromPCE(key, value string, t *testing.T) func() {
	return func() {
		conf := TestAccProvider.Meta().(Config)
		illumioClient := conf.IllumioClient

		endpoint := fmt.Sprintf("/orgs/%d/labels", illumioClient.OrgID)
		_, data, err := illumioClient.Get(endpoint, &map[string]string{
			"key":   key,
			"value": value,
		})
		if err != nil {
			t.Fatal("Failed to get label from PCE")
		}

		href := data.S("0", "href").Data().(string)

		_, err = illumioClient.Delete(href)
		if err != nil {
			t.Fatal("Failed to delete label from PCE")
		}
	}
}
