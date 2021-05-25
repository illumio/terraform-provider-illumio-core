---
layout: "illumio-core"
page_title: "illumio-core_service Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-service"
subcategory: ""
description: |-
  Manages Illumio Security Service
---

# illumio-core_service (Resource)

Manages Illumio Security Service

Example Usage
------------

```hcl
resource "illumio-core_service" "example_service" {
  name = "example name service ports"
  
  service_ports {
    proto = -1
  }

  service_ports {
    proto = 6
    port = 15
  }

  service_ports {
    proto = 1
    icmp_type = 5
	  icmp_code = 5
  }

}

resource "illumio-core_service" "example_windows" {
  name="example name windows service"
  
  windows_services {
    service_name = "example"
  }

  windows_services {
    proto = 6
    process_name = "example name"
  }

  windows_services {
    proto = 1
    icmp_type = 5
    icmp_code = 5
  }
}

```

## Schema

### Required

- **name** (String) Name of the Service (does not need to be unique). The name should be between 1 to 255 characters

### Optional

- **description** (String) Long description of the Service
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **process_name** (String) The process name. The name should be between 1 to 255 characters
- **service_ports** (Block Set) Service ports (see [below for nested schema](#nestedblock--service_ports))
- **windows_services** (Block Set) Windows services (see [below for nested schema](#nestedblock--windows_services))

### Read-Only

- **created_at** (String) Timestamp when this Service was first created
- **created_by** (Map of String) User who created this Service
- **deleted_at** (String) Timestamp when this Service was deleted
- **deleted_by** (Map of String) User who deleted this Service
- **description_url** (String) Description URL Read-only to prevent XSS attacks
- **href** (String) URI of the service
- **update_type** (String) Type of update
- **updated_at** (String) Timestamp when this Service was last updated
- **updated_by** (Map of String) User who last updated this Service

<a id="nestedblock--service_ports"></a>
### Nested Schema for `service_ports`

Required:

- **proto** (String) Transport protocol. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94

Optional:

- **icmp_code** (String) ICMP Code. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 15
- **icmp_type** (String) ICMP Type. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 255
- **port** (String) Port Number. Also, the starting port when specifying a range. Allowed when value of proto is 6 or 17. Allowed range is 0 - 65535
- **to_port** (String) High end of port range if specifying a range. Allowed range is 0 - 65535


<a id="nestedblock--windows_services"></a>
### Nested Schema for `windows_services`

Optional:

- **icmp_code** (String) ICMP Code. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range 0 - 15
- **icmp_type** (String) ICMP Type. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 255
- **port** (String) Port Number. Also, the starting port when specifying a range. Allowed when value of proto is 6 or 17. Allowed range is 0 - 65535
- **process_name** (String) Name of running process
- **proto** (String) Transport protocol. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94
- **service_name** (String) Name of Windows Service
- **to_port** (String) High end of port range if specifying a range. Allowed range is 0 - 65535


