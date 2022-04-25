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

# NOTE: the `illumio-core_firewall_settings` resource cannot be created.
# For this example to work, the PCE firewall settings must be imported into terraform with
#
# terraform import illumio-core_firewall_settings.current "/orgs/$ILLUMIO_PCE_ORG_ID/sec_policy/draft/firewall_settings"
resource "illumio-core_firewall_settings" "current" {
  ike_authentication_type = "psk"
}

data "illumio-core_firewall_settings" "current" {
  href = illumio-core_firewall_settings.current.href
}
