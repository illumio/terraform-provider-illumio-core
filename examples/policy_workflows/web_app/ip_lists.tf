# Define IP lists

resource "illumio-core_ip_list" "lab_internal" {
  name        = "IPL-LAB"
  description = "Lab VPC IPs"

  ip_ranges {
    from_ip = "172.22.0.0/19"
    description = "Lab IP subnet"
  }

  fqdns {
    fqdn = "*.lab.illum.io"
    description = "Lab domains"
  }
}

data "illumio-core_ip_lists" "default" {
  # all PCE instances define a default global IP list
  name = "Any (0.0.0.0/0 and ::/0)"
  max_results = 1
}

locals {
  any_ip_list_href = one(data.illumio-core_ip_lists.default.items[*].href)
}
