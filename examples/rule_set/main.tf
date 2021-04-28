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

data "illumio-core_rule_set" "d_rs" {
  rule_set_id = 122
}

resource "illumio-core_rule_set" "rs" {
  name = "terraform-rs4"

  ip_tables_rule {
    description = "some des"
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

  scope {
    label {
      href = "/orgs/1/labels/69"
    }    
  }

  scope {
    label {
      href = "/orgs/1/labels/94"
    }
  }

  scope {
    label_group {
      href = "/orgs/1/sec_policy/draft/label_groups/65d0ad0f-329a-4ddc-8919-bd0220051fc7"
    }
  }

  rule {
    enabled = false
    resolve_labels_as {
      consumers = ["workloads", "workloads"]
      providers = ["workloads"]
    }

    consumer {
      actors = "ams"
    }

    illumio_provider {
      label {
        href = "/orgs/1/labels/715"
      }
    }

    illumio_provider {
      label {
        href = "/orgs/1/labels/294"
      }
    }

    ingress_service {
      proto = 6
      port  = 4
    }
  }
}