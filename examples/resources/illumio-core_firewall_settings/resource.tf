# NOTE: the `illumio-core_firewall_settings` resource cannot be created.
# For this example to work, the PCE firewall settings must be imported into terraform with
#
# terraform import illumio-core_firewall_settings.current "/orgs/$ILLUMIO_PCE_ORG_ID/sec_policy/draft/firewall_settings"
resource "illumio-core_firewall_settings" "current" {
  ike_authentication_type = "psk"
}
