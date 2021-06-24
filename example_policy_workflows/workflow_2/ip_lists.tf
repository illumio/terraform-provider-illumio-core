resource "illumio-core_ip_list" "policy_2_iplist1" {
  name        = "policy_2_iplist1"
  description = "Ip list testing."
  ip_ranges {
    from_ip = illumio-core_workload_interface.vm10_intf_1.address
  }
  ip_ranges {
    from_ip = illumio-core_workload_interface.vm10_intf_2.address
  }
  ip_ranges {
    from_ip = illumio-core_workload_interface.vm10_intf_3.address
  }
}

resource "illumio-core_ip_list" "policy_2_iplist2" {
  name        = "policy_2_iplist2"
  description = "Ip list testing."
  ip_ranges {
    from_ip = illumio-core_workload_interface.vm11_intf_1.address
  }
  ip_ranges {
    from_ip = illumio-core_workload_interface.vm11_intf_2.address
  }
  ip_ranges {
    from_ip = illumio-core_workload_interface.vm11_intf_3.address
  }
}

