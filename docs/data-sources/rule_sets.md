---
layout: "illumio-core"
page_title: "illumio-core_rule_sets Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-rule-sets"
subcategory: ""
description: |-
  Represents Illumio Rulesets
---

# illumio-core_rule_sets (Data Source)

Represents Illumio Rulesets

```hcl
data "illumio-core_rule_sets" "example" {
  max_results = "5"
  labels = jsonencode([
    [
      {
        href = "/orgs/1/labels/12"
      }
    ]
  ])
}
```

## Schema

### Optional

- **description** (String) Description of Ruleset(s) to return. Supports partial matches
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **labels** (String) List of lists of label URIs, encoded as a JSON string
- **max_results** (String) Maximum number of Rulesets to return. The integer should be a non-zero positive integer
- **name** (String) Name of Ruleset(s) to return. Supports partial matches
- **pversion** (String) pversion of the security policy. Allowed values are "draft", "active", and numbers greater than 0. Default value: "draft"

### Read-Only

- **items** (List of Object) list of Rulesets (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **created_at** (String) Timestamp when this ruleset was first created
- **created_by** (Map of String) User who created this resource
- **deleted_at** (String) Timestamp when this ruleset was deleted
- **deleted_by** (Map of String) User who deleted this resource
- **description** (String) Description of Ruleset
- **enabled** (Boolean) Enabled flag. Determines whether the Ruleset is enabled or not
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **href** (String) URI of ruleset
- **ip_tables_rules** (List of Object) (see [below for nested schema](#nestedobjatt--items--ip_tables_rules))
- **name** (String)
- **rules** (List of Object) Collection of IP Tables Rules (see [below for nested schema](#nestedobjatt--items--rules))
- **scopes** (List of List of Object) scopes of Ruleset
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this ruleset was last updated
- **updated_by** (Map of String) User who last updated this resource

<a id="nestedobjatt--items--ip_tables_rules"></a>
### Nested Schema for `items.ip_tables_rules`

Read-Only:

- **actors** (List of Object) actors of IP Table Rule (see [below for nested schema](#nestedobjatt--items--ip_tables_rules--actors))
- **created_at** (String) Timestamp when this IP Table Rule was first created
- **created_by** (Map of String) User who created this IP Table Rule
- **deleted_at** (String) Timestamp when this IP Table Rule was deleted
- **deleted_by** (Map of String) User who deleted this IP Table Rule
- **description** (String) Description of the Ip Tables Rule
- **enabled** (Boolean) Enabled flag. Determines whether this IP Tables Rule is enabled or not
- **href** (String) URI of the Ip Tables Rule
- **ip_version** (String) IP version for the rules to be applied to
- **statements** (List of Object) statements for IP Tables Rule (see [below for nested schema](#nestedobjatt--items--ip_tables_rules--statements))
- **update_type** (String) Type of update for IP Table Rule
- **updated_at** (String) Timestamp when this IP Table Rule was last updated
- **updated_by** (Map of String) User who last updated this IP Table Rule

<a id="nestedobjatt--items--ip_tables_rules--actors"></a>
### Nested Schema for `items.ip_tables_rules.actors`

Read-Only:

- **actors** (String) actors of IP table Rule actors
- **label** (Map of String) Href of Label
- **label_group** (Map of String) Href of Label Group
- **workload** (Map of String) Href of Workload


<a id="nestedobjatt--items--ip_tables_rules--statements"></a>
### Nested Schema for `items.ip_tables_rules.statements`

Read-Only:

- **chain_name** (String) chain name of statement
- **parameters** (String) parameters of statement
- **table_name** (String) table name of statement 



<a id="nestedobjatt--items--rules"></a>
### Nested Schema for `items.rules`

Read-Only:

- **consumers** (List of Object) consumers of Security Rule (see [below for nested schema](#nestedobjatt--items--rules--consumers))
- **created_at** (String) Timestamp when this security rule was first created
- **created_by** (Map of String) User who created this security rule
- **deleted_at** (String) Timestamp when this security rule was deleted
- **deleted_by** (Map of String) User who deleted this security rule
- **description** (String) Description of Security Rule
- **enabled** (Boolean) Enabled flag. Determines whether this rule will be enabled in ruleset or not
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **href** (String) URI of Security Rule
- **ingress_services** (List of Object) Collection of Ingress Services (see [below for nested schema](#nestedobjatt--items--rules--ingress_services))
- **machine_auth** (Boolean)
- **providers** (List of Object) providers of Security Rule (see [below for nested schema](#nestedobjatt--items--rules--providers))
- **resolve_labels_as** (List of Object) resolve_label_as of Security rule (see [below for nested schema](#nestedobjatt--items--rules--resolve_labels_as))
- **sec_connect** (Boolean) Determines whether a secure connection is established
- **stateless** (Boolean) Determines whether packet filtering is stateless for the rule
- **unscoped_consumers** (Boolean) Set the scope for rule consumers to All
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this security rule was last updated
- **updated_by** (Map of String) User who last updated this security rule

<a id="nestedobjatt--items--rules--consumers"></a>
### Nested Schema for `items.rules.consumers`

Read-Only:

- **actors** (String) actors of consumers actors
- **ip_list** (Map of String) Href of IP List
- **label** (Map of String) Href of Label
- **label_group** (Map of String) Href of Label Group
- **virtual_service** (Map of String) Href of Virtual Service
- **workload** (Map of String) Href of Workload


<a id="nestedobjatt--items--rules--ingress_services"></a>
### Nested Schema for `items.rules.ingress_services`

Read-Only:

- **href** (String) URI of service
- **port** (Number) Protocol number
- **proto** (Number) Port number used with protocol. Also, the starting port when specifying a range
- **to_port** (Number) High end of port range inclusive if specifying a range


<a id="nestedobjatt--items--rules--providers"></a>
### Nested Schema for `items.rules.providers`

Read-Only:

- **actors** (String) actors of provider
- **ip_list** (Map of String) Href of IP List
- **label** (Map of String) Href of Label
- **label_group** (Map of String) Href of Label Group
- **virtual_server** (Map of String) Href of Virtual Server
- **virtual_service** (Map of String) Href of Virtual Service
- **workload** (Map of String) Href of Workload


<a id="nestedobjatt--items--rules--resolve_labels_as"></a>
### Nested Schema for `items.rules.resolve_labels_as`

Read-Only:

- **consumers** (List of String) consumers of resolve_labels_as
- **providers** (List of String) providers of resolve_labels_as


