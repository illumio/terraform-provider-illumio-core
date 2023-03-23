resource "illumio-core_workload_settings" "example" {
  workload_disconnected_timeout_seconds {
    value = 3600
  }

  workload_goodbye_timeout_seconds {
    value = 900
  }
}
