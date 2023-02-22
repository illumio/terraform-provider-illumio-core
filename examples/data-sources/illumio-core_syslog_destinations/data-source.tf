locals {
  # split on : to strip method and port, then strip the leading //
  pce_hostname = substr(split(":", var.pce_url)[1], 2, -1)
}

resource "illumio-core_syslog_destination" "local" {
  type        = "local_syslog"
  pce_scope   = [local.pce_hostname]
  description = "Local syslog destination config"

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
}

data "illumio-core_syslog_destinations" "all" {
  depends_on = [
    illumio-core_syslog_destination.local,
  ]
}
