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

var providerVul *schema.Provider

func TestAccIllumioVulnerabilities_CreateUpdate(t *testing.T) {
	serAttr := map[string]interface{}{}

	vulID := "test-ref-id"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerVul),
		// CheckDestroy is ignored as illumio-core_vulnerabilities does not support delete operation
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVulnerabilitiesConfig_basic(vulID, "creation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioVulnerabilitiesExists("illumio-core_vulnerabilities.test", vulID, serAttr),
					testAccCheckIllumioVulnerabilitiesAttributes("creation from terraform", serAttr),
				),
			},
			{
				Config: testAccCheckIllumioVulnerabilitiesConfig_basic(vulID, "updation from terraform"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioVulnerabilitiesExists("illumio-core_vulnerabilities.test", vulID, serAttr),
					testAccCheckIllumioVulnerabilitiesAttributes("updation from terraform", serAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioVulnerabilitiesConfig_basic(vulID, val string) string {
	return fmt.Sprintf(`
	resource "illumio-core_vulnerabilities" "test" {
		vulnerability {
		  reference_id = "%v"
		  name         = "test"
		  score        = 2
		  cve_ids      = ["someid"]
		  description  = "%s"
		}
	}
	`, vulID, val)
}

func testAccCheckIllumioVulnerabilitiesExists(name, vulID string, serAttr map[string]interface{}) resource.TestCheckFunc {
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

		// Fetching single vulnerability <vulID> for testing
		_, cont, err := illumioClient.Get(fmt.Sprintf("%v/%v", rs.Primary.ID, vulID), nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"name",
			"description",
		} {
			serAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}

		serAttr["score"] = int(cont.S("score").Data().(float64))

		return nil
	}
}

func testAccCheckIllumioVulnerabilitiesAttributes(val string, serAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"name":        "test",
			"score":       2,
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
