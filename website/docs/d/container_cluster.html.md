---
layout: "illumio-core"
page_title: "illumio-core_container_cluster Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-container-cluster"
subcategory: ""
description: |-
  Represents Illumio Container Cluster
---

# illumio-core_container_cluster (Data Source)

Represents Illumio Container Cluster

Example Usage
------------

```hcl
data "illumio-core_container_cluster" "test" {
    href = "/orgs/1/container_clusters/bd37cbdd-82bd-4f49-b52f-9405ba236a43"
}
```

## Schema

### Required

- **href** (String) URI of the Cluster

### Read-Only

- **caps** (List of String) Permission types
- **container_runtime** (String) The Container Runtime used in this Cluster
- **description** (String) Description of the Cluster
- **errors** (List of Object) Errors for Cluster (see [below for nested schema](#nestedatt--errors))
- **kubelink_version** (String) Kubelink software version string for Cluster
- **last_connected** (String) Time the Cluster last connected to
- **manager_type** (String) Manager for this Cluster (and version)
- **name** (String) Name of the Cluster
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


