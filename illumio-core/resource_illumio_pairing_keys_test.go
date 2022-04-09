// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixPK string = "TF-ACC-PK"

func TestAccIllumioPairingKeys_CreateUpdate(t *testing.T) {
	resourceName := "illumio-core_pairing_keys.pk_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioPairingKeysConfig_basic(1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "activation_tokens.#", "1"),
				),
			},
			{
				Config: testAccCheckIllumioPairingKeysConfig_basic(2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "activation_tokens.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioPairingKeysConfig_basic(val int) string {
	rName1 := acctest.RandomWithPrefix(prefixPK)

	if e := os.Getenv("ILLUMIO_AES_GCM_KEY"); e == "" {
		keyBytes := make([]byte, 32)
		_, _ = rand.Read(keyBytes)

		pairingKey := hex.EncodeToString(keyBytes)
		os.Setenv("ILLUMIO_AES_GCM_KEY", pairingKey)
	}

	return fmt.Sprintf(`
resource "illumio-core_pairing_profile" "pk_test" {
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

resource "illumio-core_pairing_keys" "pk_test" {
	pairing_profile_href = illumio-core_pairing_profile.pk_test.href
	token_count          = %[2]d
}
`, rName1, val)
}
