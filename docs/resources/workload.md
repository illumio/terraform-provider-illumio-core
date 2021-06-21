---
layout: "illumio-core"
page_title: "illumio-core_workload Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resoruce-workload"
subcategory: ""
description: |-
  Manages Illumio Workload
---

# illumio-core_workload (Resource)

Manages Illumio Workload


Example Usage
------------

```hcl
resource "illumio-core_workload" "example" {
    name = "example workload name"
    description = "example Desc"
    external_data_set ="example set"
    external_data_reference = "example reference"
    hostname = "example hostname"
    service_principal_name = "example spn"
    public_ip = "0.0.0.0"
    service_provider = "example service provider"
    data_center = "example data center"
    data_center_zone = "example data center zone"
    os_detail = "example os details"
    os_id = "example os id"
    online = false
    labels{
      href = "/orgs/1/labels/1"
    }
    enforcement_mode = "visibility_only"
}  
```
## Schema

### Optional

- **agent_to_pce_certificate_authentication_id** (String) PKI Certificate identifier to be used by the PCE for authenticating the VEN. The ID should be between 1 to 255 characters
- **data_center** (String) Data center for Workload. The data_center should be up to 255 characters
- **data_center_zone** (String) Data center Zone for Workload. The data_center_zone should be up to 255 characters
- **description** (String) The long description of the workload
- **distinguished_name** (String) X.509 Subject distinguished name. The name should be up to 255 characters
- **enforcement_mode** (String) Enforcement mode of workload(s) to return. Allowed values for enforcement modes are "idle","visibility_only","full", and "selective". Default value: "visibility_only"
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **hostname** (String) The hostname of this workload. The hostname should be up to 255 characters
<!-- - **interfaces** (Block Set) Workload network interfaces (see [below for nested schema](#nestedblock--interfaces)) -->
- **labels** (Block Set) Assigned labels for workload (see [below for nested schema](#nestedblock--labels))
- **name** (String) Name of the Workload. The name should be up to 255 characters
- **online** (Boolean) Determines if this workload is online. Default value: false
- **os_detail** (String) Additional OS details - just displayed to end-user. The os_details should be up to 255 characters
- **os_id** (String) OS identifier for Workload. The os_id should be up to 255 characters
- **public_ip** (String) The public IP address of the server. The public IP should in the IPv4 or IPv6 format
- **service_principal_name** (String) The Kerberos Service Principal Name (SPN). The SPN should be between 1 to 255 characters
- **service_provider** (String) Service provider for Workload. The service_provider should be up to 255 characters

### Read-Only

- **blocked_connection_action** (String) Blocked Connection Action for Workload
- **caps** (List of String) CAPS for Workload
- **container_cluster** (List of Object) Container Cluster for Workload (see [below for nested schema](#nestedatt--container_cluster))
- **containers_inherit_host_policy** (Boolean) This workload will apply the policy it receives both to itself and the containers hosted by it
- **created_at** (String) Timestamp when this label group was first created
- **created_by** (Map of String) User who created this label group
- **deleted** (Boolean) This indicates that the workload has been deleted
- **deleted_at** (String) Timestamp when this label group was last deleted
- **deleted_by** (Map of String) User who deleted this label group
- **detected_vulnerabilities** (List of Object) Detected Vulnerabilities for Workload (see [below for nested schema](#nestedatt--detected_vulnerabilities))
- **firewall_coexistence** (List of Object) Firewall coexistence mode for Workload (see [below for nested schema](#nestedatt--firewall_coexistence))
- **href** (String) URI of the Workload
- **ignored_interface_names** (List of String) Ignored Interface Names for Workload
- **ike_authentication_certificate** (Map of String) IKE authentication certificate for certificate-based Secure Connect and Machine Auth
- **selectively_enforced_services** (List of Object) Selectively Enforced Services for Workload (see [below for nested schema](#nestedatt--selectively_enforced_services))
- **services** (List of Object) Service report for Workload (see [below for nested schema](#nestedatt--services))
- **updated_at** (String) Timestamp when this label group was last updated
- **updated_by** (Map of String) User who last updated this label group
- **ven** (Map of String) VENS for Workload
- **visibility_level** (String) Visibility Level of workload(s) to return
- **vulnerabilities_summary** (List of Object) Vulnerabilities summary associated with the workload (see [below for nested schema](#nestedatt--vulnerabilities_summary))

<!-- <a id="nestedblock--interfaces"></a>
### Nested Schema for `interfaces`

Required:

- **link_state** (String) Link State for the workload Interface. Allowed values are "up", "down", and "unknown"
- **name** (String) Name of Interface. The name should be up to 255 characters

Optional:

- **address** (String) The Address to assign to this interface. The address should in the IPv4 or IPv6 format
- **cidr_block** (Number) CIDR BLOCK of the Interface
- **default_gateway_address** (String) Default Gateway Address of the Interface. The Default Gateway Address should in the IPv4 or IPv6 format
- **friendly_name** (String) User-friendly name for interface. The name should be up to 255 characters

Read-Only:

- **loopback** (Boolean) Loopback for Workload Interface
- **network** (Map of String) Href of Network of the Interface
- **network_detection_mode** (String) Network Detection Mode of the Interface -->


<a id="nestedblock--labels"></a>
### Nested Schema for `labels`

Required:

- **href** (String) URI of label


<a id="nestedatt--container_cluster"></a>
### Nested Schema for `container_cluster`

Read-Only:

- **href** (String) URI of container cluster
- **name** (String) Name of container cluster

<a id="nestedatt--detected_vulnerabilities"></a>
### Nested Schema for `detected_vulnerabilities`

Read-Only:

- **ip_address** (String) The IP address of the host where the vulnerability is found
- **port** (Number) The port that is associated with the vulnerability
- **port_exposure** (Number) The exposure of the port based on the current policy
- **port_wide_exposure** (List of Object) High end of an IP range(see [below for nested schema](#nestedobjatt--detected_vulnerabilities--port_wide_exposure))
- **proto** (Number) The protocol that is associated with the vulnerability
- **vulnerability** (Set of Object) Vulnerability for Workload (see [below for nested schema](#nestedobjatt--detected_vulnerabilities--vulnerability))
- **vulnerability_report** (Set of Object) Vulnerability Report for Workload(see [below for nested schema](#nestedobjatt--detected_vulnerabilities--vulnerability_report))
- **workload** (List of Object) URI of Workload (see [below for nested schema](#nestedobjatt--detected_vulnerabilities--workload))

<a id="nestedobjatt--detected_vulnerabilities--port_wide_exposure"></a>
### Nested Schema for `detected_vulnerabilities.port_wide_exposure`

Read-Only:

- **any** (Boolean) The boolean value representing if at least one port is exposed to internet (any rule) on the workload
- **ip_list** (Boolean) The boolean value representing if at least one port is exposed to ip_list(s) on the workload


<a id="nestedobjatt--detected_vulnerabilities--vulnerability"></a>
### Nested Schema for `detected_vulnerabilities.vulnerability`

Read-Only:

- **href** (String) The URI of the workload to which this vulnerability belongs to
- **name** (String) The title/name of the vulnerability
- **score** (Number) The normalized score of the vulnerability within the range of 0 to 100


<a id="nestedobjatt--detected_vulnerabilities--vulnerability_report"></a>
### Nested Schema for `detected_vulnerabilities.vulnerability_report`

Read-Only:

- **href** (String) The URI of the report to which this vulnerability belongs to


<a id="nestedobjatt--detected_vulnerabilities--workload"></a>
### Nested Schema for `detected_vulnerabilities.workload`

Read-Only:

- **href** (String) The URI of the workload to which this vulnerability belongs to



<a id="nestedatt--firewall_coexistence"></a>
### Nested Schema for `firewall_coexistence`

Read-Only:

- **illumio_primary** (Boolean) Illumio is the primary firewall if set to true


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
- **package** (String) The RPM/DEB package that the program is part of
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


