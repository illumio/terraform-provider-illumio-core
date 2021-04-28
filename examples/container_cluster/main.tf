terraform {
  required_providers {
    illumio = {
      version = "0.1"
      source  = "illumio.com/labs/illumio"
    }
  }
}


provider "illumio" {
  //  pce_host              = "https://pce.my-company.com:8443"
  //  api_username          = "api_xxxxxx"
  //  api_secret            = "big-secret"
  request_timeout = 30
  org_id          = 1
}

resource "illumio-core_container_cluster" "test" {
    name = "test cc"
    description = "test desc"
}

data "illumio-core_container_cluster" "test" {
    container_cluster_id = "e9af6d1c-ce13-4c4d-8aa8-5ff0a3a1f378"
}
