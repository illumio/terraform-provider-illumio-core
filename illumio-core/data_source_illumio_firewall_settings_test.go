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

var providerDSFS *schema.Provider

func TestAccIllumioFS_Read(t *testing.T) {
	sAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSFS),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioFSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceFSExists("data.illumio-core_firewall_settings.test", sAttr),
					testAccCheckIllumioFSDataSourceAttributes(sAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioFSDataSourceConfig_basic() string {
	return `
	data "illumio-core_firewall_settings" "test" {
		href = "/orgs/1/sec_policy/draft/firewall_settings"
	}
	`
}

func testAccCheckIllumioDataSourceFSExists(name string, sAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Firewall Settings %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSFS).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"ipv6_mode",
			"network_detection_mode",
			"ike_authentication_type",
			"static_policy_scopes.0.0.label.href",
			"allow_dhcp_client",
			"log_dropped_multicast",
			"log_dropped_broadcast",
			"allow_traceroute",
			"allow_ipv6",
		} {
			sAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioFSDataSourceAttributes(sAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"ipv6_mode":                           "policy_based",
			"network_detection_mode":              "single_private_brn",
			"ike_authentication_type":             "psk",
			"static_policy_scopes.0.0.label.href": "/orgs/1/labels/14",
			"allow_dhcp_client":                   true,
			"log_dropped_multicast":               true,
			"log_dropped_broadcast":               false,
			"allow_traceroute":                    true,
			"allow_ipv6":                          true,
		}
		for k, v := range expectation {
			if sAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, sAttr[k], v)
			}
		}

		return nil
	}
}
