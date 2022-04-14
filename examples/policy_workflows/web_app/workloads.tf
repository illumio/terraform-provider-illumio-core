# Define unmanaged workloads to represent application servers

resource "illumio-core_unmanaged_workload" "web_dev" {
  name             = "web-dev"
  hostname         = "web.lab.illum.io"
  public_ip        = "172.22.28.161"
  description      = "Web server - Lab - dev"
  enforcement_mode = "visibility_only"
  online           = true

  labels {
    href = illumio-core_label.role_tomcat.href
  }

  labels {
    href = illumio-core_label.app_web.href
  }

  labels {
    href = illumio-core_label.env_dev.href
  }

  labels {
    href = illumio-core_label.loc_lab.href
  }
}

resource "illumio-core_unmanaged_workload" "db_dev" {
  name             = "db-dev"
  hostname         = "db.lab.illum.io"
  public_ip        = "172.22.28.162"
  description      = "DB server - Lab - dev"
  enforcement_mode = "visibility_only"
  online           = true

  labels {
    href = illumio-core_label.role_db.href
  }

  labels {
    href = illumio-core_label.app_web.href
  }

  labels {
    href = illumio-core_label.env_dev.href
  }

  labels {
    href = illumio-core_label.loc_lab.href
  }
}

resource "illumio-core_unmanaged_workload" "web01_prod" {
  name             = "web01-prod"
  hostname         = "web01.illum.io"
  public_ip        = "10.10.0.13"
  description      = "Web server - US West - prod"
  enforcement_mode = "full"
  online           = true

  labels {
    href = illumio-core_label.role_tomcat.href
  }

  labels {
    href = illumio-core_label.app_web.href
  }

  labels {
    href = illumio-core_label.env_prod.href
  }

  labels {
    href = illumio-core_label.loc_us_west.href
  }
}

resource "illumio-core_unmanaged_workload" "web02_prod" {
  name             = "web02-prod"
  hostname         = "web02.illum.io"
  public_ip        = "10.10.0.14"
  description      = "Web server - US East - prod"
  enforcement_mode = "full"
  online           = true

  labels {
    href = illumio-core_label.role_tomcat.href
  }

  labels {
    href = illumio-core_label.app_web.href
  }

  labels {
    href = illumio-core_label.env_prod.href
  }

  labels {
    href = illumio-core_label.loc_us_east.href
  }
}

resource "illumio-core_unmanaged_workload" "db02_prod" {
  name             = "db02-prod"
  hostname         = "db02.illum.io"
  public_ip        = "10.10.0.15"
  description      = "DB server - US West - prod"
  enforcement_mode = "full"
  online           = true

  labels {
    href = illumio-core_label.role_db.href
  }

  labels {
    href = illumio-core_label.app_web.href
  }

  labels {
    href = illumio-core_label.env_prod.href
  }

  labels {
    href = illumio-core_label.loc_us_west.href
  }
}

resource "illumio-core_unmanaged_workload" "db01_prod" {
  name             = "db01-prod"
  hostname         = "db01.illum.io"
  public_ip        = "10.10.0.16"
  description      = "DB server - US East - prod"
  enforcement_mode = "full"
  online           = true

  labels {
    href = illumio-core_label.role_db.href
  }

  labels {
    href = illumio-core_label.app_web.href
  }

  labels {
    href = illumio-core_label.env_prod.href
  }

  labels {
    href = illumio-core_label.loc_us_east.href
  }
}
