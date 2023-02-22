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

resource "illumio-core_label" "app_hrm" {
  key   = "app"
  value = "A-HRM"
}

resource "illumio-core_label" "env_qa" {
  key   = "env"
  value = "E-QA"
}

resource "illumio-core_label" "loc_azure" {
  key   = "loc"
  value = "L-AZURE"
}

resource "illumio-core_virtual_service" "hrm_db" {
  name        = "VS-HRM-DB"
  description = "HRM Application database virtual service"
  apply_to    = "host_only"

  service_ports {
    # Illumio uses the IANA protocol numbers to identify the service proto
    proto = "6"  # TCP
    port  = "5432"
  }

  labels {
    href = illumio-core_label.role_db.href
  }

  labels {
    href = illumio-core_label.app_hrm.href
  }

  labels {
    href = illumio-core_label.env_qa.href
  }

  labels {
    href = illumio-core_label.loc_azure.href
  }
}

output "role_label_href" {
  value = illumio-core_label.role_db.href
}

output "app_label_href" {
  value = illumio-core_label.app_hrm.href
}

output "env_label_href" {
  value = illumio-core_label.env_qa.href
}

output "loc_label_href" {
  value = illumio-core_label.loc_azure.href
}

output "virtual_service_href" {
  value = replace(illumio-core_virtual_service.hrm_db.href, "draft", "active")
}
