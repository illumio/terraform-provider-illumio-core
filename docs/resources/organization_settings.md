---
layout: "illumio-core"
page_title: "illumio-core_organization_settings Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-organization-settings"
subcategory: ""
description: |-
  Manages Illumio Organization Settings
---

# illumio-core_organization_settings (Resource)

Manages Illumio Organization Settings (***Global Organization Owner access required***)

## Importing  

The `organization_settings` resource cannot be created and must be imported using the command below. This resource determines the URI for import from the provider configuration.  

```sh
$ terraform import illumio-core_firewall_settings.example placeholder
```

After import, configuration changes can be planned and applied as normal.  

Example Usage
------------

```hcl
resource "illumio-core_organization_settings" "example" {
  audit_event_retention_seconds = 2592000  # 30 days
  audit_event_min_severity = "informational"
  format = "JSON"
}
```

## Schema

### Required

- `audit_event_min_severity` (String) Minimum severity level of audit event messages. Allowed values are "error", "warning", and "informational"
- `audit_event_retention_seconds` (Number) The time in seconds an audit event is stored in the database. The value should be between 86400 and 17280000
- `format` (String) The log format (JSON, CEF, LEEF), which applies to all remote Syslog destinations. Allowed values are "JSON", "CEF", and "LEEF"
