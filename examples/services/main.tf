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

resource "illumio-core_service" "win_rdp" {
  name = "S-WIN-RDP"

  windows_services {
    service_name = "TermService"
    process_name = "svchost.exe"
    # Illumio uses the IANA protocol numbers to identify the service proto
    proto        = "6"  # TCP
    port         = "3389"
  }
}

resource "illumio-core_service" "win_kerb" {
  name = "S-WIN-KERB"

  windows_services {
    service_name = "kerberos"
    proto        = "6"
    port         = "88"
  }

  windows_services {
    service_name = "kerberos"
    proto        = "17"  # UDP
    port         = "88"
  }
}

data "illumio-core_services" "windows_services" {
	# supports partial match lookups
  name = "S-WIN"

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_service.win_rdp,
    illumio-core_service.win_kerb,
  ]
}
