resource "illumio-core_container_cluster" "kube" {
  name        = "CC-KUBE"
  description = "Kubernetes Container Cluster"
}

data "illumio-core_container_cluster_service_backends" "kube_service_backends" {
  container_cluster_href = illumio-core_container_cluster.kube.href
}
