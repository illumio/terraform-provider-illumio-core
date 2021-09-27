terraform {
  required_providers {
    illumio-core = {
      version = "0.1.0"
      source  = "illumio/illumio-core"
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

resource "illumio-core_service_binding" "example" {
  virtual_service {
    href = "/orgs/1/sec_policy/active/virtual_services/69f1fcc7-94f0-4e42-b9a8-e722038e6dda"
  }
  workload {
    href = "/orgs/1/workloads/673c3148-a419-4ed2-b0e2-30eb538695e7"
  }
}

data "illumio-core_service_binding" "example" {
   href = "/orgs/1/service_bindings/ad730105-d913-4859-b240-857ac4b8621d"
}