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


```hcl
resource "illumio-core_enforcement_boundary" "test" {
    name = "testing eb 122"
    ingress_service {
      href = "/orgs/1/sec_policy/draft/services/3"
    }
    consumer {
      ip_list {
        href = "/orgs/1/sec_policy/draft/ip_lists/1"
      }
    }
    illumio_provider {
      label {
        href = "/orgs/1/labels/1"
      }
    }
}
```

## Schema

### Required

- **consumer** (Block Set, Min: 1) Consumers for Enforcement Boundary. Only one actor can be specified in one consumer block (see [below for nested schema](#nestedblock--consumer))
- **illumio_provider** (Block Set, Min: 1) providers for Enforcement Boundary. Only one actor can be specified in one illumio_provider block (see [below for nested schema](#nestedblock--illumio_provider))
- **ingress_service** (Block Set, Min: 1) Collection of Ingress Service. Only one of the {"href"} or {"proto", "port", "to_port"} parameter combination is allowed (see [below for nested schema](#nestedblock--ingress_service))
- **name** (String) Name of the Enforcement Boundary

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **caps** (List of String) CAPS for Enforcement Boundary
- **created_at** (String) Timestamp when this Enforcement Boundary was first created
- **created_by** (Map of String) User who originally created this Enforcement Boundary
- **deleted_at** (String) Timestamp when this Enforcement Boundary was last deleted
- **deleted_by** (Map of String) User who last deleted this Enforcement Boundary
- **href** (String) URI of this Enforcement Boundary
- **updated_at** (String) Timestamp when this Enforcement Boundary was last updated
- **updated_by** (Map of String) User who last updated this Enforcement Boundary

<a id="nestedblock--consumer"></a>
### Nested Schema for `consumer`

Optional:

- **actors** (String) actors for consumers parameter. Allowed values are "ams" and "container_host"
- **ip_list** (Block List, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--consumer--ip_list))
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--consumer--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--consumer--label_group))

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



<a id="nestedblock--illumio_provider"></a>
### Nested Schema for `illumio_provider`

Optional:

- **actors** (String) actors for illumio_provider. Valid value is "ams"
- **ip_list** (Block List, Max: 1) Href of IP List (see [below for nested schema](#nestedblock--illumio_provider--ip_list))
- **label** (Block List, Max: 1) Href of Label (see [below for nested schema](#nestedblock--illumio_provider--label))
- **label_group** (Block List, Max: 1) Href of Label Group (see [below for nested schema](#nestedblock--illumio_provider--label_group))

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



<a id="nestedblock--ingress_service"></a>
### Nested Schema for `ingress_service`

Optional:

- **href** (String) URI of Service
- **port** (String) Port number used with protocol or starting port when specifying a range. Valid range is 0-65535
- **proto** (String) Protocol number. Allowed values are 6 and 17
- **to_port** (String) Upper end of port range. Valid range (0-65535)


