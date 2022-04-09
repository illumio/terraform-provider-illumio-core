resource "illumio-core_unmanaged_workload" "policy-1-workload-1" {
    name = "vm1"
    labels {
      href = illumio-core_label.policy-1-role1.href
    }
    labels {
      href = illumio-core_label.policy-1-app2.href
    }
    labels {
      href = illumio-core_label.policy-1-env2.href
    }
    labels {
      href = illumio-core_label.policy-1-loc2.href
    }
    enforcement_mode = "full"
    #enforcement_mode doesnot have enforced as a value
}

resource "illumio-core_unmanaged_workload" "policy-1-workload-2" {
    name = "vm2"
    labels {
      href = illumio-core_label.policy-1-role1.href
    }
    labels {
      href = illumio-core_label.policy-1-app4.href
    }
    labels {
      href = illumio-core_label.policy-1-env4.href
    }
    labels {
      href = illumio-core_label.policy-1-loc1.href
    }
    enforcement_mode = "full"
    #enforcement_mode doesnot have enforced as a value
}

resource "illumio-core_unmanaged_workload" "policy-1-workload-3" {
    name = "vm3"
    labels {
      href = illumio-core_label.policy-1-role2.href
    }
    labels {
      href = illumio-core_label.policy-1-app2.href
    }
    labels {
      href = illumio-core_label.policy-1-env2.href
    }
    labels {
      href = illumio-core_label.policy-1-loc2.href
    }
    enforcement_mode = "idle"
    #enforcement_mode doesnot have enforced as a value
}

resource "illumio-core_unmanaged_workload" "policy-1-workload-4" {
    name = "vm4"
    labels {
      href = illumio-core_label.policy-1-role2.href
    }
    labels {
      href = illumio-core_label.policy-1-app4.href
    }
    labels {
      href = illumio-core_label.policy-1-env4.href
    }
    labels {
      href = illumio-core_label.policy-1-loc1.href
    }
    enforcement_mode = "idle"
    #enforcement_mode doesnot have enforced as a value
}

resource "illumio-core_unmanaged_workload" "policy-1-workload-5" {
    name = "vm5"
    labels {
      href = illumio-core_label.policy-1-role1.href
    }
    labels {
      href = illumio-core_label.policy-1-app1.href
    }
    labels {
      href = illumio-core_label.policy-1-env1.href
    }
    labels {
      href = illumio-core_label.policy-1-loc2.href
    }
    enforcement_mode = "selective"
    #enforcement_mode doesnot have enforced as a value
}

resource "illumio-core_unmanaged_workload" "policy-1-workload-6" {
    name = "vm6"
    labels {
      href = illumio-core_label.policy-1-role1.href
    }
    labels {
      href = illumio-core_label.policy-1-app4.href
    }
    labels {
      href = illumio-core_label.policy-1-env3.href
    }
    labels {
      href = illumio-core_label.policy-1-loc1.href
    }
    enforcement_mode = "selective"
    #enforcement_mode doesnot have enforced as a value
}

resource "illumio-core_unmanaged_workload" "policy-1-workload-7" {
    name = "vm7"
    labels {
      href = illumio-core_label.policy-1-role2.href
    }
    labels {
      href = illumio-core_label.policy-1-app1.href
    }
    labels {
      href = illumio-core_label.policy-1-env1.href
    }
    labels {
      href = illumio-core_label.policy-1-loc2.href
    }
    enforcement_mode = "visibility_only"
    #enforcement_mode doesnot have enforced as a value
}

resource "illumio-core_unmanaged_workload" "policy-1-workload-8" {
    name = "vm8"
    labels {
      href = illumio-core_label.policy-1-role2.href
    }
    labels {
      href = illumio-core_label.policy-1-app4.href
    }
    labels {
      href = illumio-core_label.policy-1-env3.href
    }
    labels {
      href = illumio-core_label.policy-1-loc1.href
    }
    enforcement_mode = "visibility_only"
    #enforcement_mode doesnot have enforced as a value
}