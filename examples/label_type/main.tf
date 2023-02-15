terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  pce_host     = var.pce_url
  org_id       = var.pce_org_id
  api_username = var.pce_api_key
  api_secret   = var.pce_api_secret
}

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
