terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  pce_host     = var.pce_url
  org_id       = var.pce_org_id
  api_username = var.pce_api_key
  api_secret   = var.pce_api_secret
}

resource "illumio-core_workload_settings" "current" {
  workload_disconnected_timeout_seconds {
    value = 3600
  }

  workload_goodbye_timeout_seconds {
    value = 900
  }
}

data "illumio-core_workload_settings" "current" {}
