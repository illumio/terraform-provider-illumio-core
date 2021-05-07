terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
  # pce_host              = "https://2x2devtest59.ilabs.io:8443"
  # api_username          = ""
  # api_secret            = ""
  request_timeout = 30
  org_id          = 1
}

data "illumio-core_syslog_destinations" "name" {

}