resource "illumio-core_organization_settings" "example" {
  audit_event_retention_seconds = 2592000
  audit_event_min_severity = "informational"
  format = "JSON"
}
