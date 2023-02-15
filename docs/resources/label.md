---
layout: "illumio-core"
page_title: "illumio-core_label Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-label"
subcategory: ""
description: |-
  Manages Illumio Label
---
# illumio-core_label (Resource)

Manages Illumio Label

Example Usage
------------

```hcl
resource "illumio-core_label" "role_example" {
  key   = "role"
  value = "R-EXAMPLE"
}
```

## Schema

### Required

- `key` (String) Key in key-value pair. The value must be a string between 1 and 64 characters long
- `value` (String) Value in key-value pair

### Optional

- `external_data_reference` (String) A unique identifier within the external data source
- `external_data_set` (String) The data source from which a resource originates

### Read-Only

- `created_at` (String) Timestamp when this label was first created
- `created_by` (Map of String) User who created this label
- `deleted` (Boolean) Flag to indicate whether deleted or not
- `href` (String) URI of this label
- `updated_at` (String) Timestamp when this label was last updated
- `updated_by` (Map of String) User who last updated this label
