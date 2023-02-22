resource "illumio-core_label" "role_tomcat" {
  key   = "role"
  value = "R-TOMCAT"
}

resource "illumio-core_label" "app_web" {
  key   = "app"
  value = "A-WEB"
}

resource "illumio-core_label" "env_dev" {
  key   = "env"
  value = "E-DEV"
}

resource "illumio-core_label" "loc_us_east" {
  key   = "loc"
  value = "L-US-EAST"
}

resource "illumio-core_unmanaged_workload" "tomcat_sandbox" {
  name             = "tomcat_dev"
  hostname         = "tomcat.lab.illum.io"
  public_ip        = "172.22.8.224"
  description      = "Tomcat Sandbox - US East - dev"
  enforcement_mode = "visibility_only"
  online           = true

  labels {
    href = illumio-core_label.role_tomcat.href
  }

  labels {
    href = illumio-core_label.app_web.href
  }

  labels {
    href = illumio-core_label.env_dev.href
  }

  labels {
    href = illumio-core_label.loc_us_east.href
  }
}

data "illumio-core_workload" "tomcat" {
  href = illumio-core_unmanaged_workload.tomcat_sandbox.href
}