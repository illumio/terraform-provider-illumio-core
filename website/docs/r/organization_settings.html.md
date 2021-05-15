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


```hcl
# INFO: cherry-picked attributes from terraform show after import
resource "illumio-core_organization_settings" "example" {
  audit_event_retention_seconds = 7776000
  format = "JSON"
  audit_event_min_severity = "informational"
}
```

## Schema

### Required

- **audit_event_min_severity** (String) Minimum severity level of audit event messages. Allowed values : "error", "warning", and "informational"
- **audit_event_retention_seconds** (Number) The time in seconds an audit event is stored in the database. The value should be in between 86400 and 17280000
- **format** (String) The log format (JSON, CEF, LEEF), which applies to all remote syslog destinations. Allowed values : "JSON", "CEF", and "LEEF"

## Importing ##

This resource can only be imported and can not be created. Use below command to import resource.  This resource auto determines URI based on provider config. So no need of providing URI while importing. 

After importing, Cherry pick the configurable parameters from `terraform show` and paste it into .tf file.

Ref: https://www.terraform.io/docs/import/index.html


```
terraform import illumio-core_organization_settings.example <ANYTHING>
```

