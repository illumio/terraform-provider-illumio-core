# use the ip_lists data source to search against the /ip_lists endpoint by name
data "illumio-core_ip_lists" "default" {
  # all PCE instances define a special default IP list covering all addresses
  name = "Any (0.0.0.0/0 and ::/0)"
  max_results = 1
}

locals {
  default_ip_list_href = one(data.illumio-core_ip_lists.default.items[*].href)
}

resource "illumio-core_service" "smb" {
  name        = "S-SMB"
  description = "UDP Ports used for Server Message Block communication."

  service_ports {
    # Illumio uses the IANA protocol numbers to identify the service proto
    proto = 6  # UDP
    port  = 445
  }
}

resource "illumio-core_service" "netbios" {
  name        = "S-NETBIOS"
  description = "UDP Ports used for NetBIOS."

  service_ports {
    proto   = "17"  # UDP
    port    = "137"
    to_port = "138"
  }

  service_ports {
    proto = "6"  # TCP
    port  = "139"
  }
}

resource "illumio-core_enforcement_boundary" "block_smb" {
  name        = "EB-WIN-SMB"

  ingress_services {
    href = illumio-core_service.smb.href
  }

  consumers {
    ip_list {
      href = local.default_ip_list_href
    }
  }

  providers {
    actors = "ams"  # special notation meaning "all managed systems" - affects all workloads
  }
}

resource "illumio-core_enforcement_boundary" "block_netbios" {
  name        = "EB-WIN-NETBIOS"

  ingress_services {
    href = illumio-core_service.netbios.href
  }

  consumers {
    ip_list {
      href = local.default_ip_list_href
    }
  }

  providers {
    actors = "ams"
  }
}

data "illumio-core_enforcement_boundaries" "block_windows_services" {
	# supports partial match lookups
  name = "EB-WIN-"

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_enforcement_boundary.block_smb,
    illumio-core_enforcement_boundary.block_netbios,
  ]
}
