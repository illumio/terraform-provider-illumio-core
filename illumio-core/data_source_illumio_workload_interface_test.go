package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerDSWorkloadInterface *schema.Provider

func TestAccIllumioWorkloadInterface_Read(t *testing.T) {
	workloadInterfaceAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerDSWorkloadInterface),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadInterfaceDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioDataSourceWorkloadInterfaceExists("data.illumio-core_workload_interface.test", workloadInterfaceAttr),
					testAccCheckIllumioWorkloadInterfaceDataSourceAttributes(workloadInterfaceAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadInterfaceDataSourceConfig_basic() string {
	return `
	data "illumio-core_workload_interface" "test" {
		href = "/orgs/1/workloads/d42a430e-b20b-4b2d-853f-2d39fa4cea22/interfaces/acc-test-WI"
	}
	`
}

func testAccCheckIllumioDataSourceWorkloadInterfaceExists(name string, workloadInterfaceAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Workload Interface %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerDSWorkloadInterface).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"link_state",
			"friendly_name",
			"name",
		} {
			workloadInterfaceAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioWorkloadInterfaceDataSourceAttributes(workloadInterfaceAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"link_state":    "up",
			"friendly_name": "acc-test friendly name",
			"name":          "acc-test-WI",
		}
		for k, v := range expectation {
			if workloadInterfaceAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, workloadInterfaceAttr[k], v)
			}
		}

		return nil
	}
}
