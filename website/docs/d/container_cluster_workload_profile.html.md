---
layout: "illumio-core"
page_title: "illumio-core_container_cluster_workload_profile Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-container-cluster-workload-profile "
subcategory: ""
description: |-
  Represents Illumio Container Cluster Workload Profile
---

# illumio-core_container_cluster_workload_profile (Data Source)

Represents Illumio Container Cluster Workload Profile

Example Usage
------------

```hcl
data "illumio-core_container_cluster_workload_profile" "test" {
    container_cluster_id = "deb48c70-e9d2-4101-ab7e-1f48de922ff4"
    container_workload_profile_id = "0a7ed380-bc2e-4be6-99ad-741baf77fb91"
}
```

## Schema

### Required

- **container_cluster_id** (String) Numerical ID of Container Cluster
- **container_workload_profile_id** (String) Numerical ID of Container Workload Profile


### Read-Only

- **assign_labels** (Set of Object) Assigned labels container workload profile (see [below for nested schema](#nestedatt--assign_labels))
- **created_at** (String) Timestamp when this label group was first created
- **created_by** (Map of String) User who originally created this label group
- **description** (String) Description of the container workload profile
- **enforcement_mode** (String) Enforcement mode of container workload profiles to return.
- **href** (String) URI of the container workload profile
- **labels** (Set of Object) Labels to assign to the workload that matches the namespace (see [below for nested schema](#nestedatt--labels))
- **linked** (Boolean) True if the namespace exists in the cluster and is reported by kubelink.
- **managed** (Boolean) If the namespace is managed or not.
- **name** (String) A friendly name given to a profile if the namespace is not user friendly
- **namespace** (String) Namespace name of the container workload profile
- **updated_at** (String) Timestamp when this label group was last updated
- **updated_by** (Map of String) User who last updated this label group
- **visibility_level** (String) Visibility Level of the container cluster workload profile.

<a id="nestedatt--assign_labels"></a>
### Nested Schema for `assign_labels`

Read-Only:

- **href** (String) URI of the assigned label

<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Read-Only:

- **assignment** (Set of Object) The label href to set. Single element list (see [below for nested schema](#nestedobjatt--labels--assignment))
- **key** (String) Key of the Label.
- **restriction** (Set of Object) The list of allowed label hrefs (see [below for nested schema](#nestedobjatt--labels--restriction))

<a id="nestedobjatt--labels--assignment"></a>
### Nested Schema for `labels.assignment`

Read-Only:

- **href** (String) URI of label
- **value** (String) Name of label

<a id="nestedobjatt--labels--restriction"></a>
### Nested Schema for `labels.restriction`

Read-Only:

- **href** (String) URI of label
- **value** (String) Name of label


