// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixPP string = "TF-ACC-PP"

func TestAccIllumioPP_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_pairing_profile.pp_test"
	resourceName := "illumio-core_pairing_profile.pp_test"

	ppName := acctest.RandomWithPrefix(prefixPP)
	labelName := acctest.RandomWithPrefix(prefixPP)

	updatedName := acctest.RandomWithPrefix(prefixPP)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioPPDataSourceConfig_basic(ppName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
					resource.TestCheckResourceAttrPair(dataSourceName, "labels", resourceName, "labels"),
					resource.TestCheckResourceAttrPair(dataSourceName, "allowed_uses_per_key", resourceName, "allowed_uses_per_key"),
					resource.TestCheckResourceAttrPair(dataSourceName, "role_label_lock", resourceName, "role_label_lock"),
					resource.TestCheckResourceAttrPair(dataSourceName, "app_label_lock", resourceName, "app_label_lock"),
					resource.TestCheckResourceAttrPair(dataSourceName, "env_label_lock", resourceName, "env_label_lock"),
					resource.TestCheckResourceAttrPair(dataSourceName, "loc_label_lock", resourceName, "loc_label_lock"),
					resource.TestCheckResourceAttrPair(dataSourceName, "log_traffic", resourceName, "log_traffic"),
					resource.TestCheckResourceAttrPair(dataSourceName, "log_traffic_lock", resourceName, "log_traffic_lock"),
					resource.TestCheckResourceAttrPair(dataSourceName, "visibility_level", resourceName, "visibility_level"),
					resource.TestCheckResourceAttrPair(dataSourceName, "visibility_level_lock", resourceName, "visibility_level_lock"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enforcement_mode", resourceName, "enforcement_mode"),
				),
			},
			{
				Config: testAccCheckIllumioPPResource_updateAddLabel(ppName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioPPResource_updateRemoveLabels(ppName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "enforcement_mode", "idle"),
					resource.TestCheckResourceAttr(resourceName, "labels.#", "0"),
				),
			},
			{
				Config: testAccCheckIllumioPPResource_updateNameAndDesc(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
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

func ppRoleLabel(labelName string) string {
	return fmt.Sprintf(`
resource "illumio-core_label" "pp_role" {
	key   = "role"
	value = %[1]q
}
`, labelName)
}

func testAccCheckIllumioPPDataSourceConfig_basic(ppName, labelName string) string {
	return ppRoleLabel(labelName) + fmt.Sprintf(`
resource "illumio-core_pairing_profile" "pp_test" {
	name        = %[1]q
	description = "Terraform Pairing Profile test"
	enabled     = false

	labels {
		href = illumio-core_label.pp_role.href
	}

	allowed_uses_per_key  = "unlimited"
	role_label_lock       = true
	app_label_lock        = true
	env_label_lock        = false
	loc_label_lock        = true
	log_traffic           = false
	log_traffic_lock      = true
	visibility_level      = "flow_off"
	visibility_level_lock = false
	enforcement_mode      = "visibility_only"

	lifecycle {
		create_before_destroy = true
	}
}

data "illumio-core_pairing_profile" "pp_test" {
	href = illumio-core_pairing_profile.pp_test.href
}
`, ppName)
}

func testAccCheckIllumioPPResource_updateAddLabel(ppName, labelName string) string {
	appLabelName := acctest.RandomWithPrefix(prefixPP)

	return ppRoleLabel(labelName) + fmt.Sprintf(`
resource "illumio-core_label" "pp_app" {
	key   = "app"
	value = %[1]q
}

resource "illumio-core_pairing_profile" "pp_test" {
	name        = %[2]q
	description = "Terraform Pairing Profile test"
	enabled     = true

	labels {
		href = illumio-core_label.pp_role.href
	}

	labels {
		href = illumio-core_label.pp_app.href
	}

	allowed_uses_per_key  = "unlimited"
	role_label_lock       = true
	app_label_lock        = true
	env_label_lock        = false
	loc_label_lock        = true
	log_traffic           = false
	log_traffic_lock      = true
	visibility_level      = "flow_off"
	visibility_level_lock = false
	enforcement_mode      = "visibility_only"

	lifecycle {
		create_before_destroy = true
	}
}
`, appLabelName, ppName)
}

func testAccCheckIllumioPPResource_updateRemoveLabels(ppName string) string {
	return fmt.Sprintf(`
resource "illumio-core_pairing_profile" "pp_test" {
	name        = %[1]q
	description = "Terraform Pairing Profile test"
	enabled     = false

	allowed_uses_per_key  = "unlimited"
	role_label_lock       = true
	app_label_lock        = true
	env_label_lock        = false
	loc_label_lock        = true
	log_traffic           = false
	log_traffic_lock      = true
	visibility_level      = "flow_off"
	visibility_level_lock = false
	enforcement_mode      = "idle"
}
`, ppName)
}

func testAccCheckIllumioPPResource_updateNameAndDesc(updatedName string) string {
	return fmt.Sprintf(`
resource "illumio-core_pairing_profile" "pp_test" {
	name        = %[1]q
	description = ""
	enabled     = false

	allowed_uses_per_key  = "unlimited"
	role_label_lock       = true
	app_label_lock        = true
	env_label_lock        = false
	loc_label_lock        = true
	log_traffic           = false
	log_traffic_lock      = true
	visibility_level      = "flow_off"
	visibility_level_lock = false
	enforcement_mode      = "idle"
}
`, updatedName)
}
