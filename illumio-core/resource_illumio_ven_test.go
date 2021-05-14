package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerVEN *schema.Provider

func TestAccIllumioVEN_CreateUpdate(t *testing.T) {
	srAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerVEN),
		// CheckDestroy is ignored as illumio-core_ven does not support delete operation
		Steps: []resource.TestStep{
			{
				Config:             testAccCheckIllumioVENConfig_basic("creation from terraform"),
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:            "illumio-core_ven.test",
				ImportStateId:           "/orgs/1/vens/61ea9747-8f09-439a-9541-726733daa758",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"test_href"},
			},
			{
				Config: testAccCheckIllumioVENConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioVENExists("illumio-core_ven.test", srAttr),
					testAccCheckIllumioVENAttributes("updation from terraform", srAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioVENConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_ven" "test" {
		test_href = "/orgs/1/vens/61ea9747-8f09-439a-9541-726733daa758"
		status = "active"
		description = "%s"
	  }
	`, val)
}

func testAccCheckIllumioVENExists(name string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("VEN %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerVEN).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"status",
			"description",
		} {
			lgAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioVENAttributes(val string, lgAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"status":      "active",
			"description": val,
		}
		for k, v := range expectation {
			if lgAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, lgAttr[k], v)
			}
		}

		return nil
	}
}
