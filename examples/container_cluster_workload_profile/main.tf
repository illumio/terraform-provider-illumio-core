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


resource "illumio-core_container_cluster_workload_profile" "example" {
  container_cluster_href = "/orgs/1/container_clusters/bd37cbdd-82bd-4f49-b52f-9405ba236a43"
  name = "testing it"
  managed = true
  assign_labels {
    href = "/orgs/1/labels/1"
  }
}


data "illumio-core_container_cluster_workload_profile" "example" {
    href = "/orgs/1/container_clusters/bd37cbdd-82bd-4f49-b52f-9405ba236a43/container_workload_profiles/598888c7-a625-4507-a5c8-14f4a3c4c1d6"
}