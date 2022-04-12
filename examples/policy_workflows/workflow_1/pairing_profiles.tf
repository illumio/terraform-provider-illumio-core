resource "illumio-core_pairing_profile" "policy-1-paring-profile-1" {
    name = "pairing_policy_1"
    description = "This pairing policy will be used to test the feature."
    enforcement_mode = "full" #cannot set illuminated to enforcement_mode 
    enabled  = true
    allowed_uses_per_key  = 10
    key_lifespan  = 604800
    enforcement_mode_lock = false
    env_label_lock  = false
}