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

resource "illumio-core_container_cluster" "kube" {
  name        = "CC-KUBE"
  description = "Kubernetes Container Cluster"
}

resource "illumio-core_label" "location_kubernetes" {
  key   = "loc"
  value = "L-KUBE"
}

resource "illumio-core_label" "env_dev" {
  key   = "env"
  value = "E-DEV"
}

resource "illumio-core_label" "app_core_services" {
  key   = "app"
  value = "A-CORE-SERVICES"
}

resource "illumio-core_container_cluster_workload_profile" "kube_sandbox" {
  container_cluster_href = illumio-core_container_cluster.kube.href
  name                   = "CCWP-KUBE-SANDBOX"
  description            = "Workload profile for devtest pods"
  managed                = true
  enforcement_mode       = "selective"

  assign_labels {
    href = illumio-core_label.location_kubernetes.href
  }

  assign_labels {
    href = illumio-core_label.env_dev.href
  }
}

resource "illumio-core_container_cluster_workload_profile" "kube_core_services" {
  container_cluster_href = illumio-core_container_cluster.kube.href
  name                   = "CCWP-KUBE-CORE-SERVICES"
  description            = "Workload profile for core-services pods"
  managed                = true
  enforcement_mode       = "visibility_only"

  assign_labels {
    href = illumio-core_label.location_kubernetes.href
  }

  assign_labels {
    href = illumio-core_label.app_core_services.href
  }
}

data "illumio-core_container_cluster_workload_profiles" "kube_profiles" {
	container_cluster_href = illumio-core_container_cluster.kube.href
  max_results            = 2

  # despite the explicit dependency on the container cluster itself,
  # the implicit dependencies on its workload profiles need to
  # be explicitly defined here to ensure the profiles are created
  # before the data source is populated
	depends_on = [
		illumio-core_container_cluster_workload_profile.kube_sandbox,
		illumio-core_container_cluster_workload_profile.kube_core_services,
	]
}
