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

data "illumio-core_labels" "loc_cloud" {
	# supports partial match lookups
  value = "L-CLOUD"

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_label.loc_cloud_aws,
    illumio-core_label.loc_cloud_gcp,
    illumio-core_label.loc_cloud_azure,
  ]
}
