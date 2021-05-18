// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSWorkloadSettings *schema.Provider

func TestAccIllumioWorkloadSettings_Read(t *testing.T) {
	workloadSettingsAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSWorkloadSettings),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadSettingsDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceWorkloadSettingsExists("data.illumio-core_workload_settings.test", workloadSettingsAttr),
					testAccCheckIllumioWorkloadSettingsDataSourceAttributes(workloadSettingsAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadSettingsDataSourceConfig_basic() string {
	return `
	data "illumio-core_workload_settings" "test" {}
	`
}

func testAccCheckIllumioDataSourceWorkloadSettingsExists(name string, workloadSettingsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Workload Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSWorkloadSettings).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"workload_disconnected_timeout_seconds.0.value",
			"workload_goodbye_timeout_seconds.0.value",
		} {
			workloadSettingsAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioWorkloadSettingsDataSourceAttributes(workloadSettingsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"workload_disconnected_timeout_seconds.0.value": float64(3600),
			"workload_goodbye_timeout_seconds.0.value":      float64(900),
		}
		for k, v := range expectation {
			if workloadSettingsAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, workloadSettingsAttr[k], v)
			}
		}

		return nil
	}
}
