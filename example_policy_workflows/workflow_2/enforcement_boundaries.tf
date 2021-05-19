resource "illumio-core_enforcement_boundary" "policy_2_se_rule_1" {
  name = "policy_2_se_rule_1"
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
      href = illumio-core_label.policy_2_env1.href
    }
  }
  providers {
    label {
      href = illumio-core_label.policy_2_loc1.href
    }
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc1.href
  }
  ingress_services {
    port  = 12000
    proto = 6
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_2" {
  name = "policy_2_se_rule_2"
  providers {
    actors = "ams"
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  ingress_services {
    port    = 13000
    to_port = 13005
    proto   = 6
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_3" {
  name = "policy_2_se_rule_3"
  providers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
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
    port  = 15100
    proto = 6
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_4" {
  name = "policy_2_se_rule_4"
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
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
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
    port  = 16100
    proto = 6
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc2.href
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_5" {
  name = "policy_2_se_rule_5"

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
      href = illumio-core_label.policy_2_env1.href
    }
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  ingress_services {
    port    = 17010
    to_port = 17050
    proto   = 6
  }
  ingress_services {
    port    = 17010
    to_port = 17050
    proto   = 17
  }
  ingress_services {
    port  = 17200
    proto = 6
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_6" {
  name = "policy_2_se_rule_6"
  providers {
    label {
      href = illumio-core_label.policy_2_app2.href
    }
  }
  providers {
    label {
      href = illumio-core_label.policy_2_loc2.href
    }
  }
  providers {
    label {
      href = illumio-core_label.policy_2_env2.href
    }
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  ingress_services {
    port  = 18101
    proto = 6
  }
  ingress_services {
    port  = 18101
    proto = 17
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc3.href
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_7" {
  name = "policy_2_se_rule_7"
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_role_lg1.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_app_lg1.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_loc_lg1.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_env_lg1.href
    }
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  ingress_services {
    port  = 19100
    proto = 6
  }
  ingress_services {
    port  = 19100
    proto = 17
  }

  ingress_services {
    href = illumio-core_service.policy_2_svc4.href
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_8" {
  name = "policy_2_se_rule_8"
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_app_lg1.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_loc_lg1.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_env_lg1.href
    }
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc11.href
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc5.href
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_9" {
  name = "policy_2_se_rule_9"
  providers {
    label {
      href = illumio-core_label.policy_2_role1.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_loc_lg1.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_env_lg1.href
    }
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc6.href
  }
  ingress_services {
    port  = 21100
    proto = 6
  }
  ingress_services {
    port  = 21100
    proto = 17
  }

}

resource "illumio-core_enforcement_boundary" "policy_2_se_rule_10" {
  name = "policy_2_se_rule_10"
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
    label_group {
      href = illumio-core_label_group.policy_2_loc_lg1.href
    }
  }
  providers {
    label_group {
      href = illumio-core_label_group.policy_2_env_lg1.href
    }
  }
  consumers {
    ip_list {
      href = "/orgs/1/sec_policy/draft/ip_lists/1"
    }
  }
  ingress_services {
    href = illumio-core_service.policy_2_svc7.href
  }
  ingress_services {
    port  = 22100
    proto = 6
  }
  ingress_services {
    port  = 22100
    proto = 17
  }
}
