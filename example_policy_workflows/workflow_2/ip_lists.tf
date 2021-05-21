resource "illumio-core_ip_list" "policy_2_iplist1" {
    name = "policy_2_iplist1"
    description = "Ip list testing."

    dynamic "ip_ranges" {
      for_each = illumio-core_workload.vm10.interfaces
        content {
            from_ip = ip_ranges.value.address
        }
    } 
}

resource "illumio-core_ip_list" "policy_2_iplist2" {
    name = "policy_2_iplist2"
    description = "Ip list testing."
    dynamic "ip_ranges" {
      for_each = illumio-core_workload.vm11.interfaces
        content {
            from_ip = ip_ranges.value.address
        }
    }  
}

