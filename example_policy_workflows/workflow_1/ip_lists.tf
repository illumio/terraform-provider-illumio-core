resource "illumio-core_ip_list" "policy-1-iplist1" {
    name = "policy-1-iplist1"
    description = "Ip list testing."
    ip_ranges {
        from_ip = illumio-core_workload_interface.policy-1-workload-interface-1.address
    }
    ip_ranges {
        from_ip = illumio-core_workload_interface.policy-1-workload-interface-2.address
    }
    ip_ranges {
        from_ip = illumio-core_workload_interface.policy-1-workload-interface-3.address
    }
}