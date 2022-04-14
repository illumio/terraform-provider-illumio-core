---
layout: "illumio-core"
page_title: "illumio-core_vens_upgrade Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-vens-upgrade"
subcategory: ""
description: |-
  Manages Illumio VENs Upgrade
---

# illumio-core_vens_upgrade (Resource)

Manages Illumio VENs Upgrade


Example Usage
------------

```hcl
resource "illumio-core_vens_upgrade" "example" {
  release = "21.2.0-7828"

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

- `release` (String) The software release to upgrade to
- `vens` (Block Set, Min: 1) List of VENs to unpair. Max Items allowed: 25000 (see [below for nested schema](#nestedblock--vens))

<a id="nestedblock--vens"></a>
### Nested Schema for `vens`

Required:

- `href` (String) URI of VEN
