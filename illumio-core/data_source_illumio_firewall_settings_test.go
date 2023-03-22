// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIllumioFS_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_firewall_settings.fs_test"
	resourceName := "illumio-core_firewall_settings.fs_test"

	ikeAuthenticationType := "psk"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t); testAccPreCheckSaaSPCE(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioFSDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "ike_authentication_type", resourceName, "ike_authentication_type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "firewall_coexistence", resourceName, "firewall_coexistence"),
					resource.TestCheckResourceAttrPair(dataSourceName, "static_policy_scopes", resourceName, "static_policy_scopes"),
					resource.TestCheckResourceAttrPair(dataSourceName, "containers_inherit_host_policy_scopes", resourceName, "containers_inherit_host_policy_scopes"),
					resource.TestCheckResourceAttrPair(dataSourceName, "loopback_interfaces_in_policy_scopes", resourceName, "loopback_interfaces_in_policy_scopes"),
					resource.TestCheckResourceAttrPair(dataSourceName, "blocked_connection_reject_scopes", resourceName, "blocked_connection_reject_scopes"),
				),
			},
			{
				Config: testAccCheckIllumioFSResource_updateIKEAuthType(ikeAuthenticationType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "ike_authentication_type", ikeAuthenticationType),
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

func testAccCheckIllumioFSDataSourceConfig_basic() string {
	return `
resource "illumio-core_firewall_settings" "fs_test" {}

data "illumio-core_firewall_settings" "fs_test" {}
`
}

func testAccCheckIllumioFSResource_updateIKEAuthType(ikeAuthenticationType string) string {
	return fmt.Sprintf(`
resource "illumio-core_firewall_settings" "fs_test" {
	ike_authentication_type = %[1]q
}
`, ikeAuthenticationType)
}
