data "terraform_remote_state" "virtual_service" {
  backend = "local"

  config = {
    path = "virtual_service/terraform.tfstate"
  }
}

resource "illumio-core_label" "role_web" {
  key   = "role"
  value = "R-WEB"
}

resource "illumio-core_unmanaged_workload" "hrm_db01" {
  name             = "hrm_db01"
  hostname         = "db01.hrm.qa.illum.io"
  public_ip        = "172.29.132.12"
  description      = "HRM Postgres database - Azure - QA"
  enforcement_mode = "visibility_only"
  online           = true

  labels {
    href = data.terraform_remote_state.virtual_service.outputs.role_label_href
  }

  labels {
    href = data.terraform_remote_state.virtual_service.outputs.app_label_href
  }

  labels {
    href = data.terraform_remote_state.virtual_service.outputs.env_label_href
  }

  labels {
    href = data.terraform_remote_state.virtual_service.outputs.loc_label_href
  }
}

resource "illumio-core_unmanaged_workload" "hrm_db02" {
  name             = "hrm_db02"
  hostname         = "db02.hrm.qa.illum.io"
  public_ip        = "172.29.132.13"
  description      = "HRM Postgres database - Azure - QA"
  enforcement_mode = "visibility_only"
  online           = true

  labels {
    href = data.terraform_remote_state.virtual_service.outputs.role_label_href
  }

  labels {
    href = data.terraform_remote_state.virtual_service.outputs.app_label_href
  }

  labels {
    href = data.terraform_remote_state.virtual_service.outputs.env_label_href
  }

  labels {
    href = data.terraform_remote_state.virtual_service.outputs.loc_label_href
  }
}

resource "illumio-core_service_binding" "hrm_db01" {
  virtual_service {
    href = data.terraform_remote_state.virtual_service.outputs.virtual_service_href
  }

  workload {
    href = illumio-core_unmanaged_workload.hrm_db01.href
  }
}

resource "illumio-core_service_binding" "hrm_db02" {
  virtual_service {
    href = data.terraform_remote_state.virtual_service.outputs.virtual_service_href
  }

  workload {
    href = illumio-core_unmanaged_workload.hrm_db02.href
  }
}

data "illumio-core_service_bindings" "example" {
  virtual_service = data.terraform_remote_state.virtual_service.outputs.virtual_service_href
}