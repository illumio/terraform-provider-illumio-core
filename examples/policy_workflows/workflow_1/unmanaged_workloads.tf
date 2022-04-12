resource "illumio-core_unmanaged_workload" "policy-1-workload-9" {
    name = "vm9"
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
    #enforcement_mode doesnot have illuminated as a value
}

resource "illumio-core_unmanaged_workload" "policy-1-workload-10" {
    name = "vm10"
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
    enforcement_mode = "full"
    #enforcement_mode doesnot have illuminated as a value
}

resource "illumio-core_workload_interface" "policy-1-workload-interface-1" {
  workload_href = illumio-core_unmanaged_workload.policy-1-workload-10.href
  name = "policy-1-interface-1"
  link_state = "up"
  address = "10.20.30.0"
}
resource "illumio-core_workload_interface" "policy-1-workload-interface-2" {
  workload_href = illumio-core_unmanaged_workload.policy-1-workload-10.href
  name = "policy-1-interface-2"
  link_state = "up"
  address = "fd00::200:a:1:a21"
}
resource "illumio-core_workload_interface" "policy-1-workload-interface-3" {
  workload_href = illumio-core_unmanaged_workload.policy-1-workload-10.href
  name = "policy-1-interface-3"
  link_state = "up"
  address = "10.1.10.34"
}