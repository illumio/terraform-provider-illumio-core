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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLabelTypeDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "key", resourceName, "key"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_name", resourceName, "display_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "display_info", resourceName, "display_info"),
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

func testAccCheckIllumioLabelTypeDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixLabelType)
	rName2 := acctest.RandomWithPrefix(prefixLabelType)

	return fmt.Sprintf(`
resource "illumio-core_label_type" "label_type_test" {
	key          = %[1]q
	display_name = %[2]q

	display_info {
		initial = "TS"
	}
}

data "illumio-core_label_type" "label_type_test" {
	href = illumio-core_label_type.label_type_test.href
}
`, rName1, rName2)
}
