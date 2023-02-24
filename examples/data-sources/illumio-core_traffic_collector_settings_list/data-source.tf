resource "illumio-core_traffic_collector_settings" "drop_http" {
  transmission = "broadcast"
  action       = "drop"

  # drop HTTP traffic for all hosts in the 172.22.0.0/16 subnet
  target {
    dst_ip   = "172.22.0.0/16"
    dst_port = "80"
    proto    = "6"
  }
}

data "illumio-core_traffic_collector_settings_list" "all" {
  depends_on = [
    illumio-core_traffic_collector_settings.drop_http,
  ]
}