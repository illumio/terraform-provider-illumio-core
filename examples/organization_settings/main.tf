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

resource "illumio-core_organization_settings" "test" {
  audit_event_retention_seconds = 7776000
  format = "JSON"
  audit_event_min_severity = "informational"
}

data "illumio-core_organization_settings" "test" {
  
}

