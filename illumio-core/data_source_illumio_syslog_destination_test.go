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

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSDDataSourceConfig_basic(),
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
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIllumioSDDataSourceConfig_basic() string {
	u, _ := url.Parse(os.Getenv("ILLUMIO_PCE_HOST"))
	hostName := u.Hostname()

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
`, hostName)
}
