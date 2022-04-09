// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixWLL string = "TF-ACC-WLL"

func TestAccIllumioWLL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_workloads.wll"

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWLLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioWLLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixWLL)
	rName2 := acctest.RandomWithPrefix(prefixWLL)

	return fmt.Sprintf(`
resource "illumio-core_workload" "wll1" {
	name               = %[1]q
	description        = "Terraform Workloads test 1"
	hostname           = "jumpbox1"
}

resource "illumio-core_workload" "wll2" {
	name               = %[1]q
	description        = "Terraform Workloads test 2"
	hostname           = "jumpbox2"
}

data "illumio-core_workloads" "wll" {
	# enforce dependencies
	depends_on = [
		illumio-core_workload.wll1,
		illumio-core_workload.wll2,
	]
}
`, rName1, rName2)
}
