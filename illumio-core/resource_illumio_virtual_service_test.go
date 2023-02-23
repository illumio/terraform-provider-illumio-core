// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testResourceIllumioVirtualServiceStateDataV1() map[string]any {
	return map[string]any{
		"name": "VS-CRM-DB",
		"href": "/orgs/1/sec_policy/draft/virtual_services/1edaa88f-e905-4682-8623-df55e3d2a966",
		"id":   "/orgs/1/sec_policy/draft/virtual_services/1edaa88f-e905-4682-8623-df55e3d2a966",
		// other fields omitted for brevity, having name/href/id
		// shows that other fields are kept as-is
		"service_addresses": []map[string]any{
			{
				"ip":           "127.0.0.1",
				"network_href": "/orgs/1/networks/a2f3fdfc-c179-4700-8525-133727c3bab4",
			},
		},
	}
}

func testResourceIllumioVirtualServiceStateDataV2() map[string]any {
	return map[string]any{
		"name": "VS-CRM-DB",
		"href": "/orgs/1/sec_policy/draft/virtual_services/1edaa88f-e905-4682-8623-df55e3d2a966",
		"id":   "/orgs/1/sec_policy/draft/virtual_services/1edaa88f-e905-4682-8623-df55e3d2a966",
		"service_addresses": []map[string]any{
			{
				"ip":      "127.0.0.1",
				"network": []map[string]any{{"href": "/orgs/1/networks/a2f3fdfc-c179-4700-8525-133727c3bab4"}},
			},
		},
	}
}

func TestResourceIllumioVirtualServiceStateUpgradeV1(t *testing.T) {
	ctx := context.Background()

	expected := testResourceIllumioVirtualServiceStateDataV2()
	actual, err := resourceIllumioVirtualServiceStateUpgradeV1(ctx, testResourceIllumioVirtualServiceStateDataV1(), nil)
	if err != nil {
		t.Fatalf("error migrating state: %s", err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
	}
}

func TestAccResourceIllumioVS_StateUpgradeV1(t *testing.T) {
	resourceName := "illumio-core_virtual_service.vs_test"
	vsName := acctest.RandomWithPrefix(prefixVS)

	// XXX: these setup steps need to be run before the TestCase
	// PreChecks due to Config functions being executed before the
	// PreCheck function
	testAccPreCheck(t)
	networkID := setupNetworkID(t)

	resource.ParallelTest(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				// use an ExternalProvider to get the version prior to the
				// change requiring a schema version update
				// DEV NOTE: you will need to disable dev overrides for this
				// acceptance test to run
				ExternalProviders: map[string]resource.ExternalProvider{
					"illumio-core": {
						VersionConstraint: "1.0.3",
						Source:            "illumio/illumio-core",
					},
				},
				Config: testAccCheckIllumioVSResourceConfig_v1(vsName, networkID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "href"),
					resource.TestCheckResourceAttr(resourceName, "apply_to", "host_only"),
					resource.TestCheckResourceAttr(resourceName, "service_addresses.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "service_addresses.0.network_href"),
				),
			},
			{
				ProviderFactories: TestAccProviderFactories,
				Config:            testAccCheckIllumioVSResourceConfig_v2(vsName, networkID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "href"),
					resource.TestCheckResourceAttr(resourceName, "apply_to", "host_only"),
					resource.TestCheckResourceAttr(resourceName, "service_addresses.#", "1"),
					resource.TestCheckResourceAttrSet(resourceName, "service_addresses.0.network.0.href"),
				),
			},
		},
	})
}

func setupNetworkID(t *testing.T) string {
	conf := TestAccProvider.Meta().(Config)
	illumioClient := conf.IllumioClient

	endpoint := fmt.Sprintf("/orgs/%d/networks", illumioClient.OrgID)
	_, data, err := illumioClient.Get(endpoint, &map[string]string{
		// make sure we aren't getting a link_local network as
		// they are not compatible with Virtual Services
		"data_center": "Corporate",
		"max_results": "1",
	})
	if err != nil {
		t.Fatal("Failed to get network for virtual service StateUpgrader test")
	}

	return data.S("0", "href").Data().(string)
}

func testAccCheckIllumioVSResourceConfig_v1(vsName, networkID string) string {
	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "vs_test" {
	name        = %[1]q
	description = "Terraform Virtual Service test"
	apply_to    = "host_only"

	service_ports {
		proto = 6
		port  = 1234
	}

	service_addresses {
		ip           = "127.0.0.1"
		network_href = %[2]q
	}
}
`, vsName, networkID)
}

func testAccCheckIllumioVSResourceConfig_v2(vsName, networkID string) string {
	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "vs_test" {
	name        = %[1]q
	description = "Terraform Virtual Service test"
	apply_to    = "host_only"

	service_ports {
		proto = 6
		port  = 1234
	}

	service_addresses {
		ip = "127.0.0.1"

		network {
			href = %[2]q
		}
	}
}
`, vsName, networkID)
}
