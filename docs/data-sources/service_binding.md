---
layout: "illumio-core"
page_title: "illumio-core_service_binding Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-service-binding"
subcategory: ""
description: |-
  Represents Illumio Service Binding
---

# illumio-core_service_binding (Data Source)

Represents Illumio Service Binding

-> the `illumio-core_service_binding` data source must reference a virtual service that has already been provisioned and is in the `active` state  

Example Usage
------------

```hcl
data "illumio-core_service_binding" "example" {
   href = illumio-core_service_binding.example.href
}

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

- `href` (String) URI of the Service Binding

### Read-Only

- `bound_service` (Map of String) Bound service href
- `container_workload` (Map of String) Container Workload href
- `external_data_reference` (String) External Data reference identifier
- `external_data_set` (String) External Data Set identifier
- `port_overrides` (List of Object) Port Overrides for Service Bindings (see [below for nested schema](#nestedatt--port_overrides))
- `virtual_service` (Map of String) Virtual service href
- `workload` (List of Object) Workload Object for Service Bindings (see [below for nested schema](#nestedatt--workload))

<a id="nestedatt--port_overrides"></a>
### Nested Schema for `port_overrides`

Read-Only:

- `new_port` (Number) Overriding port number (or starting point when specifying a range)
- `new_to_port` (Number) Overriding port range ending port
- `port` (Number) Port Number in the original service which to override (integer 0-65535). Starting port when specifying a range
- `proto` (Number) Transport protocol in the original service which to override


<a id="nestedatt--workload"></a>
### Nested Schema for `workload`

Read-Only:

- `deleted` (Boolean) Determines whether the workload is deleted
- `hostname` (String) Workload Hostname
- `href` (String) Workload URI
- `name` (String) Workload Name
