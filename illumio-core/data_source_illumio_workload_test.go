// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var prefixWorkload string = "TF-ACC-WL"

func TestAccIllumioWorkload_Read(t *testing.T) {
	dataSourceName := "data.illumio-core_workload.wl_test"
	resourceName := "illumio-core_unmanaged_workload.wl_test"

	wkldName := acctest.RandomWithPrefix(prefixWorkload)
	labelName := acctest.RandomWithPrefix(prefixWorkload)
	updatedName := acctest.RandomWithPrefix(prefixWorkload)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadDataSourceConfig_basic(wkldName, labelName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "href", resourceName, "href"),
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "description", resourceName, "description"),
					resource.TestCheckResourceAttrPair(dataSourceName, "hostname", resourceName, "hostname"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_principal_name", resourceName, "service_principal_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "enforcement_mode", resourceName, "enforcement_mode"),
					resource.TestCheckResourceAttrPair(dataSourceName, "visibility_level", resourceName, "visibility_level"),
					resource.TestCheckResourceAttrPair(dataSourceName, "distinguished_name", resourceName, "distinguished_name"),
					resource.TestCheckResourceAttrPair(dataSourceName, "public_ip", resourceName, "public_ip"),
					resource.TestCheckResourceAttrPair(dataSourceName, "data_center", resourceName, "data_center"),
					resource.TestCheckResourceAttrPair(dataSourceName, "data_center_zone", resourceName, "data_center_zone"),
					resource.TestCheckResourceAttrPair(dataSourceName, "os_id", resourceName, "os_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "os_detail", resourceName, "os_detail"),
					resource.TestCheckResourceAttrPair(dataSourceName, "service_provider", resourceName, "service_provider"),
					resource.TestCheckResourceAttrPair(dataSourceName, "online", resourceName, "online"),
					resource.TestCheckResourceAttrPair(dataSourceName, "labels", resourceName, "labels"),
					resource.TestCheckResourceAttrPair(dataSourceName, "interfaces", resourceName, "interfaces"),
				),
			},
			{
				Config: testAccCheckIllumioWorkloadResource_updateRemoveLabels(wkldName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "labels.#", "0"),
				),
			},
			{
				Config: testAccCheckIllumioWorkloadResource_updateRemoveInterfaces(wkldName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "interfaces.#", "0"),
				),
			},
			{
				Config: testAccCheckIllumioWorkloadResource_updateNameAndEmptyStringFields(updatedName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", updatedName),
					resource.TestCheckResourceAttr(resourceName, "description", ""),
					resource.TestCheckResourceAttr(resourceName, "hostname", ""),
					resource.TestCheckResourceAttr(resourceName, "distinguished_name", ""),
					resource.TestCheckResourceAttr(resourceName, "service_provider", ""),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIllumioWorkload_Delete(t *testing.T) {
	workloadHref := new(string)
	newWorkloadHref := new(string)

	resourceName := "illumio-core_unmanaged_workload.wl_test"
	workloadName := acctest.RandomWithPrefix(prefixLabel)
	labelName := acctest.RandomWithPrefix(prefixWorkload)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIllumioWorkloadResourceConfig_basic(workloadName, labelName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(resourceName, workloadHref),
					resource.TestCheckResourceAttr(resourceName, "hostname", "example.workload"),
					resource.TestCheckResourceAttr(resourceName, "name", workloadName),
				),
			},
			{
				// check that an apply called after a workload has been deleted
				// correctly destroys and recreates the resource
				PreConfig: deleteFromPCE(workloadHref, t),
				Config:    testAccCheckIllumioWorkloadResourceConfig_basic(workloadName, labelName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(resourceName, newWorkloadHref),
					testAccCheckCompareRefs(workloadHref, newWorkloadHref, false),
					resource.TestCheckResourceAttr(resourceName, "hostname", "example.workload"),
					resource.TestCheckResourceAttr(resourceName, "name", workloadName),
				),
			},
			{
				// check that a destroy called after a workload has been deleted
				// doesn't throw an error
				PreConfig: deleteFromPCE(workloadHref, t),
				Destroy:   true,
				Config:    testAccCheckIllumioWorkloadResourceConfig_basic(workloadName, labelName),
			},
		},
	})
}

func workloadRoleLabel(labelName string) string {
	return fmt.Sprintf(`
resource "illumio-core_label" "wl_test" {
	key   = "role"
	value = %[1]q
}
`, labelName)
}

func testAccCheckIllumioWorkloadResourceConfig_basic(wkldName, labelName string) string {
	return workloadRoleLabel(labelName) + fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" "wl_test" {
	name               = %[1]q
	description        = "Terraform Workload test"
	hostname           = "example.workload"
	distinguished_name = "cn=Terraform Lab,ou=integrations,o=Illumio,c=CA"
	service_provider   = "SPN"

	interfaces {
		name       = "lo0"
		link_state = "unknown"
		address    = "127.0.0.1"
	}

	interfaces {
		name       = "lo0"
		link_state = "unknown"
		address    = "::ffff:127.0.0.1"
	}

	labels {
		href = illumio-core_label.wl_test.href
	}

	lifecycle {
		create_before_destroy = true
	}
}
`, wkldName)
}

func testAccCheckIllumioWorkloadDataSourceConfig_basic(wkldName, labelName string) string {
	return testAccCheckIllumioWorkloadResourceConfig_basic(wkldName, labelName) + `
data "illumio-core_workload" "wl_test" {
	href = illumio-core_unmanaged_workload.wl_test.href
}`
}

func testAccCheckIllumioWorkloadResource_updateRemoveLabels(wkldName string) string {
	return fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" "wl_test" {
	name               = %[1]q
	description        = "Terraform Workload test"
	hostname           = "example.workload"
	distinguished_name = "cn=Terraform Lab,ou=integrations,o=Illumio,c=CA"
	service_provider   = "SPN"

	interfaces {
		name       = "lo0"
		link_state = "unknown"
		address    = "127.0.0.1"
	}

	interfaces {
		name       = "lo0"
		link_state = "unknown"
		address    = "::ffff:127.0.0.1"
	}
}
`, wkldName)
}

func testAccCheckIllumioWorkloadResource_updateRemoveInterfaces(wkldName string) string {
	return fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" "wl_test" {
	name               = %[1]q
	description        = "Terraform Workload test"
	hostname           = "example.workload"
	distinguished_name = "cn=Terraform Lab,ou=integrations,o=Illumio,c=CA"
	service_provider   = "SPN"
}
`, wkldName)
}

func testAccCheckIllumioWorkloadResource_updateNameAndEmptyStringFields(updatedName string) string {
	return fmt.Sprintf(`
resource "illumio-core_unmanaged_workload" "wl_test" {
	name               = %[1]q
	description        = ""
	hostname           = ""
	distinguished_name = ""
	service_provider   = ""
}
`, updatedName)
}
