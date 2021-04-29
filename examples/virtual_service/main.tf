terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
//  pce_host              = "https://pce.my-company.com:8443"
//  api_username          = "api_xxxxxx"
//  api_secret            = "big-secret"
  request_timeout       = 30
  org_id                = 1
}

data "illumio-core_virtual_service" "vs_1"{
  virtual_service_id = "69f1fcc7-94f0-4e42-b9a8-e722038e6dda"
}

resource "illumio-core_virtual_service" "vs_2" {
  name = "Virtual Service 1 1"
  description = "Virtual Service to apply on host"
  apply_to = "host_only"
  service_ports {
    proto = 6
  }
  service_ports {
    proto = 17
    port = 80
    to_port = 443
  }
  service_addresses {
    fqdn = "*.illumio.com"
  }
  service_addresses {
    ip = "1.1.1.1"
    port = "80"
  }
  service_addresses {
    ip = "1.1.1.2"
    network_href = "/orgs/1/networks/b8007bd8-4b16-41b5-b500-5ea236d49d61"
  }
  labels {
    href = "/orgs/1/labels/1"
  }
  ip_overrides = [ "1.2.3.4" ]
}

