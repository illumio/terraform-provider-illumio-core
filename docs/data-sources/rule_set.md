---
layout: "illumio-core"
page_title: "illumio-core_rule_set Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-rule-set"
subcategory: ""
description: |-
  Represents Illumio Ruleset
---

# illumio-core_rule_set (Data Source)

Represents Illumio Ruleset

Example Usage
------------

```hcl
data "illumio-core_rule_set" "example" {
  href = illumio-core_rule_set.example.href
}

resource "illumio-core_rule_set" "example" {
  ...
}
```

## Schema

### Required

- `href` (String) URI of Ruleset

### Read-Only

- `created_at` (String) Timestamp when this ruleset was first created
- `created_by` (Map of String) User who created this resource
- `deleted_at` (String) Timestamp when this ruleset was deleted
- `deleted_by` (Map of String) User who deleted this resource
- `description` (String) Description of Ruleset
- `enabled` (Boolean) Enabled flag. Determines whether the Ruleset is enabled or not
- `external_data_reference` (String) External data reference identifier
- `external_data_set` (String) External data set identifier
- `ip_tables_rules` (List of Object) Collection of IP Tables Rules (see [below for nested schema](#nestedatt--ip_tables_rules))
- `name` (String) Name of Ruleset
- `rules` (List of Object) Collection of Security Rules (see [below for nested schema](#nestedatt--rules))
- `scopes` (List of List of Object) scopes for Ruleset
- `update_type` (String) Type of update
- `updated_at` (String) Timestamp when this ruleset was last updated
- `updated_by` (Map of String) User who last updated this resource

<a id="nestedatt--ip_tables_rules"></a>
### Nested Schema for `ip_tables_rules`

Read-Only:

- `actors` (List of Object) actors for IP Table Rule (see [below for nested schema](#nestedobjatt--ip_tables_rules--actors))
- `created_at` (String) Timestamp when this ruleset was first created
- `created_by` (Map of String) User who created this ruleset
- `deleted_at` (String) Timestamp when this ruleset was deleted
- `deleted_by` (Map of String) User who deleted this ruleset
- `description` (String) Description of the Ip Tables Rules
- `enabled` (Boolean) Enabled flag. Determines whether this IP Tables Rule is enabled or not
- `href` (String) URI of the Ip Tables Rules
- `ip_version` (String) IP version for the rules to be applied to
- `statements` (List of Object) statements for in this IP Tables Rule (see [below for nested schema](#nestedobjatt--ip_tables_rules--statements))
- `update_type` (String) Type of update
- `updated_at` (String) Timestamp when this ruleset was last updated
- `updated_by` (Map of String) User who last updated this ruleset


<a id="nestedobjatt--ip_tables_rules--actors"></a>
### Nested Schema for `ip_tables_rules.actors`

Read-Only:

- `actors` (String) actors for IP table Rule actors
- `label` (Map of String) Href of Label 
- `label_group` (Map of String) Href of Label Group
- `workload` (Map of String) Href of Workload


<a id="nestedobjatt--ip_tables_rules--statements"></a>
### Nested Schema for `ip_tables_rules.statements`

Read-Only:

- `chain_name` (String) Chain name for statement
- `parameters` (String) Parameters of statements
- `table_name` (String) Name of the table

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Read-Only:

- `consumers` (List of Object) Consumers for Security Rule (see [below for nested schema](#nestedobjatt--rules--consumers))
- `created_at` (String) Timestamp when this security rule was first created
- `created_by` (Map of String) User who created this security rule
- `deleted_at` (String) Timestamp when this security rule was deleted
- `deleted_by` (Map of String) User who deleted this security rule
- `description` (String) Description of Security Rule
- `enabled` (Boolean) Enabled flag. Determines whether this rule will be enabled in ruleset or not
- `external_data_reference` (String) External data reference identifier
- `external_data_set` (String) External data set identifier
- `href` (String) URI of Security Rule
- `ingress_services` (List of Object) Collection of Ingress Services (see [below for nested schema](#nestedobjatt--rules--ingress_services))
- `machine_auth` (Boolean) Determines whether machine authentication is enabled
- `providers` (List of Object) providers for Security Rule (see [below for nested schema](#nestedobjatt--rules--providers))
- `resolve_labels_as` (List of Object) resolve label as for Security rule (see [below for nested schema](#nestedobjatt--rules--resolve_labels_as))
- `sec_connect` (Boolean) Determines whether a secure connection is established
- `stateless` (Boolean) Determines whether packet filtering is stateless for the rule
- `unscoped_consumers` (Boolean) Set the scope for rule consumers to All
- `update_type` (String) Type of update
- `updated_at` (String) Timestamp when this security rule was last updated
- `updated_by` (Map of String) User who last updated this security rule

<a id="nestedobjatt--rules--consumers"></a>
### Nested Schema for `rules.consumers`

Read-Only:

- `actors` (String) actors for consumers
- `ip_list` (Map of String) Href of IP List
- `label` (Map of String) Href of Label
- `label_group` (Map of String) Href of Label Group
- `virtual_service` (Map of String) Href of Virtual Service
- `workload` (Map of String) Href of Workload

<a id="nestedobjatt--rules--ingress_services"></a>
### Nested Schema for `rules.ingress_services`

Read-Only:

- `href` (String) URI of service
- `port` (Number) Protocol number
- `proto` (Number) Port number used with protocol. Also, the starting port when specifying a range
- `to_port` (Number) High end of port range inclusive if specifying a range


<a id="nestedobjatt--rules--providers"></a>
### Nested Schema for `rules.providers`

Read-Only:

- `actors` (String) actors for illumio_provider
- `ip_list` (Map of String) Href of IP List
- `label` (Map of String) Href of Label
- `label_group` (Map of String) Href of Label Group
- `virtual_server` (Map of String) Href of Virtual Server
- `virtual_service` (Map of String) Href of Virtual Service
- `workload` (Map of String) Href of Workload


<a id="nestedobjatt--rules--resolve_labels_as"></a>
### Nested Schema for `rules.resolve_labels_as`

Read-Only:

- `consumers` (List of String) consumers for resolve_labels_as
- `providers` (List of String) providers for resolve_labels_as
