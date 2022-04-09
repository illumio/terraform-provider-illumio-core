// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixSBL string = "TF-ACC-SBL"

func TestAccIllumioSBL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_service_bindings.sbl_test"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSBLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioSBLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixSBL)
	rName2 := acctest.RandomWithPrefix(prefixSBL)
	rName3 := acctest.RandomWithPrefix(prefixSBL)

	return fmt.Sprintf(`
resource "illumio-core_virtual_service" "sbl_test" {
	name = %[1]q
	description = "Terraform Service Bindings test"
	apply_to = "host_only"

	service_ports {
		proto = 6
		port = 3389
	}

	service_ports {
		proto = 17
		port = 3389
	}
}

resource "illumio-core_workload" "sbl_test1" {
	name               = %[2]q
	description        = "Terraform Service Bindings test"
	hostname           = "example.workload1"
}

resource "illumio-core_workload" "sbl_test2" {
	name               = %[3]q
	description        = "Terraform Service Bindings test"
	hostname           = "example.workload2"
}

resource "illumio-core_service_binding" "sbl_test1" {
	virtual_service {
		href = illumio-core_virtual_service.sbl_test.href
	}
	workload {
		href = illumio-core_workload.sbl_test1.href
	}
}

resource "illumio-core_service_binding" "sbl_test2" {
	virtual_service {
		href = illumio-core_virtual_service.sbl_test.href
	}
	workload {
		href = illumio-core_workload.sbl_test2.href
	}
}

data "illumio-core_service_bindings" "sbl_test" {
	virtual_service = illumio-core_virtual_service.sbl_test.href

	# enforce dependencies
	depends_on = [
		illumio-core_service_binding.sbl_test1,
		illumio-core_service_binding.sbl_test2,
	]
}
`, rName1, rName2, rName3)
}
