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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioPPDataSourceConfig_basic(),
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
		},
	})
}

func testAccCheckIllumioPPDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixPP)
	rName2 := acctest.RandomWithPrefix(prefixPP)

	return fmt.Sprintf(`
resource "illumio-core_label" "pp_test" {
	key   = "role"
	value = %[1]q
}

resource "illumio-core_pairing_profile" "pp_test" {
	name    = %[2]q
	enabled = false

	labels {
		href = illumio-core_label.pp_test.href
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
}

data "illumio-core_pairing_profile" "pp_test" {
	href = illumio-core_pairing_profile.pp_test.href
}
`, rName1, rName2)
}
