---
layout: "illumio-core"
page_title: "illumio-core_workloads_unpair Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-workloads-upgrade"
subcategory: ""
description: |-
  Manages Illumio Workloads Unpair
---

# illumio-core_workloads_unpair (Resource)

Manages Illumio Workloads Unpair

!> This resource is deprecated as of v0.2.0 and will be removed in v1.0.0. VENs should be unpaired by deleting the associated `managed_workload` resource

Example Usage
------------

```hcl
resource "illumio-core_workloads_unpair" "example" {
  workloads {
    href = illumio-core_managed_workload.example.href
  }
}

resource "illumio-core_managed_workload" "example" {
  ...
}
```

## Schema

### Required

- `workloads` (Block Set, Min: 1) List of Workloads to unpair. Max Items allowed: 1000 (see [below for nested schema](#nestedblock--workloads))

### Optional

- `ip_table_restore` (String) The desired state of IP tables after the agent is uninstalled. Allowed values are "saved", "default" and "disable". Default value: "default"

<a id="nestedblock--workloads"></a>
### Nested Schema for `workloads`

Required:

- `href` (String) URI of Workload
