terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  pce_host     = var.pce_url
  org_id       = var.pce_org_id
  api_username = var.pce_api_key
  api_secret   = var.pce_api_secret
}

# NOTE: the `illumio-core_organization_settings` resource cannot be created.
# For this example to work, the PCE organization settings must be imported into terraform with
#
# terraform import illumio-core_organization_settings.current "/orgs/$ILLUMIO_PCE_ORG_ID/settings/events"
resource "illumio-core_organization_settings" "current" {
  audit_event_retention_seconds = 2592000
  audit_event_min_severity = "informational"
  format = "JSON"
}

data "illumio-core_organization_settings" "current" {}
