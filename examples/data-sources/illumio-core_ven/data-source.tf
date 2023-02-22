resource "illumio-core_ven" "example" {
  status = "active"
}

data "illumio-core_ven" "example" {
  href = illumio-core_ven.example.href
}
