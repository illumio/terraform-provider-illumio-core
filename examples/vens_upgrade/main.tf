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

resource "illumio-core_vens_upgrade" "example" {
  release = "1.2"
  vens {
    href = "/orgs/1/vens/e6eaccb3-39b0-44db-907d-d61c6ff1f8f6"
  }
  vens {
    href = "/orgs/1/vens/8754058f-819f-4c50-91f1-da6e9af28918"
  }
  vens {
    href = "/orgs/1/vens/11635f19-625f-436c-a299-43d1883145d5"
  }
}
