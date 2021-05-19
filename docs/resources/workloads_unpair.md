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

Example Usage
------------

```hcl
resource "illumio-core_workloads_unpair" "example" {
  workloads {
    href = "/orgs/1/workloads/7c3789ea-661b-49c3-b3ba-8eb180f5f3d2"
  }
  workloads {
    href = "/orgs/1/workloads/1d230cf4-6718-44b8-8ffa-64383a4dbee1"
  }
  workloads {
    href = "/orgs/1/workloads/11635f19-625f-436c-a299-43d1883145d5"
  }
}

```

## Schema

### Required

- **workloads** (Block Set, Min: 1) List of Workloads to unpair. Max Items allowed: 1000. (see [below for nested schema](#nestedblock--workloads))

### Optional

- **ip_table_restore** (String) The desired state of IP tables after the agent is uninstalled. Allowed values are "saved", "default" and "disable". Default value: "default"

<a id="nestedblock--workloads"></a>
### Nested Schema for `workloads`

Required:

- **href** (String) URI of Workload


