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

resource "illumio-core_label" "env_prod" {
  key   = "env"
  value = "E-PROD"
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

# use the services data source to search against the /services endpoint by name
data "illumio-core_services" "all_services" {
  # all PCE instances define a default Service covering all service ports
  name = "All Services"
  max_results = 1
}

resource "illumio-core_security_rule" "core_services_ringfence" {
  rule_set_href = illumio-core_rule_set.core_services_prod.href

  enabled = true

  resolve_labels_as {
    consumers = ["workloads"]
    providers = ["workloads"]
  }

  consumers {
    actors = "ams"  # special notation meaning "all managed systems" - affects all workloads
  }

  providers {
    actors = "ams"
  }

  ingress_services {
    href = one(data.illumio-core_services.all_services.items[*].href)
  }
}

data "illumio-core_security_rule" "core_services_ringfence" {
  href = illumio-core_security_rule.core_services_ringfence.href
}
