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

resource "illumio-core_container_cluster" "test" {
    name = "test cc"
    description = "test desc"
}

data "illumio-core_container_cluster" "test" {
    href = "/orgs/1/container_clusters/bd37cbdd-82bd-4f49-b52f-9405ba236a43"
}
