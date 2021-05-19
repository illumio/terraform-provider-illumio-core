resource "illumio-core_virtual_service" "policy-1-vservice1" {
    name = "policy-1-vservice1"
    labels {
      href = illumio-core_label.policy-1-role2.href
    }
    labels {
      href = illumio-core_label.policy-1-app2.href
    }
    labels {
      href = illumio-core_label.policy-1-env3.href
    }
    labels {
      href = illumio-core_label.policy-1-loc1.href
    }
    service_ports {
      port = "16000"
      proto = 6
    }
    service_ports {
      port = "16000"
      proto = 17
    }
    apply_to = "host_only"
}

resource "illumio-core_virtual_service" "policy-1-vservice2" {
    name = "policy-1-vservice2"
    labels {
      href = illumio-core_label.policy-1-role1.href
    }
    labels {
      href = illumio-core_label.policy-1-app4.href
    }
    labels {
      href = illumio-core_label.policy-1-env4.href
    }
    labels {
      href = illumio-core_label.policy-1-loc1.href
    }
    service_ports {
      port = "17000"
      proto = 6
    }
    service_ports {
      port = "17000"
      proto = 17
    }
    apply_to = "host_only"
}