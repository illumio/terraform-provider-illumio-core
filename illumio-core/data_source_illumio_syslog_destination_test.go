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

var providerDSSyslogDestination *schema.Provider

func TestAccIllumioSyslogDestination_Read(t *testing.T) {
	sysDestAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSSyslogDestination),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSyslogDestinationDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceSyslogDestinationExists("data.illumio-core_syslog_destination.test", sysDestAttr),
					testAccCheckIllumioSyslogDestinationDataSourceAttributes(sysDestAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioSyslogDestinationDataSourceConfig_basic() string {
	return `
	data "illumio-core_syslog_destination" "test" {
		href = "/orgs/1/settings/syslog/destinations/11a4cfdf-a78e-4144-bbbc-67faec728df1"
	}
	`
}

func testAccCheckIllumioDataSourceSyslogDestinationExists(name string, sysDestAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Syslog Destination %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSSyslogDestination).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"type",
			"description",
			"node_status_logger.node_status_included",
		} {
			sysDestAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioSyslogDestinationDataSourceAttributes(sysDestAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"type":        "remote_syslog",
			"description": "splunk-dev3",
			"node_status_logger.node_status_included": true,
		}
		for k, v := range expectation {
			if sysDestAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, sysDestAttr[k], v)
			}
		}

		return nil
	}
}
