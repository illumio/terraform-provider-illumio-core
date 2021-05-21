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


data "illumio-core_firewall_settings" "current" {
    href = "/orgs/1/sec_policy/draft/firewall_settings"
}


# INFO: cherry-picked attributes from terraform show after import 
resource "illumio-core_firewall_settings" "current" {

    ike_authentication_type = "psk"
    blocked_connection_reject_scopes {
        label {
            href = "/orgs/1/labels/1"
        }
        label {
            href = "/orgs/1/labels/12"
        }
        label {
            href = "/orgs/1/labels/14"
        }
        label {
            href = "/orgs/1/labels/787"
        }
    }

    containers_inherit_host_policy_scopes {
        label {
            href = "/orgs/1/labels/1"
        }
        label {
            href = "/orgs/1/labels/11"
        }
        label {
            href = "/orgs/1/labels/14"
        }
        label {
            href = "/orgs/1/labels/787"
        }
    }

    firewall_coexistence {
        illumio_primary = true

        scope {
            href = "/orgs/1/labels/1"
        }
        scope {
            href = "/orgs/1/labels/787"
        }
        scope {
            href = "/orgs/1/labels/788"
        }
        scope {
            href = "/orgs/1/labels/11"
        }

        workload_mode = "enforced"
    }

    loopback_interfaces_in_policy_scopes {
        label {
            href = "/orgs/1/labels/1"
        }
        label {
            href = "/orgs/1/labels/12"
        }
        label {
            href = "/orgs/1/labels/787"
        }
        label {
            href = "/orgs/1/labels/788"
        }
    }

    static_policy_scopes {
        label {
            href = "/orgs/1/labels/1"
        }
        label {
            href = "/orgs/1/labels/12"
        }
        label {
            href = "/orgs/1/labels/14"
        }

        label_group {
            href = "/orgs/1/sec_policy/draft/label_groups/a715cd8f-04f3-4bc1-82bf-d650b01453a5"
        }
    }
}
