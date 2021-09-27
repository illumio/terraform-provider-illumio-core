terraform {
  required_providers {
    illumio-core = {
      version = "0.1.0"
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  # pce_host        = ""
  # api_username    = "api_xxxxxxxxxxxxxxxxxxx"
  # api_secret      = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  request_timeout = 30
  org_id          = 1
}

resource "illumio-core_workload_settings" "example" {
  workload_disconnected_timeout_seconds {
    value = -1 
    }
  workload_goodbye_timeout_seconds {
    value = -1
  }
}

data "illumio-core_workload_settings" "example" {

}
