---
layout: "illumio-core"
page_title: "illumio-core_service_bindings Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-service-bindings"
subcategory: ""
description: |-
  Represents Illumio Service Bindings
---

# illumio-core_service_bindings (Data Source)

Represents Illumio Service Bindings

Example Usage
------------

```hcl
data "illumio-core_service_bindings" "test" {
   max_results = "5"
}
```

## Schema

### Optional

- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **max_results** (String) Maximum number of virtual service bindings to return.
- **virtual_service** (String) Virtual service URI
- **workload** (String) Workload URI

### Read-Only

- **items** (List of Object) list of Service Bindings (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **bound_service** (Map of String) Bound service href
- **container_workload** (Map of String) Container Workload href
- **external_data_reference** (String) External Data reference identifier
- **external_data_set** (String) External Data Set identifier
- **href** (String) URI of Service Binding
- **port_overrides** (Set of Object) Port Overrides for Service Bindings (see [below for nested schema](#nestedobjatt--items--port_overrides))
- **virtual_service** (Map of String) Virtual service href
- **workload** (Set of Object) Workload Object for Service Bindings (see [below for nested schema](#nestedobjatt--items--workload))

<a id="nestedobjatt--items--port_overrides"></a>
### Nested Schema for `items.port_overrides`

Read-Only:

- **new_port** (Number) Overriding port number (or starting point when specifying a range)
- **new_to_port** (Number) Overriding port range ending port
- **port** (Number) Port Number in the original service which to override (integer 0-65535). Starting port when specifying a range
- **proto** (Number) Transport protocol in the original service which to override


<a id="nestedobjatt--items--workload"></a>
### Nested Schema for `items.workload`

Read-Only:

- **deleted** (Boolean) Determines whether the workload is deleted
- **hostname** (String) Workload Hostname
- **href** (String) Workload URI
- **name** (String) Workload Name


