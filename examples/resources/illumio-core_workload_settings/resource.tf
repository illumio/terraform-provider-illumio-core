resource "illumio-core_workload_settings" "current" {
  workload_disconnected_timeout_seconds {
    value = 3600
  }

  workload_goodbye_timeout_seconds {
    value = 900
  }
}
