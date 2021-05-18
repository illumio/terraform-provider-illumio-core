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

var providerTCS *schema.Provider

func TestAccIllumioTrafficCollectorSettings_CreateUpdate(t *testing.T) {
	serAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerTCS),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerTCS, "illumio-core_traffic_collector_settings", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioTrafficCollectorSettingsConfig_basic("1.1.1.2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioTrafficCollectorSettingsExists("illumio-core_traffic_collector_settings.test", serAttr),
					testAccCheckIllumioTrafficCollectorSettingsAttributes("1.1.1.2", serAttr),
				),
			},
			{
				Config: testAccCheckIllumioTrafficCollectorSettingsConfig_basic("10.0.0.0/8"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioTrafficCollectorSettingsExists("illumio-core_traffic_collector_settings.test", serAttr),
					testAccCheckIllumioTrafficCollectorSettingsAttributes("10.0.0.0/8", serAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioTrafficCollectorSettingsConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_traffic_collector_settings" "test" {
		action       = "drop"
		transmission = "broadcast"
		target {
		  dst_ip   = "%s"
		  dst_port = 10
		  proto    = 6
		}
	  }
	`, val)
}

func testAccCheckIllumioTrafficCollectorSettingsExists(name string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Traffic Collector Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerTCS).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"action",
			"transmission",
			"target.proto",
			"target.dst_ip",
			"target.dst_port",
		} {
			serAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioTrafficCollectorSettingsAttributes(val string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"action":          "drop",
			"transmission":    "B", // API returns B for broadcast
			"target.proto":    float64(6),
			"target.dst_ip":   val,
			"target.dst_port": float64(10),
		}

		for k, v := range expectation {
			if serAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, serAttr[k], v)
			}
		}

		return nil
	}
}
