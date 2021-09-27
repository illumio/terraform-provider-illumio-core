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

resource "illumio-core_pairing_keys" "example" {
    pairing_profile_href = "/orgs/1/pairing_profiles/1"
    token_count = 3
}