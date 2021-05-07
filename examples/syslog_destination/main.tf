terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
  # pce_host              = "https://2x2devtest59.ilabs.io:8443"
  # api_username          = ""
  # api_secret            = ""
  request_timeout = 30
  org_id          = 1
}

data "illumio-core_syslog_destination" "example" {
  href = "/orgs/1/settings/syslog/destinations/11a4cfdf-a78e-4144-bbbc-67faec728df1"
}

resource "illumio-core_syslog_destination" "syslog" {
  type        = "remote_syslog"
  pce_scope   = ["crest-mnc.ilabs.io"]
  description = "test"

  audit_event_logger {
    configuration_event_included = false
    system_event_included        = true
    min_severity                 = "warning"
  }

  traffic_event_logger {
    traffic_flow_allowed_event_included             = true
    traffic_flow_potentially_blocked_event_included = false
    traffic_flow_blocked_event_included             = false
  }

  node_status_logger {
    node_status_included = false
  }

  remote_syslog {
    protocol        = 6
    address         = "35.164.106.210"
    port            = 5141
    tls_enabled     = false
    tls_verify_cert = false
  }
}
