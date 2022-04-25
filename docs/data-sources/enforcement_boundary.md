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

Example Usage
------------

```hcl
data "illumio-core_enforcement_boundary" "example" {
  href = illumio-core_enforcement_boundary.example.href
}

resource "illumio-core_enforcement_boundary" "example" {
  ...
}
```

## Schema

### Required

- `href` (String) URI of this Enforcement Boundary

### Read-Only

- `caps` (List of String) CAPS for Enforcement Boundary
- `consumers` (List of Object) Consumers for Enforcement Boundary. Only one actor can be specified in one consumer block (see [below for nested schema](#nestedatt--consumers))
- `created_at` (String) Timestamp when this Enforcement Boundary was first created
- `created_by` (Map of String) User who created this Enforcement Boundary
- `deleted_at` (String) Timestamp when this Enforcement Boundary was last deleted
- `deleted_by` (Map of String) User who last deleted this Enforcement Boundary
- `ingress_services` (List of Object) Collection of Ingress Service. Only one of the {"href"} or {"proto", "port", "to_port"} parameter combination is allowed (see [below for nested schema](#nestedatt--ingress_services))
- `name` (String) Name of the Enforcement Boundary
- `providers` (Block List) providers for Enforcement Boundary. Only one actor can be specified in one providers block (see [below for nested schema](#nestedblock--providers))
- `updated_at` (String) Timestamp when this Enforcement Boundary was last updated
- `updated_by` (Map of String) User who last updated this Enforcement Boundary

<a id="nestedatt--consumers"></a>
### Nested Schema for `consumers`

Read-Only:

- `actors` (String) actors for consumers parameter
- `ip_list` (List of Object) Href of IP List (see [below for nested schema](#nestedobjatt--consumers--ip_list))
- `label` (List of Object) Href of Label (see [below for nested schema](#nestedobjatt--consumers--label))
- `label_group` Href of Label Group (List of Object) (see [below for nested schema](#nestedobjatt--consumers--label_group))

<a id="nestedobjatt--consumers--ip_list"></a>
### Nested Schema for `consumers.ip_list`

Read-Only:

- `href` (String)


<a id="nestedobjatt--consumers--label"></a>
### Nested Schema for `consumers.label`

Read-Only:

- `href` (String)


<a id="nestedobjatt--consumers--label_group"></a>
### Nested Schema for `consumers.label_group`

Read-Only:

- `href` (String) Href of Label Group




<a id="nestedatt--ingress_services"></a>
### Nested Schema for `ingress_services`

Read-Only:

- `href` (String) URI of Ingress Service
- `port` (String) Port number used with protocol or starting port when specifying a range
- `proto` (String) Protocol number
- `to_port` (String) Upper end of port range


<a id="nestedblock--providers"></a>
### Nested Schema for `providers`

Read-Only:

- `actors` (String) actors for providers
- `ip_list` (List of Object) Href of IP List (see [below for nested schema](#nestedatt--providers--ip_list))
- `label` (List of Object) Href of Label (see [below for nested schema](#nestedatt--providers--label))
- `label_group` (List of Object) Href of Label Group (see [below for nested schema](#nestedatt--providers--label_group))

<a id="nestedatt--providers--ip_list"></a>
### Nested Schema for `providers.ip_list`

Read-Only:

- `href` (String) Href of IP List


<a id="nestedatt--providers--label"></a>
### Nested Schema for `providers.label`

Read-Only:

- `href` (String) Href of Label


<a id="nestedatt--providers--label_group"></a>
### Nested Schema for `providers.label_group`

Read-Only:

- `href` (String) Href of Label Group
