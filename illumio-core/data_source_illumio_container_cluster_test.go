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

var providerDSCC *schema.Provider

func TestAccIllumioCC_Read(t *testing.T) {
	ccAttr := map[string]interface{}{}
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSCC),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceCCExists("data.illumio-core_container_cluster.test", ccAttr),
					testAccCheckIllumioCCDataSourceAttributes(ccAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioCCDataSourceConfig_basic() string {
	return `
	data "illumio-core_container_cluster" "test" {
		href = "/orgs/1/container_clusters/bd37cbdd-82bd-4f49-b52f-9405ba236a43"
	}
	`
}

func testAccCheckIllumioDataSourceCCExists(name string, ccAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Container Cluster %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSCC).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
			"manager_type",
			"online",
		} {
			ccAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioCCDataSourceAttributes(ccAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":         "Acc. test name",
			"description":  "Acc. test description",
			"manager_type": nil,
			"online":       false,
		}
		for k, v := range expectation {
			if ccAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ccAttr[k], v)
			}
		}

		return nil
	}
}
