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

resource "illumio-core_label" "policy_2_role1" {
  key   = "role"
  value = "policy_2_role1"
}

resource "illumio-core_label" "policy_2_app1" {
  key   = "app"
  value = "policy_2_app1"
}

resource "illumio-core_label" "policy_2_env1" {
  key   = "env"
  value = "policy_2_env1"
}

resource "illumio-core_label" "policy_2_loc1" {
  key   = "loc"
  value = "policy_2_loc1"
}

resource "illumio-core_label" "policy_2_role2" {
  key   = "role"
  value = "policy_2_role2"
}

resource "illumio-core_label" "policy_2_app2" {
  key   = "app"
  value = "policy_2_app2"
}

resource "illumio-core_label" "policy_2_env2" {
  key   = "env"
  value = "policy_2_env2"
}

resource "illumio-core_label" "policy_2_loc2" {
  key   = "loc"
  value = "policy_2_loc2"
}

resource "illumio-core_label" "policy_2_role3" {
  key   = "role"
  value = "policy_2_role3"
}
