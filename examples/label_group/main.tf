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

data "illumio-core_label_group" "example" {
  href = "/orgs/1/sec_policy/draft/label_groups/db3fc597-e0ee-4391-8a8b-31a0d1acb3b5"
}

resource "illumio-core_label_group" "example" {
  key         = "role"
  name        = "example name"
  description = "example description"
  labels {
    href = "/orgs/1/labels/2"
  }
  labels {
    href = "/orgs/1/labels/1"
  }
}

