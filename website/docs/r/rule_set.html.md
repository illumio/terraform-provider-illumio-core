---
layout: "illumio-core"
page_title: "illumio-core_rule_set Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-rule-set"
subcategory: ""
description: |-
  Manages Illumio Rule Set
---

# illumio-core_rule_set (Resource)

Manages Illumio Rule Set

Example Usage
------------

```hcl
resource "illumio-core_rule_set" "example" {
  name = "example name"

  ip_tables_rule {
    description = "some desc"
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

  scope {
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

  rule {
    enabled       = true

    resolve_labels_as {
      consumers = ["workloads"]
      providers = ["workloads"]
    }

    consumer {
      actors = "ams"
    }

    illumio_provider {
      label {
        href = "/orgs/1/labels/715"
      }
    }

    illumio_provider {
      label {
        href = "/orgs/1/labels/294"
      }
    }

    ingress_service {
      proto = 6
      port  = "1"
    }
  }
}
```

## Schema

### Required

- **name** (String) Name of Rule Set. Valid name should be in between 1 to 255 characters
- **scope** (Block Set, Min: 1) scopes for Rule Set. At most 3 blocks of label/label_group can be specified inside each scope block (see [below for nested schema](#nestedblock--scope))

### Optional

- **description** (String) Description of Rule Set
- **enabled** (Boolean) Enabled flag. Determines wheter the Rule Set is enabled or not. Default value: true
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **ip_tables_rule** (Block Set) Collection of IP Tables Rules (see [below for nested schema](#nestedblock--ip_tables_rule))
- **rule** (Block Set) Collection of Security Rules (see [below for nested schema](#nestedblock--rule))

### Read-Only

- **created_at** (String) Timestamp when this rule set was first created
- **created_by** (Map of String) User who originally created this rule set
- **deleted_at** (String) Timestamp when this rule set was deleted
- **deleted_by** (Map of String) User who deleted this rule set
- **href** (String) URI of Rule Set
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this rule set was last updated
- **updated_by** (Map of String) User who last updated this rule set

<a id="nestedblock--scope"></a>
### Nested Schema for `scope`

Optional:

- **label** (Block Set) Href of Label (see [below for nested schema](#nestedblock--scope--label))
- **label_group** (Block Set) Href of Label Group (see [below for nested schema](#nestedblock--scope--label_group))

<a id="nestedblock--scope--label"></a>
### Nested Schema for `scope.label`

Optional:

- **href** (String) URI of Label


<a id="nestedblock--scope--label_group"></a>
### Nested Schema for `scope.label_group`

Optional:

- **href** (String) URI of Label Group



<a id="nestedblock--ip_tables_rule"></a>
### Nested Schema for `ip_tables_rule`

Required:

- **actors** (Block Set, Min: 1) actors for IP Table Rule (see [below for nested schema](#nestedblock--ip_tables_rule--actors))
- **enabled** (Boolean) Enabled flag. Determines whether this IP Tables Rule is enabled or not
- **ip_version** (String) IP version for the rules to be applied to. Allowed values are "4" and "6"
- **statements** (Block Set, Min: 1) statements for this IP Tables Rule (see [below for nested schema](#nestedblock--ip_tables_rule--statements))

Optional:

- **description** (String) Description of the IP Tables Rules

Read-Only:

- **created_at** (String) Timestamp when this IP Table Rule was first created
- **created_by** (Map of String) User who originally created this IP Table Rule
- **deleted_at** (String) Timestamp when this IP Table Rule was deleted
- **deleted_by** (Map of String) User who deleted this IP Table Rule
- **href** (String) URI of the Ip Tables Rules
- **update_type** (String) Type of update for IP Table Rule
- **updated_at** (String) Timestamp when this IP Table Rule was last updated
- **updated_by** (Map of String) User who last updated this IP Table Rule

<a id="nestedblock--ip_tables_rule--actors"></a>
### Nested Schema for `ip_tables_rule.actors`

Optional:

- **actors** (String) Set this if rule actors are all workloads. Allowed value: "ams"
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--ip_tables_rule--actors--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--ip_tables_rule--actors--label_group))
- **workload** (Block List, Max: 1) Href of Worklaod (see [below for nested schema](#nestedblock--ip_tables_rule--actors--workload))

<a id="nestedblock--ip_tables_rule--actors--label"></a>
### Nested Schema for `ip_tables_rule.actors.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--ip_tables_rule--actors--label_group"></a>
### Nested Schema for `ip_tables_rule.actors.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--ip_tables_rule--actors--workload"></a>
### Nested Schema for `ip_tables_rule.actors.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--ip_tables_rule--statements"></a>
### Nested Schema for `ip_tables_rule.statements`

Required:

- **chain_name** (String) Chain name for statement. Allowed values are "PREROUTING", "INPUT" and "OUTPUT"
- **parameters** (String) Parameters of statements
- **table_name** (String) Name of the table. Allowed values are "nat", "mangle" and "filter"



<a id="nestedblock--rule"></a>
### Nested Schema for `rule`

Required:

- **consumer** (Block Set, Min: 1) Consumers for Security Rule. Only one actor can be specified in one consumer block (see [below for nested schema](#nestedblock--rule--consumer))
- **enabled** (Boolean) Enabled flag. Determines whether the rule will be enabled in rule set or not
- **illumio_provider** (Block Set, Min: 1) providers for Security Rule. Only one actor can be specified in one illumio_provider block (see [below for nested schema](#nestedblock--rule--illumio_provider))
- **resolve_labels_as** (Block List, Min: 1, Max: 1) resolve label as for Security rule (see [below for nested schema](#nestedblock--rule--resolve_labels_as))

Optional:

- **description** (String) Description of Security Rule
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **ingress_service** (Block Set) Collection of Ingress Services. If resolve_label_as.providers list includes "workloads" then ingress_service is required. Only one of the {"href"} or {"proto", "port", "to_port"} parameter combination is allowed (see [below for nested schema](#nestedblock--rule--ingress_service))
- **machine_auth** (Boolean) Determines whether machine authentication is enabled
- **sec_connect** (Boolean) Determines whether a secure connection is established. Defaule Value: false
- **stateless** (Boolean) Determines whether packet filtering is stateless for the rule
- **unscoped_consumers** (Boolean) Set the scope for rule consumers to All. Defaule Value: false

Read-Only:

- **created_at** (String) Timestamp when this security rule was first created
- **created_by** (Map of String) User who originally created this security rule
- **deleted_at** (String) Timestamp when this security rule was deleted
- **deleted_by** (Map of String) User who deleted this security rule
- **href** (String) URI of Security Rule
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this security rule was last updated
- **updated_by** (Map of String) User who last updated this security rule

<a id="nestedblock--rule--consumer"></a>
### Nested Schema for `rule.consumer`

Optional:

- **actors** (String) actors for consumers parameter. Allowed values are "ams" and "container_host"
- **ip_list** (Block List, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--rule--consumer--ip_list))
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--rule--consumer--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--rule--consumer--label_group))
- **virtual_service** (Block List, Max: 1) Href of Virtual Service (see [below for nested schema](#nestedblock--rule--consumer--virtual_service))
- **workload** (Block List, Max: 1) Href of Worklaod (see [below for nested schema](#nestedblock--rule--consumer--workload))

<a id="nestedblock--rule--consumer--ip_list"></a>
### Nested Schema for `rule.consumer.ip_list`

Required:

- **href** (String) URI of IP List


<a id="nestedblock--rule--consumer--label"></a>
### Nested Schema for `rule.consumer.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--rule--consumer--label_group"></a>
### Nested Schema for `rule.consumer.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--rule--consumer--virtual_service"></a>
### Nested Schema for `rule.consumer.virtual_service`

Required:

- **href** (String) URI of Virtual Service


<a id="nestedblock--rule--consumer--workload"></a>
### Nested Schema for `rule.consumer.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--rule--illumio_provider"></a>
### Nested Schema for `rule.illumio_provider`

Optional:

- **actors** (String) actors for illumio_provider. Valid value is "ams"
- **ip_list** (Block List, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--rule--illumio_provider--ip_list))
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--rule--illumio_provider--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--rule--illumio_provider--label_group))
- **virtual_server** (Block List, Max: 1) Href of Virtual Server (see [below for nested schema](#nestedblock--rule--illumio_provider--virtual_server))
- **virtual_service** (Block List, Max: 1) Href of Virtual Service (see [below for nested schema](#nestedblock--rule--illumio_provider--virtual_service))
- **workload** (Block List, Max: 1) Href of Worklaod (see [below for nested schema](#nestedblock--rule--illumio_provider--workload))

<a id="nestedblock--rule--illumio_provider--ip_list"></a>
### Nested Schema for `rule.illumio_provider.ip_list`

Required:

- **href** (String) URI of IP List


<a id="nestedblock--rule--illumio_provider--label"></a>
### Nested Schema for `rule.illumio_provider.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--rule--illumio_provider--label_group"></a>
### Nested Schema for `rule.illumio_provider.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--rule--illumio_provider--virtual_server"></a>
### Nested Schema for `rule.illumio_provider.virtual_server`

Required:

- **href** (String) URI of Virtual Server


<a id="nestedblock--rule--illumio_provider--virtual_service"></a>
### Nested Schema for `rule.illumio_provider.virtual_service`

Required:

- **href** (String) URI of Virtual Service


<a id="nestedblock--rule--illumio_provider--workload"></a>
### Nested Schema for `rule.illumio_provider.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--rule--resolve_labels_as"></a>
### Nested Schema for `rule.resolve_labels_as`

Required:

- **consumers** (List of String) consumers for resolve_labels_as. Allowed values are "workloads", "virtual_services"
- **providers** (List of String) providers for resolve_labels_as. Allowed values are "workloads", "virtual_services"


<a id="nestedblock--rule--ingress_service"></a>
### Nested Schema for `rule.ingress_service`

Optional:

- **href** (String) URI of Service
- **port** (String) Port number used with protocol or starting port when specifying a range. Valid range is 0-65535
- **proto** (String) Protocol number. Allowed values are 6 and 17
- **to_port** (String) Upper end of port range. Valid range is 0-65535

