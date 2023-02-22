terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  pce_host     = "https://pce.my-company.com:8443"
  api_username = "api_xxxxxx"
  api_secret   = "xxxxxxxxxx"
  org_id       = 1
}
