---
layout: "illumio-core"
page_title: "illumio-core_virtual_service Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-virtual-service"
subcategory: ""
description: |-
  Represents Illumio Virtual Services
---

# illumio-core_virtual_service (Data Source)

Represents Illumio Virtual Services

Example Usage
------------

```hcl
data "illumio-core_virtual_service" "example"{
  href = "/orgs/1/sec_policy/draft/virtual_services/e2e82190-350c-4034-8096-b67e30123baf"
}

```

## Schema

### Required

- **href** (String) URI of the virtual service

### Read-Only

- **apply_to** (String) Firewall rule target for workloads bound to this virtual service: host_only or internal_bridge_network
- **caps** (List of String) CAPS
- **created_at** (String) Timestamp when this Virtual Service was first created
- **created_by** (Map of String) User who originally created this Virtual Service
- **deleted_at** (String) Timestamp when this Virtual Service was deleted
- **deleted_by** (Map of String) User who deleted this Virtual Service
- **description** (String) The long description of this virtual service
- **external_data_reference** (String) A unque identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **ip_overrides** (List of String) Array of IPs or CIDRs as IP overrides
- **labels** (List of Object) Assigned labels (see [below for nested schema](#nestedatt--labels))
- **name** (String) Name of the virtual service
- **pce_fqdn** (String) PCE FQDN for this container cluster. Used in Supercluster only
- **service** (List of Object) URI of associated service (see [below for nested schema](#nestedatt--service))
- **service_addresses** (List of Object) Serivce Addresses (see [below for nested schema](#nestedatt--service_addresses))
- **service_ports** (List of Object) Service Ports (see [below for nested schema](#nestedatt--service_ports))
- **update_type** (String) Update Type
- **updated_at** (String) Timestamp when this Virtual Service was last updated
- **updated_by** (Map of String) User who last updated this Virtual Service

<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Read-Only:

- **href** (String) URI of label
- **key** (String) Key in key-value pair
- **value** (String) Value in key-value pair

<a id="nestedatt--service"></a>
### Nested Schema for `service`

Read-Only:

- **href** (String) URI of associated service


<a id="nestedatt--service_addresses"></a>
### Nested Schema for `service_addresses`

Read-Only:

- **description** (String) Description of FQDN
- **fqdn** (String) FQDN to assign to the virtual service
- **ip** (String) IP address to assign to the virtual service
- **network** (Map of String) Network Object
- **port** (Number) Port associated with the IP address for the service.


<a id="nestedatt--service_ports"></a>
### Nested Schema for `service_ports`

Read-Only:

- **port** (Number) Port Number. Also the starting port when specifying a range.
- **proto** (Number) High end of port range inclusive if specifying a range. If not specifying a range then don't send this.
- **to_port** (Number) Transport protocol


