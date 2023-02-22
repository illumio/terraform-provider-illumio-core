resource "illumio-core_label_type" "os" {
  key          = "os"
  display_name = "OS"

  display_info {
    initial = "OS"
  }
}

resource "illumio-core_label" "os_windows" {
  # create an implicit dependency on the label type
  # to ensure it's created before the label
  key   = illumio-core_label_type.os.key
  value = "OS_Windows"
}

data "illumio-core_label" "os_windows" {
  href = illumio-core_label.os_windows.href
}
