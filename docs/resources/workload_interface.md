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

~> Any updates to this resource once created will be ignored.

Example Usage
------------

```hcl
resource "illumio-core_workload_interface" "example" {
    workload_href = illumio-core_unmanaged_workload.example.href
    name          = "eth0"
    friendly_name = "Wired Netwrok (Ethernet)"
    link_state    = "up"
}

resource "illumio-core_unamanaged_workload" "example" {
  ...
}
```

## Schema

### Required

- `link_state` (String) Link State for Workload Interface. Allowed values are "up", "down", and "unknown"
- `name` (String) Name of the Workload Interface. The name should be between 1 to 255 characters
- `workload_href` (String) URI of Workload

### Optional

- `address` (String) The IP Address to assign to this interface. The address should in the IPv4 or IPv6 format
- `cidr_block` (Number) CIDR BLOCK of the Workload Interface
- `default_gateway_address` (String) The IP Address of the default gateway. The Default Gateway Address should in the IPv4 or IPv6 format
- `friendly_name` (String) User-friendly name for Workload Interface

### Read-Only

- `href` (String) URI of the Workload Interface
- `loopback` (Boolean) Loopback for Workload Interface
- `network` (Map of String) Href of Network for the Workload Interface
- `network_detection_mode` (String) Network Detection Mode for Workload Interface
