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
	labelGroupName := acctest.RandomWithPrefix(prefixLG)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioLGDataSourceConfig_basic(labelGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "labels", resourceName, "labels"),
					resource.TestCheckResourceAttrPair(dataSourceName, "sub_groups", resourceName, "sub_groups"),
				),
			},
			{
				Config: testAccCheckIllumioLGResource_removeLabels(labelGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "labels.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "subgroups.#", "0"),
				),
			},
			{
				Config: testAccCheckIllumioLGResource_emptyDescription(labelGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "subgroups.#", "0"),
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

func testAccCheckIllumioLGDataSourceConfig_basic(labelGroupName string) string {
	rName1 := acctest.RandomWithPrefix(prefixLG)
	rName2 := acctest.RandomWithPrefix(prefixLG)

	// create_before_destroy is needed here to avoid a reference
	// error when removing the referenced objects in the subsequent
	// step as the apply will attempt to delete them before the
	// label group is updated
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

	lifecycle {
		create_before_destroy = true
	}
}

data "illumio-core_label_group" "lg_test" {
	href = illumio-core_label_group.lg_test.href
}
`, rName1, rName2, labelGroupName)
}

func testAccCheckIllumioLGResource_removeLabels(labelGroupName string) string {
	return fmt.Sprintf(`
resource "illumio-core_label_group" "lg_test" {
	key           = "role"
	name          = %[1]q
	description   = "Terraform Label Group test"
}
`, labelGroupName)
}

func testAccCheckIllumioLGResource_emptyDescription(labelGroupName string) string {
	return fmt.Sprintf(`
resource "illumio-core_label_group" "lg_test" {
	key           = "role"
	name          = %[1]q
	description   = ""
}
`, labelGroupName)
}
