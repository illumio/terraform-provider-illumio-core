resource "illumio-core_workload" "vm9" {
  name = "vm9"
  labels {
    href = illumio-core_label.policy_2_role2.href
  }
  enforcement_mode = "selective"
}

resource "illumio-core_workload" "vm8" {
  name = "vm8"
  labels {
    href = illumio-core_label.policy_2_role2.href
  }
  labels {
    href = illumio-core_label.policy_2_app2.href
  }
  labels {
    href = illumio-core_label.policy_2_env2.href
  }
  labels {
    href = illumio-core_label.policy_2_loc2.href
  }
  enforcement_mode = "idle"
}

resource "illumio-core_workload" "vm7" {
  name = "vm7"
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
  enforcement_mode = "idle"
}

resource "illumio-core_workload" "vm6" {
  name = "vm6"
  labels {
    href = illumio-core_label.policy_2_role2.href
  }
  labels {
    href = illumio-core_label.policy_2_app2.href
  }
  labels {
    href = illumio-core_label.policy_2_env2.href
  }
  labels {
    href = illumio-core_label.policy_2_loc2.href
  }
  enforcement_mode = "full"
}

resource "illumio-core_workload" "vm5" {
  name = "vm5"
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
  enforcement_mode = "full"
}

resource "illumio-core_workload" "vm4" {
  name = "vm4"
  labels {
    href = illumio-core_label.policy_2_role2.href
  }
  labels {
    href = illumio-core_label.policy_2_app2.href
  }
  labels {
    href = illumio-core_label.policy_2_env2.href
  }
  labels {
    href = illumio-core_label.policy_2_loc2.href
  }
  enforcement_mode = "visibility_only"
}

resource "illumio-core_workload" "vm3" {
  name = "vm3"
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
  enforcement_mode = "visibility_only"
}

resource "illumio-core_workload" "vm2" {
  name = "vm2"
  labels {
    href = illumio-core_label.policy_2_role2.href
  }
  labels {
    href = illumio-core_label.policy_2_app2.href
  }
  labels {
    href = illumio-core_label.policy_2_env2.href
  }
  labels {
    href = illumio-core_label.policy_2_loc2.href
  }
  enforcement_mode = "selective"
}


resource "illumio-core_workload" "vm1" {
  name = "vm1"
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
  enforcement_mode = "selective"
}
