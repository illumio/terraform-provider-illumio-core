resource "illumio-core_traffic_collector_settings" "drop_local_tcp" {
  transmission = "broadcast"
  action       = "drop"

  # drop all localhost TCP traffic
  target {
    dst_ip = "127.0.0.1"
    proto  = "6"
  }
}
