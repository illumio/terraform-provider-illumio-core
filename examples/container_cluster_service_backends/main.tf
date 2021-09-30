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

data "illumio-core_container_cluster_service_backends" "example" {
  container_cluster_href = "/orgs/1/container_clusters/f959d2d0-fe56-4bd9-8132-b7a31d1cbdde"
}

