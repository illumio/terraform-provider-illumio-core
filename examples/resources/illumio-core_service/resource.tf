resource "illumio-core_service" "example" {
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
