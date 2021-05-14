package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSTrafficCollectorSettings *schema.Provider

func TestAccIllumioTrafficCollectorSettings_Read(t *testing.T) {
	tcsAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSTrafficCollectorSettings),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioTrafficCollectorSettingsDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceTrafficCollectorSettingsExists("data.illumio-core_traffic_collector_settings.test", tcsAttr),
					testAccCheckIllumioTrafficCollectorSettingsDataSourceAttributes(tcsAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioTrafficCollectorSettingsDataSourceConfig_basic() string {
	return `
	data "illumio-core_traffic_collector_settings" "test" {
		href = "/orgs/1/settings/traffic_collector/2d9d2170-520e-42c4-92bd-cdf2216a1dab"
	}
	`
}

func testAccCheckIllumioDataSourceTrafficCollectorSettingsExists(name string, tcsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Traffic Collector Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSTrafficCollectorSettings).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"transmission",
			"target.proto",
			"action",
		} {
			tcsAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioTrafficCollectorSettingsDataSourceAttributes(tcsAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"transmission": "B",
			"target.proto": float64(6),
			"action":       "drop",
		}
		for k, v := range expectation {
			if tcsAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, tcsAttr[k], v)
			}
		}

		return nil
	}
}
