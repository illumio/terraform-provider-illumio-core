---
layout: "illumio-core"
page_title: "illumio-core_ip_list Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-ip-list"
subcategory: ""
description: |-
  Represents Illumio IP List
---
# illumio-core_ip_list (Data Source)

Represents Illumio IP List

Example Usage
------------

```hcl
data "illumio-core_ip_list" "example" {
  ip_list_id  = 15
}
```

## Schema

### Required

- **ip_list_id** (Number) Numerical ID of IP List

### Read-Only

- **created_at** (String) Timestamp when this IP List was first created
- **created_by** (Map of String) User who originally created this IP List
- **deleted_at** (String) Timestamp when this IP List was deleted
- **deleted_by** (Map of String) User who deleted this IP List
- **description** (String) URI of this IPList
- **external_data_reference** (String) A unque identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **fqdns** (List of Object) Collection of Fully Qualified Domain Names for IP List (see [below for nested schema](#nestedatt--fqdns))
- **href** (String) URI of the IPList
- **ip_ranges** (List of Object) IP addresses or ranges for IP List (see [below for nested schema](#nestedatt--ip_ranges))
- **name** (String) Name of the IPList
- **updated_at** (String) Timestamp when this IP List was last updated
- **updated_by** (Map of String) User who last updated this IP List


<a id="nestedatt--fqdns"></a>
### Nested Schema for `fqdns`

Read-Only:

- **description** (String) Desciption of FQDN
- **fqdn** (String) Full Qualified Domain Name

<a id="nestedatt--ip_ranges"></a>
### Nested Schema for `ip_ranges`

Read-Only:

- **description** (String) Desciption of IP Range
- **exclusion** (Boolean) Whether this IP address is an exclusion. Exclusions must be a strict subset of inclusive IP addresses.
- **from_ip** (String) IP adress or a low end of IP range. Might be specified with CIDR notation
- **to_ip** (String) High end of an IP range


