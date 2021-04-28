terraform {
  required_providers {
    illumio = {
      version = "0.1"
      source  = "illumio.com/labs/illumio"
    }
  }
}

provider "illumio" {
  # pce_host              = "https://2x2devtest59.ilabs.io:8443"
  # api_username          = ""
  # api_secret            = ""
  request_timeout = 30
  org_id          = 1
}

data "illumio-core_service" "example" {
  service_id = 49
}

resource "illumio-core_service" "example_with_service_port_1" {
  name = "example"
  service_port {
    proto = "-1"
  }

  service_port {
    proto = "6"
    port  = "15"
  }

  service_port {
    proto     = "1"
    icmp_type = "5"
    icmp_code = "5"
  }
}

resource "illumio-core_service" "example_with_windows_service" {
  name         = "example"
  process_name = "value"
  windows_service {
    proto = "-1"
  }

  windows_service {
    proto        = "6"
    process_name = "example"
  }

  windows_service {
    proto     = "1"
    icmp_type = "5"
    icmp_code = "5"
  }

}
