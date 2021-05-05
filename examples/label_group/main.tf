terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
//  pce_host              = "https://pce.my-company.com:8443"
//  api_username          = "api_xxxxxx"
//  api_secret            = "big-secret"
  request_timeout       = 30
  org_id                = 1
}

# data "illumio-core_label_group" "datalg1"{
#   label_group_id = "731fca37-7f57-4852-96ac-753ddfd359d3"
# }

# data "illumio-core_label" "label_1" {
#   label_id  = 1
# }

# data "illumio-core_label" "label_2" {
#   label_id  = 2
# }

# resource "illumio-core_label_group" "role_lg_a" {
#   key           = "role"
#   name          = "test label group - a"
#   description   = "Update Desc"
#   labels {
#     href = data.illumio-core_label.label_1.href
#   }
#   labels {
#     href = data.illumio-core_label.label_2.href
#   }
# }

resource "illumio-core_label" "l1" {
  key = "app"
  value = "everyone-got-app"
}

resource "illumio-core_label_group" "lg1" {
  key = "app"
  name = "LabelGroupTestingDep"
  description = "Amazing Apps"
  provision_on_delete = true
  labels {
    href = illumio-core_label.l1.href
  }
}

# resource "illumio-core_label_group" "lg2" {
#   key = "app"
#   name = "LabelGroupTestingDep1"
#   description = "Amazing Apps"
#   provision_on_delete = true
#   sub_groups {
#     href = illumio-core_label_group.lg1.href
#   }
# }