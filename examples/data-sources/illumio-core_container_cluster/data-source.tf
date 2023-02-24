resource "illumio-core_container_cluster" "kube" {
  name        = "CC-KUBE"
  description = "Kubernetes Container Cluster"
}

data "illumio-core_container_cluster" "kube" {
  href = illumio-core_container_cluster.kube.href
}
