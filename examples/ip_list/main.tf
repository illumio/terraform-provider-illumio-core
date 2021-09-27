terraform {
  required_providers {
    illumio-core = {
      version = "0.1.0"
      source  = "illumio/illumio-core"
    }
  }
}


provider "illumio-core" {
  //  pce_host              = "https://pce.my-company.com:8443"
  //  api_username          = "api_xxxxxx"
  //  api_secret            = "big-secret"
  request_timeout = 30
  org_id          = 1
}

resource "illumio-core_ip_list" "example" {
  name        = "example name"
  description = "example desc"
  ip_ranges {
    from_ip = "1.1.0.0/24"
    // from_ip = "1.1.0.0"
    // to_ip = "1.1.0.254"
    description = "example ip_ranges description"
    exclusion = false
  }
  fqdns {
    fqdn = "app.example.com"
    description = "example fqdn description"
  }
}

data "illumio-core_ip_list" "example" {
  href = "/orgs/1/sec_policy/draft/ip_lists/1"
}
