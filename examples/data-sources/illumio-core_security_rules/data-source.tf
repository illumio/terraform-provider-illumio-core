resource "illumio-core_label" "app_web" {
  key   = "app"
  value = "A-WEB"
}

resource "illumio-core_label" "role_database" {
  key   = "role"
  value = "R-DB"
}

resource "illumio-core_label" "role_web_server" {
  key   = "role"
  value = "R-WEB-SERVER"
}

resource "illumio-core_rule_set" "web_apps" {
  name = "RS-WEB"

  scopes {
    label {
      href = illumio-core_label.app_web.href
    }
  }
}

resource "illumio-core_service" "mysql" {
  name        = "S-MYSQL"
  description = "MySQL default service port"

  service_ports {
    # Illumio uses the IANA protocol numbers to identify the service proto
    proto = "6"  # TCP
    port  = "3306"
  }
}

resource "illumio-core_service" "http" {
  name        = "S-HTTP"
  description = "HTTP(S) default ports"

  service_ports {
    proto = "6"
    port  = "80"
  }

  service_ports {
    proto = "6"
    port  = "443"
  }
}

resource "illumio-core_security_rule" "web_mysql" {
  rule_set_href = illumio-core_rule_set.web_apps.href

  enabled = true

  resolve_labels_as {
    consumers = ["workloads"]
    providers = ["workloads"]
  }

  consumers {
    label {
      href = illumio-core_label.role_web_server.href
    }
  }

  providers {
    label {
      href = illumio-core_label.role_database.href
    }
  }

  ingress_services {
    href = illumio-core_service.mysql.href
  }
}

# use the ip_lists data source to search against the /ip_lists endpoint by name
data "illumio-core_ip_lists" "default" {
  # all PCE instances define a special default IP list covering all addresses
  name = "Any (0.0.0.0/0 and ::/0)"
  max_results = 1
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
      href = one(data.illumio-core_ip_lists.default.items[*].href)
    }
  }

  providers {
    label {
      href = illumio-core_label.role_web_server.href
    }
  }

  ingress_services {
    href = illumio-core_service.http.href
  }
}

data "illumio-core_security_rules" "web_application_rules" {
  rule_set_href = illumio-core_rule_set.web_apps.href

  # despite the explicit dependency on the ruleset itself,
  # the implicit dependencies on the rules it contains need
  # to be explicitly defined here to ensure they're created
  # before the data source is populated
  depends_on = [
    illumio-core_security_rule.web_mysql,
    illumio-core_security_rule.web_http_inbound,
  ]
}
