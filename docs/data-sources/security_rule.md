---
layout: "illumio-core"
page_title: "illumio-core_security_rule Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-security-rule"
subcategory: ""
description: |-
  Represents Illumio Security Rule
---

# illumio-core_security_rule (Data Source)

Represents Illumio Security Rule

Example Usage
------------

```hcl
data "illumio-core_security_rule" "example" {
  href = "/orgs/1/sec_policy/draft/rule_sets/6/sec_rules/7"
}
```

## Schema

### Required

- **href** (String) URI of security rule

### Read-Only

- **consumers** (Set of Object) consumers of Security Rule (see [below for nested schema](#nestedatt--consumers))
- **created_at** (String) Timestamp when this security rule was first created
- **created_by** (Map of String) User who created this security rule
- **deleted_at** (String) Timestamp when this security rule was deleted
- **deleted_by** (Map of String) User who deleted this security rule
- **description** (String) Description of Security Rule
- **enabled** (Boolean) Enabled flag. Determines whether this rule will be enabled in ruleset or not
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **ingress_services** (List of Object) Collection of Ingress Services (see [below for nested schema](#nestedatt--ingress_services))
- **machine_auth** (Boolean) Determines whether machine authentication is enabled
- **providers** (Set of Object) providers of Security Rule (see [below for nested schema](#nestedatt--providers))
- **resolve_labels_as** (List of Object) resolve_label_as of Security rule (see [below for nested schema](#nestedatt--resolve_labels_as))
- **sec_connect** (Boolean) Determines whether a secure connection is established
- **stateless** (Boolean) Determines whether packet filtering is stateless for the rule
- **unscoped_consumers** (Boolean) Set the scope for rule consumers to All
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this security rule was last updated
- **updated_by** (Map of String) User who last updated this security rule

<a id="nestedatt--consumers"></a>
### Nested Schema for `consumers`

Read-Only:

- **actors** (String) actors of consumers actors
- **ip_list** (Map of String) Href of IP List
- **label** (Map of String) Href of Label
- **label_group** (Map of String) Href of Label Group
- **virtual_service** (Map of String) Href of Virtual Service
- **workload** (Map of String) Href of Workload


<a id="nestedatt--ingress_services"></a>
### Nested Schema for `ingress_services`

Read-Only:

- **href** (String) URI of service
- **port** (Number) Protocol number
- **proto** (Number) Port number used with protocol. Also, the starting port when specifying a range
- **to_port** (Number) High end of port range inclusive if specifying a range


<a id="nestedatt--providers"></a>
### Nested Schema for `providers`

Read-Only:

- **actors** (String) actors for illumio_provider
- **ip_list** (Map of String) Href of IP List
- **label** (Map of String) Href of Label
- **label_group** (Map of String) Href of Label Group
- **virtual_server** (Map of String) Href of Virtual Server
- **virtual_service** (Map of String) Href of Virtual Service
- **workload** (Map of String) Href of Workload


<a id="nestedatt--resolve_labels_as"></a>
### Nested Schema for `resolve_labels_as`

Read-Only:

- **consumers** (List of String) consumers of resolve_labels_as
- **providers** (List of String) providers of resolve_labels_as
