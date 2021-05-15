---
layout: "illumio-core"
page_title: "illumio-core_workload_interface Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-workload-interface"
subcategory: ""
description: |-
  Manages Illumio Workload Interface
---

# illumio-core_workload_interface (Resource)

Manages Illumio Workload Interface


Example Usage
------------

```hcl
resource "illumio-core_workload_interface" "example" {
    workload_href = "/orgs/1/workloads/d42a430e-b20b-4b2d-853f-2d39fa4cea22"
    name = "example name"
    link_state = "up"
    friendly_name = "example friendly name"
}
```

## Schema

### Required

- **link_state** (String) Link State for Workload Interface. Allowed values are "up", "down", and "unknown"
- **name** (String) Name of the Workload Interface
- **workload_href** (String) URI of Workload

### Optional

- **address** (String) The IP Address to assign to this interface. The address should in the IPv4 or IPv6 format.
- **cidr_block** (Number) CIDR BLOCK of the Workload Interface.
- **default_gateway_address** (String) The IP Address of the default gateway. The Default Gateaway Address should in the IPv4 or IPv6 format.
- **friendly_name** (String) User-friendly name for Workload Interface

### Read-Only

- **href** (String) URI of the Workload Interface
- **loopback** (Boolean) Loopback for Workload Interface
- **network** (Map of String) Network for the Workload Interface.
- **network_detection_mode** (String) Network Detection Mode for Workload Interface


