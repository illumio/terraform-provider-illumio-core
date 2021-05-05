terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
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

resource "illumio-core_ip_list" "iplist_1" {
  name        = "test iplist"
  description = "desc test"
  ip_ranges {
    from_ip = "1.1.0.0/24"
    to_ip = "0.0.0.0/0"
    description = "test ip_ranges description"
    exclusion = false
  }
  fqdns {
    fqdn = "app.example.com"
    description = "test fqdn description"
  }
}

data "illumio-core_ip_list" "iplist_1" {
  href = "/orgs/1/sec_policy/draft/ip_lists/1"
}