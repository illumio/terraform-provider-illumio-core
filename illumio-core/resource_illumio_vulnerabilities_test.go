package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerVul *schema.Provider

func TestAccIllumioVulnerabilities_CreateUpdate(t *testing.T) {
	serAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerVul),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerVul, "illumio-core_vulnerabilities", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVulnerabilitiesConfig_basic("creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioVulnerabilitiesExists("illumio-core_vulnerabilities.test", serAttr),
					testAccCheckIllumioVulnerabilitiesAttributes("creation from terraform", serAttr),
				),
			},
			{
				Config: testAccCheckIllumioVulnerabilitiesConfig_basic("updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioVulnerabilitiesExists("illumio-core_vulnerabilities.test", serAttr),
					testAccCheckIllumioVulnerabilitiesAttributes("updation from terraform", serAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioVulnerabilitiesConfig_basic(val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_vulnerabilities" "test" {
		vulnerability {
		  reference_id = "go-test-id"
		  name         = "test"
		  score        = 2
		  cve_ids      = ["someid"]
		  description  = "%s"
		}
	  }
	`, val)
}

func testAccCheckIllumioVulnerabilitiesExists(name string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Vulnerabilities %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerVul).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		vulId := "/orgs/1/vulnerabilities/go-test-id"

		_, cont, err := illumioClient.Get(vulId, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"score",
			"description",
		} {
			serAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		return nil
	}
}

func testAccCheckIllumioVulnerabilitiesAttributes(val string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":        "test",
			"score":       float64(2),
			"description": val,
		}

		for k, v := range expectation {
			if serAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, serAttr[k], v)
			}
		}

		return nil
	}
}
