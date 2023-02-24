resource "illumio-core_container_cluster" "kube" {
  name        = "CC-KUBE"
  description = "Kubernetes Container Cluster"
}

resource "illumio-core_label" "location_kubernetes" {
  key   = "loc"
  value = "L-KUBE"
}

resource "illumio-core_label" "app_core_services" {
  key   = "app"
  value = "A-CORE-SERVICES"
}

resource "illumio-core_container_cluster_workload_profile" "example" {
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
