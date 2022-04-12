resource "illumio-core_label_group" "policy-1-role_lg1" {
    name = "policy-1-role_lg1"
    key = "role"
    description = "policy-1-role_lg1"
    labels {
      href = illumio-core_label.policy-1-role1.href
    }
}

resource "illumio-core_label_group" "policy-1-app_lg1" {
    name = "policy-1-app_lg1"
    key = "app"
    description = "policy-1-app_lg1"
    labels {
      href = illumio-core_label.policy-1-app1.href
    }
}

resource "illumio-core_label_group" "policy-1-env_lg1" {
    name = "policy-1-env_lg1"
    key = "env"
    description = "policy-1-env_lg1"
    labels {
      href = illumio-core_label.policy-1-env1.href
    }
}

resource "illumio-core_label_group" "policy-1-loc_lg1" {
    name = "policy-1-loc_lg1"
    key = "loc"
    description = "policy-1-loc_lg1"
    labels {
      href = illumio-core_label.policy-1-loc1.href
    }
}


resource "illumio-core_label_group" "policy-1-role_lg2" {
    name = "policy-1-role_lg2"
    key = "role"
    description = "policy-1-role_lg2"
    labels {
      href = illumio-core_label.policy-1-role2.href
    }
    labels {
      href = illumio-core_label.policy-1-role3.href
    }
}

resource "illumio-core_label_group" "policy-1-app_lg2" {
    name = "policy-1-app_lg2"
    key = "app"
    description = "policy-1-app_lg2"
    labels {
      href = illumio-core_label.policy-1-app2.href
    }
    labels {
      href = illumio-core_label.policy-1-app3.href
    }
}

resource "illumio-core_label_group" "policy-1-env_lg2" {
    name = "policy-1-env_lg2"
    key = "env"
    description = "policy-1-env_lg2"
    labels {
      href = illumio-core_label.policy-1-env2.href
    }
    labels {
      href = illumio-core_label.policy-1-env3.href
    }
}

resource "illumio-core_label_group" "policy-1-loc_lg2" {
    name = "policy-1-loc_lg2"
    key = "loc"
    description = "policy-1-loc_lg2"
    labels {
      href = illumio-core_label.policy-1-loc2.href
    }
    labels {
      href = illumio-core_label.policy-1-loc3.href
    }
}

resource "illumio-core_label_group" "policy-1-role_lg3" {
    name = "policy-1-role_lg3"
    key = "role"
    description = "policy-1-role_lg3"
    labels {
      href = illumio-core_label.policy-1-role4.href
    }
    sub_groups {
      href = illumio-core_label_group.policy-1-role_lg1.href
    }
}

resource "illumio-core_label_group" "policy-1-app_lg3" {
    name = "policy-1-app_lg3"
    key = "app"
    description = "policy-1-app_lg3"
    labels {
      href = illumio-core_label.policy-1-app4.href
    }
    sub_groups {
      href = illumio-core_label_group.policy-1-app_lg1.href
    }
}

resource "illumio-core_label_group" "policy-1-env_lg3" {
    name = "policy-1-env_lg3"
    key = "env"
    description = "policy-1-env_lg3"
    labels {
      href = illumio-core_label.policy-1-env4.href
    }
    sub_groups {
      href = illumio-core_label_group.policy-1-env_lg1.href
    }
}

resource "illumio-core_label_group" "policy-1-loc_lg3" {
    name = "policy-1-loc_lg3"
    key = "loc"
    description = "policy-1-loc_lg3"
    labels {
      href = illumio-core_label.policy-1-loc4.href
    }
    sub_groups {
      href = illumio-core_label_group.policy-1-loc_lg1.href
    }
}