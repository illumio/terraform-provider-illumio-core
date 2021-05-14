---
layout: "illumio-core"
page_title: "illumio-core_security_rule Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-security-rule"
subcategory: ""
description: |-
  Manages Illumio Security Rule
---

# illumio-core_security_rule (Resource)

Manages Illumio Security Rule

Example Usage
------------

```hcl
resource "illumio-core_security_rule" "example" {
  rule_set_href = "/orgs/1/sec_policy/draft/rule_sets/70"

  enabled = true
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
    href = "/orgs/1/sec_policy/draft/services/19"
  }

  ingress_services {
    proto = 6
    port  = 1
  }

  ingress_services {
    proto   = 6
    port    = 1
    to_port = 12
  }
}
```

## Schema

### Required

- **consumers** (Block Set, Min: 1) Consumers for Security Rule. Only one actor can be specified in one consumers block (see [below for nested schema](#nestedblock--consumers))
- **enabled** (Boolean) Enabled flag. Determines whether the rule will be enabled in rule set or not
- **providers** (Block Set, Min: 1) providers for Security Rule. Only one actor can be specified in one providers block (see [below for nested schema](#nestedblock--providers))
- **resolve_labels_as** (Block List, Min: 1, Max: 1) resolve label as for Security rule (see [below for nested schema](#nestedblock--resolve_labels_as))
- **rule_set_href** (String) URI of Rule set, in which security rule will be added.

### Optional

- **description** (String) Description of Security Rule
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **ingress_services** (Block Set) Collection of Ingress Services. If resolve_label_as.providers list includes "workloads" then ingress_services is required. Only one of the {"href"} or {"proto", "port", "to_port"} parameter combination is allowed (see [below for nested schema](#nestedblock--ingress_services))
- **machine_auth** (Boolean) Determines whether machine authentication is enabled
- **sec_connect** (Boolean) Determines whether a secure connection is established. Defaule Value: false
- **stateless** (Boolean) Determines whether packet filtering is stateless for the rule
- **unscoped_consumers** (Boolean) Set the scope for rule consumers to All. Defaule Value: false

### Read-Only

- **created_at** (String) Timestamp when this security rule was first created
- **created_by** (Map of String) User who originally created this security rule
- **deleted_at** (String) Timestamp when this security rule was deleted
- **deleted_by** (Map of String) User who deleted this security rule
- **href** (String) URI of Security Rule
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this security rule was last updated
- **updated_by** (Map of String) User who last updated this security rule

<a id="nestedblock--consumers"></a>
### Nested Schema for `consumers`

Optional:

- **actors** (String) actors for consumers parameter. Allowed values are "ams" and "container_host"
- **ip_list** (Block List, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--consumers--ip_list))
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--consumers--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--consumers--label_group))
- **virtual_service** (Block List, Max: 1) Href of Virtual Service (see [below for nested schema](#nestedblock--consumers--virtual_service))
- **workload** (Block List, Max: 1) Href of Workload (see [below for nested schema](#nestedblock--consumers--workload))

<a id="nestedblock--consumers--ip_list"></a>
### Nested Schema for `consumers.ip_list`

Required:

- **href** (String) URI of IP List


<a id="nestedblock--consumers--label"></a>
### Nested Schema for `consumers.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--consumers--label_group"></a>
### Nested Schema for `consumers.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--consumers--virtual_service"></a>
### Nested Schema for `consumers.virtual_service`

Required:

- **href** (String) URI of Virtual Service


<a id="nestedblock--consumers--workload"></a>
### Nested Schema for `consumers.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--providers"></a>
### Nested Schema for `providers`

Optional:

- **actors** (String) actors for providers. Valid value is "ams"
- **ip_list** (Block List, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--providers--ip_list))
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--providers--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--providers--label_group))
- **virtual_server** (Block List, Max: 1) Href of Virtual Server (see [below for nested schema](#nestedblock--providers--virtual_server))
- **virtual_service** (Block List, Max: 1) Href of Virtual Service (see [below for nested schema](#nestedblock--providers--virtual_service))
- **workload** (Block List, Max: 1) Href of Workload (see [below for nested schema](#nestedblock--providers--workload))

<a id="nestedblock--providers--ip_list"></a>
### Nested Schema for `providers.ip_list`

Required:

- **href** (String) URI of IP List


<a id="nestedblock--providers--label"></a>
### Nested Schema for `providers.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--providers--label_group"></a>
### Nested Schema for `providers.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--providers--virtual_server"></a>
### Nested Schema for `providers.virtual_server`

Required:

- **href** (String) URI of Virtual Server


<a id="nestedblock--providers--virtual_service"></a>
### Nested Schema for `providers.virtual_service`

Required:

- **href** (String) URI of Virtual Service


<a id="nestedblock--providers--workload"></a>
### Nested Schema for `providers.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--resolve_labels_as"></a>
### Nested Schema for `resolve_labels_as`

Required:

- **consumers** (List of String) consumers for resolve_labels_as. Allowed values are "workloads", "virtual_services"
- **providers** (List of String) providers for resolve_labels_as. Allowed values are "workloads", "virtual_services"


<a id="nestedblock--ingress_services"></a>
### Nested Schema for `ingress_services`

Optional:

- **href** (String) URI of Service
- **port** (String) Port number used with protocol or starting port when specifying a range. Valid range is 0-65535
- **proto** (String) Protocol number. Allowed values are 6 and 17
- **to_port** (String) Upper end of port range. Valid range is 0-65535

