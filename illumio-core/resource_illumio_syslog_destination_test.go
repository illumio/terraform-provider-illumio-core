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

var providerSyslogD *schema.Provider

func TestAccIllumioSyslogDestination_CreateUpdate(t *testing.T) {
	serAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerSyslogD),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerSyslogD, "illumio-core_syslog_destination", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSyslogDestinationConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioSyslogDestinationExists("illumio-core_syslog_destination.test", serAttr),
					testAccCheckIllumioSyslogDestinationAttributes("creation from terraform", serAttr),
				),
			},
			{
				Config: testAccCheckIllumioSyslogDestinationConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioSyslogDestinationExists("illumio-core_syslog_destination.test", serAttr),
					testAccCheckIllumioSyslogDestinationAttributes("updation from terraform", serAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioSyslogDestinationConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_syslog_destination" "test" {
		type        = "local_syslog"
		pce_scope   = ["crest-mnc.ilabs.io"]
		description = "%s"
	  
		audit_event_logger {
		  configuration_event_included = true
		  system_event_included        = true
		  min_severity                 = "informational"
		}
	  
		traffic_event_logger {
		  traffic_flow_allowed_event_included             = true
		  traffic_flow_potentially_blocked_event_included = false
		  traffic_flow_blocked_event_included             = false
		}
	  
		node_status_logger {
		  node_status_included = true
		}
	  }
	  
	`, val)
}

func testAccCheckIllumioSyslogDestinationExists(name string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Syslog Destination %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerSyslogD).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"type",
			"description",
			"audit_event_logger.configuration_event_included",
			"traffic_event_logger.traffic_flow_allowed_event_included",
			"node_status_logger.node_status_included",
		} {
			serAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioSyslogDestinationAttributes(val string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"type":        "local_syslog",
			"description": val,
			"audit_event_logger.configuration_event_included":          true,
			"traffic_event_logger.traffic_flow_allowed_event_included": true,
			"node_status_logger.node_status_included":                  true,
		}

		for k, v := range expectation {
			if serAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, serAttr[k], v)
			}
		}

		return nil
	}
}
