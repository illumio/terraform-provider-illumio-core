resource "illumio-core_label" "role_db" {
  key   = "role"
  value = "R-DB"
}

resource "illumio-core_label" "app_crm" {
  key   = "app"
  value = "A-CRM"
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

resource "illumio-core_virtual_service" "example" {
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
