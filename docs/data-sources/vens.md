---
layout: "illumio-core"
page_title: "illumio-core_vens Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-vens"
subcategory: ""
description: |-
  Represents Illumio VENs
---

# illumio-core_vens (Data Source)

Represents Illumio VENs

Example Usage
------------

```hcl
data "illumio-core_vens" "example" {}
```

## Schema

### Optional

- `activation_type` (String) The method in which the VEN was activated. Allowed values are "pairing_key", "kerberos" and "certificate"
- `active_pce_fqdn` (String) FQDN of the PCE
- `condition` (String) A specific error condition to filter by. Allowed values are "agent.upgrade_time_out", "agent.missing_heartbeats_after_upgrade", "agent.clone_detected" and "agent.missed_heartbeats"
- `container_clusters` (String) Array of container cluster URIs, encoded as a JSON string
- `description` (String) Description of VEN(s) to return. Supports partial matches
- `disconnected_before` (String) Return VENs that have been disconnected since the given time
- `health` (String) The overall health (condition) of the VEN. Allowed values are  "healthy", "unhealthy", "error" and "warning"
- `hostname` (String) Hostname of VEN(s) to return. Supports partial matches
- `ip_address` (String) IP address of VEN(s) to return. Supports partial matches
- `labels` (String) 2D Array of label URIs, encoded as a JSON string
- `last_goodbye_at_gte` (String) Greater than or equal to value for last goodbye at timestamp
- `last_goodbye_at_lte` (String) Greater than or equal to value for last goodbye at timestamp
- `last_heartbeat_at_gte` (String) Greater than or equal to value for last heartbeat timestamp
- `last_heartbeat_at_lte` (String) Less than or equal to value for last heartbeat timestamp
- `max_results` (String) Maximum number of VENs to return. The integer should be a non-zero positive integer
- `name` (String) Name of VEN(s) to return. Supports partial matches
- `os` (String) Operating System of VEN(s) to return. Supports partial matches
- `status` (String) The current status of the VEN. Allowed values are "active", "suspended", "stopped" and "uninstalled"
- `upgrade_pending` (String) Only return VENs with/without a pending upgrade
- `version_gte` (String) Greater than or equal to value for version
- `version_lte` (String) Less than or equal to value for version

### Read-Only

- `items` (List of Object) list of VENs (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `activation_type` (String) The method by which the VEN was activated
- `active_pce_fqdn` (String) The FQDN of the PCE that the VEN last connected to
- `caps` (List of String) Permission types
- `conditions` (List of Object) (see [below for nested schema](#nestedobjatt--items--conditions))
- `container_cluster` (List of Object) container_cluster details for ven. Single element list(see [below for nested schema](#nestedobjatt--items--container_cluster))
- `created_at` (String) The time (rfc3339 timestamp) at which this VEN was created
- `created_by` (Map of String) The href of the user who created this VEN
- `description` (String) The description of the VEN
- `hostname` (String) The hostname of the host managed by the VEN
- `href` (String) URI of VEN
- `interfaces` (List of Object) Network interfaces of the host managed by the VEN (see [below for nested schema](#nestedobjatt--items--interfaces))
- `labels` (List of Object) Labels assigned to the host managed by the VEN (see [below for nested schema](#nestedobjatt--items--labels))
- `last_goodbye_at` (String) The time (rfc3339 timestamp) of the last goodbye from the VEN
- `last_heartbeat_at` (String) The last time (rfc3339 timestamp) a heartbeat was received from this VEN
- `name` (String) Friendly name for the VEN
- `os_detail` (String) Additional OS details from the host managed by the VEN
- `os_id` (String) OS identifier of the host managed by the VEN
- `os_platform` (String) OS platform of the host managed by the VEN
- `secure_connect` (List of Object) secure_connect details for vens (see [below for nested schema](#nestedobjatt--items--secure_connect))
- `status` (String) Status of the VEN
- `target_pce_fqdn` (String) The FQDN of the PCE that the VEN will use for future connections
- `uid` (String) The unique ID of the host managed by the VEN
- `unpair_allowed` (Boolean)
- `updated_at` (String) The time (rfc3339 timestamp) at which this VEN was last updated
- `updated_by` (Map of String) The href of the user who last updated this VEN
- `version` (String) Software version of the VEN
- `workloads` (List of Object) collection of Workloads (see [below for nested schema](#nestedobjatt--items--workloads))

<a id="nestedobjatt--items--conditions"></a>
### Nested Schema for `items.conditions`

Read-Only:

- `first_reported_timestamp` (String) The timestamp of the first event that reported this condition
- `latest_event` (List of Object) The latest notification event that was generated for the corresponding condition. Single element list (see [below for nested schema](#nestedobjatt--items--conditions--latest_event))

<a id="nestedobjatt--items--conditions--latest_event"></a>
### Nested Schema for `items.conditions.latest_event`

Read-Only:

- `href` (String) The href of the event
- `info` (List of Object) (see [below for nested schema](#nestedobjatt--items--conditions--latest_event--info))
- `notification_type` (String) The information from the notification event that was generated by the condition
- `severity` (String) Severity of the condition, same as the event
- `timestamp` (String) RFC 3339 timestamp at which this event was created

<a id="nestedobjatt--items--conditions--latest_event--info"></a>
### Nested Schema for `items.conditions.latest_event.timestamp`

Read-Only:

- `agent` (Map of String) Agent info

<a id="nestedobjatt--items--container_cluster"></a>
### Nested Schema for `items.container_cluster`

Read-Only:

- `href` (String) The URI of the container cluster managed by this VEN
- `name` (String) The name of the container cluster managed by this VEN, only present in expanded representations


<a id="nestedobjatt--items--interfaces"></a>
### Nested Schema for `items.interfaces`

Read-Only:

- `address` (String) The IP Address to assign to this interface
- `cidr_block` (Number) The number of bits in the subnet
- `default_gateway_address` (String) The IP Address of the default gateway
- `friendly_name` (String) User-friendly name for interface
- `href` (String) Interface URI
- `link_state` (String) Link State
- `loopback` (Boolean) loopback for interface
- `name` (String) Interface name
- `network` (Map of String) Network that the interface belongs to
- `network_detection_mode` (String) Network Detection Mode

<a id="nestedobjatt--items--labels"></a>
### Nested Schema for `items.labels`

Read-Only:

- `href` (String) Label URI
- `key` (String) Key of the label
- `value` (String) Value of the label

<a id="nestedobjatt--items--secure_connect"></a>
### Nested Schema for `items.secure_connect`

Read-Only:

- `matching_issuer_name` (String) Issuer name match criteria for certificate used during establishing secure connections

<a id="nestedobjatt--items--workloads"></a>
### Nested Schema for `items.workloads`

Read-Only:

- `enforcement_mode` (String) Policy enforcement mode
- `hostname` (String) The hostname of this workload
- `href` (String) URI of the Workload
- `interfaces` (List of Object) Network interfaces of the workload (see [below for nested schema](#nestedobjatt--items--workloads--interfaces))
- `labels` (List of Object) Labels assigned to the host managed by the VEN (see [below for nested schema](#nestedobjatt--items--workloads--labels))
- `mode` (String) Policy enforcement mode
- `name` (String) The short friendly name of the workload
- `online` (Boolean) If this workload is online and present in policy
- `os_detail` (String) Additional OS details
- `os_id` (String) OS identifier for the workload
- `public_ip` (String) The public IP address of the server
- `security_policy_applied_at` (String) Last reported time when policy was applied to the workload (UTC)
- `security_policy_received_at` (String) Last reported time when policy was received by the workload (UTC)
- `visibility_level` (String) Visibility level of the workload

<a id="nestedobjatt--items--workloads--interfaces"></a>
### Nested Schema for `items.workloads.interfaces`

Read-Only:

- `address` (String) The IP Address to assign to this interface
- `cidr_block` (Number) The number of bits in the subnet
- `default_gateway_address` (String) The IP Address of the default gateway
- `friendly_name` (String) User-friendly name for interface
- `link_state` (String) Link State
- `loopback` (Boolean) loopback for interface
- `name` (String) Interface name
- `network` (Map of String) Network that the interface belongs to
- `network_detection_mode` (String) Network Detection Mode

<a id="nestedobjatt--items--workloads--labels"></a>
### Nested Schema for `items.workloads.labels`

Read-Only:

- `href` (String) Label URI
- `key` (String) Key of the label
- `value` (String) Value of the label
