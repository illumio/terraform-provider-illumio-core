terraform {
  required_providers {
    illumio = {
      version = "0.1"
      source = "illumio.com/labs/illumio"
    }
  }
}

provider "illumio" {
//  pce_host              = "https://pce.my-company.com:8443"
//  api_username          = "api_xxxxxx"
//  api_secret            = "big-secret"
    request_timeout       = 30
    org_id                = 1
}

resource "illumio-core_workload_interface" "test" {
    workload_id = "d42a430e-b20b-4b2d-853f-2d39fa4cea22"
    name = "testing it itit"
    link_state = "up"
    friendly_name = "test friendly name"
}

data "illumio-core_workload_interface" "test" {
    workload_id = "d42a430e-b20b-4b2d-853f-2d39fa4cea22"
    workload_interface_name = "testing workload"
}
