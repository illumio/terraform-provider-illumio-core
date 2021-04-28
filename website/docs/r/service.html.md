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
resource "illumio-core_service" "example_with_service_port" {
  name = "example"
  
  service_port {
    proto = "-1"
  }

  service_port {
    proto = "6"
    port = "15"
  }

  service_port {
    proto = "1"
    icmp_type = "5"
	  icmp_code = "5"
  }

}

resource "illumio-core_service" "example_with_windows_service" {
  name="example"
  
  windows_service {
    service_name = "example"
  }

  windows_service {
    proto = "6"
    process_name = "example"
  }

  windows_service {
    proto="1"
    icmp_type="5"
    icmp_code="5"
  }

}

```

## Schema

### Required

- **name** (String) Name of the servcie (does not need to be unique)

### Optional

- **description** (String) Long description of the servcie
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **process_name** (String) The process name
- **service_port** (Block Set) Service ports (see [below for nested schema](#nestedblock--service_port))
- **windows_service** (Block Set) Windows services (see [below for nested schema](#nestedblock--windows_service))

### Read-Only

- **created_at** (String) Time stamp when this Service was first created
- **created_by** (Map of String) User who originally created this Service
- **deleted_at** (String) Time stamp when this Service was deleted
- **deleted_by** (Map of String) User who deleted this Service
- **description_url** (String) Description URL Read-only to prevent XSS attacks
- **href** (String) URI of the service
- **update_type** (String) Type of update
- **updated_at** (String) Time stamp when this Service was last updated
- **updated_by** (Map of String) User who last updated this Service

<a id="nestedblock--service_port"></a>
### Nested Schema for `service_port`

Required:

- **proto** (String) Transport protocol. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94

Optional:

- **icmp_code** (String) ICMP Code. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 15 inclusive
- **icmp_type** (String) ICMP Type. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 255 inclusive
- **port** (String) Port Number. Also the starting port when specifying a range. Allowed when value of proto is 6 or 17. Allowed range is 0 - 65535 inclusive
- **to_port** (String) High end of port range inclusive if specifying a range. Allowed range is 0 - 65535 inclusive


<a id="nestedblock--windows_service"></a>
### Nested Schema for `windows_service`

Optional:

- **icmp_code** (String) ICMP Code. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range 0 - 15 inclusive
- **icmp_type** (String) ICMP Type. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 255 inclusive
- **port** (String) Port Number. Also the starting port when specifying a range. Allowed when value of proto is 6 or 17. Allowed range is 0 - 65535 inclusive
- **process_name** (String) Name of running process
- **proto** (String) Transport protocol. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94
- **service_name** (String) Name of Windows Service
- **to_port** (String) High end of port range inclusive if specifying a range. Allowed range is 0 - 65535 inclusive


