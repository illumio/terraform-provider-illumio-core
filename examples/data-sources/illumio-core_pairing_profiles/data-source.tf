resource "illumio-core_label" "role_db" {
  key   = "role"
  value = "R-DB"
}

resource "illumio-core_label" "app_core_services" {
  key   = "app"
  value = "A-CORE-SERVICES"
}

resource "illumio-core_label" "app_web" {
  key   = "app"
  value = "A-WEB"
}

resource "illumio-core_label" "env_dev" {
  key   = "env"
  value = "E-DEV"
}

resource "illumio-core_pairing_profile" "dev_core_services" {
  name                  = "PP-DEV-CORE-SERVICES"
  enabled               = true

  allowed_uses_per_key  = "unlimited"  # infinite uses per key
  key_lifespan          = "3600"       # keys are valid for 1 hour

  role_label_lock       = false
  app_label_lock        = true
  env_label_lock        = true
  loc_label_lock        = false

  log_traffic           = false
  log_traffic_lock      = true

  visibility_level      = "flow_off"
  visibility_level_lock = false

  labels {
    href = illumio-core_label.app_core_services.href
  }

  labels {
    href = illumio-core_label.env_dev.href
  }
}

resource "illumio-core_pairing_profile" "dev_web_db" {
  name                  = "PP-DEV-WEB-DB"
  enabled               = true

  allowed_uses_per_key  = "2"    # limit to 2 uses per key
  key_lifespan          = "3600" # keys are valid for 1 hour

  role_label_lock       = true
  app_label_lock        = true
  env_label_lock        = true
  loc_label_lock        = false

  log_traffic           = false
  log_traffic_lock      = true

  visibility_level      = "flow_off"
  visibility_level_lock = false

  labels {
    href = illumio-core_label.role_db.href
  }

  labels {
    href = illumio-core_label.app_core_services.href
  }

  labels {
    href = illumio-core_label.env_dev.href
  }
}

data "illumio-core_pairing_profiles" "dev" {
  name = "PP-DEV"

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_pairing_profile.dev_core_services,
    illumio-core_pairing_profile.dev_web_db,
  ]
}
