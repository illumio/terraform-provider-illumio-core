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
resource "illumio-core_security_rule" "resoruce_example" {
  rule_set_href = "/orgs/1/sec_policy/draft/rule_sets/70"

  enabled = true
  resolve_labels_as {
    consumers = ["workloads", "virtual_services"]
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
    href = "/orgs/1/sec_policy/draft/services/19"
  }

  ingress_service {
    proto = 6
    port  = 1
  }

  ingress_service {
    proto   = 6
    port    = 1
    to_port = 12
  }
}
```

## Schema

### Required

- **consumer** (Block Set, Min: 1) Consumers for Security Rule. Only one actor can be specified in one consumer block (see [below for nested schema](#nestedblock--consumer))
- **enabled** (Boolean) Enabled flag. Determines whether the rule will be enabled in rule set or not
- **illumio_provider** (Block Set, Min: 1) providers for Security Rule. Only one actor can be specified in one illumio_provider block (see [below for nested schema](#nestedblock--illumio_provider))
- **resolve_labels_as** (Block List, Min: 1, Max: 1) resolve label as for Security rule (see [below for nested schema](#nestedblock--resolve_labels_as))
- **rule_set_href** (String) URI of Rule set, in which security rule will be added.

### Optional

- **description** (String) Description of Security Rule
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **ingress_service** (Block Set) Collection of Ingress Services. If resolve_label_as.providers list includes "workloads" then ingress_service is required. Only one of the {"href"} or {"proto", "port", "to_port"} parameter combination is allowed (see [below for nested schema](#nestedblock--ingress_service))
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

<a id="nestedblock--consumer"></a>
### Nested Schema for `consumer`

Optional:

- **actors** (String) actors for consumers parameter. Allowed values are "ams" and "container_host"
- **ip_list** (Block List, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--consumer--ip_list))
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--consumer--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--consumer--label_group))
- **virtual_service** (Block List, Max: 1) Href of Virtual Service (see [below for nested schema](#nestedblock--consumer--virtual_service))
- **workload** (Block List, Max: 1) Href of Worklaod (see [below for nested schema](#nestedblock--consumer--workload))

<a id="nestedblock--consumer--ip_list"></a>
### Nested Schema for `consumer.ip_list`

Required:

- **href** (String) URI of IP List


<a id="nestedblock--consumer--label"></a>
### Nested Schema for `consumer.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--consumer--label_group"></a>
### Nested Schema for `consumer.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--consumer--virtual_service"></a>
### Nested Schema for `consumer.virtual_service`

Required:

- **href** (String) URI of Virtual Service


<a id="nestedblock--consumer--workload"></a>
### Nested Schema for `consumer.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--illumio_provider"></a>
### Nested Schema for `illumio_provider`

Optional:

- **actors** (String) actors for illumio_provider. Valid value is "ams"
- **ip_list** (Block List, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--illumio_provider--ip_list))
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--illumio_provider--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--illumio_provider--label_group))
- **virtual_server** (Block List, Max: 1) Href of Virtual Server (see [below for nested schema](#nestedblock--illumio_provider--virtual_server))
- **virtual_service** (Block List, Max: 1) Href of Virtual Service (see [below for nested schema](#nestedblock--illumio_provider--virtual_service))
- **workload** (Block List, Max: 1) Href of Worklaod (see [below for nested schema](#nestedblock--illumio_provider--workload))

<a id="nestedblock--illumio_provider--ip_list"></a>
### Nested Schema for `illumio_provider.ip_list`

Required:

- **href** (String) URI of IP List


<a id="nestedblock--illumio_provider--label"></a>
### Nested Schema for `illumio_provider.label`

Required:

- **href** (String) URI of Label


<a id="nestedblock--illumio_provider--label_group"></a>
### Nested Schema for `illumio_provider.label_group`

Required:

- **href** (String) URI of Label Group


<a id="nestedblock--illumio_provider--virtual_server"></a>
### Nested Schema for `illumio_provider.virtual_server`

Required:

- **href** (String) URI of Virtual Server


<a id="nestedblock--illumio_provider--virtual_service"></a>
### Nested Schema for `illumio_provider.virtual_service`

Required:

- **href** (String) URI of Virtual Service


<a id="nestedblock--illumio_provider--workload"></a>
### Nested Schema for `illumio_provider.workload`

Required:

- **href** (String) URI of Workload



<a id="nestedblock--resolve_labels_as"></a>
### Nested Schema for `resolve_labels_as`

Required:

- **consumers** (List of String) consumers for resolve_labels_as. Allowed values are "workloads", "virtual_services"
- **providers** (List of String) providers for resolve_labels_as. Allowed values are "workloads", "virtual_services"


<a id="nestedblock--ingress_service"></a>
### Nested Schema for `ingress_service`

Optional:

- **href** (String) URI of Service
- **port** (String) Port number used with protocol or starting port when specifying a range. Valid range is 0-65535
- **proto** (String) Protocol number. Allowed values are 6 and 17
- **to_port** (String) Upper end of port range. Valid range is 0-65535

