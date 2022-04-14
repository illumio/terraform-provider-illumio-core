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

resource "illumio-core_label" "role_db" {
  key   = "role"
  value = "R-DB"
}

data "illumio-core_label" "role_db" {
  href = illumio-core_label.role_db.href
}
