---
layout: "illumio-core"
page_title: "illumio-core_services Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-services"
subcategory: ""
description: |-
  Represents Illumio Services
---

# illumio-core_services (Data Source)

Represents Illumio Services

Example Usage
------------

```hcl
data "illumio-core_services" "example" {
  max_results = "5"
}
```

## Schema

### Optional

- **description** (String) Long description of the Service
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **max_results** (String) Maximum number of Services to return. The integer should be a non-zero positive integer
- **name** (String) Name of the Service (does not need to be unique)
- **port** (String) Specify port or port range to filter results. The range is from -1 to 65535 (0 is not supported)
- **proto** (String) Protocol to filter on. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94
- **pversion** (String) pversion of the security policy. Allowed values are "draft", "active", and numbers greater than 0. Default value: "draft"

### Read-Only

- **items** (List of Object) List of services (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **created_at** (String) Timestamp when this Service was first created
- **created_by** (Map of String) User who created this Service
- **deleted_at** (String) Timestamp when this Service was deleted
- **deleted_by** (Map of String) User who deleted this Service
- **description** (String) Long Description of Service
- **description_url** (String) Description URL Read-only to prevent XSS attacks
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **href** (String) URI of Service
- **name** (String) The short friendly name of the service
- **process_name** (String) The process name
- **service_ports** (List of Object) Service ports of Illumio Service (see [below for nested schema](#nestedobjatt--items--service_ports))
- **update_type** (Map of String) Type of update
- **updated_at** (String) Timestamp when this Service was last updated
- **updated_by** (Map of String) User who last updated this Service
- **windows_services** (List of Object) windows_services for Services (see [below for nested schema](#nestedobjatt--items--windows_services))

<a id="nestedobjatt--items--service_ports"></a>
### Nested Schema for `items.service_ports`

Read-Only:

- **icmp_code** (String) ICMP Code
- **icmp_type** (String) ICMP Type
- **port** (String) Port Number. Also the starting port when specifying a range
- **proto** (String) Transport protocol
- **to_port** (String) High end of port range inclusive if specifying a range


<a id="nestedobjatt--items--windows_services"></a>
### Nested Schema for `items.windows_services`

Read-Only:

- **icmp_code** (String) ICMP Code
- **icmp_type** (String) ICMP Type
- **port** (String) Port Number. Also the starting port when specifying a range
- **process_name** (String) Name of running process
- **proto** (String) Transport protocol
- **service_name** (String) Name of Windows Service
- **to_port** (String) High end of port range inclusive if specifying a range


