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

resource "illumio-core_label" "role_cluster_leader" {
	key   = "role"
	value = "R-CLUSTER-LEADER"
}

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

resource "illumio-core_unmanaged_workload" "jenkins_leader" {
  name             = "jenkins_w01"
  hostname         = "jenkins.lab.illum.io"
  public_ip        = "172.22.8.210"
  description      = "Jenkins cluster leader - EU - dev"
  enforcement_mode = "visibility_only"
  online           = true

  labels {
    href = illumio-core_label.role_cluster_leader.href
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

resource "illumio-core_unmanaged_workload" "jenkins_worker01" {
  name             = "jenkins_w01"
  hostname         = "w01.jenkins.lab.illum.io"
  public_ip        = "172.22.8.211"
  description      = "Jenkins worker - EU - dev"
  enforcement_mode = "visibility_only"
  online           = true

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

resource "illumio-core_unmanaged_workload" "jenkins_worker02" {
  name             = "jenkins_w02"
  hostname         = "w02.jenkins.lab.illum.io"
  public_ip        = "172.22.8.212"
  description      = "Jenkins worker - EU - dev"
  enforcement_mode = "visibility_only"
  online           = true

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

data "illumio-core_workloads" "jenkins" {
	# supports partial match lookups
  hostname = "jenkins.lab.illum.io"

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_unmanaged_workload.jenkins_leader,
    illumio-core_unmanaged_workload.jenkins_worker01,
    illumio-core_unmanaged_workload.jenkins_worker02,
  ]
}