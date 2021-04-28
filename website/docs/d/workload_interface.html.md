---
layout: "illumio-core"
page_title: "illumio-core_workload_interface Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-workload-interface"
subcategory: ""
description: |-
  Represents Illumio Workload Interface
---

# illumio-core_workload_interface (Data Source)

Represents Illumio Workload Interface

Example Usage
------------

```hcl
data "illumio-core_workload_interface" "test" {
    workload_id = "d42a430e-b20b-4b2d-853f-2d39fa4cea22"
    workload_interface_name = "testing workload"
}
```


## Schema

### Required

- **workload_id** (String) Numerical ID of Workload
- **workload_interface_name** (String) Alphanumeric name of Workload Interface


### Read-Only

- **address** (String) The IP Address to assign to this interface
- **cidr_block** (Number) The number of bits in the subnet /24 is 255.255.255.0
- **default_gateway_address** (String) The IP Address of the default gateway
- **friendly_name** (String) User-friendly name for Workload Interface
- **href** (String) URI of the Workload Interface
- **link_state** (String) Link State for Workload Interface
- **loopback** (Boolean) Loopback for Workload Interface
- **name** (String) Name of the Workload Interface
- **network** (Map of String) Network for the Workload Interface.
- **network_detection_mode** (String) Network Detection Mode for Workload Interface


