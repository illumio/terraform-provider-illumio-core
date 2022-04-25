// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixSRL string = "TF-ACC-SRL"

func TestAccIllumioSRL_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_security_rules.srl_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSRLDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "items.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIllumioSRLDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixSRL)
	rName2 := acctest.RandomWithPrefix(prefixSRL)

	return fmt.Sprintf(`
resource "illumio-core_label" "srl_test" {
	key   = "env"
	value = %[1]q
}

resource "illumio-core_rule_set" "srl_test" {
	name = %[2]q
	description = "Terraform Security Rules test"

	scopes {
		label {
			href = illumio-core_label.srl_test.href
		}
	}
}

resource "illumio-core_security_rule" "srl_test1" {
	rule_set_href = illumio-core_rule_set.srl_test.href
	enabled = true
	description = "Terraform Security Rules test"

	resolve_labels_as {
		consumers = ["workloads"]
		providers = ["workloads"]
	}

	consumers {
		actors = "ams"
	}

	providers {
		label {
			href = illumio-core_label.srl_test.href
		}
	}

	ingress_services {
		proto = 6
		port  = 80
	}
}

resource "illumio-core_security_rule" "srl_test2" {
	rule_set_href = illumio-core_rule_set.srl_test.href
	enabled = true
	description = "Terraform Security Rules test"

	resolve_labels_as {
		consumers = ["workloads"]
		providers = ["workloads"]
	}

	consumers {
		actors = "ams"
	}

	providers {
		label {
			href = illumio-core_label.srl_test.href
		}
	}

	ingress_services {
		proto = 6
		port  = 443
	}
}

data "illumio-core_security_rules" "srl_test" {
	rule_set_href = illumio-core_rule_set.srl_test.href

	# enforce dependencies
	depends_on = [
		illumio-core_security_rule.srl_test1,
		illumio-core_security_rule.srl_test2,
	]
}
`, rName1, rName2)
}
