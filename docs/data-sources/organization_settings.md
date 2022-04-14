---
layout: "illumio-core"
page_title: "illumio-core_organization_settings Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-organization-settings"
subcategory: ""
description: |-
  Represents Illumio Organization Settings
---

# illumio-core_organization_settings (Data Source)

Represents Illumio Organization Settings

~> the `illumio-core_organization_settings` resource cannot be created. For the example below to work, the PCE organization settings must be imported into terraform with

```sh
$ terraform import illumio-core_organization_settings.example placeholder
```

Example Usage
------------

```hcl
data "illumio-core_organization_settings" "example" {}

resource "illumio-core_organization_settings" "example" {
  audit_event_retention_seconds = 2592000  # 30 days
  audit_event_min_severity = "informational"
  format = "JSON"
}
```

## Schema

### Read-Only

- `audit_event_min_severity` (String) Minimum severity level of audit event messages
- `audit_event_retention_seconds` (Number) The time in seconds an audit event is stored in the database
- `format` (String) The log format (JSON, CEF, LEEF), which applies to all remote Syslog destinations
