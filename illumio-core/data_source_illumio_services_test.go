// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixSL string = "TF-ACC-SL"

func TestAccIllumioSL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_services.sl_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioSLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixSL)
	rName2 := acctest.RandomWithPrefix(prefixSL)

	return fmt.Sprintf(`
resource "illumio-core_service" "sl_test1" {
	name          = %[1]q
	description   = "Terraform Services test"

	service_ports {
		proto = 6
		port = 137
	}

	service_ports {
		proto = 6
		port = 138
	}
}

resource "illumio-core_service" "sl_test2" {
	name          = %[2]q
	description   = "Terraform Services test"

	service_ports {
		proto = 17
		port = 137
	}

	service_ports {
		proto = 17
		port = 138
	}
}

data "illumio-core_services" "sl_test" {
	# lookup based on partial match
	name = %[3]q

	# enforce dependencies
	depends_on = [
		illumio-core_service.sl_test1,
		illumio-core_service.sl_test2,
	]
}
`, rName1, rName2, prefixSL)
}
