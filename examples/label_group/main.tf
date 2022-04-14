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

resource "illumio-core_label_group" "loc_cloud" {
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

data "illumio-core_label_group" "loc_cloud" {
  href = illumio-core_label_group.loc_cloud.href
}
