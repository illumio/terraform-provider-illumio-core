// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceIllumioSR_Read(t *testing.T) {
	resourceName := "illumio-core_security_rule.sr_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				SkipFunc: skipIfPCEVersionBelow("22.3.0"),
				Config:   testAccCheckIllumioSRResourceConfig_useWorkloadSubnets(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "use_workload_subnets.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioSRResourceConfig_useWorkloadSubnets() string {
	return testAccCheckIllumioSRRuleSet() + `
resource "illumio-core_security_rule" "sr_test" {
	rule_set_href = illumio-core_rule_set.sr_test.href
	enabled = true
	description = "Terraform Security Rule test"

	resolve_labels_as {
		consumers = ["workloads"]
		providers = ["workloads"]
	}

	consumers {
		actors = "ams"
	}

	providers {
		actors = "ams"
	}

	ingress_services {
		proto = 6
		port  = 1234
	}

	use_workload_subnets = ["providers", "consumers"]
}`
}
