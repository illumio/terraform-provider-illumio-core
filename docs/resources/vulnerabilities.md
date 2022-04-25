---
layout: "illumio-core"
page_title: "illumio-core_vulnerabilities Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-vulnerabilities"
subcategory: ""
description: |-
  Manages Illumio Vulnerabilities
---

# illumio-core_vulnerabilities (Resource)

Manages Illumio Vulnerabilities

Example Usage
------------

!> The `illumio-core_vulnerabilities` resource is not fully implemented in the current version (v0.2.0)

## Schema

### Required

- `vulnerability` (Block List, Min: 1) Collection of Vulenerabilites (see [below for nested schema](#nestedblock--vulnerability))

<a id="nestedblock--vulnerability"></a>
### Nested Schema for `vulnerability`

Required:

- `name` (String) The title/name of the vulnerability
- `reference_id` (String) reference id of vulnerability
- `score` (Number) The normalized score of the vulnerability within the range of 0 to 100. CVSS Score can be used here with a 10x multiplier

Optional:

- `cve_ids` (Set of String) The cve_ids for the vulnerability
- `description` (String) An arbitrary field to store some details of the vulnerability class
