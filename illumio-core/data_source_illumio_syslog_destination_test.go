// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIllumioSD_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_syslog_destination.sd_test"
	resourceName := "illumio-core_syslog_destination.sd_test"

	u, _ := url.Parse(os.Getenv("ILLUMIO_PCE_HOST"))
	hostname := u.Hostname()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t); testAccPreCheckSaaSPCE(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSDDataSourceConfig_basic(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "type", resourceName, "type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "pce_scope", resourceName, "pce_scope"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "audit_event_logger", resourceName, "audit_event_logger"),
					resource.TestCheckResourceAttrPair(dataSourceName, "traffic_event_logger", resourceName, "traffic_event_logger"),
					resource.TestCheckResourceAttrPair(dataSourceName, "node_status_logger", resourceName, "node_status_logger"),
				),
			},
			{
				Config: testAccCheckIllumioSDResource_updateLoggerSettings(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "audit_event_logger.0.configuration_event_included", "false"),
					resource.TestCheckResourceAttr(resourceName, "traffic_event_logger.0.traffic_flow_allowed_event_included", "false"),
					resource.TestCheckResourceAttr(resourceName, "traffic_event_logger.0.traffic_flow_potentially_blocked_event_included", "true"),
					resource.TestCheckResourceAttr(resourceName, "traffic_event_logger.0.traffic_flow_blocked_event_included", "true"),
					resource.TestCheckResourceAttr(resourceName, "node_status_logger.0.node_status_included", "false"),
				),
			},
			{
				Config: testAccCheckIllumioSDResource_updateDesc(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIllumioSDDataSourceConfig_basic(hostname string) string {
	return fmt.Sprintf(`
resource "illumio-core_syslog_destination" "sd_test" {
	type        = "local_syslog"
	pce_scope   = [%[1]q]
	description = "Terraform Syslog Destination test"

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

data "illumio-core_syslog_destination" "sd_test" {
	href = illumio-core_syslog_destination.sd_test.href
}
`, hostname)
}

func testAccCheckIllumioSDResource_updateLoggerSettings(hostname string) string {
	return fmt.Sprintf(`
resource "illumio-core_syslog_destination" "sd_test" {
	type        = "local_syslog"
	pce_scope   = [%[1]q]
	description = "Terraform Syslog Destination test"

	audit_event_logger {
	  configuration_event_included = false
	  system_event_included        = true
	  min_severity                 = "informational"
	}

	traffic_event_logger {
	  traffic_flow_allowed_event_included             = false
	  traffic_flow_potentially_blocked_event_included = true
	  traffic_flow_blocked_event_included             = true
	}

	node_status_logger {
	  node_status_included = false
	}
}
`, hostname)
}

func testAccCheckIllumioSDResource_updateDesc(hostname string) string {
	return fmt.Sprintf(`
resource "illumio-core_syslog_destination" "sd_test" {
	type        = "local_syslog"
	pce_scope   = [%[1]q]
	description = ""

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
`, hostname)
}
