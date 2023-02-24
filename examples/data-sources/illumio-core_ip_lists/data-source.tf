resource "illumio-core_ip_list" "aws_vpc" {
  name        = "IPL-AWS-VPC"
  description = "AWS VPC application IPs"

  ip_ranges {
    // from_ip can be a CIDR range or individual IP
    from_ip = "10.10.0.0/18"
    description = "AWS VPC block"
  }

  ip_ranges {
    from_ip = "10.10.0.1"
    to_ip   = "10.10.15.255"
    description = "Designated management subnet"
    exclusion = true
  }

  fqdns {
    fqdn = "*.aws.illum.io"
    description = "Wildcard domain for AWS instances"
  }
}

resource "illumio-core_ip_list" "aws_vpc_mgmt" {
  name        = "IPL-AWS-VPC-MGMT"
  description = "AWS VPC management IPs"

  ip_ranges {
    from_ip = "10.10.0.0/20"
    description = "AWS VPC management subnet"
  }
}

resource "illumio-core_ip_list" "aws_vpc_dev" {
  name        = "IPL-AWS-VPC-DEV"
  description = "AWS VPC devtest IPs"

  ip_ranges {
    from_ip = "10.10.16.0/20"
    description = "AWS VPC devtest subnet"
  }
}

resource "illumio-core_ip_list" "aws_vpc_staging" {
  name        = "IPL-AWS-VPC-STAGING"
  description = "AWS VPC staging IPs"

  ip_ranges {
    from_ip = "10.10.32.0/20"
    description = "AWS VPC staging subnet"
  }
}

resource "illumio-core_ip_list" "aws_vpc_prod" {
  name        = "IPL-AWS-VPC-PROD"
  description = "AWS VPC prod IPs"

  ip_ranges {
    from_ip = "10.10.48.0/20"
    description = "AWS VPC production subnet"
  }
}

data "illumio-core_ip_lists" "aws_vpc" {
	# supports partial match lookups
  name = "IPL-AWS-VPC"

  # explicitly define the dependencies to ensure the resources
  # are created before the data source is populated
  depends_on = [
    illumio-core_ip_list.aws_vpc,
    illumio-core_ip_list.aws_vpc_mgmt,
    illumio-core_ip_list.aws_vpc_dev,
    illumio-core_ip_list.aws_vpc_staging,
    illumio-core_ip_list.aws_vpc_prod,
  ]
}
