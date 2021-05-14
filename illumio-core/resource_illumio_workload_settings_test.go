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

	var err error

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerWorkloadSettings),
		Steps: []resource.TestStep{
			{
				Config:             testAccCheckIllumioWorkloadSettingsConfig_basic(),
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      "illumio-core_workload_settings.test",
				ImportStateId:     "/orgs/1/settings/workloads",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"test_href",
				},
			},
			{
				Config: testAccCheckIllumioWorkloadSettingsConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWorkloadSettingsExists("illumio-core_workload_settings.test", wsAttr),
					testAccCheckIllumioWorkloadSettingsAttributes(wsAttr, &err),
				),
			},
			{
				Config: testAccCheckIllumioWorkloadSettingsConfig_revert(),
				Check: resource.ComposeTestCheckFunc(
					throwErrorOccuredInValidationStep(err),
				),
			},
		},
	})
}

func throwErrorOccuredInValidationStep(err error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if err != nil {
			return fmt.Errorf("Attribute Validation error: %v", err)
		}

		return nil
	}
}

func testAccCheckIllumioWorkloadSettingsConfig_revert() string {
	return `
	resource "illumio-core_workload_settings" "test" {
		test_href = "/orgs/1/settings/workloads"
		workload_disconnected_timeout_seconds {
		  value = 3600
		}
		workload_goodbye_timeout_seconds {
		  value = 900
		}
	  }
	`
}

func testAccCheckIllumioWorkloadSettingsConfig_basic() string {
	return `
	resource "illumio-core_workload_settings" "test" {
		test_href = "/orgs/1/settings/workloads"
		workload_disconnected_timeout_seconds {
		  value = 3399
		}
		workload_goodbye_timeout_seconds {
		  value = 3399
		}
	  }
	`
}

func testAccCheckIllumioWorkloadSettingsExists(name string, wsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Workload Settings %s not found", name)
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

func testAccCheckIllumioWorkloadSettingsAttributes(wsAttr map[string]interface{}, err *error) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"workload_disconnected_timeout_seconds.0.value": float64(3399),
			"workload_goodbye_timeout_seconds.0.value":      float64(3399),
		}
		for k, v := range expectation {
			if wsAttr[k] != v {
				*err = fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, wsAttr[k], v)
				return nil
			}
		}

		return nil
	}
}
