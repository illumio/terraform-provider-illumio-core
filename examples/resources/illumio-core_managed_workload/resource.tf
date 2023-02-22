resource "illumio-core_label" "role_example" {
  key   = "role"
  value = "R-EXAMPLE"
}

resource "illumio-core_label" "app_web" {
  key   = "app"
  value = "A-WEB"
}

resource "illumio-core_label" "env_dev" {
  key   = "env"
  value = "E-DEV"
}

resource "illumio-core_label" "loc_aws" {
  key   = "loc"
  value = "L-AWS"
}

resource "illumio-core_managed_workload" "example" {
  data_center      = "us-west-1.amazonaws.com"
  data_center_zone = "us-west-1"

  service_provider = "amazonaws.com"

  enforcement_mode = "visibility_only"

  labels {
    href = illumio-core_label.role_example.href
  }

  labels {
    href = illumio-core_label.app_web.href
  }

  labels {
    href = illumio-core_label.env_dev.href
  }

  labels {
    href = illumio-core_label.loc_aws.href
  }
}
