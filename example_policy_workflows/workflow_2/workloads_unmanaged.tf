resource "illumio-core_unmanaged_workload" "vm10" {
  name = "vm10"
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


resource "illumio-core_workload_interface" "vm10_intf_1" {
  workload_href = illumio-core_unmanaged_workload.vm10.href
  name          = "policy-1-interface-1"
  link_state    = "up"
  address       = "10.20.30.0"
}
resource "illumio-core_workload_interface" "vm10_intf_2" {
  workload_href = illumio-core_unmanaged_workload.vm10.href
  name          = "policy-1-interface-2"
  link_state    = "up"
  address       = "fd00::200:a:1:a21"
}
resource "illumio-core_workload_interface" "vm10_intf_3" {
  workload_href = illumio-core_unmanaged_workload.vm10.href
  name          = "policy-1-interface-3"
  link_state    = "up"
  address       = "10.1.10.34"
}

resource "illumio-core_unmanaged_workload" "vm11" {
  name = "vm11"
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
}

resource "illumio-core_workload_interface" "vm11_intf_1" {
  workload_href = illumio-core_unmanaged_workload.vm11.href
  name          = "policy-1-interface-1"
  link_state    = "up"
  address       = "10.20.30.0"
}
resource "illumio-core_workload_interface" "vm11_intf_2" {
  workload_href = illumio-core_unmanaged_workload.vm11.href
  name          = "policy-1-interface-2"
  link_state    = "up"
  address       = "fd00::200:a:1:a21"
}
resource "illumio-core_workload_interface" "vm11_intf_3" {
  workload_href = illumio-core_unmanaged_workload.vm11.href
  name          = "policy-1-interface-3"
  link_state    = "up"
  address       = "10.1.10.34"
}
