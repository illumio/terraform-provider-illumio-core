# NOTE: the `illumio-core_organization_settings` resource cannot be created.
# For this example to work, the PCE organization settings must be imported into terraform with
#
# terraform import illumio-core_organization_settings.current "/orgs/$ILLUMIO_PCE_ORG_ID/settings/events"
resource "illumio-core_organization_settings" "current" {
  audit_event_retention_seconds = 2592000
  audit_event_min_severity = "informational"
  format = "JSON"
}
