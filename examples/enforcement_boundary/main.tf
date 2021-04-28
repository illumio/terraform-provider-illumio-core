terraform {
  required_providers {
    illumio = {
      version = "0.1"
      source  = "illumio.com/labs/illumio"
    }
  }
}


provider "illumio" {
  //  pce_host              = "https://pce.my-company.com:8443"
  //  api_username          = "api_xxxxxx"
  //  api_secret            = "big-secret"
  request_timeout = 30
  org_id          = 1
}

resource "illumio-core_enforcement_boundary" "test" {
    name = "testing eb"
    ingress_service {
      href = "/orgs/1/sec_policy/draft/services/3"
    }
    consumer {
      ip_list {
        href = "/orgs/1/sec_policy/draft/ip_lists/1"
      }
    }
    illumio_provider {
      label {
        href = "/orgs/1/labels/1"
      }
    }
}

data "illumio-core_enforcement_boundary" "test" {
  enforcement_boundary_id = 1
}
