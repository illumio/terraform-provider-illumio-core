---
layout: "illumio-core"
page_title: "illumio-core_enforcement_boundaries Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-enforcement-boundary"
subcategory: ""
description: |-
  Represents Illumio Enforcement Boundaries
---

# illumio-core_enforcement_boundaries (Data Source)

Represents Illumio Enforcement Boundaries

Example Usage
------------

```hcl
data "illumio-core_enforcement_boundaries" "example" {}
```

## Schema

### Optional

- `labels` (String) List of lists of label URIs, encoded as a JSON string
- `max_results` (String) Maximum number of enforcement boundaries to return. The integer should be a non-zero positive integer
- `name` (String) Filter by name supports partial matching
- `pversion` (String) pversion of the security policy. Allowed values are "draft", "active", and numbers greater than 0. Default value: "draft"
- `service` (String) Service URI
- `service_ports_port` (String) Specify port or port range to filter results. The range is from -1 to 65535
- `service_ports_proto` (String) Protocol to filter on. Allowed values are 6 and 17

### Read-Only

- `items` (List of Object) List of Enforcement Boundary (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `caps` (List of String) CAPS for Enforcement Boundary
- `consumer` (List of Object) Consumers for Enforcement Boundary (see [below for nested schema](#nestedobjatt--items--consumer))
- `created_at` (String) Timestamp when this Enforcement Boundary was first created
- `created_by` (Map of String) User who created this Enforcement Boundary
- `deleted_at` (String) Timestamp when this Enforcement Boundary was last deleted
- `deleted_by` (Map of String) User who last deleted this Enforcement Boundary
- `href` (String) Href of Enforcement Boundary
- `ingress_service` (List of Object)  Collection of Ingress Service (see [below for nested schema](#nestedobjatt--items--ingress_service))
- `name` (String) Name of the Enforcement Boundary
- `providers` (List of Object) Providers for Enforcement Boundary (see [below for nested schema](#nestedobjatt--items--providers))
- `updated_at` (String) Timestamp when this Enforcement Boundary was last updated
- `updated_by` (Map of String) User who last updated this Enforcement Boundary

<a id="nestedobjatt--items--consumers"></a>
### Nested Schema for `items.consumers`

Read-Only:

- `actors` (String) actors for consumers parameter
- `ip_list` (List of Object) Href of IP List  (see [below for nested schema](#nestedobjatt--items--consumer--ip_list))
- `label` (List of Object) Href of Label  (see [below for nested schema](#nestedobjatt--items--consumer--label))
- `label_group` (List of Object) Href of Label Group (see [below for nested schema](#nestedobjatt--items--consumer--label_group))

<a id="nestedobjatt--items--consumers--ip_list"></a>
### Nested Schema for `items.consumers.ip_list`

Read-Only:

- `href` (String) Href of IP List


<a id="nestedobjatt--items--consumers--label"></a>
### Nested Schema for `items.consumers.label`

Read-Only:

- `href` (String) Href of Label


<a id="nestedobjatt--items--consumers--label_group"></a>
### Nested Schema for `items.consumers.label_group`

Read-Only:

- `href` (String) Href of Label Group



<a id="nestedobjatt--items--ingress_service"></a>
### Nested Schema for `items.ingress_service`

Read-Only:

- `href` (String) URI of Ingress Service
- `port` (String) Port number used with protocol or starting port when specifying a range
- `proto` (String) Protocol number
- `to_port` (String) Upper end of port range


<a id="nestedobjatt--items--providers"></a>
### Nested Schema for `items.providers`

Read-Only:

- `actors` (String) actors for providers
- `ip_list` (List of Object) Href of IP List (see [below for nested schema](#nestedobjatt--items--providers--ip_list))
- `label` (List of Object) Href of Label (see [below for nested schema](#nestedobjatt--items--providers--label))
- `label_group` (List of Object) Href of Label Group (see [below for nested schema](#nestedobjatt--items--providers--label_group))

<a id="nestedobjatt--items--providers--ip_list"></a>
### Nested Schema for `items.providers.ip_list`

Read-Only:

- `href` (String) Href of IP List


<a id="nestedobjatt--items--providers--label"></a>
### Nested Schema for `items.providers.label`

Read-Only:

- `href` (String) Href of Label 


<a id="nestedobjatt--items--providers--label_group"></a>
### Nested Schema for `items.providers.label_group`

Read-Only:

- `href` (String) Href of Label Group
