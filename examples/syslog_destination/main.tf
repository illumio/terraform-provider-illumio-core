terraform {
  required_providers {
    illumio = {
      version = "0.1"
      source  = "illumio.com/labs/illumio"
    }
  }
}

provider "illumio" {
  # pce_host              = "https://2x2devtest59.ilabs.io:8443"
  # api_username          = ""
  # api_secret            = ""
  request_timeout = 30
  org_id          = 1
}

data "illumio-core_syslog_destination" "name" {
  syslog_destination_id = "04bdf118-c898-4dcb-8c74-b44a59cc1e02"
}

# output "name" {
#   value = data.illumio-core_syslog_destination.name
# }

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
