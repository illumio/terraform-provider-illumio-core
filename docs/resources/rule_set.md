---
layout: "illumio-core"
page_title: "illumio-core_rule_set Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-rule-set"
subcategory: ""
description: |-
  Manages Illumio Ruleset
---

# illumio-core_rule_set (Resource)

Manages Illumio Ruleset

Example Usage
------------

```hcl
resource "illumio-core_rule_set" "example" {
  name = "RS-EXAMPLE"

  scopes {
    label {
      href = illumio-core_label.example.href
    }
  }
}

resource "illumio-core_label" "example" {
  ...
}
```

## Schema

### Required

- `name` (String) Name of Ruleset. Valid name should be between 1 to 255 characters
- `scopes` (Block List, Min: 1) scopes for Ruleset. At most 3 blocks of label/label_group can be specified inside each scope block (see [below for nested schema](#nestedblock--scopes))

### Optional

- `description` (String) Description of Ruleset
- `enabled` (Boolean) Enabled flag. Determines whether the Ruleset is enabled or not. Default value: true
- `external_data_reference` (String) External data reference identifier
- `external_data_set` (String) External data set identifier
- `ip_tables_rules` (Block Set) Collection of IP Tables Rules (see [below for nested schema](#nestedblock--ip_tables_rules))

### Read-Only

- `created_at` (String) Timestamp when this ruleset was first created
- `created_by` (Map of String) User who created this ruleset
- `deleted_at` (String) Timestamp when this ruleset was deleted
- `deleted_by` (Map of String) User who deleted this ruleset
- `href` (String) URI of Ruleset
- `update_type` (String) Type of update
- `updated_at` (String) Timestamp when this ruleset was last updated
- `updated_by` (Map of String) User who last updated this ruleset

<a id="nestedblock--scopes"></a>
### Nested Schema for `scopes`

Optional:

- `label` (Block Set) Href of Label (see [below for nested schema](#nestedblock--scopes--label))
- `label_group` (Block Set) Href of Label Group (see [below for nested schema](#nestedblock--scopes--label_group))

<a id="nestedblock--scopes--label"></a>
### Nested Schema for `scopes.label`

Required:

- `href` (String) URI of Label

<a id="nestedblock--scopes--label_group"></a>
### Nested Schema for `scopes.label_group`

Required:

- `href` (String) URI of Label Group

<a id="nestedblock--ip_tables_rules"></a>
### Nested Schema for `ip_tables_rules`

Required:

- `actors` (Block Set, Min: 1) actors for IP Table Rule (see [below for nested schema](#nestedblock--ip_tables_rules--actors))
- `enabled` (Boolean) Enabled flag. Determines whether this IP Tables Rule is enabled or not
- `ip_version` (String) IP version for the rules to be applied to. Allowed values are "4" and "6"
- `statements` (Block Set, Min: 1) statements for this IP Tables Rule (see [below for nested schema](#nestedblock--ip_tables_rules--statements))

Optional:

- `description` (String) Description of the IP Tables Rules

Read-Only:

- `created_at` (String) Timestamp when this IP Table Rule was first created
- `created_by` (Map of String) User who created this IP Table Rule
- `deleted_at` (String) Timestamp when this IP Table Rule was deleted
- `deleted_by` (Map of String) User who deleted this IP Table Rule
- `href` (String) URI of the Ip Tables Rules
- `update_type` (String) Type of update for IP Table Rule
- `updated_at` (String) Timestamp when this IP Table Rule was last updated
- `updated_by` (Map of String) User who last updated this IP Table Rule

<a id="nestedblock--ip_tables_rules--actors"></a>
### Nested Schema for `ip_tables_rules.actors`

Optional:

- `actors` (String) Set this if rule actors are all workloads. Allowed value is "ams"
- `label` (Block Set, Max: 1) Href of Label (see [below for nested schema](#nestedblock--ip_tables_rules--actors--label))
- `label_group` (Block Set, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--ip_tables_rules--actors--label_group))
- `workload` (Block Set, Max: 1) Href of Workload (see [below for nested schema](#nestedblock--ip_tables_rules--actors--workload))

<a id="nestedblock--ip_tables_rules--actors--label"></a>
### Nested Schema for `ip_tables_rules.actors.label`

Required:

- `href` (String) URI of Label

<a id="nestedblock--ip_tables_rules--actors--label_group"></a>
### Nested Schema for `ip_tables_rules.actors.label_group`

Required:

- `href` (String) URI of Label Group

<a id="nestedblock--ip_tables_rules--actors--workload"></a>
### Nested Schema for `ip_tables_rules.actors.workload`

Required:

- `href` (String) URI of Workload

<a id="nestedblock--ip_tables_rules--statements"></a>
### Nested Schema for `ip_tables_rules.statements`

Required:

- `chain_name` (String) Chain name for statement. Allowed values are "PREROUTING", "INPUT" and "OUTPUT"
- `parameters` (String) Parameters of statements
- `table_name` (String) Name of the table. Allowed values are "nat", "mangle" and "filter"
