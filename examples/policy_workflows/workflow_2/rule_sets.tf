resource "illumio-core_rule_set" "name" {
  name        = "scoped_rset_1"
  description = "Individual scope test"
  scopes {
    label {
      href = illumio-core_label.policy_2_app1.href
    }
    label {
      href = illumio-core_label.policy_2_env1.href
    }
    label {
      href = illumio-core_label.policy_2_loc1.href
    }
  }
}

resource "illumio-core_security_rule" "name" {
  rule_set_href = illumio-core_rule_set.name.href

  enabled     = true
  sec_connect = false
  providers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  consumers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  ingress_services {
    port  = 11500
    proto = 6
  }
  ingress_services {
    port  = 11500
    proto = 17
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc8.href

  }
  resolve_labels_as {
    providers = ["workloads"]
    consumers = ["workloads"]
  }
  unscoped_consumers = false
  description        = "r1"
}

resource "illumio-core_rule_set" "policy_2_all_scoped_rset_1" {
  name        = "all_scoped_rset_1"
  description = "All scope test"
  scopes {}
}

resource "illumio-core_security_rule" "policy_2_all_scoped_rset_1_sr" {
  rule_set_href = illumio-core_rule_set.policy_2_all_scoped_rset_1.href
  enabled       = true
  sec_connect   = false
  providers {
    label {
      href = illumio-core_label.policy_2_app1.href
    }
  }
  providers {
    label {
      href = illumio-core_label.policy_2_app2.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_env_lg2.href
    }
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy_2_role_lg1.href
    }
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy_2_loc_lg2.href
    }
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc9.href

  }
  ingress_services {
    port  = 14000
    proto = 6
  }
  ingress_services {
    port  = 15000
    proto = 17
  }
  ingress_services {
    port  = 13500
    proto = 6
  }
  ingress_services {
    port  = 13500
    proto = 17
  }
  resolve_labels_as {
    providers = ["workloads"]
    consumers = ["workloads"]
  }
  unscoped_consumers = false
  description        = "r3"

}

resource "illumio-core_rule_set" "policy_2_partial_all_scoped_rset_1" {
  name        = "partial_all_scoped_rset_1"
  description = "Partial scope test"
  scopes {
    label {
      href = illumio-core_label.policy_2_app1.href
    }
  }
}

resource "illumio-core_security_rule" "policy_2_partial_all_scoped_rset_1_sr" {
  rule_set_href = illumio-core_rule_set.policy_2_partial_all_scoped_rset_1.href
  enabled       = true
  sec_connect   = false
  providers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  consumers {
    actors = "ams"
  }
  ingress_services {
    port    = 16000
    to_port = 16005
    proto   = 6
  }
  ingress_services {
    port    = 16000
    to_port = 16005
    proto   = 17
  }
  ingress_services {
    port  = 16500
    proto = 6
  }
  ingress_services {
    port  = 16500
    proto = 17
  }
  resolve_labels_as {
    providers = ["workloads"]
    consumers = ["workloads"]
  }
  unscoped_consumers = false
  description        = "r4"

}

resource "illumio-core_rule_set" "policy_2_partial_all_scoped_rset_2" {
  name        = "policy_2_partial_all_scoped_rset_2"
  description = "Partial scope test2"
  scopes {
    label {
      href = illumio-core_label.policy_2_app1.href
    }

    label {
      href = illumio-core_label.policy_2_env1.href
    }
  }
}

resource "illumio-core_security_rule" "policy_2_partial_all_scoped_rset_2_sr" {
  rule_set_href = illumio-core_rule_set.policy_2_partial_all_scoped_rset_2.href
  enabled       = true
  sec_connect   = false
  providers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  providers {
    label {
      href = illumio-core_label.policy_2_loc1.href
    }
  }
  consumers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  consumers {
    ip_list {
      href = illumio-core_ip_list.policy_2_iplist1.href
    }
  }
  ingress_services {
    port  = 17500
    proto = 6
  }
  ingress_services {
    port  = 17500
    proto = 17
  }
  ingress_services {
    port    = 17100
    to_port = 17105
    proto   = 6
  }
  ingress_services {
    port    = 17100
    to_port = 17105
    proto   = 17
  }
  resolve_labels_as {
    providers = ["workloads"]
    consumers = ["workloads"]
  }
  unscoped_consumers = false
  description        = "r5"
}

resource "illumio-core_rule_set" "policy_2_scoped_rset_2" {
  name        = "policy_2_scoped_rset_2"
  description = "Individual scope test 2"
  scopes {
    label {
      href = illumio-core_label.policy_2_app2.href
    }

    label {
      href = illumio-core_label.policy_2_env2.href
    }

    label {
      href = illumio-core_label.policy_2_loc2.href
    }
  }
}

resource "illumio-core_security_rule" "policy_2_scoped_rset_2_sr" {
  rule_set_href = illumio-core_rule_set.policy_2_scoped_rset_2.href
  enabled       = true
  sec_connect   = false
  providers {
    label {
      href = illumio-core_label.policy_2_role2.href
    }
  }
  consumers {
    label_group {
      href = illumio-core_label_group.policy_2_role_lg2.href
    }
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc3.href

  }
  ingress_services {
    port  = 18500
    proto = 6
  }
  ingress_services {
    port  = 18500
    proto = 17
  }
  resolve_labels_as {
    providers = ["workloads"]
    consumers = ["workloads"]
  }
  unscoped_consumers = false
  description        = "r6"

}

resource "illumio-core_rule_set" "policy_2_lg_scoped_rset_1" {
  name        = "lg_scoped_rset_1"
  description = "Label group scope test"
  scopes {
    label_group {
      href = illumio-core_label_group.policy_2_app_lg1.href
    }
    label_group {
      href = illumio-core_label_group.policy_2_env_lg2.href
    }
    label_group {
      href = illumio-core_label_group.policy_2_loc_lg2.href
    }
  }
}

resource "illumio-core_security_rule" "policy_2_lg_scoped_rset_1_sr" {
  rule_set_href = illumio-core_rule_set.policy_2_lg_scoped_rset_1.href
  enabled       = true
  sec_connect   = false
  providers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  consumers {
    actors = "ams"
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc5.href

  }
  ingress_services {
    href = illumio-core_service.policy_2_svc10.href

  }
  ingress_services {
    href = illumio-core_service.policy_2_svc4.href

  }
  resolve_labels_as {
    providers = ["workloads"]
    consumers = ["workloads"]
  }
  unscoped_consumers = false
  description        = "r7"

}

resource "illumio-core_rule_set" "policy_2_lg_partial_scoped_rset_1" {
  name        = "policy_2_lg_partial_scoped_rset_1"
  description = "Partial lg scope test"
  scopes {
    label_group {
      href = illumio-core_label_group.policy_2_env_lg2.href
    }
    label_group {
      href = illumio-core_label_group.policy_2_loc_lg2.href
    }
  }
}

resource "illumio-core_security_rule" "policy_2_lg_partial_scoped_rset_1_sr" {
  rule_set_href = illumio-core_rule_set.policy_2_lg_partial_scoped_rset_1.href
  enabled       = true
  sec_connect   = false
  providers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  providers {
    label {
      href = illumio-core_label.policy_2_app1.href
    }
  }
  providers {
    label {
      href = illumio-core_label.policy_2_app2.href
    }
  }

  consumers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  consumers {
    label {
      href = illumio-core_label.policy_2_role2.href
    }
  }
  consumers {
    label {
      href = illumio-core_label.policy_2_app1.href
    }
  }

  ingress_services {
    href = illumio-core_service.policy_2_svc6.href
  }
  ingress_services {
    port  = 21001
    proto = 6
  }
  ingress_services {
    port  = 21500
    proto = 6
  }

  resolve_labels_as {
    providers = ["workloads"]
    consumers = ["workloads"]
  }
  unscoped_consumers = false
  description        = "r8"

}

resource "illumio-core_rule_set" "policy_2_mixed_scoped_rset_2" {
  name        = "policy_2_mixed_scoped_rset_2"
  description = "Mixed label and lg test"
  scopes {
    label {
      href = illumio-core_label.policy_2_app1.href
    }
    label_group {
      href = illumio-core_label_group.policy_2_env_lg2.href
    }
    label {
      href = illumio-core_label.policy_2_loc2.href
    }
  }
}

resource "illumio-core_security_rule" "policy_2_mixed_scoped_rset_2_sr" {
  rule_set_href = illumio-core_rule_set.policy_2_mixed_scoped_rset_2.href
  enabled       = true
  sec_connect   = false
  providers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  consumers {
    label {
      href = illumio-core_label.policy_2_loc1.href
    }
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc6.href
  }
  ingress_services {
    port  = 21001
    proto = 6
  }
  ingress_services {
    port  = 21500
    proto = 6
  }
  resolve_labels_as {
    providers = ["workloads"]
    consumers = ["workloads"]
  }
  unscoped_consumers = true
  description        = "r9"
}
