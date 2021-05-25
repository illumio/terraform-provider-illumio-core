---
layout: "illumio-core"
page_title: "illumio-core_label_group Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-label-group"
subcategory: ""
description: |-
  Manages Illumio Label Group
---

# illumio-core_label_group (Resource)

Manages Illumio Label Group

```hcl
resource "illumio-core_label_group" "example" {
  key           = "role"
  name          = "example role"
  description   = "example desc"
  labels {
    href = "/orgs/1/labels/1"
  }
  labels {
    href = "/orgs/1/labels/7"
  }
}
```


## Schema

### Required

- **description** (String) The long description of the label group
- **key** (String) Key in key-value pair of contained labels or label groups. Allowed values are "role", "loc", "app" and "env"
- **name** (String) Name of the label group

### Optional

- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **labels** (Block Set) Contained labels (see [below for nested schema](#nestedblock--labels))
- **sub_groups** (Block Set) Contained label groups (see [below for nested schema](#nestedblock--sub_groups))

### Read-Only

- **created_at** (String) Timestamp when this label group was first created
- **created_by** (Map of String) User who created this label group
- **deleted_at** (String) Timestamp when this label group was last deleted
- **deleted_by** (Map of String) User who deleted this label group
- **href** (String) URI of this label group
- **updated_at** (String) Timestamp when this label group was last updated
- **updated_by** (Map of String) User who last updated this label group

<a id="nestedblock--labels"></a>
### Nested Schema for `labels`

Required:

- **href** (String) URI of label

Read-Only:

- **key** (String) Key in key-value pair
- **value** (String) Value in key-value pair


<a id="nestedblock--sub_groups"></a>
### Nested Schema for `sub_groups`

Required:

- **href** (String) URI of label group

Read-Only:

- **name** (String) Key in key-value pair