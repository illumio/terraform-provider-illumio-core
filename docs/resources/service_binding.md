---
layout: "illumio-core"
page_title: "illumio-core_service_binding Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-service-binding"
subcategory: ""
description: |-
  Manages Illumio Service Binding
---

# illumio-core_service_binding (Resource)

Manages Illumio Service Binding

-> the `illumio-core_service_binding` resource must reference a virtual service that has already been provisioned and is in the `active` state  

!> Any updates to this resource once created will be ignored  

Example Usage
------------

```hcl
resource "illumio-core_service_binding" "example" {
  virtual_service {
    href = illumio-core_virtual_service.example.href
  }

  workload {
    href = illumio-core_unmanaged_workload.example.href
  }
}

resource "illumio-core_virtual_service" "example" {
  ...
}

resource "illumio-core_unmanaged_workload" "example" {
  ...
}
```

## Schema

### Required

- `virtual_service` (Block List, Min: 1, Max: 1) Virtual service href (see [below for nested schema](#nestedblock--virtual_service))

### Optional

- `container_workload` (Block List, Max: 1) Container Workload href (see [below for nested schema](#nestedblock--container_workload))
- `external_data_reference` (String) External Data reference identifier
- `external_data_set` (String) External Data Set identifier
- `port_overrides` (Block Set) Port Overrides for Service Bindings (see [below for nested schema](#nestedblock--port_overrides))
- `workload` (Block List, Max: 1) Workload Object for Service Bindings (see [below for nested schema](#nestedblock--workload))

### Read-Only

- `bound_service` (Map of String) Bound service href
- `href` (String) URI of the Service Binding

<a id="nestedblock--virtual_service"></a>
### Nested Schema for `virtual_service`

Required:

- `href` (String) Virtual Service URI

<a id="nestedblock--container_workload"></a>
### Nested Schema for `container_workload`

Required:

- `href` (String) Container Workload URI

<a id="nestedblock--port_overrides"></a>
### Nested Schema for `port_overrides`

Required:

- `new_port` (Number) Overriding port number (or starting point when specifying a range). Allowed range is 0 - 65535

Optional:

- `new_to_port` (Number) Overriding port range ending port. Allowed range is 0 - 65535
- `port` (Number) Port Number in the original service which to override (integer 0-65535). Starting port when specifying a range
- `proto` (Number) Transport protocol in the original service which to override. Allowed values are 6 (TCP) and 17 (UDP)

<a id="nestedblock--workload"></a>
### Nested Schema for `workload`

Required:

- `href` (String) Workload URI

Read-Only:

- `deleted` (Boolean) Determines whether the workload is deleted
- `hostname` (String) Workload Hostname
- `name` (String) Workload Name
