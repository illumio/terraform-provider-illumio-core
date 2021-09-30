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
  request_timeout       = 30
  org_id                = 1
}

data "illumio-core_label" "example" {
  href = "/orgs/1/labels/1"
}

resource "illumio-core_label" "example" {
  key     = "role"
  value   = "example role"
}
