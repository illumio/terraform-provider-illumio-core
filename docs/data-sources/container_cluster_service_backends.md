---
layout: "illumio-core"
page_title: "illumio-core_container_cluster_service_backends Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-container-cluster-service-backends"
subcategory: ""
description: |-
  Represents Illumio Container Cluster Service Backends
---

# illumio-core_container_cluster_service_backends (Data Source)

Represents Illumio Container Cluster Service Backends

Example Usage
------------

```hcl
data "illumio-core_container_cluster_service_backends" "example" {
  container_cluster_href = illumio-core_container_cluster.example.href
}

resource "illumio-core_container_cluster" "example" {
  ...
}
```

## Schema

### Required

- `container_cluster_href` (String) URI of Container Cluster

### Read-Only

- `items` (List of Object) Collection of Container Cluster Service Backends (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- `created_at` (String) The time (rfc339 timestamp) in which the Container Cluster Backend was created
- `kind` (String) The type (or kind) of Container Cluster Backend
- `name` (String) The name of the Container Cluster Backend
- `namespace` (String) The namespace of the Container Cluster Backend
- `updated_at` (String) The time (rfc339 timestamp) at which the Container Cluster Backend was last updated
- `virtual_service` (List of Object) Associated virtual service. Single element list (see [below for nested schema](#nestedobjatt--items--virtual_service))

<a id="nestedobjatt--items--virtual_service"></a>
### Nested Schema for `items.virtual_service`

Read-Only:

- `href` (String) The URI to the associated virtual service
- `name` (String) The name of virtual service
