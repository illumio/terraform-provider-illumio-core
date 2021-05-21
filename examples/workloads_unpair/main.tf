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

resource "illumio-core_workloads_unpair" "example" {
  workloads {
    href = "/orgs/1/workloads/7c3789ea-661b-49c3-b3ba-8eb180f5f3d2"
  }
  workloads {
    href = "/orgs/1/workloads/1d230cf4-6718-44b8-8ffa-64383a4dbee1"
  }
  workloads {
    href = "/orgs/1/workloads/11635f19-625f-436c-a299-43d1883145d5"
  }
}
