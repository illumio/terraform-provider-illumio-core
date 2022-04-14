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

resource "illumio-core_label" "app_core_services" {
  key   = "app"
  value = "A-CORE-SERVICES"
}

resource "illumio-core_label" "env_dev" {
  key   = "env"
  value = "E-DEV"
}

resource "illumio-core_rule_set" "core_services_dev" {
  name = "RS-CORE-SERVICES-DEV"

  scopes {
    label {
      href = illumio-core_label.app_core_services.href
    }

    label {
      href = illumio-core_label.env_dev.href
    }
  }
}

data "illumio-core_rule_set" "core_services_dev" {
  href = illumio-core_rule_set.core_services_dev.href
}
