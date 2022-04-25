# Define a Rule Set to set policy rules for all web workloads

resource "illumio-core_rule_set" "web_apps" {
  name = "RS-WEB"

  scopes {
    label {
      href = illumio-core_label.app_web.href
    }
  }
}

# Define the security rules that will allow traffic to and between web workloads

resource "illumio-core_security_rule" "web_mysql" {
  rule_set_href = illumio-core_rule_set.web_apps.href

  enabled = true

  resolve_labels_as {
    consumers = ["workloads"]
    providers = ["workloads"]
  }

  consumers {
    label {
      href = illumio-core_label.role_tomcat.href
    }
  }

  providers {
    label {
      href = illumio-core_label.role_db.href
    }
  }

  ingress_services {
    href = illumio-core_service.mysql.href
  }
}

resource "illumio-core_security_rule" "web_http_inbound" {
  rule_set_href = illumio-core_rule_set.web_apps.href

  enabled = true

  resolve_labels_as {
    consumers = ["workloads"]
    providers = ["workloads"]
  }

  consumers {
    ip_list {
      href = local.any_ip_list_href
    }
  }

  providers {
    label {
      href = illumio-core_label.role_tomcat.href
    }
  }

  providers {
    label {
      href = illumio-core_label.env_prod.href
    }
  }

  ingress_services {
    href = illumio-core_service.http.href
  }
}

resource "illumio-core_security_rule" "web_dev_inbound" {
  rule_set_href = illumio-core_rule_set.web_apps.href

  enabled = true

  resolve_labels_as {
    consumers = ["workloads"]
    providers = ["workloads"]
  }

  consumers {
    actors = "ams"  # special notation meaning "all managed systems" - affects all workloads
  }

  providers {
    label {
      href = illumio-core_label.role_tomcat.href
    }
  }

  providers {
    label {
      href = illumio-core_label.env_dev.href
    }
  }

  ingress_services {
    href = local.all_services_href
  }
}
