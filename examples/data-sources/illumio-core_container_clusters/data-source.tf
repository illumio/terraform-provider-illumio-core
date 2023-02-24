resource "illumio-core_container_cluster" "kube_aws" {
  name        = "CC-KUBE-AWS"
  description = "Kubernetes Cluster on AWS"
}

resource "illumio-core_container_cluster" "kube_gcp" {
  name        = "CC-KUBE-GCP"
  description = "Kubernetes Cluster on GCP"
}

data "illumio-core_container_clusters" "kube_clusters" {
	# lookup all Kube clusters by name with a partial match
  name = "CC-KUBE-"
  max_results = 2

  # the implicit dependency on the container cluster resources
  # needs to be made explicit in order for the data source to
  # populate correctly
  depends_on = [
    illumio-core_container_cluster.kube_aws,
    illumio-core_container_cluster.kube_gcp,
  ]
}
