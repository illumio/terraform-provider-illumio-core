---
layout: "illumio-core"
page_title: "illumio-core_label Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-label"
subcategory: ""
description: |-
  Represents Illumio Label
---

# illumio-core_label (Data Source)

Represents Illumio Label

Example Usage
------------

```hcl
data "illumio-core_label" "example" {
  href = illumio-core_label.example.href
}

resource "illumio-core_label" "example" {
  ...
}
```

## Schema

### Required

- `href` (String) URI of this label

### Read-Only

- `created_at` (String) Timestamp when this label was first created
- `created_by` (Map of String) User who created this label
- `deleted` (Boolean) Flag to indicate whether deleted or not
- `external_data_reference` (String) A unique identifier within the external data source
- `external_data_set` (String) The data source from which a resource originates
- `key` (String) Key in key-value pair
- `updated_at` (String) Timestamp when this label was last updated
- `updated_by` (Map of String) User who last updated this label
- `value` (String) Value in key-value pair
