terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
  pce_host        = "https://2x2devtest59.ilabs.io:8443/"
  api_username    = "api_15ee563a713f544b2"
  api_secret      = "a4a5a60d8d68bb22d9ab2e5b68337d3b12784e2b8a45c13118e5098a9d4638c9"
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
