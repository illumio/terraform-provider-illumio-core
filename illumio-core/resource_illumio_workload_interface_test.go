package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerWorkloadInterface *schema.Provider

func TestAccIllumioWorkloadInterface_CreateUpdate(t *testing.T) {
	ipAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerWorkloadInterface),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerWorkloadInterface, "illumio_workload_interface", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadInterfaceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWorkloadInterfaceExists("illumio_workload_interface.test", ipAttr),
					testAccCheckIllumioWorkloadInterfaceAttributes(ipAttr),
				),
			},
			{
				Config: testAccCheckIllumioWorkloadInterfaceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioWorkloadInterfaceExists("illumio_workload_interface.test", ipAttr),
					testAccCheckIllumioWorkloadInterfaceAttributes(ipAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioWorkloadInterfaceConfig_basic() string {
	return `
	resource "illumio_workload_interface" "test" {
		workload_id = "d42a430e-b20b-4b2d-853f-2d39fa4cea22"
		name = "acc. test Workload Interface"
		link_state = "up"
		friendly_name = "test friendly name"
	  }
	`
}

func testAccCheckIllumioWorkloadInterfaceExists(name string, ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("IP List %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerWorkloadInterface).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"link_state",
			"friendly_name",
		} {
			ipAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioWorkloadInterfaceAttributes(ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":          "acc. test Workload Interface",
			"link_state":    "up",
			"friendly_name": "test friendly name",
		}
		for k, v := range expectation {
			if ipAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ipAttr[k], v)
			}
		}

		return nil
	}
}
