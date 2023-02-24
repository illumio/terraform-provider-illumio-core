resource "illumio-core_label_type" "example" {
  key          = "os"
  display_name = "OS"

  display_info {
    initial = "OS"
  }
}

resource "illumio-core_label" "example" {
  # create an implicit dependency on the label type
  # to ensure it's created before the label
  key   = illumio-core_label_type.example.key
  value = "OS_Windows"
}
