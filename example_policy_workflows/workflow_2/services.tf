resource "illumio-core_service" "policy_2_svc1" {
  name         = "policy_2_svc1"
  description  = "itt"
  process_name = "itt"
  service_ports {
    port  = 11000
    proto = 6
  }
  service_ports {
    port  = 11000
    proto = 17
  }
  service_ports {
    port    = 11005
    to_port = 11010
    proto   = 6
  }
  service_ports {
    port    = 11005
    to_port = 11010
    proto   = 17
  }
}

resource "illumio-core_service" "policy_2_svc2" {
  name         = "policy_2_svc2"
  description  = "itt"
  process_name = "itt"
  service_ports {
    port  = 17000
    proto = 6
  }
}

resource "illumio-core_service" "policy_2_svc3" {
  name         = "policy_2_svc3"
  description  = "itt"
  process_name = "itt"
  service_ports {
    port  = 18000
    proto = 6
  }
  service_ports {
    port  = 18000
    proto = 17
  }
  service_ports {
    port  = 18100
    proto = 6
  }
  service_ports {
    port  = 18100
    proto = 17
  }
}

resource "illumio-core_service" "policy_2_svc4" {
  name         = "policy_2_svc4"
  description  = "itt"
  process_name = "itt_java"
  service_ports {
    port  = 19000
    proto = 6
  }
  service_ports {
    port  = 19000
    proto = 17
  }
}

resource "illumio-core_service" "policy_2_svc5" {
  name         = "policy_2_svc5"
  description  = "itt"
  process_name = "itt"
  service_ports {
    port  = 20000
    proto = 6
  }
  service_ports {
    port  = 20000
    proto = 17
  }
}

resource "illumio-core_service" "policy_2_svc6" {
  name         = "policy_2_svc6"
  description  = "itt"
  process_name = "itt"
  service_ports {
    port  = 21000
    proto = 6
  }
  service_ports {
    port  = 21000
    proto = 17
  }
  service_ports {
    port  = 22100
    proto = 6
  }
}

resource "illumio-core_service" "policy_2_svc7" {
  name         = "policy_2_svc7"
  description  = "itt"
  process_name = "itt"
  service_ports {
    port  = 22000
    proto = 6
  }
  service_ports {
    port  = 22000
    proto = 17
  }
}

resource "illumio-core_service" "policy_2_svc8" {
  name         = "policy_2_svc8"
  description  = "itt"
  process_name = "itt"
  service_ports {
    port    = 11005
    to_port = 11008
    proto   = 6
  }
  service_ports {
    port    = 11005
    to_port = 11008
    proto   = 17
  }
}

resource "illumio-core_service" "policy_2_svc9" {
  name         = "policy_2_svc9"
  description  = "itt"
  process_name = "itt"
  service_ports {
    port    = 13000
    to_port = 13003
    proto   = 6
  }
  service_ports {
    port    = 13000
    to_port = 13003
    proto   = 17
  }


}

resource "illumio-core_service" "policy_2_svc10" {
  name         = "policy_2_svc10"
  description  = "itt"
  process_name = "itt_java"
  service_ports {
    port  = 20002
    proto = 6
  }
  service_ports {
    port  = 20002
    proto = 17
  }
  service_ports {
    port  = 20500
    proto = 6
  }
  service_ports {
    port  = 20500
    proto = 17
  }
}

resource "illumio-core_service" "policy_2_svc11" {
  name         = "policy_2_svc11"
  description  = "itt"
  process_name = "itt_java"
  service_ports {
    port  = 20100
    proto = 6
  }
  service_ports {
    port  = 20100
    proto = 17
  }
}

resource "illumio-core_service" "policy_2_E5" {
  name         = "policy_2_E5"
  description  = "itt"
  process_name = "itt_java"
  service_ports {
    port  = 16000
    proto = 6
  }
  service_ports {
    port  = 16001
    proto = 6
  }
  service_ports {
    port  = 16000
    proto = 17
  }
  service_ports {
    port  = 16001
    proto = 17
  }
}

resource "illumio-core_service" "policy_2_EBS1" {
  name         = "EBS1"
  description  = "itt"
  process_name = "itt_java"
  service_ports {
    port  = 12001
    proto = 6
  }
  service_ports {
    port  = 12006
    proto = 6
  }
  service_ports {
    port  = 12001
    proto = 17
  }
  service_ports {
    port  = 12006
    proto = 17
  }
}

resource "illumio-core_service" "policy_2_EBS2" {
  name         = "policy_2_EBS2"
  description  = "itt"
  process_name = "itt_java"
  service_ports {
    port    = 13000
    to_port = 13005
    proto   = 6
  }
  service_ports {
    port    = 13000
    to_port = 13005
    proto   = 17
  }
}

resource "illumio-core_service" "policy_2_EBS3" {
  name         = "policy_2_EBS3"
  description  = "itt"
  process_name = "itt_java"
  service_ports {
    port    = 15001
    to_port = 15003
    proto   = 6
  }
  service_ports {
    port    = 15001
    to_port = 15003
    proto   = 17
  }
}

resource "illumio-core_service" "policy_2_EBS4" {
  name         = "policy_2_EBS4"
  description  = "itt"
  process_name = "itt_java"
  service_ports {
    port    = 17002
    to_port = 17006
    proto   = 6
  }
  service_ports {
    port    = 17002
    to_port = 17006
    proto   = 17
  }
}
