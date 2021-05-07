---
layout: "illumio-core"
page_title: "illumio-core_workload_settings Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-workload-settings"
subcategory: ""
description: |-
  Represents Illumio Workload Settings
---

# illumio-core_workload_settings (Data Source)

Represents Illumio Workload Settings

Example Usage
------------

```hcl
data "illumio-core_workload_settings" "example" {

}
```


## Schema


### Read-Only

- **workload_disconnected_timeout_seconds** (Set of Object) Workload Disconnected Timeout Seconds for Workload Settings (see [below for nested schema](#nestedatt--workload_disconnected_timeout_seconds))
- **workload_goodbye_timeout_seconds** (Set of Object) Workload Goodbye Timeout Seconds for Workload Settings (see [below for nested schema](#nestedatt--workload_goodbye_timeout_seconds))

<a id="nestedatt--workload_disconnected_timeout_seconds"></a>
### Nested Schema for `workload_disconnected_timeout_seconds`

Read-Only:

- **scope** (Set of Object) Assigned labels for Workload Disconnected Timeout Seconds (see [below for nested schema](#nestedobjatt--workload_disconnected_timeout_seconds--scope))
- **value** (Number) Property value associated with the scope

<a id="nestedobjatt--workload_disconnected_timeout_seconds--scope"></a>
### Nested Schema for `workload_disconnected_timeout_seconds.scope`

Read-Only:

- **href** (String) Label URI

<a id="nestedatt--workload_goodbye_timeout_seconds"></a>
### Nested Schema for `workload_goodbye_timeout_seconds`

Read-Only:

- **scope** (Set of Object) Assigned labels for Workload Goodbye Timeout Seconds (see [below for nested schema](#nestedobjatt--workload_goodbye_timeout_seconds--scope))
- **value** (Number) Property value associated with the scope

<a id="nestedobjatt--workload_goodbye_timeout_seconds--scope"></a>
### Nested Schema for `workload_goodbye_timeout_seconds.scope`

Read-Only:

- **href** (String) Label URI


