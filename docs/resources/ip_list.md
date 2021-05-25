---
layout: "illumio-core"
page_title: "illumio-core_ip_list Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-ip-list"
subcategory: ""
description: |-
  Manages Illumio IP List
---


# illumio-core_ip_list (Resource)

Manages Illumio IP List

Example Usage
------------

```hcl
resource "illumio-core_ip_list" "example" {
    name = "example name"
    description = "example desc"
    ip_ranges{
      	from_ip = "0.0.0.0/0"
        to_ip = "0.0.0.0/0"
        description = "example ip_ranges description"
        exclusion = false
    }
    fqdns {
        fqdn = "app.example.com"
        description = "example fqdn description"
    }
    external_data_reference = "example reference"
    external_data_set = "example set"
}
```


## Schema

### Required

- **name** (String) Name of the IP List. The name should be between 1 to 255 characters


- **description** (String) Description of the IP List
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **fqdns** (Block Set) Collection of Fully Qualified Domain Names for IP List (see [below for nested schema](#nestedblock--fqdns))
- **ip_ranges** (Block Set) IP addresses or ranges for IP List (see [below for nested schema](#nestedblock--ip_ranges))

### Read-Only

- **created_at** (String) Timestamp when this IP List was first created
- **created_by** (Map of String) User who created this IP List
- **href** (String) URI of this IP List
- **deleted_at** (String) Timestamp when this IP List was last deleted
- **deleted_by** (Map of String) User who last deleted this IP List
- **updated_at** (String) Timestamp when this IP List was last updated
- **updated_by** (Map of String) User who last updated this IP List

<a id="nestedatt--fqdns"></a>
### Nested Schema for `fqdns`

Required:

- **fqdn** (String) Fully Qualified Domain Name for IP List. Supported formats are hostname, IP, and URI

Optional:

- **description** (String) Description of FQDN

<a id="nestedatt--ip_ranges"></a>
### Nested Schema for `ip_ranges`

Required:

- **from_ip** (String) IP address or a low end of IP range. Might be specified with CIDR notation. The IP given should be in CIDR format example "0.0.0.0/0"

Optional:

- **description** (String) Description of IP Range
- **exclusion** (Boolean) Whether this IP address is an exclusion. Exclusions must be a strict subset of inclusive IP addresses
- **to_ip** (String) High end of an IP range. The IP given should be in CIDR form i.e. "0.0.0.0/0"


