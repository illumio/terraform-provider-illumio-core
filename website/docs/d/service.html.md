---
layout: "illumio-core"
page_title: "illumio-core_service Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-service"
subcategory: ""
description: |-
  Represents Illumio Service
---

# illumio-core_service (Data Source)

Represents Illumio Service

Example Usage
------------

```hcl
data "illumio-core_service" "example" {
  href = "/orgs/1/sec_policy/draft/services/3"
}
```

## Schema

### Required

- **href** (String) URI of Service

### Read-Only

- **created_at** (String) Timestamp when this Service was first created
- **created_by** (Map of String) User who created this Service
- **deleted_at** (String) Timestamp when this Service was deleted
- **deleted_by** (Map of String) User who deleted this Service
- **description** (String) Long Description of Service
- **description_url** (String) Description URL Read-only to prevent XSS attacks
- **external_data_reference** (String) External data reference identifier
- **external_data_set** (String) External data set identifier
- **name** (String) The short friendly name of the service
- **process_name** (String) The process name
- **service_ports** (List of Object) Service ports of Illumio Service (see [below for nested schema](#nestedatt--service_ports))
- **update_type** (Map of String) Type of update
- **updated_at** (String) Timestamp when this Service was last updated
- **updated_by** (Map of String) User who last updated this Service
- **windows_services** (List of Object) (see [below for nested schema](#nestedatt--windows_services))

<a id="nestedatt--service_ports"></a>
### Nested Schema for `service_ports`

Read-Only:

- **icmp_code** (String) ICMP Code
- **icmp_type** (String) ICMP Type
- **port** (String) Port Number. Also, the starting port when specifying a range.
- **proto** (String) Transport protocol.
- **to_port** (String) High end of port range inclusive if specifying a range.


<a id="nestedatt--windows_services"></a>
### Nested Schema for `windows_services`

Read-Only:

- **icmp_code** (String) ICMP Code. 
- **icmp_type** (String) ICMP Type. 
- **port** (String) Port Number. Also, the starting port when specifying a range.
- **process_name** (String) Name of running process
- **proto** (String) Transport protocol.
- **service_name** (String) Name of Windows Service
- **to_port** (String) High end of port range inclusive if specifying a range.


