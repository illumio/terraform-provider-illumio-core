# use the ip_lists data source to search against the /ip_lists endpoint by name
data "illumio-core_ip_lists" "default" {
  # all PCE instances define a special default IP list covering all addresses
  name = "Any (0.0.0.0/0 and ::/0)"
  max_results = 1
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

resource "illumio-core_enforcement_boundary" "block_rdp" {
  name        = "EB-RDP"

  ingress_services {
    href = illumio-core_service.rdp.href
  }

  consumers {
    ip_list {
      href = one(data.illumio-core_ip_lists.default.items[*].href)
    }
  }

  providers {
    actors = "ams"  # special notation meaning "all managed systems" - affects all workloads
  }
}

data "illumio-core_enforcement_boundary" "block_rdp" {
  href = illumio-core_enforcement_boundary.block_rdp.href
}
