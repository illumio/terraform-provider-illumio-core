# Define labels to be used by the example application

resource "illumio-core_label" "role_db" {
  key   = "role"
  value = "R-DB"
}

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

resource "illumio-core_label" "env_prod" {
  key   = "env"
  value = "E-PROD"
}

resource "illumio-core_label" "loc_lab" {
  key   = "loc"
  value = "L-LAB"
}

resource "illumio-core_label" "loc_us_east" {
  key   = "loc"
  value = "L-US-EAST"
}

resource "illumio-core_label" "loc_us_west" {
  key   = "loc"
  value = "L-US-WEST"
}

resource "illumio-core_label_group" "loc_us" {
  key         = "loc"
  name        = "L-LG-US"
  description = "US locations Label Group"

  labels {
    href = illumio-core_label.loc_us_west.href
  }

  labels {
    href = illumio-core_label.loc_us_east.href
  }
}
