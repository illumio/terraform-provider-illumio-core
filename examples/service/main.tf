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

resource "illumio-core_service" "rdp" {
  name        = "S-RDP"
  description = "TCP and UDP Remote Desktop Protocol ports"

  service_ports {
    # Illumio uses the IANA protocol numbers to identify the service proto
    proto = "6"  # TCP
    port  = "3389"
  }

  service_ports {
    proto = "17"  # UDP
    port  = "3389"
  }
}

data "illumio-core_service" "rdp" {
  href = illumio-core_service.rdp.href
}
