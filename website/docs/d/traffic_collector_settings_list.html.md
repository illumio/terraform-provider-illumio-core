---
layout: "illumio-core"
page_title: "illumio-core_traffic_collector_settings_list Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-traffic-collector-settings-list"
subcategory: ""
description: |-
  Represents List of Illumio Traffic Collector Settings
---

# illumio-core_traffic_collector_settings_list (Data Source)

Represents List of Illumio Traffic Collector Settings 

```hcl
data "illumio-core_traffic_collector_settings_list" "example" {}
```

## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **items** (List of Object) list of Traffic Collector Setting hrefs (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **action** (String) action for target traffic
- **target** (List of Object) target for traffic collector (see [below for nested schema](#nestedobjatt--items--target))
- **transmission** (String) transmission type

<a id="nestedobjatt--items--target"></a>
### Nested Schema for `items.target`

Read-Only:

- **dst_ip** (String) single ip address or CIDR
- **dst_port** (Number) destination port for target
- **proto** (Number) protocol for target


