---
layout: "illumio-core"
page_title: "illumio-core_workloads Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-workloads"
subcategory: ""
description: |-
  Represents Illumio Workloads
---

# illumio-core_workloads (Data Source)

Represents Illumio Workloads

Example Usage
------------

```hcl
data "illumio-core_workloads" "example" {}
```

## Schema

### Optional

- `agent_active_pce_fqdn` (String) FQDN of the PCE
- `container_clusters` (String) List of container cluster URIs, encoded as a JSON string
- `description` (String) Description of workload(s) to return. Supports partial matches
- `enforcement_mode` (String) Enforcement mode of workload(s) to return. Allowed values are "idle", "visibility_only", "full" and "selective"
- `external_data_reference` (String) A unique identifier within the external data source
- `external_data_set` (String) The data source from which a resource originates
- `hostname` (String) Hostname of workload(s) to return. Supports partial matches
- `include_deleted` (String) Include deleted workloads
- `ip_address` (String) IP address of workload(s) to return. Supports partial matches
- `labels` (String) List of lists of label URIs, encoded as a JSON string
- `last_heartbeat_on_gte` (String) Greater than or equal to value for last heartbeat on timestamp
- `last_heartbeat_on_lte` (String) Less than or equal to value for last heartbeat on timestamp
- `log_traffic` (String) Whether we want to log traffic events from this workload
- `managed` (String) Return managed or unmanaged workloads using this filter
- `max_results` (String) Maximum number of workloads to return. The integer should be a non-zero positive integer
- `name` (String) Name of workload(s) to return. Supports partial matches
- `online` (String) Return online/offline workloads using this filter
- `os_id` (String) Operating System of workload(s) to return. Supports partial matches
- `policy_health` (String) Policy of health of workload(s) to return. Allowed values are "active", "warning", "error" and "suspended"
- `security_policy_sync_state` (String) Advanced search option for workload based on policy sync state. Allowed value: "staged"
- `security_policy_update_mode` (String) Advanced search option for workload based on security policy update mode. Allowed values are "static" and "adaptive"
- `ven` (String) URI of VEN to filter by
- `visibility_level` (String) Filter by visibility level. Allowed values are "flow_full_detail", "flow_summary", "flow_drops", "flow_off" and "enhanced_data_collection"
- `vulnerability_summary_vulnerability_exposure_score_gte` (String) Greater than or equal to value for vulnerability_exposure_score
- `vulnerability_summary_vulnerability_exposure_score_lte` (String) Less than or equal to value for vulnerability_exposure_score

### Read-Only

- `items` (List of Object) list of workloads (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `agent_to_pce_certificate_authentication_id` (String) PKI Certificate identifier to be used by the PCE for authenticating the VEN
- `blocked_connection_action` (String) Blocked Connection Action for Workload
- `caps` (List of String) CAPS for Workload
- `container_cluster` (List of Object) Container Cluster for Workload (see [below for nested schema](#nestedobjatt--items--container_cluster))
- `containers_inherit_host_policy` (Boolean) This workload will apply the policy it receives both to itself and the containers hosted by it
- `created_at` (String) Timestamp when this Workload was first created
- `created_by` (Map of String) User who created this Workload
- `data_center` (String) Data center for Workload
- `data_center_zone` (String) Data center Zone for Workload
- `deleted` (Boolean) This indicates that the workload has been deleted
- `deleted_at` (String) Timestamp when this Workload was deleted
- `deleted_by` (Map of String) User who deleted this Workload
- `description` (String) Description of the Workload
- `detected_vulnerabilities` (List of Object) Detected Vulnerabilities (see [below for nested schema](#nestedobjatt--items--detected_vulnerabilities))
- `distinguished_name` (String) X.509 Subject distinguished name
- `enforcement_mode` (String) Enforcement mode of workload(s) to return
- `external_data_reference` (String) A unique identifier within the external data source
- `external_data_set` (String) The data source from which a resource originates
- `firewall_coexistence` (List of Object) Firewall coexistence mode (see [below for nested schema](#nestedobjatt--items--firewall_coexistence))
- `hostname` (String) The hostname of this workload
- `href` (String) URI of workload
- `ignored_interface_names` (List of String) Ignored Interface Names for Workload
- `ike_authentication_certificate` (Map of String)
- `interfaces` (List of Object) List of interfaces for workload (see [below for nested schema](#nestedobjatt--items--interfaces))
- `labels` (List of Object) List of lists of label URIs (see [below for nested schema](#nestedobjatt--items--labels))
- `name` (String) Name of the Workload
- `online` (Boolean) Determines if this workload is online
- `os_detail` (String) Additional OS details - just displayed to end-user
- `os_id` (String) OS identifier for Workload
- `public_ip` (String) The public IP address of the server
- `selectively_enforced_services` (List of Object) Selectively Enforced Services for Workload (see [below for nested schema](#nestedobjatt--items--selectively_enforced_services))
- `service_principal_name` (String) The Kerberos Service Principal Name (SPN)
- `service_provider` (String) Service provider for Workload
- `services` (List of Object) Service report for Workload (see [below for nested schema](#nestedobjatt--items--services))
- `updated_at` (String) Timestamp when this Workload was last updated
- `updated_by` (Map of String) User who last updated this Workload
- `ven` (List of String) VENS for Workload
- `visibility_level` (String) Visibility Level of workload(s) to return
- `vulnerabilities_summary` (List of Object) Vulnerabilities summary associated with the workload (see [below for nested schema](#nestedobjatt--items--vulnerabilities_summary))

<a id="nestedobjatt--items--container_cluster"></a>
### Nested Schema for `items.container_cluster`

Read-Only:

- `href` (String) URI of container cluster
- `name` (String) Name of container cluster


<a id="nestedobjatt--items--detected_vulnerabilities"></a>
### Nested Schema for `items.detected_vulnerabilities`

Read-Only:

- `ip_address` (String) The IP address of the host where the vulnerability is found
- `port` (Number) The port that is associated with the vulnerability
- `port_exposure` (Number) The exposure of the port based on the current policy
- `port_wide_exposure` (List of Object) Port Wide Exposure for detected vulnerabilities (see [below for nested schema](#nestedobjatt--items--detected_vulnerabilities--port_wide_exposure))
- `proto` (Number) The protocol that is associated with the vulnerability
- `vulnerability` (List of Object) Vulnerability for Workload (see [below for nested schema](#nestedobjatt--items--detected_vulnerabilities--vulnerability))
- `vulnerability_report` (List of Object) Vulnerability Report for Workload (see [below for nested schema](#nestedobjatt--items--detected_vulnerabilities--vulnerability_report))
- `workload` (List of Object) URI of Workload (see [below for nested schema](#nestedobjatt--items--detected_vulnerabilities--workload))

<a id="nestedobjatt--items--detected_vulnerabilities--port_wide_exposure"></a>
### Nested Schema for `items.detected_vulnerabilities.port_wide_exposure`

Read-Only:

- `any` (Boolean) The boolean value representing if the port is exposed to internet (any rule)
- `ip_list` (Boolean) The boolean value representing if the port is exposed to ip_list(s)


<a id="nestedobjatt--items--detected_vulnerabilities--vulnerability"></a>
### Nested Schema for `items.detected_vulnerabilities.vulnerability`

Read-Only:

- `href` (String) The URI of the workload to which this vulnerability belongs to
- `name` (String) The title/name of the vulnerability
- `score` (Number) The normalized score of the vulnerability within the range of 0 to 100

<a id="nestedobjatt--items--detected_vulnerabilities--vulnerability_report"></a>
### Nested Schema for `items.detected_vulnerabilities.vulnerability_report`

Read-Only:

- `href` (String) The URI of the vulnerability Report to which this vulnerability belongs to

<a id="nestedobjatt--items--detected_vulnerabilities--workload"></a>
### Nested Schema for `items.detected_vulnerabilities.workload`

Read-Only:

- `href` (String) The URI of the workload to which this vulnerability belongs to

<a id="nestedobjatt--items--firewall_coexistence"></a>
### Nested Schema for `items.firewall_coexistence`

Read-Only:

- `illumio_primary` (Boolean) Illumio is the primary firewall if set to true

<a id="nestedobjatt--items--interfaces"></a>
### Nested Schema for `items.interfaces`

Read-Only:

- `address` (String) Address of the Interface
- `cidr_block` (Number) CIDR BLOCK of the Interface. The number of bits in the subnet /24 is 255.255.255.0
- `default_gateway_address` (String) Default Gateway Address of the Interface
- `friendly_name` (String) Friendly name of the Interface
- `loopback` (Boolean) Loopback for workload interfaces
- `link_state` (String) Link State of the Interface
- `name` (String) Name of the Interface
- `network` (Map of String) Network of the Interface
- `network_detection_mode` (String) Network Detection Mode of the Interface

<a id="nestedobjatt--items--labels"></a>
### Nested Schema for `items.labels`

Read-Only:

- `href` (String) URI of the labels

<a id="nestedobjatt--items--selectively_enforced_services"></a>
### Nested Schema for `items.selectively_enforced_services`

Read-Only:

- `href` (String) URI of Selectively Enforced Services
- `port` (Integer) Port number, or the starting port of a range. If unspecified, this will apply to all ports for the given protocol. Minimum and maximum value for port is 0 and 65535 respectively
- `to_port` (Integer) Upper end of port range; this field should not be included if specifying an individual port. Minimum and maximum value for to_port is 0 and 65535 respectively
- `proto` (Integer) Transport protocol of Selectively Enforced Services

<a id="nestedobjatt--items--services"></a>
### Nested Schema for `items.services`

Read-Only:

- `created_at` (String) Timestamp when this service was first created
- `open_service_ports` (List of Object) A list of open ports (see [below for nested schema](#nestedobjatt--items--services--open_service_ports))
- `uptime_seconds` (Number)  How long since the last reboot of this box - used as a timestamp for this

<a id="nestedobjatt--items--services--open_service_ports"></a>
### Nested Schema for `items.services.open_service_ports`

Read-Only:

- `address` (String) The local address this service is bound to
- `package` (String) The RPM/DEB package that the program is part of
- `port` (Number) The local port this service is bound to
- `process_name` (String) The process name (including the full path)
- `protocol` (Number) Transport protocol for open service ports
- `user` (String) The user account that the process is running under
- `win_service_name` (String) Name of the Windows service

<a id="nestedobjatt--items--vulnerabilities_summary"></a>
### Nested Schema for `items.vulnerabilities_summary`

Read-Only:

- `max_vulnerability_score` (Number) The maximum of all the vulnerability scores associated with the detected_vulnerabilities on the workload
- `num_vulnerabilities` (Number) Number of vulnerabilities associated with the workload
- `vulnerability_exposure_score` (Number) The aggregated vulnerability exposure score of the workload across all the vulnerable ports
- `vulnerability_score` (Number) The aggregated vulnerability score of the workload across all the vulnerable ports
- `vulnerable_port_exposure` (Number) The aggregated vulnerability port exposure score of the workload across all the vulnerable ports
- `vulnerable_port_wide_exposure` (List of Object) High end of an IP range (see [below for nested schema](#nestedobjatt--items--vulnerabilities_summary--vulnerable_port_wide_exposure))

<a id="nestedobjatt--items--vulnerabilities_summary--vulnerable_port_wide_exposure"></a>
### Nested Schema for `items.vulnerabilities_summary.vulnerable_port_wide_exposure`

Read-Only:

- `any` (Boolean) The boolean value representing if at least one port is exposed to internet (any rule) on the workload
- `ip_list` (Boolean) The boolean value representing if at least one port is exposed to ip_list(s) on the workload
