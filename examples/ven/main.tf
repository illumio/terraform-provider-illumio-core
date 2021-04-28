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

data "illumio-core_ven" "ven" {
  ven_id = "be310e96-7486-4de5-8068-61cd1e2298a1"
}

output "name" {
  value = data.illumio-core_ven.ven
}

resource "illumio-core_ven" "name" {
  status = "suspended"
  name = "example name"
  description = "description"
  target_pce_fqdn = "example.fqdn"
}
