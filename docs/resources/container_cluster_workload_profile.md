---
layout: "illumio-core"
page_title: "illumio-core_container_cluster_workload_profile Resource - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-resource-container-cluster-workload-profile"
subcategory: ""
description: |-
  Manages Illumio Container Cluster
---

# illumio-core_container_cluster_workload_profile (Resource)

Manages Illumio Container Cluster

Example Usage
------------

```hcl
resource "illumio-core_container_cluster_workload_profile" "example" {
  container_cluster_href = illumio-core_container_cluster.example.href
  name                   = "CCWP-EXAMPLE"
  description            = "Container Cluster Workload profile example"
  managed                = true
  enforcement_mode       = "visibility_only"

  assign_labels {
    href = illumio-core_label.example.href
  }
}

resource "illumio-core_container_cluster" "example" {
  ...
}

resource "illumio-core_label" "example" {
  ...
}
```

## Schema

### Required

- `container_cluster_href` (String) URI of Container Cluster
- `name` (String) A friendly name given to a profile if the namespace is not user-friendly. The name should be up to 255 characters

### Optional

- `assign_labels` (Block Set) Assigned labels container workload profile (see [below for nested schema](#nestedblock--assign_labels))
- `description` (String) Description of the container workload profile
- `enforcement_mode` (String) Enforcement mode of container workload profiles to return. Allowed values for enforcement modes are "idle","visibility_only", "full", and "selective". Default value: "idle"
- `labels` (Block Set) Labels to assign to the workload that matches the namespace (see [below for nested schema](#nestedblock--labels))
- `managed` (Boolean) If the namespace is managed or not

### Read-Only

- `created_at` (String) Timestamp when this label group was first created
- `created_by` (Map of String) User who created this label group
- `href` (String) URI of the container workload profile
- `linked` (Boolean) True if the namespace exists in the cluster and is reported by Kubelink
- `namespace` (String) Namespace name of the container workload profile
- `updated_at` (String) Timestamp when this label group was last updated
- `updated_by` (Map of String) User who last updated this label group
- `visibility_level` (String) Visibility Level of the container cluster workload profile

<a id="nestedblock--assign_labels"></a>
### Nested Schema for `assign_labels`

Required:

- `href` (String) URI of the assigned label

<a id="nestedblock--labels"></a>
### Nested Schema for `labels`

Required:

- `key` (String) Key of the Label. Allowed values for key are "role", "loc", "app" and "env"

Optional:

- `assignment` (Block Set, Max: 1) The label href to set. Single element list (see [below for nested schema](#nestedblock--labels--assignment))
- `restriction` (Block Set) The list of allowed label hrefs (see [below for nested schema](#nestedblock--labels--restriction))

<a id="nestedblock--labels--assignment"></a>
### Nested Schema for `labels.assignment`

Required:

- `href` (String) URI of label

Read-Only:

- `value` (String) Name of label

<a id="nestedblock--labels--restriction"></a>
### Nested Schema for `labels.restriction`

Required:

- `href` (String) URI of label

Read-Only:

- `value` (String) Name of label
