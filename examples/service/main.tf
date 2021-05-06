terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
  # pce_host              = "https://2x2devtest59.ilabs.io:8443"
  # api_username          = ""
  # api_secret            = ""
  request_timeout = 30
  org_id          = 1
}

data "illumio-core_service" "example" {
  href = "/orgs/1/sec_policy/draft/services/3"
}

resource "illumio-core_service" "example_with_service_ports" {
  name = "example"
  service_ports {
    proto = "-1"
  }

  service_ports {
    proto = "6"
    port  = "15"
  }

  service_ports {
    proto     = "1"
    icmp_type = "5"
    icmp_code = "5"
  }
}

resource "illumio-core_service" "example_with_windows_services" {
  name         = "example"
  process_name = "value"
  windows_services {
    proto = "-1"
  }

  windows_services {
    proto        = "6"
    process_name = "example"
  }

  windows_services {
    proto     = "1"
    icmp_type = "5"
    icmp_code = "5"
  }

}
