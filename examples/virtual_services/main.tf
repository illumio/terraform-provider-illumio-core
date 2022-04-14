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

resource "illumio-core_label" "app_crm" {
  key   = "app"
  value = "A-CRM"
}

resource "illumio-core_label" "app_hrm" {
  key   = "app"
  value = "A-HRM"
}

resource "illumio-core_label" "env_qa" {
  key   = "env"
  value = "E-QA"
}

resource "illumio-core_label" "loc_au" {
  key   = "loc"
  value = "L-AU"
}

resource "illumio-core_service" "mysql" {
  name        = "S-MYSQL"
  description = "TCP and UDP Remote Desktop Protocol ports"

  service_ports {
    # Illumio uses the IANA protocol numbers to identify the service proto
    proto = "6"  # TCP
    port  = "3389"
  }
}

resource "illumio-core_virtual_service" "crm_db" {
  name        = "VS-CRM-DB"
  description = "CRM Application database virtual service"
  apply_to    = "host_only"

  service {
    href = illumio-core_service.mysql.href
  }

  labels {
    href = illumio-core_label.role_db.href
  }

  labels {
    href = illumio-core_label.app_crm.href
  }

  labels {
    href = illumio-core_label.env_qa.href
  }

  labels {
    href = illumio-core_label.loc_au.href
  }
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
    href = illumio-core_label.loc_au.href
  }
}

data "illumio-core_virtual_services" "au_qa_db" {
  labels = jsonencode([
    [
      {
        href = illumio-core_label.role_db.href
      },
      {
        href = illumio-core_label.env_qa.href
      },
      {
        href = illumio-core_label.loc_au.href
      },
    ]
  ])

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_virtual_service.crm_db,
    illumio-core_virtual_service.hrm_db,
  ]
}
