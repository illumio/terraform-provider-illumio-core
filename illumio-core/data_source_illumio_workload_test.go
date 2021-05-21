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

var providerDSWorkload *schema.Provider

func TestAccIllumioWorkload_Read(t *testing.T) {
	workloadAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSWorkload),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceWorkloadExists("data.illumio-core_workload.test", workloadAttr),
					testAccCheckIllumioWorkloadDataSourceAttributes(workloadAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadDataSourceConfig_basic() string {
	return `
	data "illumio-core_workload" "test" {
		href = "/orgs/1/workloads/d42a430e-b20b-4b2d-853f-2d39fa4cea22"
	}
	`
}

func testAccCheckIllumioDataSourceWorkloadExists(name string, workloadAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Workload %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSWorkload).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"enforcement_mode",
			"visibility_level",
			"interfaces.1.name",
			"name",
		} {
			workloadAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioWorkloadDataSourceAttributes(workloadAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"enforcement_mode":  "visibility_only",
			"visibility_level":  "flow_summary",
			"interfaces.1.name": "acc-test-WI",
			"name":              "acc-test-Workload",
		}
		for k, v := range expectation {
			if workloadAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, workloadAttr[k], v)
			}
		}

		return nil
	}
}
