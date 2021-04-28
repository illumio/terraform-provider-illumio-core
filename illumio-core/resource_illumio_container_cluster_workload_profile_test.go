package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerContainerClusterWorkloadProfile *schema.Provider

func TestAccIllumioContainerClusterWorkloadProfileWorkloadProfile_CreateUpdate(t *testing.T) {
	ccAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerContainerClusterWorkloadProfile),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerContainerClusterWorkloadProfile, "illumio_container_cluster_workload_profile", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioContainerClusterWorkloadProfileConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioContainerClusterWorkloadProfileExists("illumio_container_cluster_workload_profile.test", ccAttr),
					testAccCheckIllumioContainerClusterWorkloadProfileAttributes("creation from terraform", ccAttr),
				),
			},
			{
				Config: testAccCheckIllumioContainerClusterWorkloadProfileConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioContainerClusterWorkloadProfileExists("illumio_container_cluster_workload_profile.test", ccAttr),
					testAccCheckIllumioContainerClusterWorkloadProfileAttributes("updation from terraform", ccAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioContainerClusterWorkloadProfileConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio_container_cluster_workload_profile" "test" {
		container_cluster_id = "deb48c70-e9d2-4101-ab7e-1f48de922ff4"
		name = "acc. test Container Cluster Workload Profile"
		description = "%s"
		managed = true
		enforcement_mode = "visibility_only"
	}
	`, val)
}

func testAccCheckIllumioContainerClusterWorkloadProfileExists(name string, ccAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("IP List %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerContainerClusterWorkloadProfile).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"managed",
			"description",
			"enforcement_mode",
		} {
			ccAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioContainerClusterWorkloadProfileAttributes(val string, ccAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":             "acc. test Container Cluster Workload Profile",
			"description":      val,
			"managed":          true,
			"enforcement_mode": "visibility_only",
		}
		for k, v := range expectation {
			if ccAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ccAttr[k], v)
			}
		}

		return nil
	}
}
