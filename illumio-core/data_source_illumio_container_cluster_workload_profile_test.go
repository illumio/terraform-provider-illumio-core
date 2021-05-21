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

var providerDSCCWorkloadProfile *schema.Provider

func TestAccIllumioCCWP_Read(t *testing.T) {
	ccsbAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSCCWorkloadProfile),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCWPDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceCCWPxists("data.illumio-core_container_cluster_workload_profile.test", ccsbAttr),
					testAccCheckIllumioCCWPDataSourceAttributes(ccsbAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioCCWPDataSourceConfig_basic() string {
	return `
	data "illumio-core_container_cluster_workload_profile" "test" {
		href = "/orgs/1/container_clusters/bd37cbdd-82bd-4f49-b52f-9405ba236a43/container_workload_profiles/598888c7-a625-4507-a5c8-14f4a3c4c1d6"
	}
	`
}

func testAccCheckIllumioDataSourceCCWPxists(name string, ccsbAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Container Cluster Workload Profile %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSCCWorkloadProfile).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"enforcement_mode",
			"managed",
			"assign_labels.0.href",
			"labels.0.key",
			"labels.0.assignment.href",
			"labels.0.assignment.value",
		} {
			ccsbAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioCCWPDataSourceAttributes(ccsbAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":                      "Acc. test name",
			"enforcement_mode":          "idle",
			"managed":                   true,
			"assign_labels.0.href":      "/orgs/1/labels/1",
			"labels.0.key":              "role",
			"labels.0.assignment.href":  "/orgs/1/labels/1",
			"labels.0.assignment.value": "Web",
		}
		for k, v := range expectation {
			if ccsbAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ccsbAttr[k], v)
			}
		}

		return nil
	}
}
