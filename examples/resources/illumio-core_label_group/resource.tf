resource "illumio-core_label" "loc_cloud_aws" {
  key   = "loc"
  value = "L-CLOUD-AWS"
}

resource "illumio-core_label" "loc_cloud_gcp" {
  key   = "loc"
  value = "L-CLOUD-GCP"
}

resource "illumio-core_label" "loc_cloud_azure" {
  key   = "loc"
  value = "L-CLOUD-AZURE"
}

resource "illumio-core_label_group" "example" {
  key         = "loc"
  name        = "LG-CLOUD"
  description = "Cloud locations Label Group"

  labels {
    href = illumio-core_label.loc_cloud_aws.href
  }

  labels {
    href = illumio-core_label.loc_cloud_gcp.href
  }

  labels {
    href = illumio-core_label.loc_cloud_azure.href
  }
}
