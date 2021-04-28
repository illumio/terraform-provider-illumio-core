package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var providerServiceBinding *schema.Provider

func TestAccIllumioServiceBinding_CreateUpdate(t *testing.T) {
	ipAttr := map[string]interface{}{}

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactoriesInternal(&providerServiceBinding),
		CheckDestroy:      testAccCheckIllumioGeneralizeDestroy(providerServiceBinding, "illumio_service_binding", false),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioServiceBindingConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioServiceBindingExists("illumio_service_binding.test", ipAttr),
					testAccCheckIllumioServiceBindingAttributes(ipAttr),
				),
			},
			{
				Config: testAccCheckIllumioServiceBindingConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIllumioServiceBindingExists("illumio_service_binding.test", ipAttr),
					testAccCheckIllumioServiceBindingAttributes(ipAttr),
				),
			},
		},
	})
}

func testAccCheckIllumioServiceBindingConfig_basic() string {
	return `
	resource "illumio_service_binding" "test" {
		virtual_service {
		  href = "/orgs/1/sec_policy/active/virtual_services/69f1fcc7-94f0-4e42-b9a8-e722038e6dda"
		}
		workload {
		  href = "/orgs/1/workloads/673c3148-a419-4ed2-b0e2-30eb538695e7"
		}
	  }			
	`
}

func testAccCheckIllumioServiceBindingExists(name string, ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Service Binding %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("ID was not set")
		}

		pConfig := (*providerServiceBinding).Meta().(Config)
		illumioClient := pConfig.IllumioClient

		_, cont, err := illumioClient.Get(rs.Primary.ID, nil)
		if err != nil {
			return err
		}

		for _, k := range []string{
			"workload.href",
			"virtual_service.href",
		} {
			ipAttr[k] = cont.S(strings.Split(k, ".")...).Data()
		}
		return nil
	}
}

func testAccCheckIllumioServiceBindingAttributes(ipAttr map[string]interface{}) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		expectation := map[string]interface{}{
			"virtual_service.href": "/orgs/1/sec_policy/active/virtual_services/69f1fcc7-94f0-4e42-b9a8-e722038e6dda",
			"workload.href":        "/orgs/1/workloads/673c3148-a419-4ed2-b0e2-30eb538695e7",
		}
		for k, v := range expectation {
			if ipAttr[k] != v {
				return fmt.Errorf("Bad %s, Actual: %v, Expected: %v", k, ipAttr[k], v)
			}
		}

		return nil
	}
}
