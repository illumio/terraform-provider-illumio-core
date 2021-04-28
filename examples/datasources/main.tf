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

data "illumio-core_workloads" "name" {
    max_results = "5"
}   

data "illumio-core_vens" "name" {
    max_results = "5"
}  

data "illumio-core_ip_lists" "name" {
    max_results = "5"
}  


data "illumio-core_labels" "name" {
  max_results = "5"
}

data "illumio-core_label_groups" "name" {
  max_results = "5"
}

data "illumio-core_services" "name" {
  max_results = "5"
}

data "illumio-core_virtual_services" "name" {
  max_results = "5"
}

