---
layout: "illumio-core"
page_title: "illumio-core_labels Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-labels"
subcategory: ""
description: |-
  Represents Illumio Labels
---

# illumio-core_labels (Data Source)

Represents Illumio Labels

Example Usage
------------

```hcl
data "illumio-core_labels" "example" {
  max_results = "5"
}
```

## Schema

### Optional

- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **include_deleted** (String) Include deleted labels
- **key** (String) Key in key-value pair. Allowed values for key are "role", "loc", "app" and "env"
- **max_results** (String) Maximum number of Labels to return. The integer should be a non-zero positive integer
- **usage** (String) Include label usage flags as well
- **value** (String) Value on which to filter. Supports partial matches

### Read-Only

- **items** (List of Object) List of labels (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **created_at** (String) Timestamp when this label was first created
- **created_by** (Map of String) User who created this label
- **deleted** (Boolean) Flag to indicate whether deleted or not
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **href** (String) URI of Label
- **key** (String) Key in key-value pair
- **updated_at** (String) Timestamp when this label was last updated
- **updated_by** (Map of String) User who last updated this label
- **value** (String) Value in key-value pair


