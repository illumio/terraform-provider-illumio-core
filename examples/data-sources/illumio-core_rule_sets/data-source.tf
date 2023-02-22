resource "illumio-core_label" "app_core_services" {
  key   = "app"
  value = "A-CORE-SERVICES"
}

resource "illumio-core_label" "env_dev" {
  key   = "env"
  value = "E-DEV"
}

resource "illumio-core_label" "env_prod" {
  key   = "env"
  value = "E-PROD"
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

resource "illumio-core_rule_set" "core_services_prod" {
  name = "RS-CORE-SERVICES-PROD"

  scopes {
    label {
      href = illumio-core_label.app_core_services.href
    }

    label {
      href = illumio-core_label.env_prod.href
    }
  }
}

data "illumio-core_rule_sets" "core_services" {
	# supports partial match lookups
  name = "RS-CORE-SERVICES-"

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_rule_set.core_services_dev,
    illumio-core_rule_set.core_services_prod,
  ]
}
