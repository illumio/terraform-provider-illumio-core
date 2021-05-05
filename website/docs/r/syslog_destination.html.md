---
layout: "illumio-core"
page_title: "illumio-core_syslog_destination Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-syslog-destination"
subcategory: ""
description: |-
  Manages Illumio SyslogDestination
---

# illumio-core_syslog_destination (Resource)

Manages Illumio SyslogDestination

Example Usage
------------

```hcl
resource "illumio-core_syslog_destination" "example" {
  type        = "remote_syslog"
  pce_scope   = ["company-mnc.ilabs.io"]
  description = "example description"

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
    address         = "36.164.106.210"
    port            = 5141
    tls_enabled     = false
    tls_verify_cert = false
  }
}
```

## Schema

### Required

- **audit_event_logger** (Block List, Min: 1, Max: 1) audit_event_logger details for destination. Single element list (see [below for nested schema](#nestedblock--audit_event_logger))
- **description** (String) Description of the destination
- **node_status_logger** (Block List, Min: 1, Max: 1) node_status_logger details for destination. Single element list (see [below for nested schema](#nestedblock--node_status_logger))
- **pce_scope** (Set of String) pce_scope for destination
- **traffic_event_logger** (Block List, Min: 1, Max: 1) traffic_event_logger details for destination. Single element list (see [below for nested schema](#nestedblock--traffic_event_logger))
- **type** (String) Destination type. Allowed values are "local_syslog" and "remote_syslog"

### Optional
- **remote_syslog** (Block List, Max: 1) remote_syslog details for destination. Single element list (see [below for nested schema](#nestedblock--remote_syslog))

### Read-Only

- **href** (String) URI of the destination

<a id="nestedblock--audit_event_logger"></a>
### Nested Schema for `audit_event_logger`

Required:

- **configuration_event_included** (Boolean) Configuration (Northbound) auditable events
- **min_severity** (String) Minimum severity level of audit event messages. Allowed values are "error", "warning" and "informational"
- **system_event_included** (Boolean) System (PCE) auditable events


<a id="nestedblock--node_status_logger"></a>
### Nested Schema for `node_status_logger`

Required:

- **node_status_included** (Boolean) Syslog messages regarding status of the nodes


<a id="nestedblock--traffic_event_logger"></a>
### Nested Schema for `traffic_event_logger`

Required:

- **traffic_flow_allowed_event_included** (Boolean) Set to enable traffic flow events
- **traffic_flow_blocked_event_included** (Boolean) Set to enable traffic flow events
- **traffic_flow_potentially_blocked_event_included** (Boolean) Set to enable traffic flow events


<a id="nestedblock--remote_syslog"></a>
### Nested Schema for `remote_syslog`

Required:

- **address** (String) The remote syslog IP or DNS address
- **port** (Number) The remote syslog port
- **protocol** (Number) The protocol for streaming syslog messages. Allowed values are 6 and 17
- **tls_enabled** (Boolean) To enable TLS
- **tls_verify_cert** (Boolean) Perform TLS verification

Optional:

- **tls_ca_bundle** (String) Trustee CA bundle


