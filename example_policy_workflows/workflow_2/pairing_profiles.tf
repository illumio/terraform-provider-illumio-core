resource "illumio-core_pairing_profile" "policy_2_selective_pp" {
  name             = "policy_2_selective_pp"
  description      = "This pairing policy will be used to test the feature."
  enforcement_mode = "selective"
  enabled          = true
  labels {
    href = illumio-core_label.policy_2_role1.href
  }
  labels {
    href = illumio-core_label.policy_2_app1.href
  }
  labels {
    href = illumio-core_label.policy_2_env1.href
  }
  labels {
    href = illumio-core_label.policy_2_loc1.href
  }
}
