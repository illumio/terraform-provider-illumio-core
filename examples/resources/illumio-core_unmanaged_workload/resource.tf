resource "illumio-core_label" "role_cluster_worker" {
	key   = "role"
	value = "R-CLUSTER-WORKER"
}

resource "illumio-core_label" "app_jenkins" {
	key   = "app"
	value = "A-JENKINS"
}

resource "illumio-core_label" "env_dev" {
	key   = "env"
	value = "E-DEV"
}

resource "illumio-core_label" "loc_eu" {
	key   = "loc"
	value = "L-EU"
}

resource "illumio-core_unmanaged_workload" "example" {
  name             = "jenkins_w01"
  hostname         = "w01.jenkins.lab.illum.io"
  public_ip        = "172.22.8.211"
  description      = "Jenkins worker - EU - dev"
  enforcement_mode = "visibility_only"
  online           = true

  interfaces {
    name       = "eth0"
    address    = "172.22.8.211"
    cidr_block = 28
    link_state = "up"
  }

  labels {
    href = illumio-core_label.role_cluster_worker.href
  }

  labels {
    href = illumio-core_label.app_jenkins.href
  }

  labels {
    href = illumio-core_label.env_dev.href
  }

  labels {
    href = illumio-core_label.loc_eu.href
  }
}
