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

data "illumio-core_ven" "example" {
  href = "/orgs/1/vens/80ec4e0a-e628-41c2-b79a-866f72a6b070"
}

resource "illumio-core_ven" "example" {
  status = "suspended"
  name = "example name"
  description = "example description"
  target_pce_fqdn = "example.fqdn"
}
