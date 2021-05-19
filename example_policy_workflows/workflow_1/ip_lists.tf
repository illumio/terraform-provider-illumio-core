resource "illumio-core_ip_list" "policy-1-iplist1" {
    name = "policy-1-iplist1"
    description = "Ip list testing."

    dynamic "ip_ranges" {
        for_each = illumio-core_workload.policy-1-workload-10.interfaces
        content {
            from_ip = ip_ranges.value.address
        }
    }
}