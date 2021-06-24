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

data "illumio-core_rule_set" "example" {
  href = "/orgs/1/sec_policy/draft/rule_sets/70"
}

resource "illumio-core_rule_set" "example" {
  name = "example"

  ip_tables_rules {
    description = "example desc"
    actors {
      actors = "ams"
    }
    actors {
      label {
        href = "/orgs/1/labels/69"
      }
    }

    enabled = false

    ip_version = 6
    statements {
      table_name = "nat"
      chain_name = "PREROUTING"
      parameters = "value"
    }
  }

  scopes {
    label {
      href = "/orgs/1/labels/69"
    }
    label_group {
      href = "/orgs/1/sec_policy/draft/label_groups/65d0ad0f-329a-4ddc-8919-bd0220051fc7"
    }
  }

  scopes {
    label {
      href = "/orgs/1/labels/94"
    }
  }
}
