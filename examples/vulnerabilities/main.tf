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

resource "illumio-core_vulnerabilities" "example1" {
  vulnerability {
    reference_id = "example2"
    name         = "name1"
    score        = 2
  }

  vulnerability {
    reference_id = "example2"
    name         = "name2"
    score        = 3
    cve_ids      = ["someid"]
    description  = "example description"
  }
}

data "illumio-core_vulnerability" "name" {
  reference_id = "exampl1e"
}
