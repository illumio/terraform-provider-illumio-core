package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerContainerCluster *schema.Provider

func TestAccIllumioContainerCluster_CreateUpdate(t *testing.T) {
	ccAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerContainerCluster),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerContainerCluster, "illumio-core_container_cluster", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioContainerClusterConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioContainerClusterExists("illumio-core_container_cluster.test", ccAttr),
					testAccCheckIllumioContainerClusterAttributes("creation from terraform", ccAttr),
				),
			},
			{
				Config: testAccCheckIllumioContainerClusterConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioContainerClusterExists("illumio-core_container_cluster.test", ccAttr),
					testAccCheckIllumioContainerClusterAttributes("updation from terraform", ccAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioContainerClusterConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_container_cluster" "test" {
		name = "acc. test Container Cluster"
		description = "%s"
	}
	`, val)
}

func testAccCheckIllumioContainerClusterExists(name string, ccAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("IP List %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerContainerCluster).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
		} {
			ccAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioContainerClusterAttributes(val string, ccAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":        "acc. test Container Cluster",
			"description": val,
		}
		for k, v := range expectation {
			if ccAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ccAttr[k], v)
			}
		}

		return nil
	}
}
