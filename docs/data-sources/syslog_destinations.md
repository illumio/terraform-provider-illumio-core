---
layout: "illumio-core"
page_title: "illumio-core_syslog_destinations Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-syslog-destinations"
subcategory: ""
description: |-
  Represents Illumio Syslog Destinations
---

# illumio-core_syslog_destinations (Data Source)

Represents Illumio Syslog Destinations

Example Usage
------------

```hcl
data "illumio-core_syslog_destinations" "example" {
  
}
```

## Schema

### Read-Only

- **items** (List of Object) list of Syslog Destinations (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **audit_event_logger** (List of Object) audit_event_logger details for destination. Single element list (see [below for nested schema](#nestedobjatt--items--audit_event_logger))
- **description** (String) Description of the destination
- **href** (String) URI of the destination
- **node_status_logger** (List of Object) node_status_logger details for destination. Single element list (see [below for nested schema](#nestedobjatt--items--node_status_logger))
- **pce_scope** (List of String) pce_scope for syslog destinations
- **remote_syslog** (List of Object) remote_syslog details for destination. Single element list (see [below for nested schema](#nestedobjatt--items--remote_syslog))
- **traffic_event_logger** (List of Object) traffic_event_logger details for destination. Single element list (see [below for nested schema](#nestedobjatt--items--traffic_event_logger))
- **type** (String) Destination type

<a id="nestedobjatt--items--audit_event_logger"></a>
### Nested Schema for `items.audit_event_logger`

Read-Only:

- **configuration_event_included** (Boolean) Configuration (Northbound) auditable events
- **min_severity** (String) Minimum severity level of audit event messages
- **system_event_included** (Boolean) System (PCE) auditable events


<a id="nestedobjatt--items--node_status_logger"></a>
### Nested Schema for `items.node_status_logger`

Read-Only:

- **node_status_included** (Boolean) Syslog messages regarding status of the nodes


<a id="nestedobjatt--items--remote_syslog"></a>
### Nested Schema for `items.remote_syslog`

Read-Only:

- **address** (String) The remote syslog IP or DNS address
- **port** (Number) The remote syslog port
- **protocol** (Number) The protocol for streaming syslog messages
- **tls_ca_bundle** (String) Trustee CA bundle
- **tls_enabled** (Boolean) To enable TLS
- **tls_verify_cert** (Boolean) Perform TLS verification


<a id="nestedobjatt--items--traffic_event_logger"></a>
### Nested Schema for `items.traffic_event_logger`

Read-Only:

- **traffic_flow_allowed_event_included** (Boolean) Set to enable traffic flow events
- **traffic_flow_blocked_event_included** (Boolean) Set to enable traffic flow events
- **traffic_flow_potentially_blocked_event_included** (Boolean) Set to enable traffic flow events


