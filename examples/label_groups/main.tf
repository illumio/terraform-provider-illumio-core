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

resource "illumio-core_label" "env_dev" {
  key   = "env"
  value = "E-DEV"
}

resource "illumio-core_label" "env_test" {
  key   = "env"
  value = "E-TEST"
}

resource "illumio-core_label" "env_stage" {
  key   = "env"
  value = "E-STAGE"
}

resource "illumio-core_label" "env_prod" {
  key   = "env"
  value = "E-PROD"
}

resource "illumio-core_label_group" "env_preprod" {
  key         = "env"
  name        = "LG-E-PREPROD"
  description = "Pre-production environments Label Group"

  labels {
    href = illumio-core_label.env_dev.href
  }

  labels {
    href = illumio-core_label.env_test.href
  }

  labels {
    href = illumio-core_label.env_stage.href
  }
}

resource "illumio-core_label_group" "env_all" {
  key         = "env"
  name        = "LG-E-ALL"
  description = "All environments Label Group"

  sub_groups {
    href = illumio-core_label_group.env_preprod.href
  }

  labels {
    href = illumio-core_label.env_prod.href
  }
}

data "illumio-core_label_groups" "env_groups" {
	# supports partial match lookups
  name = "LG-E"

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_label_group.env_all,
  ]
}
