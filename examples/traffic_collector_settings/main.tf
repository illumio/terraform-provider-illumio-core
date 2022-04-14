terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  pce_host     = var.pce_url
  org_id       = var.pce_org_id
  api_username = var.pce_api_key
  api_secret   = var.pce_api_secret
}

resource "illumio-core_traffic_collector_settings" "drop_local_tcp" {
  transmission = "broadcast"
  action       = "drop"

  # drop all localhost TCP traffic
  target {
    dst_ip = "127.0.0.1"
    proto  = "6"
  }
}

data "illumio-core_traffic_collector_settings" "drop_local_tcp" {
  href = illumio-core_traffic_collector_settings.drop_local_tcp.href
}
