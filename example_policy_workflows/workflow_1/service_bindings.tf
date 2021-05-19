resource "illumio-core_service_binding" "policy-1-service-binding-1" {
  workload {
    href = illumio-core_workload.policy-1-workload-2.href
  }
  virtual_service {
    href = replace(illumio-core_virtual_service.policy-1-vservice1.href, "draft", "active")
  }
  external_data_reference = "policy-1-vservice1_1"
  external_data_set       = "policy-1-cmdb"
}

resource "illumio-core_service_binding" "policy-1-service-binding-2" {
  workload {
    href = illumio-core_workload.policy-1-workload-7.href
  }
  virtual_service {
    href = replace(illumio-core_virtual_service.policy-1-vservice2.href, "draft", "active")
  }
  external_data_reference = "policy-1-vservice2_1"
  external_data_set       = "policy-1-cmdb"
}
