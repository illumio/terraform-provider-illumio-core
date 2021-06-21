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
  name = "example name"

  ip_tables_rules {
    description = "example desc"
    enabled = true
    ip_version = 4
    actors {
      actors = "ams"
    }
    actors {
      label {
        href = "/orgs/1/labels/69"
      }
    }
    statements {
      table_name = "nat"
      chain_name = "PREROUTING"
      parameters = "value"
    }
  }

  scopes {
    label {
      href = "/orgs/1/labels/69"
    }
    label {
      href = "/orgs/1/labels/94"
    }
    label_group {
      href = "/orgs/1/sec_policy/draft/label_groups/65d0ad0f-329a-4ddc-8919-bd0220051fc7"
    }
  }

  rules {
    enabled       = true
    resolve_labels_as {
      consumers = ["workloads"]
      providers = ["workloads"]
    }
    consumers {
      actors = "ams"
    }
    providers {
      label {
        href = "/orgs/1/labels/715"
      }
    }
    ingress_services {
      proto = 6
      port  = 1
    }
  }
}
```

## Schema

### Required

- **name** (String) Name of Ruleset. Valid name should be between 1 to 255 characters
- **scopes** (Block List, Min: 1) scopes for Ruleset. At most 3 blocks of label/label_group can be specified inside each scope block (see [below for nested schema](#nestedblock--scopes))

### Optional

- **description** (String) Description of Ruleset
- **enabled** (Boolean) Enabled flag. Determines whether the Ruleset is enabled or not. Default value: true
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **ip_tables_rules** (Block Set) Collection of IP Tables Rules (see [below for nested schema](#nestedblock--ip_tables_rules))
- **rules** (Block Set) Collection of Security Rules (see [below for nested schema](#nestedblock--rules))

### Read-Only

- **created_at** (String) Timestamp when this ruleset was first created
- **created_by** (Map of String) User who created this ruleset
- **deleted_at** (String) Timestamp when this ruleset was deleted
- **deleted_by** (Map of String) User who deleted this ruleset
- **href** (String) URI of Ruleset
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this ruleset was last updated
- **updated_by** (Map of String) User who last updated this ruleset

<a id="nestedblock--scopes"></a>
### Nested Schema for `scopes`

Optional:

- **label** (Block Set) Href of Label (see [below for nested schema](#nestedblock--scopes--label))
- **label_group** (Block Set) Href of Label Group (see [below for nested schema](#nestedblock--scopes--label_group))

<a id="nestedblock--scopes--label"></a>
### Nested Schema for `scopes.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--scopes--label_group"></a>
### Nested Schema for `scopes.label_group`

Required:

- **href** (String) URI of Label Group



<a id="nestedblock--ip_tables_rules"></a>
### Nested Schema for `ip_tables_rules`

Required:

- **actors** (Block Set, Min: 1) actors for IP Table Rule (see [below for nested schema](#nestedblock--ip_tables_rules--actors))
- **enabled** (Boolean) Enabled flag. Determines whether this IP Tables Rule is enabled or not
- **ip_version** (String) IP version for the rules to be applied to. Allowed values are "4" and "6"
- **statements** (Block Set, Min: 1) statements for this IP Tables Rule (see [below for nested schema](#nestedblock--ip_tables_rules--statements))

Optional:

- **description** (String) Description of the IP Tables Rules

Read-Only:

- **created_at** (String) Timestamp when this IP Table Rule was first created
- **created_by** (Map of String) User who created this IP Table Rule
- **deleted_at** (String) Timestamp when this IP Table Rule was deleted
- **deleted_by** (Map of String) User who deleted this IP Table Rule
- **href** (String) URI of the Ip Tables Rules
- **update_type** (String) Type of update for IP Table Rule
- **updated_at** (String) Timestamp when this IP Table Rule was last updated
- **updated_by** (Map of String) User who last updated this IP Table Rule

<a id="nestedblock--ip_tables_rules--actors"></a>
### Nested Schema for `ip_tables_rules.actors`

Optional:

- **actors** (String) Set this if rule actors are all workloads. Allowed value is "ams"
- **label** (Block Set, Max: 1) Href of Label (see [below for nested schema](#nestedblock--ip_tables_rules--actors--label))
- **label_group** (Block Set, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--ip_tables_rules--actors--label_group))
- **workload** (Block Set, Max: 1) Href of Workload (see [below for nested schema](#nestedblock--ip_tables_rules--actors--workload))

<a id="nestedblock--ip_tables_rules--actors--label"></a>
### Nested Schema for `ip_tables_rules.actors.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--ip_tables_rules--actors--label_group"></a>
### Nested Schema for `ip_tables_rules.actors.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--ip_tables_rules--actors--workload"></a>
### Nested Schema for `ip_tables_rules.actors.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--ip_tables_rules--statements"></a>
### Nested Schema for `ip_tables_rules.statements`

Required:

- **chain_name** (String) Chain name for statement. Allowed values are "PREROUTING", "INPUT" and "OUTPUT"
- **parameters** (String) Parameters of statements
- **table_name** (String) Name of the table. Allowed values are "nat", "mangle" and "filter"



<a id="nestedblock--rules"></a>
### Nested Schema for `rules`

Required:

- **consumers** (Block Set, Min: 1) Consumers for Security Rule. Only one actor can be specified in one consumers block (see [below for nested schema](#nestedblock--rules--consumers))
- **enabled** (Boolean) Enabled flag. Determines whether the rule will be enabled in ruleset or not
- **providers** (Block Set, Min: 1) providers for Security Rule. Only one actor can be specified in one providers block (see [below for nested schema](#nestedblock--rules--providers))
- **resolve_labels_as** (Block List, Min: 1, Max: 1) resolve label as for Security rule (see [below for nested schema](#nestedblock--rules--resolve_labels_as))

Optional:

- **description** (String) Description of Security Rule
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **ingress_services** (Block Set) Collection of Ingress Service. If resolve_label_as.providers list includes "workloads" then ingress_service is required. Only one of the {"href"} or {"proto", "port", "to_port"} parameter combination is allowed (see [below for nested schema](#nestedblock--rules--ingress_services))
- **machine_auth** (Boolean) Determines whether machine authentication is enabled. Defaule Value: false
- **sec_connect** (Boolean) Determines whether a secure connection is established. Defaule Value: false
- **stateless** (Boolean) Determines whether packet filtering is stateless for the rule. Defaule Value: false
- **unscoped_consumers** (Boolean) Set the scope for rule consumers to All. Defaule Value: false

Read-Only:

- **created_at** (String) Timestamp when this security rule was first created
- **created_by** (Map of String) User who created this security rule
- **deleted_at** (String) Timestamp when this security rule was deleted
- **deleted_by** (Map of String) User who deleted this security rule
- **href** (String) URI of Security Rule
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this security rule was last updated
- **updated_by** (Map of String) User who last updated this security rule

<a id="nestedblock--rules--consumers"></a>
### Nested Schema for `rules.consumers`

Optional:

- **actors** (String) actors for consumers parameter. Allowed values are "ams" and "container_host"
- **ip_list** (Block Set, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--rules--consumers--ip_list))
- **label** (Block Set, Max: 1) Href of Label (see [below for nested schema](#nestedblock--rules--consumers--label))
- **label_group** (Block Set, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--rules--consumers--label_group))
- **virtual_service** (Block Set, Max: 1) Href of Virtual Service (see [below for nested schema](#nestedblock--rules--consumers--virtual_service))
- **workload** (Block Set, Max: 1) Href of Workload (see [below for nested schema](#nestedblock--rules--consumers--workload))

<a id="nestedblock--rules--consumers--ip_list"></a>
### Nested Schema for `rules.consumers.ip_list`

Required:

- **href** (String) URI of IP List


<a id="nestedblock--rules--consumers--label"></a>
### Nested Schema for `rules.consumers.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--rules--consumers--label_group"></a>
### Nested Schema for `rules.consumers.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--rules--consumers--virtual_service"></a>
### Nested Schema for `rules.consumers.virtual_service`

Required:

- **href** (String) URI of Virtual Service


<a id="nestedblock--rules--consumers--workload"></a>
### Nested Schema for `rules.consumers.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--rules--providers"></a>
### Nested Schema for `rules.providers`

Optional:

- **actors** (String) actors for provider. Allowed value is "ams"
- **ip_list** (Block Set, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--rules--providers--ip_list))
- **label** (Block Set, Max: 1) Href of Label (see [below for nested schema](#nestedblock--rules--providers--label))
- **label_group** (Block Set, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--rules--providers--label_group))
- **virtual_server** (Block Set, Max: 1) Href of Virtual Server (see [below for nested schema](#nestedblock--rules--providers--virtual_server))
- **virtual_service** (Block Set, Max: 1) Href of Virtual Service (see [below for nested schema](#nestedblock--rules--providers--virtual_service))
- **workload** (Block Set, Max: 1) Href of Workload (see [below for nested schema](#nestedblock--rules--providers--workload))

<a id="nestedblock--rules--providers--ip_list"></a>
### Nested Schema for `rules.providers.ip_list`

Required:

- **href** (String) URI of IP List


<a id="nestedblock--rules--providers--label"></a>
### Nested Schema for `rules.providers.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--rules--providers--label_group"></a>
### Nested Schema for `rules.providers.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--rules--providers--virtual_server"></a>
### Nested Schema for `rules.providers.virtual_server`

Required:

- **href** (String) URI of Virtual Server


<a id="nestedblock--rules--providers--virtual_service"></a>
### Nested Schema for `rules.providers.virtual_service`

Required:

- **href** (String) URI of Virtual Service


<a id="nestedblock--rules--providers--workload"></a>
### Nested Schema for `rules.providers.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--rules--resolve_labels_as"></a>
### Nested Schema for `rules.resolve_labels_as`

Required:

- **consumers** (Set of String) consumers for resolve_labels_as. Allowed values are "workloads", "virtual_services"
- **providers** (Set of String) providers for resolve_labels_as. Allowed values are "workloads", "virtual_services"


<a id="nestedblock--rules--ingress_services"></a>
### Nested Schema for `rules.ingress_services`

Optional:

- **href** (String) URI of Service
- **port** (String) Port number used with protocol or starting port when specifying a range. Allowed range is 0-65535
- **proto** (String) Protocol number. Allowed values are 6 (TCP) and 17 (UDP)
- **to_port** (String) Upper end of port range. Allowed range is 0-65535

