---
layout: "illumio-core"
page_title: "illumio-core_vens_unpair Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-vens-unpair"
subcategory: ""
description: |-
  Manages Illumio VENs Unpair
---

# illumio-core_vens_unpair (Resource)

Manages Illumio VENs Unpair

!> **WARNING:** This resource is deprecated as of v0.2.0 and will be removed in v1.0.0

Example Usage
------------

```hcl
resource "illumio-core_vens_unpair" "example" {
  vens {
    href = "/orgs/1/vens/de7c1705-6e9d-46e7-b3b1-5ef5a638c0f8"
  }
  vens {
    href = "/orgs/1/vens/32254e8b-eddc-428d-96fa-f8625416a0d6"
  }
  vens {
    href = "/orgs/1/vens/11416eb7-43df-4acc-a4b0-c17c1e2b1b77"
  }
}
```

## Schema

### Required

- **vens** (Block Set, Min: 1) List of VENs to unpair. Max Items allowed: 1000 (see [below for nested schema](#nestedblock--vens))

### Optional

- **firewall_restore** (String) The strategy to use to restore the firewall state after the VEN is uninstalled. Allowed values are "saved", "default" and "disable". Default value: "default"

<a id="nestedblock--vens"></a>
### Nested Schema for `vens`

Required:

- **href** (String) URI of VEN


