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

data "illumio-core_workload_interfaces" "example" {
  workload_href = "/orgs/1/workloads/63bf19d1-1efa-49ec-b712-c51d5c0aa552"
}