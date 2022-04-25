terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  pce_host     = var.pce_url
  org_id       = var.pce_org_id
  api_username = var.pce_api_key
  api_secret   = var.pce_api_secret
}

resource "illumio-core_unmanaged_workload" "lab_gtm" {
  name             = "lab_gtm"
  hostname         = "gtm.lab.illum.io"
  public_ip        = "172.22.1.14"
  description      = "Lab Global Traffic Manager"
  enforcement_mode = "full"
  online           = true
}

resource "illumio-core_workload_interface" "eth0" {
    workload_href = illumio-core_unmanaged_workload.lab_gtm.href
    name          = "eth0"
    friendly_name = "Wired Network (Ethernet)"
    link_state    = "up"
}

resource "illumio-core_workload_interface" "bridge0" {
    workload_href = illumio-core_unmanaged_workload.lab_gtm.href
    name          = "bridge0"
    friendly_name = "Bridge Network"
    link_state    = "up"
}

data "illumio-core_workload_interfaces" "example" {
  workload_href = illumio-core_unmanaged_workload.lab_gtm.href

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_workload_interface.eth0,
    illumio-core_workload_interface.bridge0,
  ]
}
