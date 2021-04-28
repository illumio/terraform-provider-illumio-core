---
layout: "illumio-core"
page_title: "illumio-core_pairing_profile Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-pairing-profile"
subcategory: ""
description: |-
  Represents Illumio Pairing Profile
---

# illumio-core_pairing_profile (Data Source)

Represents Illumio Pairing Profile

Example Usage
------------

```hcl
data "illumio-core_pairing_profile" "example"{
    pairing_profile_id = 1
}
```

## Schema

### Required

- **pairing_profile_id** (Number) Numerical ID of pairing profile

### Read-Only

- **agent_software_release** (String) Agent software release associated with this paring profile
- **allowed_uses_per_key** (String) The number of times the pairing profile can be use
- **app_label_lock** (Boolean) Flag that controls whether app label can be overridden from pairing script
- **caps** (List of String) CAP
- **created_at** (String) Timestamp when this pairing profile was first created
- **created_by** (Map of String) User who originally created this pairing profile
- **description** (String) The long description of the pairing profile
- **enabled** (Boolean) The enabled flag of the pairing profile
- **enforcement_mode** (String) Filter by mode
- **enforcement_mode_lock** (Boolean) Flag that controls whether app label can be overridden from pairing script
- **env_label_lock** (Boolean) Flag that controls whether env label can be overridden from pairing script
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **href** (String) URI of this pairing profile
- **is_default** (Boolean) Flag indicating this is default auto-created pairing profile
- **key_lifespan** (String) Number of seconds pairing profile keys will be valid for
- **labels** (List of Object) Assigned labels (see [below for nested schema](#nestedatt--labels))
- **last_pairing_at** (String) Timestamp when this pairing profile was last used for pairing a workload
- **loc_label_lock** (Boolean) Flag that controls whether loc label can be overridden from pairing script
- **log_traffic** (Boolean) State of VEN
- **log_traffic_lock** (Boolean) Flag that controls whether log_traffic can be overridden from pairing script
- **name** (String) The short friendly name of the pairing profile
- **role_label_lock** (Boolean) Flag that controls whether role label can be overridden from pairing script
- **status** (String) State of VEN
- **status_lock** (Boolean) Flag that controls whether status can be overridden from pairing script
- **total_use_count** (Number) The number of times the pairing profile has been used
- **updated_at** (String) Timestamp when this pairing profile was last updated
- **updated_by** (Map of String) User who last updated this pairing profile
- **visibility_level** (String) Visibility level of the workload
- **visibility_level_lock** (Boolean) Flag that controls whether visibility_level can be overridden from pairing script

<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Read-Only:

- **href** (String) Label URI


