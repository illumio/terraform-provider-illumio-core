resource "illumio-core_service_binding" "policy_2_sb1" {
  virtual_service {
    href = replace(illumio-core_virtual_service.policy_2_vservice1.href, "draft", "active")

  }
  workload {
    href = illumio-core_unmanaged_workload.vm9.href
  }
  external_data_reference = "vservice1_1"
  external_data_set       = "cmdb"
}

resource "illumio-core_service_binding" "policy_2_sb2" {
  virtual_service {
    href = replace(illumio-core_virtual_service.policy_2_vservice2.href, "draft", "active")
  }
  workload {
    href = illumio-core_unmanaged_workload.vm2.href
  }
  external_data_reference = "vservice2_1"
  external_data_set       = "cmdb"
}
