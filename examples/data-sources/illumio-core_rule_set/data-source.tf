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
