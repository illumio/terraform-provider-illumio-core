// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixCC string = "TF-ACC-CC"

func TestAccIllumioCC_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_container_cluster.cc_test"
	resourceName := "illumio-core_container_cluster.cc_test"

	ccName := acctest.RandomWithPrefix(prefixCC)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioCCDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "manager_type", resourceName, "manager_type"),
					resource.TestCheckResourceAttrPair(dataSourceName, "online", resourceName, "online"),
					resource.TestCheckResourceAttrPair(dataSourceName, "container_cluster_id", resourceName, "container_cluster_id"),
					resource.TestCheckResourceAttrSet(resourceName, "container_cluster_token"),
				),
			},
			{
				Config: testAccCheckIllumioCCResource_updateNameAndDesc(ccName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", ccName),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"container_cluster_token"},
			},
		},
	})
}

func testAccCheckIllumioCCDataSourceConfig_basic() string {
	rName := acctest.RandomWithPrefix(prefixCC)

	return fmt.Sprintf(`
resource "illumio-core_container_cluster" "cc_test" {
	name = %[1]q
	description = "Terraform Container Cluster test"
}

data "illumio-core_container_cluster" "cc_test" {
	href = illumio-core_container_cluster.cc_test.href
}
`, rName)
}

func testAccCheckIllumioCCResource_updateNameAndDesc(ccName string) string {
	return fmt.Sprintf(`
resource "illumio-core_container_cluster" "cc_test" {
	name = %[1]q
	description = ""
}
`, ccName)
}
