---
layout: "illumio-core"
page_title: "illumio-core_selective_enforcement_rule Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-selective-enforcement-rule"
subcategory: ""
description: |-
  Manages Illumio Selective Enforcement Rule
---

# illumio-core_selective_enforcement_rule (Resource)

Manages Illumio Selective Enforcement Rule

Example Usage
------------

```hcl
resource "illumio-core_selective_enforcement_rule" "example" {
  name = "SER rule 1"
  scope {
    label {
      href = "/orgs/1/labels/69"
    }
    label {
      href = "/orgs/1/labels/294"
    }
    label_group {
      href = "/orgs/1/sec_policy/draft/label_groups/523f5cd0-2126-4b30-bb40-a9fd19429dd7"
    }

  }

  enforced_services {
    href = "/orgs/1/sec_policy/draft/services/3"
  }
}
```


## Schema





### Required

- **enforced_services** (Block Set, Min: 1) Collection of services that are enforced (see [below for nested schema](#nestedblock--enforced_services))
- **name** (String) Name of the selective enforcement rule
- **scope** (Block Set, Max: 1) Scope of Selective Enforcement Rule (see [below for nested schema](#nestedblock--scope))

### Optional

### Read-Only

- **created_at** (String) Timestamp when this rule set was first created
- **created_by** (Map of String) User who originally created this rule set
- **href** (String) URI of the selective enforcement rule
- **updated_at** (String) Timestamp when this rule set was last updated
- **updated_by** (Map of String) User who last updated this rule set

<a id="nestedblock--enforced_services"></a>
### Nested Schema for `enforced_services`

Required:

- **href** (String) URI of enforced service.


<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- **label** (Map of String) Href of Label
- **label_group** (Map of String) Href of Label Group


