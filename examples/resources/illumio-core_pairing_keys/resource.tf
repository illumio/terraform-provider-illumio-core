resource "illumio-core_pairing_profile" "example" {
  name                  = "PP-EXAMPLE"
  enabled               = true

  allowed_uses_per_key  = "unlimited"
  key_lifespan          = "3600"

  role_label_lock       = false
  app_label_lock        = false
  env_label_lock        = false
  loc_label_lock        = false

  log_traffic           = false
  log_traffic_lock      = true

  visibility_level      = "flow_off"
  visibility_level_lock = false
}

resource "illumio-core_pairing_keys" "example" {
  pairing_profile_href = illumio-core_pairing_profile.example.href
  token_count = 1
}
