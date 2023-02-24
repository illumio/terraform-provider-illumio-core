data "terraform_remote_state" "virtual_service" {
  backend = "local"

  config = {
    path = "virtual_service/terraform.tfstate"
  }
}

resource "illumio-core_unmanaged_workload" "hrm_db" {
  name             = "hrm_db"
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

resource "illumio-core_service_binding" "hrm_db" {
  virtual_service {
    href = data.terraform_remote_state.virtual_service.outputs.virtual_service_href
  }

  workload {
    href = illumio-core_unmanaged_workload.hrm_db.href
  }
}

data "illumio-core_service_binding" "hrm_db" {
   href = illumio-core_service_binding.hrm_db.href
}
