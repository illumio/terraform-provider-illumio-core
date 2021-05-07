---
layout: "illumio-core"
page_title: "illumio-core_virtual_services Data Source - terraform-provider-illumio-core"
subcategory: ""
sidebar_current: "docs-illumio-core-data-source-virtual-services"
description: |-
  Represents Illumio Virtual Services
---

# illumio-core_virtual_services (Data Source)

Represents Illumio Virtual Services


Example Usage
------------

```hcl
data "illumio-core_virtual_services" "vs_1"{
  max_results = 5
}
```


## Schema

### Optional

- **description** (String) Description on which to filter. Supports partial matches
- **external_data_reference** (String) A unique identifier within the external data source
- **external_data_set** (String) The data source from which a resource originates
- **labels** (String) List of lists of label URIs, encoded as a JSON string
- **max_results** (String) Maximum number of Virtual Services to return.
- **name** (String) Name on which to filter. Supports partial matches
- **pversion** (String) pversion of the security policy. Allowed values are "draft", "active" and numbers greater than 0. Default value: "draft"
- **service** (String) Service URI
- **service_address_fqdn** (String) FQDN configured under service_address property, supports partial matches
- **service_address_ip** (String) IP address configured under service_address property, supports partial matches
- **service_ports_port** (String) Specify port or port range to filter results. The range is from -1 to 65535.
- **service_ports_proto** (String) Protocol to filter on
- **usage** (String) Include Virtual Service usage flags

### Read-Only

- **items** (List of Object) list of virtual services (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

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
- **labels** (List of Object) Assigned labels (see [below for nested schema](#nestedobjatt--items--labels))
- **name** (String) Name of the virtual service
- **pce_fqdn** (String) PCE FQDN for this container cluster. Used in Supercluster only
- **service** (List of Object) URI of associated service (see [below for nested schema](#nestedobjatt--items--service))
- **service_addresses** (List of Object) Serivce Addresses (see [below for nested schema](#nestedobjatt--items--service_addresses))
- **service_ports** (List of Object) Service Ports (see [below for nested schema](#nestedobjatt--items--service_ports))
- **update_type** (String) Update Type
- **updated_at** (String) Timestamp when this Virtual Service was last updated
- **updated_by** (Map of String) User who last updated this Virtual Service

<a id="nestedobjatt--items--labels"></a>
### Nested Schema for `items.labels`

Read-Only:

- **href** (String) URI of label
- **key** (String) Key in key-value pair
- **value** (String) Value in key-value pair


<a id="nestedobjatt--items--service"></a>
### Nested Schema for `items.service`

Read-Only:

- **href** (String) URI of associated service


<a id="nestedobjatt--items--service_addresses"></a>
### Nested Schema for `items.service_addresses`

Read-Only:

- **description** (String) Description of FQDN
- **fqdn** (String) FQDN to assign to the virtual service
- **ip** (String) IP address to assign to the virtual service
- **network** (Map of String) Network Object
- **port** (Number) Port associated with the IP address for the service.


<a id="nestedobjatt--items--service_ports"></a>
### Nested Schema for `items.service_ports`

Read-Only:

- **port** (Number) Port Number. Also the starting port when specifying a range.
- **proto** (Number) High end of port range inclusive if specifying a range. If not specifying a range then don't send this.
- **to_port** (Number) Transport protocol


