---
layout: "illumio-core"
page_title: "illumio-core_traffic_collector_settings Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-traffic-collector-settings"
subcategory: ""
description: |-
  Represents Illumio Traffic Collector Settings
---

# illumio-core_traffic_collector_settings (Data Source)

Represents Illumio Traffic Collector Settings

Example Usage
------------

```hcl
data "illumio-core_traffic_collector_settings" "example" {
  href = "/orgs/1/settings/traffic_collector/2d9d2170-520e-42c4-92bd-cdf2216a1dab"
}
```

## Schema

### Required

- **href** (String) URI of traffic collecter settings

### Read-Only

- **action** (String) action for target traffic
- **target** (List of Object) target for traffic collector settings (see [below for nested schema](#nestedatt--target))
- **transmission** (String) transmission type

<a id="nestedatt--target"></a>
### Nested Schema for `target`

Read-Only:

- **dst_ip** (String) single IP address or CIDR
- **dst_port** (Number) destination port for target
- **proto** (Number) protocol for target


