---
layout: "illumio-core"
page_title: "illumio-core_enforcement_boundary Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-enforcement-boundary"
subcategory: ""
description: |-
  Represents Illumio Enforcement Boundary
---

# illumio-core_enforcement_boundary (Data Source)

Represents Illumio Enforcement Boundary

```hcl
data "illumio-core_enforcement_boundary" "test" {
  href = "/orgs/1/sec_policy/draft/enforcement_boundaries/57"
}
```

## Schema

### Required

- **href** (String) URI of this Enforcement Boundary

### Read-Only

- **caps** (List of String) CAPS for Enforcement Boundary
- **consumer** (Set of Object) Consumers for Enforcement Boundary. Only one actor can be specified in one consumer block (see [below for nested schema](#nestedatt--consumer))
- **created_at** (String) Timestamp when this Enforcement Boundary was first created
- **created_by** (Map of String) User who originally created this Enforcement Boundary
- **deleted_at** (String) Timestamp when this Enforcement Boundary was last deleted
- **deleted_by** (Map of String) User who last deleted this Enforcement Boundary
- **illumio_provider** (Block Set) providers for Enforcement Boundary. Only one actor can be specified in one illumio_provider block (see [below for nested schema](#nestedblock--illumio_provider))
- **ingress_service** (Set of Object) Collection of Ingress Service. Only one of the {"href"} or {"proto", "port", "to_port"} parameter combination is allowed (see [below for nested schema](#nestedatt--ingress_service))
- **name** (String) Name of the Enforcement Boundary
- **updated_at** (String) Timestamp when this Enforcement Boundary was last updated
- **updated_by** (Map of String) User who last updated this Enforcement Boundary

<a id="nestedatt--consumer"></a>
### Nested Schema for `consumer`

Read-Only:

- **actors** (String) actors for consumers parameter
- **ip_list** (List of Object) Href of IP List (see [below for nested schema](#nestedobjatt--consumer--ip_list))
- **label** (List of Object) Href of Label (see [below for nested schema](#nestedobjatt--consumer--label))
- **label_group** Href of Label Group (List of Object) (see [below for nested schema](#nestedobjatt--consumer--label_group))

<a id="nestedobjatt--consumer--ip_list"></a>
### Nested Schema for `consumer.ip_list`

Read-Only:

- **href** (String) Href of IP List


<a id="nestedobjatt--consumer--label"></a>
### Nested Schema for `consumer.label`

Read-Only:

- **href** (String) Href of Label


<a id="nestedobjatt--consumer--label_group"></a>
### Nested Schema for `consumer.label_group`

Read-Only:

- **href** (String) Href of Label Group



<a id="nestedblock--illumio_provider"></a>
### Nested Schema for `illumio_provider`

Read-Only:

- **actors** (String) actors for illumio_provider.
- **ip_list** (List of Object) Href of IP List (see [below for nested schema](#nestedatt--illumio_provider--ip_list))
- **label** (List of Object) Href of Label (see [below for nested schema](#nestedatt--illumio_provider--label))
- **label_group** (List of Object) Href of Label Group (see [below for nested schema](#nestedatt--illumio_provider--label_group))

<a id="nestedatt--illumio_provider--ip_list"></a>
### Nested Schema for `illumio_provider.ip_list`

Read-Only:

- **href** (String) Href of IP List


<a id="nestedatt--illumio_provider--label"></a>
### Nested Schema for `illumio_provider.label`

Read-Only:

- **href** (String) Href of Label


<a id="nestedatt--illumio_provider--label_group"></a>
### Nested Schema for `illumio_provider.label_group`

Read-Only:

- **href** (String) Href of Label Group



<a id="nestedatt--ingress_service"></a>
### Nested Schema for `ingress_service`

Read-Only:

- **href** (String) URI of Ingress Service
- **port** (String) Port number used with protocol or starting port when specifying a range
- **proto** (String) Protocol number
- **to_port** (String) Upper end of port range


