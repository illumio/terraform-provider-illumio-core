resource "illumio-core_label_type" "os" {
  key          = "os"
  display_name = "OS"

  display_info {
    initial = "OS"
  }
}

resource "illumio-core_label_type" "device_type" {
  key          = "device"
  display_name = "Device Type"

  display_info {
    initial = "DT"
  }
}

data "illumio-core_label_types" "example" {
  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_pairing_profile.dev_core_services,
    illumio-core_pairing_profile.dev_web_db,
  ]
}
