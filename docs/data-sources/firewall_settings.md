---
layout: "illumio-core"
page_title: "illumio-core_firewall_settings Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-firewall-settings"
subcategory: ""
description: |-
  Represents Illumio Firewall Settings
---

# illumio-core_firewall_settings (Data Source)

Represents Illumio Firewall Settings


Example Usage
------------

```hcl
data "illumio-core_firewall_settings" "example" {
    href = "/orgs/1/sec_policy/draft/firewall_settings"
}

```

## Schema

### Required

- **href** (String) URI of Firewall Settings

### Read-Only

- **blocked_connection_reject_scopes** (List of List of Object) scopes for reject connections
- **containers_inherit_host_policy_scopes** (List of List of Object) scopes for container inherit host policy
- **created_at** (String) Timestamp when these firewall settings were first created
- **created_by** (Map of String) User who created this resource
- **deleted_at** (String) Timestamp when these firewall settings were deleted
- **deleted_by** (Map of String) User who deleted this resource
- **firewall_coexistence** (List of Object) Firewall coexistence configuration (see [below for nested schema](#nestedatt--firewall_coexistence))
- **ike_authentication_type** (String) IKE authentication type to use for IPsec (SecureConnect and Machine Authentication)
- **loopback_interfaces_in_policy_scopes** (List of List of Object) scopes for loopback interfaces
- **static_policy_scopes** (List of List of Object) scopes for static policy
- **update_type** (String) Type of Update
- **updated_at** (String) Timestamp when these firewall settings were last updated
- **updated_by** (Map of String) User who last updated this resource

<a id="nestedatt--firewall_coexistence"></a>
### Nested Schema for `firewall_coexistence`

Read-Only:

- **illumio_primary** (Boolean) Whether Illumio is primary firewall or not
- **scope** (List of Object) (see [below for nested schema](#nestedobjatt--firewall_coexistence--scope))
- **workload_mode** (String) Match criteria to select workload(s)

<a id="nestedobjatt--firewall_coexistence--scope"></a>
### Nested Schema for `firewall_coexistence.scope`

Read-Only:

- **href** (String) Href of Label


