// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixSBL string = "TF-ACC-SBL"

func TestAccIllumioSBL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_service_bindings.sbl_test"

	virtualServiceName := acctest.RandomWithPrefix(prefixSB)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		CheckDestroy:      testAccCheckIllumioSBDataSource_VSDestroyed,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSBDataSourceConfig_VSSetup(virtualServiceName),
				Check:  testAccCheckIllumioSBDataSource_VSExists("illumio-core_virtual_service.test"),
			},
			{
				Config: testAccCheckIllumioSBLDataSourceConfig_basic(virtualServiceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioSBLDataSourceConfig_basic(virtualServiceName string) string {
	rName1 := acctest.RandomWithPrefix(prefixSBL)
	rName2 := acctest.RandomWithPrefix(prefixSBL)

	var sb strings.Builder

	sb.WriteString(testAccCheckIllumioSBDataSourceConfig_VSSetup(virtualServiceName))
	sb.WriteString(fmt.Sprintf(`
locals {
	virtual_service_active_href = replace(illumio-core_virtual_service.test.href, "draft", "active")
}

resource "illumio-core_unmanaged_workload" "sbl_test1" {
	name               = %[1]q
	description        = "Terraform Service Bindings test"
	hostname           = "example.workload1"
}

resource "illumio-core_unmanaged_workload" "sbl_test2" {
	name               = %[2]q
	description        = "Terraform Service Bindings test"
	hostname           = "example.workload2"
}

resource "illumio-core_service_binding" "sbl_test1" {
	virtual_service {
		href = local.virtual_service_active_href
	}
	workload {
		href = illumio-core_unmanaged_workload.sbl_test1.href
	}
}

resource "illumio-core_service_binding" "sbl_test2" {
	virtual_service {
		href = local.virtual_service_active_href
	}
	workload {
		href = illumio-core_unmanaged_workload.sbl_test2.href
	}
}

data "illumio-core_service_bindings" "sbl_test" {
	virtual_service = local.virtual_service_active_href

	# enforce dependencies
	depends_on = [
		illumio-core_service_binding.sbl_test1,
		illumio-core_service_binding.sbl_test2,
	]
}
`, rName1, rName2))
	return sb.String()
}
