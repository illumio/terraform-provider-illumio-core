resource "illumio-core_ip_list" "example" {
  name        = "IPL-LOCAL"
  description = "Local addresses"

  ip_ranges {
    # from_ip can be a CIDR range or individual IP
    from_ip = "192.168.0.0/16"
    description = "Internal network range"
  }

  ip_ranges {
    from_ip = "127.0.0.1"
    description = "Loopback address"
  }

  fqdns {
    fqdn = "*.localdomain"
    description = "Default localdomain VBox hostnames"
  }
}
