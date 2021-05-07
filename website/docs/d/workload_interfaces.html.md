---
layout: "illumio-core"
page_title: "illumio-core_workload_interfaces Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-workload-interfaces"
subcategory: ""
description: |-
  Represents Illumio Workload Interfaces
---

# illumio-core_workload_interfaces (Data Source)

Represents Illumio Workload Interfaces


Example Usage
------------

```hcl
data "illumio-core_workload_interface" "example" {
    workload_href = "/orgs/1/workloads/d42a430e-b20b-4b2d-853f-2d39fa4cea22"
}
```

## Schema

### Required

- **workload_href** (String) URI of Workload

### Read-Only

- **items** (List of Object) list of Workload Interfaces (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **address** (String) The IP Address to assign to this interface
- **cidr_block** (Number) The number of bits in the subnet /24 is 255.255.255.0
- **default_gateway_address** (String) The IP Address of the default gateway
- **friendly_name** (String) User-friendly name for Workload Interface
- **link_state** (String) Link State for Workload Interface
- **loopback** (Boolean) Loopback for Workload Interface
- **name** (String) Name of the Workload Interface
- **network** (Map of String) Network for the Workload Interface.
- **network_detection_mode** (String) Network Detection Mode for Workload Interface



