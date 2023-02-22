resource "illumio-core_label" "role_db" {
  key   = "role"
  value = "R-DB"
}

data "illumio-core_label" "role_db" {
  href = illumio-core_label.role_db.href
}
