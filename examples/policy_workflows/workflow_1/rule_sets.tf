resource "illumio-core_rule_set" "policy-1-rule_set_1" {
    name = "policy-1-rule_set_1"
    description = "Container cluster Test"
    scopes {
      label_group {
        href = illumio-core_label_group.policy-1-app_lg2.href
      }
      label_group {
        href = illumio-core_label_group.policy-1-env_lg2.href
      }
      label_group {
        href = illumio-core_label_group.policy-1-loc_lg2.href
      }
    }
    scopes {
      label_group {
        href = illumio-core_label_group.policy-1-app_lg3.href
      }
      label_group {
        href = illumio-core_label_group.policy-1-env_lg3.href
      }
      label_group {
        href = illumio-core_label_group.policy-1-loc_lg3.href
      }
    }
}

resource "illumio-core_security_rule" "policy-1-sec_rule_1" {
  rule_set_href = illumio-core_rule_set.policy-1-rule_set_1.href
  enabled = true
  sec_connect = false
  providers {
    label {
      href = illumio-core_label.policy-1-role1.href
    }
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy-1-role_lg2.href
    }
  }
  ingress_services {
    port = "12000"
    proto = "6"
  }
  ingress_services {
    port = "12000"
    proto = "17"
  }
  resolve_labels_as {
    providers = [ "workloads" ]
    consumers = [ "workloads" ]
  }
  unscoped_consumers = false
  description = "policy-1-r1"
}

resource "illumio-core_security_rule" "policy-1-sec_rule_2" {
    rule_set_href = illumio-core_rule_set.policy-1-rule_set_1.href
 enabled = true
  sec_connect = false
  providers {
    label {
      href = illumio-core_label.policy-1-role2.href
    }
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy-1-role_lg2.href
    }
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy-1-role_lg3.href
    }
  }
  ingress_services {
    port = "13000"
    proto = "6"
  }
  ingress_services {
    port = "13000"
    proto = "17"
  }
  resolve_labels_as {
    providers = [ "virtual_services", "workloads" ]
    consumers = [ "virtual_services", "workloads" ]
  }
  unscoped_consumers = false
  description = "policy-1-r2"
}

resource "illumio-core_security_rule" "policy-1-sec_rule_3" {
    rule_set_href = illumio-core_rule_set.policy-1-rule_set_1.href
  enabled = true
  sec_connect = false
  providers {
    actors = "ams"
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy-1-app_lg1.href
    }
  }
  consumers {
    label {
      href = illumio-core_label.policy-1-env1.href
    }
  }
  ingress_services {
    port = "14000"
    proto = "6"
  }
  ingress_services {
    port = "14000"
    proto = "17"
  }
  resolve_labels_as {
    providers = [ "workloads" ]
    consumers = [ "workloads" ]
  }
  unscoped_consumers = true
  description = "policy-1-r3"
}

resource "illumio-core_security_rule" "policy-1-sec_rule_4" {
    rule_set_href = illumio-core_rule_set.policy-1-rule_set_1.href
  enabled = true
  sec_connect = false
  providers {
    label {
      href = illumio-core_label.policy-1-role2.href
    }
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy-1-app_lg3.href
    }
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy-1-env_lg2.href
    }
  }
  ingress_services {
    port = "15000"
    proto = "6"
  }
  ingress_services {
    port = "15000"
    proto = "17"
  }
  resolve_labels_as {
    providers = [ "workloads" ]
    consumers = [ "workloads" ]
  }
  unscoped_consumers = true
  description = "policy-1-r4"
}
