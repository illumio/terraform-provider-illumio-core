---
layout: "illumio-core"
page_title: "illumio-core_pairing_profile Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-pairing-profile"
subcategory: ""
description: |-
  Manages Illumio Pairing Profile
---

# illumio-core_pairing_profile (Resource)

Manages Illumio Pairing Profile

Example Usage
------------

```hcl
resource "illumio-core_pairing_profile" "example" {
    name = "example name"
    enabled = false
    labels {
      href = "/orgs/1/labels/1"
    }
    labels{
      href = "/orgs/1/labels/7"
    }
    allowed_uses_per_key = "50"
    key_lifespan = "50"
    env_label_lock = false
    loc_label_lock = true
    role_label_lock = true
    app_label_lock = true
    log_traffic = false
    log_traffic_lock = true
    visibility_level = "flow_off"
    visibility_level_lock = false 
}

```
## Schema

### Required

- **enabled** (Boolean) The enabled flag of the pairing profile
- **name** (String) The short friendly name of the pairing profile

### Optional

- **agent_software_release** (String) Agent software release associated with this paring profile. Default value: "Default ()"
- **allowed_uses_per_key** (String) The number of times pairing profile keys can be used. Allowed values are range(1-2147483647) and "unlimited". Default value: "unlimited"
- **app_label_lock** (Boolean) Flag that controls whether app label can be overridden from pairing script. Default value: "true"
- **description** (String) The long description of the pairing profile
- **enforcement_mode** (String) Flag that controls whether mode can be overridden from pairing script. Allowed values are "idle", "visibility_only", "full" and "selective". Default value: "visibility_only"
- **enforcement_mode_lock** (Boolean) Flag that controls whether enforcement mode can be overridden from pairing script, Default value: "true"
- **env_label_lock** (Boolean) Flag that controls whether env label can be overridden from pairing script. Default value: "true"
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **key_lifespan** (String) Number of seconds pairing profile keys will be valid for. Allowed values are range(1-2147483647) and "unlimited". Default value: "unlimited"
- **labels** (Block Set) Assigned labels (see [below for nested schema](#nestedblock--labels))
- **loc_label_lock** (Boolean) Flag that controls whether loc label can be overridden from pairing script. Default value: "true"
- **log_traffic** (Boolean) Status of VEN(alternative of status). Default value: false
- **log_traffic_lock** (Boolean) Flag that controls whether log_traffic can be overridden from pairing script. Default value: true
- **role_label_lock** (Boolean) Flag that controls whether role label can be overridden from pairing script. Default value: "true"
- **visibility_level** (String) Visibility level of the agent. Allowed values are "flow_full_detail", "flow_summary", "flow_drops", "flow_off" and "enhanced_data_collection". Default value: "flow_summary"
- **visibility_level_lock** (Boolean) Flag that controls whether visibility_level can be overridden from pairing script. Default value: "true"

### Read-Only

- **created_at** (String) Timestamp when this pairing profile was first created
- **created_by** (Map of String) User who originally created this pairing profile
- **href** (String) URI of this pairing profile
- **is_default** (Boolean) Flag indicating this is default auto-created pairing profile
- **last_pairing_at** (String) Timestamp when this pairing profile was last used for pairing a workload
- **status** (String) State of VEN
- **status_lock** (Boolean) Flag that controls whether status can be overridden from pairing script
- **total_use_count** (Number) The number of times the pairing profile has been used
- **updated_at** (String) Timestamp when this pairing profile was last updated
- **updated_by** (Map of String) User who last updated this pairing profile

<a id="nestedblock--labels"></a>
### Nested Schema for `labels`

Required:

- **href** (String) Label URI


