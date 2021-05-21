---
layout: "illumio-core"
page_title: "illumio-core_ip_lists Data Source - terraform-provider-illumio-core"
subcategory: ""
sidebar_current: "docs-illumio-core-data-source-ip-lists"
description: |-
  Represents Illumio IP Lists
---

# illumio-core_ip_lists (Data Source)

Represents Illumio IP Lists

Example Usage
------------

```hcl
data "illumio-core_ip_lists" "example" {
  max_results = "5"
}  
```

## Schema

### Optional

- **description** (String) Description of IP list(s) to return. Supports partial matches
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **fqdn** (String) IP lists matching FQDN. Supports partial matches
- **ip_address** (String) IP address matching IP list(s) to return. Supports partial matches
- **max_results** (String) Maximum number of IP Lists to return. The integer should be a non-zero positive integer.
- **name** (String) Name of IP list(s) to return. Supports partial matches
- **pversion** (String) pversion of the security policy. Allowed values are "draft", "active", and numbers greater than 0. Default value: "draft"

### Read-Only

- **items** (List of Object) list of IP Lists (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **created_at** (String) Timestamp when this IP List was first created
- **created_by** (Map of String) User who created this IP List
- **deleted_at** (String) Timestamp when this IP List was deleted
- **deleted_by** (Map of String) User who deleted this IP List
- **description** (String) description of this IP List
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **fqdns** (List of Object) Collection of Fully Qualified Domain Names for IP List (see [below for nested schema](#nestedobjatt--items--fqdns))
- **href** (String) URI of the IP List
- **ip_ranges** (List of Object) (see [below for nested schema](#nestedobjatt--items--ip_ranges))
- **name** (String) Name of the IP List
- **updated_at** (String) Timestamp when this IP List was last updated
- **updated_by** (Map of String) User who last updated this IP List

<a id="nestedobjatt--items--fqdns"></a>
### Nested Schema for `items.fqdns`

Read-Only:

- **description** (String) Description of FQDN
- **fqdn** (String) Fully Qualified Domain Name


<a id="nestedobjatt--items--ip_ranges"></a>
### Nested Schema for `items.ip_ranges`

Read-Only:

- **description** (String) Description of IP Range
- **exclusion** (Boolean) Whether this IP address is an exclusion. Exclusions must be a strict subset of inclusive IP addresses.
- **from_ip** (String) IP address or a low end of IP range. Might be specified with CIDR notation
- **to_ip** (String) High end of an IP range

