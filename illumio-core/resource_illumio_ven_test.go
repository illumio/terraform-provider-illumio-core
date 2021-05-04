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
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerVEN, "illumio-core_ven", false),
		Steps: []resource.TestStep{
			{
				// Config:            `resource "illumio-core_ven" "test" { status = "active" }`,
				ResourceName:  "illumio-core_ven.test",
				ImportStateId: "/orgs/1/vens/e6eec907-85c0-4ca7-8607-1b35c27501d7",
				ImportState:   true,
				// ImportStateVerify: true,
			},
			// {
			// 	Check: resource.ComposeTestCheckFunc(
			// 		testAccCheckIllumioVENExists("illumio-core_ven.test", srAttr),
			// 		testAccCheckIllumioVENAttributes("creation from terraform", srAttr),
			// 	),
			// },
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
		status = "suspended"
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
			"status":      "suspended",
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
