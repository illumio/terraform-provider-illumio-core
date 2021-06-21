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

resource "illumio-core_workload" "example" {
  name        = "example workload name"
  description = "example Desc"
  external_data_set       = "example set"
  external_data_reference = "example reference"
  hostname               = "example hostname"
  service_principal_name = "example spn 99"
  service_provider = "example service provider"
  data_center      = "example data center"
  data_center_zone = "example data center zone"
  os_detail        = "example os details"
  os_id            = "example os id"
  online           = false
  labels {
    href = "/orgs/1/labels/1"
  }
  enforcement_mode = "visibility_only"
}

data "illumio-core_workload" "example" {
  href = "/orgs/1/workloads/e683b686-8afe-4675-88a1-4463395f0482"
}
