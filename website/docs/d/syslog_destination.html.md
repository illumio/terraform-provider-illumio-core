---
layout: "illumio-core"
page_title: "illumio-core_syslog_destination Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-syslog-destination"
subcategory: ""
description: |-
  Represents Illumio Syslog Destination
---

# illumio-core_syslog_destination (Data Source)

Represents Illumio Syslog Destination

Example Usage
------------

```hcl
data "illumio-core_syslog_destination" "example" {
  href = "/orgs/1/settings/syslog/destinations/11a4cfdf-a78e-4144-bbbc-67faec728df1"
}
```

## Schema

### Required

- **href** (String) URI of the destination

### Read-Only

- **audit_event_logger** (List of Object) audit_event_logger details for destination. Single element list (see [below for nested schema](#nestedatt--audit_event_logger))
- **description** (String) Description of the destination
- **node_status_logger** (List of Object) node_status_logger details for destination. Single element list (see [below for nested schema](#nestedatt--node_status_logger))
- **pce_scope** (List of String)
- **remote_syslog** (List of Object) remote_syslog details for destination. Single element list (see [below for nested schema](#nestedatt--remote_syslog))
- **traffic_event_logger** (List of Object) traffic_event_logger details for destination. Single element list (see [below for nested schema](#nestedatt--traffic_event_logger))
- **type** (String) Destination type

<a id="nestedatt--audit_event_logger"></a>
### Nested Schema for `audit_event_logger`

Read-Only:

- **configuration_event_included** (Boolean) Configuration (Northbound) auditable events
- **min_severity** (String) Minimum severity level of audit event messages
- **system_event_included** (Boolean) System (PCE) auditable events


<a id="nestedatt--node_status_logger"></a>
### Nested Schema for `node_status_logger`

Read-Only:

- **node_status_included** (Boolean) Syslog messages regarding status of the nodes


<a id="nestedatt--remote_syslog"></a>
### Nested Schema for `remote_syslog`

Read-Only:

- **address** (String) The remote syslog IP or DNS address
- **port** (Number) The remote syslog port
- **protocol** (Number) The protocol for streaming syslog messages
- **tls_ca_bundle** (String) Trustee CA bundle
- **tls_enabled** (Boolean) To enable TLS
- **tls_verify_cert** (Boolean) Perform TLS verification


<a id="nestedatt--traffic_event_logger"></a>
### Nested Schema for `traffic_event_logger`

Read-Only:

- **traffic_flow_allowed_event_included** (Boolean) Set to enable traffic flow events
- **traffic_flow_blocked_event_included** (Boolean) Set to enable traffic flow events
- **traffic_flow_potentially_blocked_event_included** (Boolean) Set to enable traffic flow events
