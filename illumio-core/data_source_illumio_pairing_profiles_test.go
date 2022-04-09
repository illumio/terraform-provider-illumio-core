// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixPPL string = "TF-ACC-PPL"

func TestAccIllumioPPL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_pairing_profiles.ppl_test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioPPLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioPPLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixPPL)
	rName2 := acctest.RandomWithPrefix(prefixPPL)

	return fmt.Sprintf(`
resource "illumio-core_pairing_profile" "ppl_test1" {
	name    = %[1]q
	enabled = false

	allowed_uses_per_key  = "unlimited"
	role_label_lock       = true
	app_label_lock        = true
	env_label_lock        = true
	loc_label_lock        = true
	log_traffic           = false
	log_traffic_lock      = true
	visibility_level      = "flow_off"
	visibility_level_lock = false
	enforcement_mode      = "visibility_only"
}

resource "illumio-core_pairing_profile" "ppl_test2" {
	name    = %[2]q
	enabled = false

	allowed_uses_per_key  = "unlimited"
	role_label_lock       = true
	app_label_lock        = true
	env_label_lock        = true
	loc_label_lock        = true
	log_traffic           = false
	log_traffic_lock      = true
	visibility_level      = "flow_off"
	visibility_level_lock = false
	enforcement_mode      = "visibility_only"
}

data "illumio-core_pairing_profiles" "ppl_test" {
	# lookup based on partial match
	name = %[3]q

	# enforce dependencies
	depends_on = [
		illumio-core_pairing_profile.ppl_test1,
		illumio-core_pairing_profile.ppl_test2,
	]
}
`, rName1, rName2, prefixPPL)
}
