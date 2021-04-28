---
layout: "illumio-core"
page_title: "illumio-core_traffic_collector_settings Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-traffic-collector-settings"
subcategory: ""
description: |-
  Manages Illumio Traffic Collector Settings
---

# illumio-core_traffic_collector_settings (Resource)

Manages Illumio Traffic Collector Settings

Example Usage
------------

```hcl
resource "illumio-core_traffic_collector_settings" "example" {
  action       = "drop"
  transmission = "broadcast"
  target {
    dst_ip   = "1.1.1.2"
    dst_port = -1
    proto    = 6
  }
}
```

## Schema

### Required

- **action** (String) action for target traffic. Allowed values are "drop" or "aggregate"
- **transmission** (String) transmission type. Allowed values are "broadcast" and "multicast"

### Optional

- **target** (Block List, Max: 1) target for traffic collector settings (see [below for nested schema](#nestedblock--target))

### Read-Only

- **href** (String) URI of traffic collecter settings

<a id="nestedblock--target"></a>
### Nested Schema for `target`

Required:

- **proto** (Number) protocol for target. Allowed values are 6 (TCP), 17 (UDP), 1 (ICMP) and 58 (ICMPv6)

Optional:

- **dst_ip** (String) single ip address or CIDR. Default value: "0.0.0.0/0"
- **dst_port** (Number) destination port for target. Allowed range is -1 to 65535. Default value: -1


