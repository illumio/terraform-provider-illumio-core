// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixVSL string = "TF-ACC-VSL"

func TestAccIllumioVSL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_virtual_services.vsl_test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioVSLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioVSLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixVSL)
	rName2 := acctest.RandomWithPrefix(prefixVSL)

	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "vsl_test1" {
	name = %[1]q
	description = "Terraform Virtual Services test"
	apply_to = "host_only"

	service_ports {
		proto = 6
		port = 8080
	}
}

resource "illumio-core_virtual_service" "vsl_test2" {
	name = %[2]q
	description = "Terraform Virtual Services test"
	apply_to = "host_only"

	service_ports {
		proto = 6
		port = 8443
	}
}

data "illumio-core_virtual_services" "vsl_test" {
	# lookup based on partial match
	name = %[3]q

	# enforce dependencies
	depends_on = [
		illumio-core_virtual_service.vsl_test1,
		illumio-core_virtual_service.vsl_test2,
	]
}
`, rName1, rName2, prefixVSL)
}
