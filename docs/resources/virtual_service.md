---
layout: "illumio-core"
page_title: "illumio-core_virtual_service Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-virtual-service"
subcategory: ""
description: |-
  Manages Illumio Virtual Service
---

# illumio-core_virtual_service (Resource)

Manages Illumio Virtual Service

Example Usage
------------

```hcl
resource "illumio-core_virtual_service" "example" {
  name        = "VS-EXAMPLE"
  description = "Virtual Service example"
  apply_to    = "host_only"

  service {
    href = illumio-core_service.example.href
  }

  labels {
    href = illumio-core_label.example.href
  }
}

resource "illumio-core_service" "example" {
  ...
}

resource "illumio-core_label" "example" {
  ...
}
```

## Schema

### Required

- `apply_to` (String) Name of the virtual service. Allowed values are "host_only" and "internal_bridge_network"
- `name` (String) Name of the virtual service. The name should be between 1 to 255 characters

### Optional

- `description` (String) The long description of the virtual service
- `external_data_reference` (String) A unique identifier within the external data source
- `external_data_set` (String) The data source from which a resource originates
- `ip_overrides` (Set of String) Array of IPs or CIDRs as IP overrides
- `labels` (Block Set) Contained labels (see [below for nested schema](#nestedblock--labels))
- `service` (Block List, Max: 1) Associated service (see [below for nested schema](#nestedblock--service))
- `service_addresses` (Block Set) List of service address. Specify one of the combination {fqdn, description, port}, {ip, network_href} or {ip, port} (see [below for nested schema](#nestedblock--service_addresses))
- `service_ports` (Block Set) URI of associated service (see [below for nested schema](#nestedblock--service_ports))

### Read-Only

- `created_at` (String) Timestamp when this virtual service was first created
- `created_by` (Map of String) User who created this virtual service
- `deleted_at` (String) Timestamp when this virtual service was last deleted
- `deleted_by` (Map of String) User who deleted this virtual service
- `href` (String) URI of Virtual Service
- `pce_fqdn` (String) PCE FQDN for this container cluster. Used in Supercluster only
- `updated_at` (String) Timestamp when this virtual service was last updated
- `updated_by` (Map of String) User who last updated this virtual service

<a id="nestedblock--labels"></a>
### Nested Schema for `labels`

Required:

- `href` (String) URI of label

Read-Only:

- `key` (String) Key in key-value pair
- `value` (String) Value in key-value pair

<a id="nestedblock--service"></a>
### Nested Schema for `service`

Required:

- `href` (String) URI of associated service

<a id="nestedblock--service_addresses"></a>
### Nested Schema for `service_addresses`

Optional:

- `description` (String) Description for given fqdn
- `fqdn` (String) FQDN to assign to the virtual service.  Allowed formats: hostname, IP, or URI
- `ip` (String) IP address to assign to the virtual service
- `network_href` (String) Network URI for this IP address
- `port` (String) Port Number. Also, the starting port when specifying a range. Allowed range is -1 - 65535

<a id="nestedblock--service_ports"></a>
### Nested Schema for `service_ports`

Required:

- `proto` (Number) Transport protocol. Allowed values are 6 (TCP) and 17 (UDP)

Optional:

- `port` (String) Port Number. Also, the starting port when specifying a range. Allowed range is -1 - 65535
- `to_port` (String) High end of port range inclusive if specifying a range. Allowed range is 0 - 65535
