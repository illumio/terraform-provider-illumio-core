resource "illumio-core_label_group" "policy_2_role_lg3" {
  name        = "policy_2_role_lg3"
  description = "policy_2_role_lg3"

  key = "role"
  labels {
    href = illumio-core_label.policy_2_role3.href
  }
}

resource "illumio-core_label_group" "policy_2_loc_lg2" {
  name        = "policy_2_loc_lg2"
  description = "policy_2_loc_lg2"

  key = "loc"
  labels {
    href = illumio-core_label.policy_2_loc2.href
  }
  sub_groups {
    href = illumio-core_label_group.policy_2_loc_lg1.href
  }
}

resource "illumio-core_label_group" "policy_2_env_lg2" {
  name        = "policy_2_env_lg2"
  description = "policy_2_env_lg2"

  key = "env"
  labels {
    href = illumio-core_label.policy_2_env2.href
  }
  sub_groups {
    href = illumio-core_label_group.policy_2_env_lg1.href
  }
}

resource "illumio-core_label_group" "policy_2_app_lg2" {
  name        = "policy_2_app_lg2"
  description = "policy_2_app_lg2"

  key = "app"
  labels {
    href = illumio-core_label.policy_2_app2.href
  }
  sub_groups {
    href = illumio-core_label_group.policy_2_app_lg1.href
  }
}

resource "illumio-core_label_group" "policy_2_role_lg2" {
  name        = "policy_2_role_lg2"
  description = "policy_2_role_lg2"
  key         = "role"
  labels {
    href = illumio-core_label.policy_2_role2.href
  }
  sub_groups {
    href = illumio-core_label_group.policy_2_role_lg1.href
  }
}

resource "illumio-core_label_group" "policy_2_loc_lg1" {
  name        = "policy_2_loc_lg1"
  description = "policy_2_loc_lg1"
  key         = "loc"
  labels {
    href = illumio-core_label.policy_2_loc1.href
  }
}

resource "illumio-core_label_group" "policy_2_env_lg1" {
  name        = "policy_2_env_lg1"
  description = "policy_2_env_lg1"
  key         = "env"
  labels {
    href = illumio-core_label.policy_2_env1.href
  }
}

resource "illumio-core_label_group" "policy_2_app_lg1" {
  name        = "policy_2_app_lg1"
  description = "policy_2_app_lg1"
  key         = "app"
  labels {
    href = illumio-core_label.policy_2_app1.href
  }
}

resource "illumio-core_label_group" "policy_2_role_lg1" {
  name        = "policy_2_role_lg1"
  description = "policy_2_role_lg1"
  key         = "role"
  labels {
    href = illumio-core_label.policy_2_role1.href
  }
}
