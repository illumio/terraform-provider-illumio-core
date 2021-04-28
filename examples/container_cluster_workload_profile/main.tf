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


resource "illumio-core_container_cluster_workload_profile" "test" {
  container_cluster_id = "deb48c70-e9d2-4101-ab7e-1f48de922ff4"
  name = "testing it"
  managed = true
  assign_labels {
    href = "/orgs/1/labels/1"
  }
}


data "illumio-core_container_cluster_workload_profile" "test" {
    container_cluster_id = "deb48c70-e9d2-4101-ab7e-1f48de922ff4"
    container_workload_profile_id = "0a7ed380-bc2e-4be6-99ad-741baf77fb91"
}