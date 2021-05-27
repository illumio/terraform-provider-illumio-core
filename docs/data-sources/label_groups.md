---
layout: "illumio-core"
page_title: "illumio-core_label_groups Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-label-groups"
subcategory: ""
description: |-
  Represents Illumio Label Groups
---

# illumio-core_label_groups (Data Source)

Represents Illumio Label Groups

Example Usage
------------

```hcl
data "illumio-core_label_groups" "example" {
  max_results = "5"
}
```

## Schema

### Optional

- **description** (String) The long description of the label group
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **key** (String) Key in key-value pair of contained labels or label groups. Allowed values for key are "role", "loc", "app" and "env"
- **max_results** (String) Maximum number of Labels to return. The integer should be a non-zero positive integer
- **name** (String) Name of Label Group(s) to return. Supports partial matches
- **pversion** (String) pversion of the security policy. Allowed values are "draft", "active", and numbers greater than 0. Default value: "draft"
- **usage** (String) Include label usage flags as well

### Read-Only

- **items** (List of Object) List of label group hrefs (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:


- **created_at** (String) Timestamp when this Label Group was first created
- **created_by** (Map of String) User who created this Label Group
- **deleted_at** (String) Timestamp when this Label Group was last deleted
- **deleted_by** (Map of String) User who deleted this Label Group
- **description** (String) The long description of the Label Group
- **external_data_reference** (String) External Data reference identifier
- **external_data_set** (String) External Data set Identifier
- **href** (String) URI of Label Group
- **key** (String) Key in key-value pair of contained labels or Label Groups
- **labels** (List of Object) Contained labels (see [below for nested schema](#nestedobjatt--items--labels))
- **name** (String) Name of the Label Group
- **sub_groups** (List of Object)Contained Label Group (see [below for nested schema](#nestedobjatt--items--sub_groups))
- **update_type** (String) Type of Update
- **updated_at** (String) Timestamp when this Label Group was last updated
- **updated_by** (Map of String) User who last updated this Label Group

<a id="nestedobjatt--items--labels"></a>
### Nested Schema for `items.labels`

Read-Only:

- **href** (String) URI of label
- **key** (String) Label Key same as label group key
- **value** (String) Label Value in key-value pair


<a id="nestedobjatt--items--sub_groups"></a>
### Nested Schema for `items.sub_groups`

Read-Only:

- **href** (String) URI of Label Group
- **name** (String) name of sub Label Group


