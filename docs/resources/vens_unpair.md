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

!> This resource is deprecated as of v0.2.0 and will be removed in v1.0.0. VENs should be unpaired by deleting the associated `managed_workload` resource

Example Usage
------------

```hcl
resource "illumio-core_vens_unpair" "example" {
  vens {
    href = illumio-core_ven.example.href
  }
}

resource "illumio-core_ven" "example" {
  ...
}
```

## Schema

### Required

- `vens` (Block Set, Min: 1) List of VENs to unpair. Max Items allowed: 1000 (see [below for nested schema](#nestedblock--vens))

### Optional

- `firewall_restore` (String) The strategy to use to restore the firewall state after the VEN is uninstalled. Allowed values are "saved", "default" and "disable". Default value: "default"

<a id="nestedblock--vens"></a>
### Nested Schema for `vens`

Required:

- `href` (String) URI of VEN
