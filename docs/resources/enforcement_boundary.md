---
layout: "illumio-core"
page_title: "illumio-core_enforcement_boundary Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-enforcement-boundary"
subcategory: ""
description: |-
  Manages Illumio Enforcement Boundary
---

# illumio-core_enforcement_boundary (Resource)

Manages Illumio Enforcement Boundary

Example Usage
------------

```hcl
resource "illumio-core_enforcement_boundary" "example" {
  name = "EB-EXAMPLE"

  ingress_services {
    href = illumio-core_service.example.href
  }

  consumers {
    ip_list {
      href = illumio-core_ip_list.example.href
    }
  }

  providers {
    actors = "ams"  # special notation meaning "all managed systems" - affects all workloads
  }
}

resource "illumio-core_service" "example" {
  ...
}

resource "illumio-core_ip_list" "example" {
  ...
}
```

## Schema

### Required

- `consumers` (Block Set, Min: 1) Consumers for Enforcement Boundary. Only one actor can be specified in one consumers block (see [below for nested schema](#nestedblock--consumers))
- `ingress_services` (Block Set, Min: 1) Collection of Ingress Service. Only one of the {"href"} or {"proto", "port", "to_port"} parameter combination is allowed (see [below for nested schema](#nestedblock--ingress_services))
- `name` (String) Name of the Enforcement Boundary
- `providers` (Block Set, Min: 1) providers for Enforcement Boundary. Only one actor can be specified in one providers block (see [below for nested schema](#nestedblock--providers))

### Read-Only

- `caps` (List of String) CAPS for Enforcement Boundary
- `created_at` (String) Timestamp when this Enforcement Boundary was first created
- `created_by` (Map of String) User who created this Enforcement Boundary
- `deleted_at` (String) Timestamp when this Enforcement Boundary was last deleted
- `deleted_by` (Map of String) User who last deleted this Enforcement Boundary
- `href` (String) URI of this Enforcement Boundary
- `updated_at` (String) Timestamp when this Enforcement Boundary was last updated
- `updated_by` (Map of String) User who last updated this Enforcement Boundary

<a id="nestedblock--consumers"></a>
### Nested Schema for `consumers`

Optional:

- `actors` (String) actors for consumers parameter. Allowed values is "ams"
- `ip_list` (Block Set, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--consumers--ip_list))
- `label` (Block Set, Max: 1) Href of Label (see [below for nested schema](#nestedblock--consumers--label))
- `label_group` (Block Set, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--consumers--label_group))

<a id="nestedblock--consumers--ip_list"></a>
### Nested Schema for `consumers.ip_list`

Required:

- `href` (String) URI of IP List

<a id="nestedblock--consumers--label"></a>
### Nested Schema for `consumers.label`

Required:

- `href` (String) URI of Label

<a id="nestedblock--consumers--label_group"></a>
### Nested Schema for `consumers.label_group`

Required:

- `href` (String) URI of Label Group

<a id="nestedblock--ingress_services"></a>
### Nested Schema for `ingress_services`

Optional:

- `href` (String) URI of Service
- `port` (String) Port number used with protocol or starting port when specifying a range. Allowed range is 0-65535
- `proto` (String) Protocol number. Allowed values are 6 (TCP) and 17 (UDP)
- `to_port` (String) Upper end of port range. Allowed range is 0-65535

<a id="nestedblock--providers"></a>
### Nested Schema for `providers`

Optional:

- `actors` (String) actors for providers. Valid value is "ams"
- `ip_list` (Block Set, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--providers--ip_list))
- `label` (Block Set, Max: 1) Href of Label (see [below for nested schema](#nestedblock--providers--label))
- `label_group` (Block Set, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--providers--label_group))

<a id="nestedblock--providers--ip_list"></a>
### Nested Schema for `providers.ip_list`

Required:

- `href` (String) URI of IP List

<a id="nestedblock--providers--label"></a>
### Nested Schema for `providers.label`

Required:

- `href` (String) URI of Label

<a id="nestedblock--providers--label_group"></a>
### Nested Schema for `providers.label_group`

Required:

- `href` (String) URI of Label Group
