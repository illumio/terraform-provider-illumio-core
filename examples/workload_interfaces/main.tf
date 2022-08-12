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

  interfaces {
    name          = "eth0"
    address       = "172.22.1.14"
    friendly_name = "Wired Network (Ethernet)"
    link_state    = "up"
  }

  interfaces {
    name          = "lo0"
    address       = "127.0.0.1"
    friendly_name = "Loopback Interface"
    link_state    = "up"
  }
}

data "illumio-core_workload_interfaces" "example" {
  workload_href = illumio-core_unmanaged_workload.lab_gtm.href
}
