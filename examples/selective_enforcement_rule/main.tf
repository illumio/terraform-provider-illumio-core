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

data "illumio-core_selective_enforcement_rule" "rule15" {
  ser_id = 15
}

resource "illumio-core_selective_enforcement_rule" "rule15r16" {
  name = "SER rule 1"
  scope {
    label {
      href = "/orgs/1/labels/69"
    }
    label {
      href = "/orgs/1/labels/294"
    }
    label_group {
      href = "/orgs/1/sec_policy/draft/label_groups/523f5cd0-2126-4b30-bb40-a9fd19429dd7"
    }

  }

  enforced_services {
    href = "/orgs/1/sec_policy/draft/services/3"
  }
}
