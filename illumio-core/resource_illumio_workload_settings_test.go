package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerWorkloadSettings *schema.Provider

func TestAccIllumioWorkloadSettings_CreateUpdate(t *testing.T) {
	wsAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerWorkloadSettings),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerWorkloadSettings, "illumio-core_workload_settings", false),
		Steps: []resource.TestStep{
			// TODO: Check for the import thingy not working without config and the resource cannot be created.
			{
				Config:        testAccCheckIllumioWorkloadSettingsConfig_basic(),
				ResourceName:  "illumio-core_workload_settings.test",
				ImportStateId: "/orgs/1/settings/workloads",
				ImportState:   true,
			},
			{
				Config: testAccCheckIllumioWorkloadSettingsConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWorkloadSettingsExists("illumio-core_workload_settings.test", wsAttr),
					testAccCheckIllumioWorkloadSettingsAttributes(wsAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadSettingsConfig_basic() string {
	return `
	resource "illumio-core_workload_settings" "test" {
		workload_disconnected_timeout_seconds {
		  value = -1
		}
		workload_goodbye_timeout_seconds {
		  value = -1
		}
	  }
	`
}

func testAccCheckIllumioWorkloadSettingsExists(name string, wsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("IP List %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerWorkloadSettings).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"workload_disconnected_timeout_seconds.0.value",
			"workload_goodbye_timeout_seconds.0.value",
		} {
			wsAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioWorkloadSettingsAttributes(wsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"workload_disconnected_timeout_seconds.0.value": -1,
			"workload_goodbye_timeout_seconds.0.value":      -1,
		}
		for k, v := range expectation {
			if wsAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, wsAttr[k], v)
			}
		}

		return nil
	}
}
