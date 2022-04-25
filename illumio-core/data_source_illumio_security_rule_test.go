// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixSR string = "TF-ACC-SR"

func TestAccIllumioSR_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_security_rule.sr_test"
	resourceName := "illumio-core_security_rule.sr_test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioSRDataSourceConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enabled", resourceName, "enabled"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "resolve_labels_as", resourceName, "resolve_labels_as"),
					resource.TestCheckResourceAttrPair(dataSourceName, "ingress_services", resourceName, "ingress_services"),
					resource.TestCheckResourceAttrPair(dataSourceName, "consumers", resourceName, "consumers"),
					resource.TestCheckResourceAttrPair(dataSourceName, "providers", resourceName, "providers"),
					resource.TestCheckResourceAttrPair(dataSourceName, "sec_connect", resourceName, "sec_connect"),
					resource.TestCheckResourceAttrPair(dataSourceName, "stateless", resourceName, "stateless"),
					resource.TestCheckResourceAttrPair(dataSourceName, "machine_auth", resourceName, "machine_auth"),
					resource.TestCheckResourceAttrPair(dataSourceName, "unscoped_consumers", resourceName, "unscoped_consumers"),
				),
			},
		},
	})
}

func testAccCheckIllumioSRDataSourceConfig_basic() string {
	rName1 := acctest.RandomWithPrefix(prefixSR)
	rName2 := acctest.RandomWithPrefix(prefixSR)

	return fmt.Sprintf(`
resource "illumio-core_label" "sr_test" {
	key   = "loc"
	value = %[1]q
}

resource "illumio-core_rule_set" "sr_test" {
	name = %[2]q
	description = "Terraform Security Rule test"

	ip_tables_rules {
		actors {
			actors = "ams"
		}

		actors {
			label {
				href = illumio-core_label.sr_test.href
			}
		}

		enabled = false

		ip_version = 6
		statements {
			table_name = "nat"
			chain_name = "PREROUTING"
			parameters = "value"
		}
	}

	scopes {
		label {
			href = illumio-core_label.sr_test.href
		}
	}
}

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
		label {
			href = illumio-core_label.sr_test.href
		}
	}

	ingress_services {
		proto = 6
		port  = 1234
	}
}

data "illumio-core_security_rule" "sr_test" {
	href = illumio-core_security_rule.sr_test.href
}
`, rName1, rName2)
}
