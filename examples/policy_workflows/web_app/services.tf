# Define services used by the web application

resource "illumio-core_service" "mysql" {
  name        = "S-MYSQL"
  description = "MySQL default service port"

  service_ports {
    # Illumio uses the IANA protocol numbers to identify the service proto
    proto = "6"  # TCP
    port  = "3306"
  }
}

resource "illumio-core_service" "http" {
  name        = "S-HTTP"
  description = "HTTP(S) default ports"

  service_ports {
    proto = "6"
    port  = "80"
  }

  service_ports {
    proto = "6"
    port  = "443"
  }
}

data "illumio-core_services" "all_services" {
  # all PCE instances define a default Service covering all service ports
  name = "All Services"
  max_results = 1
}

locals {
  all_services_href = one(data.illumio-core_services.all_services.items[*].href)
}
