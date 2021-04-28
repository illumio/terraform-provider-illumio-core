---
layout: "illumio-core"
page_title: "illumio-core_selective_enforcement_rule Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-selective-enforcement-rule"
subcategory: ""
description: |-
  Represents Illumio Selective Enforcement Rule
---

# illumio-core_selective_enforcement_rule (Data Source)

Represents Illumio Selective Enforcement Rule

Example Usage
------------

```hcl
data "illumio-core_selective_enforcement_rule" "rule15"{
    ser_id = 15
}
```


## Schema



### Required

- **ser_id** (Number) selective enforcement rule id


### Read-Only

- **created_at** (String) Timestamp when this rule set was first created
- **created_by** (Map of String) User who originally created this rule set
- **enforced_services** (List of Object) Collection of services that are enforced (see [below for nested schema](#nestedatt--enforced_services))
- **href** (String) URI of the selective enforcement rule
- **name** (String) Name of the selective enforcement rule
- **scope** (List of Object) Scope of Selective Enforcement Rule (see [below for nested schema](#nestedatt--scope))
- **updated_at** (String) Timestamp when this rule set was last updated
- **updated_by** (Map of String) User who last updated this rule set

<a id="nestedatt--enforced_services"></a>
### Nested Schema for `enforced_services`

Read-Only:

- **href** (String) URI of enforced service.


<a id="nestedatt--scope"></a>
### Nested Schema for `scope`

Read-Only:

- **label** (Map of String) Href of Label
- **label_group** (Map of String) Href of Label Group


