// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIllumioSDL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_syslog_destinations.sdl_test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSDLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceSDLExists(dataSourceName),
					testAccCheckIllumioDataSourceSDLNonEmpty(dataSourceName),
				),
			},
		},
	})
}

func testAccCheckIllumioSDLDataSourceConfig_basic() string {
	u, _ := url.Parse(os.Getenv("ILLUMIO_PCE_HOST"))
	hostName := u.Hostname()

	return fmt.Sprintf(`
resource "illumio-core_syslog_destination" "sdl_test" {
	type        = "local_syslog"
	pce_scope   = [%[1]q]
	description = "Terraform Syslog Destinations test"

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

data "illumio-core_syslog_destinations" "sdl_test" {
	# enforce dependency
	depends_on = [
		illumio-core_syslog_destination.sdl_test,
	]
}
`, hostName)
}

func testAccCheckIllumioDataSourceSDLExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Syslog Destinations %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		return nil
	}
}

func testAccCheckIllumioDataSourceSDLNonEmpty(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("List of Syslog Destinations %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		if v, ok := rs.Primary.Attributes["items.#"]; ok {
			count, _ := strconv.Atoi(v)
			if count <= 0 {
				return fmt.Errorf("Empty list of Syslog Destinations")
			}
			return nil
		}

		return fmt.Errorf("Error reading item count in Syslog Destinations data source")
	}
}
