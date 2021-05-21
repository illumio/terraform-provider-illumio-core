terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
    }
  }
}


provider "illumio-core" {
  //  pce_host              = "https://pce.my-company.com:8443"
  //  api_username          = "api_xxxxxx"
  //  api_secret            = "big-secret"
  request_timeout = 30
  org_id          = 1
}

resource "illumio-core_enforcement_boundary" "example" {
  name = "testing eb"
  ingress_services {
    href = "/orgs/1/sec_policy/draft/services/3"
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  providers {
    label {
      href = "/orgs/1/labels/1"
    }
  }
}

data "illumio-core_enforcement_boundary" "example" {
  href = "/orgs/1/sec_policy/draft/enforcement_boundaries/57"
}
