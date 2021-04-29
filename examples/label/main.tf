terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source = "illumio.com/labs/illumio-core"
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

data "illumio-core_label" "label_1" {
  label_id  = 1
}

resource "illumio-core_label" "test_label" {
  key     = "role"
  value   = "test_role_2"
}
