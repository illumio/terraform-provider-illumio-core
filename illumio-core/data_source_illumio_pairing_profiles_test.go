// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixPPL string = "TF-ACC-PPL"

func init() {
	resource.AddTestSweepers("pairing_profiles", &resource.Sweeper{
		Name: "pairing_profiles",
		F:    sweep("pairing profile", "name", prefixPPL, "/orgs/%d/pairing_profiles"),
	})
}

func TestAccIllumioPPL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_pairing_profiles.ppl_test"
	profileName := acctest.RandomWithPrefix(prefixPPL)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioPPLDataSourceConfig_basic(profileName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
			{
				Config: testAccCheckIllumioPPLDataSourceConfig_exactMatch(profileName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "1"),
					resource.TestCheckResourceAttr(dataSourceName, "items.0.name", profileName),
				),
			},
		},
	})
}

func pplConfig(profileName string) string {
	rName := acctest.RandomWithPrefix(prefixPPL)

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
`, rName, profileName)
}

func testAccCheckIllumioPPLDataSourceConfig_basic(profileName string) string {
	return pplConfig(profileName) + fmt.Sprintf(`
data "illumio-core_pairing_profiles" "ppl_test" {
	# lookup based on partial match
	name = %[1]q

	# enforce dependencies
	depends_on = [
		illumio-core_pairing_profile.ppl_test1,
		illumio-core_pairing_profile.ppl_test2,
	]
}
`, prefixPPL)
}

func testAccCheckIllumioPPLDataSourceConfig_exactMatch(profileName string) string {
	return pplConfig(profileName) + fmt.Sprintf(`
data "illumio-core_pairing_profiles" "ppl_test" {
	# lookup using exact match
	name       = %[1]q
	match_type = "exact"

	# enforce dependencies
	depends_on = [
		illumio-core_pairing_profile.ppl_test1,
		illumio-core_pairing_profile.ppl_test2,
	]
}
`, profileName)
}
