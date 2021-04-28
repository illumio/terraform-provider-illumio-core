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

# resource "illumio-core_workload" "example" {
#   name        = "test workload name"
#   description = "test Desc"
#   # external_data_set       = "test set"
#   # external_data_reference = "test reference"
#   hostname               = "test hostname"
#   service_principal_name = "test spn"
#   interfaces {
#     name       = "test interface"
#     link_state = "up"
#     address    = "10.10.3.10"
#   }
#   service_provider = "test service provider"
#   data_center      = "test data center"
#   data_center_zone = "test data center zone"
#   os_detail        = "test os details"
#   os_id            = "test os id"
#   online           = false
#   labels {
#     href = "/orgs/1/labels/1"
#   }
#   enforcement_mode = "visibility_only"
# }

resource "illumio-core_workload" "list" {
  name  = ""
  count = 0
}

data "illumio-core_workload" "test" {
  workload_id = "3e34d8f6-3255-4b53-b98e-888e93647200"
}
