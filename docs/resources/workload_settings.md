---
layout: "illumio-core"
page_title: "illumio-core_workload_settings Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-workload-settings"
subcategory: ""
description: |-
  Manages Illumio Workload Settings
---

# illumio-core_workload_settings (Resource)

Manages Illumio Workload Settings

## Importing  

The `workload_settings` resource cannot be created and must be imported using the command below. This resource determines the URI for import from the provider configuration.  

```sh
$ terraform import illumio-core_workload_settings.example placeholder
```

After import, configuration changes can be planned and applied as normal.  

Example Usage
------------

```hcl
resource "illumio-core_workload_settings" "example" {
  workload_disconnected_timeout_seconds {
    value = -1
  }
  workload_goodbye_timeout_seconds {
    value = -1
  }
}
```

## Schema

### Required

- `workload_disconnected_timeout_seconds` (Block Set, Min: 1) Workload Disconnected Timeout Seconds for Workload Settings (see [below for nested schema](#nestedblock--workload_disconnected_timeout_seconds))
- `workload_goodbye_timeout_seconds` (Block Set, Min: 1) Workload Goodbye Timeout Seconds for Workload Settings (see [below for nested schema](#nestedblock--workload_goodbye_timeout_seconds))


### Read-Only

- `href` (String) URI of the Workload Settings

<a id="nestedblock--workload_disconnected_timeout_seconds"></a>
### Nested Schema for `workload_disconnected_timeout_seconds`

Optional:

- `scope` (Block Set) Assigned labels for Workload Disconnected Timeout Seconds (see [below for nested schema](#nestedblock--workload_disconnected_timeout_seconds--scope))
- `value` (Number) Property value associated with the scope. Allowed range is 300 - 2147483647 or -1

<a id="nestedblock--workload_disconnected_timeout_seconds--scope"></a>
### Nested Schema for `workload_disconnected_timeout_seconds.scope`

Optional:

- `href` (String) Label URI



<a id="nestedblock--workload_goodbye_timeout_seconds"></a>
### Nested Schema for `workload_goodbye_timeout_seconds`

Optional:

- `scope` (Block Set) Assigned labels for Workload Goodbye Timeout Seconds (see [below for nested schema](#nestedblock--workload_goodbye_timeout_seconds--scope))
- `value` (Number) Property value associated with the scope. Allowed range is 300 - 2147483647 or -1

<a id="nestedblock--workload_goodbye_timeout_seconds--scope"></a>
### Nested Schema for `workload_goodbye_timeout_seconds.scope`

Optional:

- `href` (String) Label URI
