terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

# Configure the PCE connection information as
# TF variables and set up the provider

variable "pce_url" {
  type        = string
  description = "URL of the Illumio Policy Compute Engine to connect to"
}

variable "pce_org_id" {
  type        = number
  description = "Illumio PCE Organization ID number"
  default     = 1
}

variable "pce_api_key" {
  type        = string
  description = "Illumio PCE API key username"
  sensitive   = true
}

variable "pce_api_secret" {
  type        = string
  description = "Illumio PCE API key secret"
  sensitive   = true
}

provider "illumio-core" {
    pce_host     = var.pce_url
    org_id       = var.pce_org_id
    api_username = var.pce_api_key
    api_secret   = var.pce_api_secret
}

# Define labels to be used by the workload

resource "illumio-core_label" "role_load_balancer" {
  key   = "role"
  value = "R-LB"
}

resource "illumio-core_label" "app_core_services" {
  key   = "app"
  value = "A-CORE-SERVICES"
}

resource "illumio-core_label" "env_prod" {
  key   = "env"
  value = "E-PROD"
}

resource "illumio-core_label" "loc_new_york" {
  key   = "loc"
  value = "L-NY"
}

# Define unmanaged workloads to represent application servers

resource "illumio-core_unmanaged_workload" "nginx_lb_ny" {
  name             = "nginx"
  hostname         = "ny.lb.illum.io"
  public_ip        = "10.10.10.1"
  description      = "NGINX load balancing proxy - NY - Prod"
  enforcement_mode = "full"
  online           = true

  labels {
    href = illumio-core_label.role_load_balancer.href
  }

  labels {
    href = illumio-core_label.app_core_services.href
  }

  labels {
    href = illumio-core_label.env_prod.href
  }

  labels {
    href = illumio-core_label.loc_new_york.href
  }
}
