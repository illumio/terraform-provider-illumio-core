resource "illumio-core_virtual_service" "policy_2_vservice1" {
  name     = "policy_2_vservice1"
  apply_to = "host_only"
  labels {
    href = illumio-core_label.policy_2_role2.href
  }
  labels {
    href = illumio-core_label.policy_2_app2.href
  }
  labels {
    href = illumio-core_label.policy_2_env2.href
  }
  labels {
    href = illumio-core_label.policy_2_loc2.href
  }

  service_ports {
    port  = 26000
    proto = 6
  }

  service_ports {
    port  = 26000
    proto = 17
  }

}

resource "illumio-core_virtual_service" "policy_2_vservice2" {
  name     = "policy_2_vservice2"
  apply_to = "host_only"

  labels {
    href = illumio-core_label.policy_2_role1.href
  }
  labels {
    href = illumio-core_label.policy_2_app1.href
  }
  labels {
    href = illumio-core_label.policy_2_env1.href
  }
  labels {
    href = illumio-core_label.policy_2_loc1.href
  }

  service_ports {
    port    = 16000
    to_port = 16005
    proto   = 6
  }

  service_ports {
    port    = 16000
    to_port = 16005
    proto   = 17
  }

  service_ports {
    port  = 16500
    proto = 6
  }

  service_ports {
    port  = 16500
    proto = 17
  }

}
