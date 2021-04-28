---
layout: "illumio-core"
page_title: "illumio-core_workload Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-workload"
subcategory: ""
description: |-
  Represents Illumio Workload
---

# illumio-core_workload (Data Source)

Represents Illumio Workload

Example Usage
------------

```hcl
data "illumio-core_workload" "example"  {
  workload_id = "c078947c-15ec-4451-a758-e9a5be3d2fa7"
}
```

## Schema

### Required

- **workload_id** (String) Numerical ID of workload


### Read-Only

- **agent_to_pce_certificate_authentication_id** (String) PKI Certificate identifier to be used by the PCE for authenticating the VEN
- **blocked_connection_action** (String) Blocked Connection Action for Workload
- **caps** (List of String) CAPS for Workload
- **container_cluster** (List of Object) Container Cluster for Workload (see [below for nested schema](#nestedatt--container_cluster))
- **containers_inherit_host_policy** (Boolean) This workload will apply the policy it receives both to itself and the containers hosted by it
- **created_at** (String) Timestamp when this Workload was first created
- **created_by** (Map of String) User who originally created this Workload
- **data_center** (String) Data center for Workload
- **data_center_zone** (String) Data center Zone for Workload
- **deleted** (Boolean) This indicates that the workload has been deleted
- **deleted_at** (String) Timestamp when this Workload was deleted
- **deleted_by** (Map of String) User who deleted this Workload
- **description** (String) Description of the Workload
- **detected_vulnerabilities** (List of Object) Detected Vulnerabilities (see [below for nested schema](#nestedatt--detected_vulnerabilities))
- **distinguished_name** (String) X.509 Subject distinguished name
- **enforcement_mode** (String) Enforcement mode of workload(s) to return
- **external_data_reference** (String) A unque identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **firewall_coexistence** (List of Object) Firewall coexistence mode (see [below for nested schema](#nestedatt--firewall_coexistence))
- **hostname** (String) The hostname of this workload
- **href** (String) URI of the Workload
- **ignored_interface_names** (List of String) Ignored Interface Names for Workload
- **interfaces** (List of Object) A unque identifier within the external data source (see [below for nested schema](#nestedatt--interfaces))
- **labels** (List of Object) List of lists of label URIs (see [below for nested schema](#nestedatt--labels))
- **name** (String) Name of the Workload
- **online** (Boolean) Determines if this workload is online
- **os_detail** (String) Additional OS details - just displayed to end user
- **os_id** (String) OS identifier for Workload
- **public_ip** (String) The public IP address of the server
- **selectively_enforced_services** (List of Object) Selectively Enforced Services for Workload (see [below for nested schema](#nestedatt--selectively_enforced_services))
- **service_principal_name** (String) The Kerberos Service Principal Name (SPN)
- **service_provider** (String) Service provider for Workload
- **services** (List of Object) Service report for Workload (see [below for nested schema](#nestedatt--services))
- **updated_at** (String) Timestamp when this Workload was last updated
- **updated_by** (Map of String) User who last updated this Workload
- **ven** (List of String) VENS for Workload
- **visibility_level** (String) Visibility Level of workload(s) to return
- **vulnerabilities_summary** (List of Object) Vulnerabilities summary associated with the workload (see [below for nested schema](#nestedatt--vulnerabilities_summary))

<a id="nestedatt--container_cluster"></a>
### Nested Schema for `container_cluster`

Read-Only:

- **href** (String) URI of conatainer cluster
- **name** (String) Name of conatainer cluster

<a id="nestedatt--detected_vulnerabilities"></a>
### Nested Schema for `detected_vulnerabilities`

Read-Only:

- **ip_address** (String) The ip address of the host where the vulnerability is found
- **port** (Number) The port which is associated with the vulnerability
- **port_exposure** (Number) The exposure of the port based on the current policy.
- **proto** (Number) The protocol which is associated with the vulnerability
- **vulnerability** (Set of Object) Vulnerability for Workload (see [below for nested schema](#nestedobjatt--detected_vulnerabilities--vulnerability **vulnerability_report** (Set of Object) Vulnerability Report for Workload (see [below for nested schema](#nestedobjatt--detected_vulnerabilities--vulnerability_report))
- **workload** (List of Object) URI of Workload (see [below for nested schema](#nestedobjatt--detected_vulnerabilities--vulnerability_report))

<a id="nestedobjatt--detected_vulnerabilities--vulnerability_report"></a>
### Nested Schema for `detected_vulnerabilities.vulnerability_report`

Read-Only:

- **href** (String) The URI of the workload to which this vulnerability belongs to
- **name** (String) The title/name of the vulnerability
- **score** (Number) The normalized score of the vulnerability within the range of 0 to 100


<a id="nestedobjatt--detected_vulnerabilities--workload"></a>
### Nested Schema for `detected_vulnerabilities.workload`

Read-Only:

- **href** (String) The URI of the workload to which this vulnerability belongs to

<a id="nestedatt--firewall_coexistence"></a>
### Nested Schema for `firewall_coexistence`

Read-Only:

- **illumio_primary** (Boolean) Illumio is the primary firewall if set to true

<a id="nestedatt--interfaces"></a>
### Nested Schema for `interfaces`

Read-Only:

- **address** (String) Address of the Interface
- **cidr_block** (Number) CIDR BLOCK of the Interface. The number of bits in the subnet /24 is 255.255.255.0.
- **default_gateway_address** (String) Default Gateaway Address of the Interface
- **friendly_name** (String) Friendly name of the Interface
- **link_state** (String) Link State of the Interface
- **name** (String) Name of the Interface
- **network** (Map of String) Network of the Interface
- **network_detection_mode** (String) Network Detection Mode of the Interface

<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Read-Only:

- **href** (String) URI of the labels

<a id="nestedatt--selectively_enforced_services"></a>
### Nested Schema for `selectively_enforced_services`

Read-Only:

- **href** (String) URI of Selectively Enforced Services
- **port** (Integer) Port number, or the starting port of a range. If unspecified, this will apply to all ports for the given protocol. Minimum and maximum value for port is 0 and 65535 respectively.
- **to_port** (Integer) Upper end of port range; this field should not be included if specifying an individual port. Minimum and maximum value for to_port is 0 and 65535 respectively.
- **proto** (Integer) Transport protocol of Selectively Enforced Services

<a id="nestedatt--services"></a>
### Nested Schema for `services`

Read-Only:

- **created_at** (String) Timestamp when this service was first created
- **open_service_ports** (List of Object) A list of open ports (see [below for nested schema] (#nestedobjatt--services--open_service_ports))
- **uptime_seconds** (Number) How long since the last reboot of this box - used as a timestamp for this

<a id="nestedobjatt--services--open_service_ports"></a>

### Nested Schema for `services.open_service_ports`

Read-Only:

- **address** (String) The local address this service is bound to
- **package** (String) The RPM/DEB pacakge that the program is part of
- **port** (Number) The local port this service is bound to
- **process_name** (String) The process name (including the full path)
- **protocol** (Number) Transport protocol for open service ports
- **user** (String) The user account that the process is running under
- **win_service_name** (String) Name of the Windows service


<a id="nestedatt--vulnerabilities_summary"></a>
### Nested Schema for `vulnerabilities_summary`

Read-Only:

- **max_vulnerability_score** (Number) The maximum of all the vulnerability scores associated with the detected_vulnerabilities on the workload
- **num_vulnerabilities** (Number) Number of vulnerabilities associated with the workload
- **vulnerability_exposure_score** (Number) The aggregated vulnerability exposure score of the workload across all the vulnerable ports.
- **vulnerability_score** (Number) The aggregated vulnerability score of the workload across all the vulnerable ports.
- **vulnerable_port_exposure** (Number) The aggregated vulnerability port exposure score of the workload across all the vulnerable ports
- **vulnerable_port_wide_exposure** (List of Object) High end of an IP range (see [below for nested schema](#nestedobjatt--vulnerabilities_summary--vulnerable_port_wide_exposure))

<a id="nestedobjatt--vulnerabilities_summary--vulnerable_port_wide_exposure"></a>
### Nested Schema for `vulnerabilities_summary.vulnerable_port_wide_exposure`

Read-Only:

- **any** (Boolean) The boolean value representing if at least one port is exposed to internet (any rule) on the workload
- **ip_list** (Boolean) The boolean value representing if at least one port is exposed to ip_list(s) on the workload


