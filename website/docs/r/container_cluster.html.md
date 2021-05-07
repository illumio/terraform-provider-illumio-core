---
layout: "illumio-core"
page_title: "illumio-core_container_cluster resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-container-cluster"
subcategory: ""
description: |-
  Manages Illumio Container Cluster
---

# illumio-core_container_cluster (Resource)

Manages Illumio Container Cluster


Example Usage
------------

```hcl
resource "illumio-core_container_cluster" "example" {
    name = "name example"
    description = "description example"
}
```

## Schema

### Optional

- **description** (String) Description of the Cluster
- **name** (String) Name of the Cluster. The name should be  upto 255 characters.

### Read-Only

- **caps** (List of String) Permission types
- **container_runtime** (String) The Container Runtime used in this Cluster
- **errors** (List of Object) Errors for Cluster (see [below for nested schema](#nestedatt--errors))
- **href** (String) URI of the Cluster
- **kubelink_version** (String) Kubelink software version string for Cluster
- **last_connected** (String) Time the Cluster last connected to
- **manager_type** (String) Manager for this Cluster (and version)
- **nodes** (List of Object) Nodes of the Cluster (see [below for nested schema](#nestedatt--nodes))
- **online** (Boolean) Whether the Cluster is online or not
- **pce_fqdn** (String) PCE FQDN for this container cluster. Used in Supercluster only

<a id="nestedatt--errors"></a>
### Nested Schema for `errors`

Read-Only:

- **audit_event** (Map of String) Audit Event of Error
- **duplicate_ids** (List of String) Duplicate IDs of Error
- **error_type** (String) Error Type of Error


<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

Read-Only:

- **pod_subnet** (String) Pod Subnet of the node


