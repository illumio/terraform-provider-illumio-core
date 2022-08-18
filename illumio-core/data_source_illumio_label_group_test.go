// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixLG string = "TF-ACC-LG"

func TestAccIllumioLG_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_label_group.lg_test"
	resourceName := "illumio-core_label_group.lg_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLGDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "labels", resourceName, "labels"),
					resource.TestCheckResourceAttrPair(dataSourceName, "sub_groups", resourceName, "sub_groups"),
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

func testAccCheckIllumioLGDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixLG)
	rName2 := acctest.RandomWithPrefix(prefixLG)
	rName3 := acctest.RandomWithPrefix(prefixLG)

	return fmt.Sprintf(`
resource "illumio-core_label" "lg_test" {
	key   = "role"
	value = %[1]q
}

resource "illumio-core_label_group" "lg_subgroup" {
	key           = "role"
	name          = %[2]q
	description   = "Terraform Label Group subgroup"
}

resource "illumio-core_label_group" "lg_test" {
	key           = "role"
	name          = %[3]q
	description   = "Terraform Label Group test"
	labels {
		href = illumio-core_label.lg_test.href
	}
	sub_groups {
		href = illumio-core_label_group.lg_subgroup.href
	}
}

data "illumio-core_label_group" "lg_test" {
	href = illumio-core_label_group.lg_test.href
}
`, rName1, rName2, rName3)
}
