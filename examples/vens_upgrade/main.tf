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

resource "illumio-core_vens_upgrade" "example" {
  release = "21.2.0-7828"
  vens {
    href = "/orgs/1/vens/2820775a-b8bd-4932-a812-50ee0df0ccaf"
  }
  vens {
    href = "/orgs/1/vens/1f7cc136-40c8-4e92-b810-c70afe649291"
  }
  vens {
    href = "/orgs/1/vens/11635f19-625f-436c-a299-43d1883145d5"
  }
}
