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

resource "illumio-core_ip_list" "local" {
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

data "illumio-core_ip_list" "local" {
  href = illumio-core_ip_list.local.href
}
