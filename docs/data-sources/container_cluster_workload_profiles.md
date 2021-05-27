---
layout: "illumio-core"
page_title: "illumio-core_container_cluster_workload_profiles Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-container-cluster-workload-profiles"
subcategory: ""
description: |-
  Represents Illumio Container Cluster Workload Profiles
---

# illumio-core_container_cluster_workload_profiles (Data Source)

Represents Illumio Container Cluster Workload Profiles

Example Usage
------------

```hcl
data "illumio-core_container_cluster_workload_profiles" "example" {
  max_results = "5"
  container_cluster_href = "/orgs/1/container_clusters/f959d2d0-fe56-4bd9-8132-b7a31d1cbdde"
}
```

## Schema

### Required

- **container_cluster_href** (String) URI of the Container Cluster

### Optional

- **assign_labels** (String) List of label URIs, encoded as a JSON string
- **enforcement_mode** (String) Filter by enforcement mode. Allowed values for enforcement modes are "idle", "visibility_only", "full", and "selective"
- **linked** (String) Filter by linked container workload profiles
- **managed** (String) Filter by managed state
- **max_results** (String) Maximum number of container workloads to return. The integer should be a non-zero positive integer
- **name** (String) Name string to match.Supports partial matches
- **namespace** (String) Namespace string to match.Supports partial matches
- **visibility_level** (String) Filter by visibility level. Allowed values are "flow_full_detail", "flow_summary", "flow_drops", "flow_off" and "enhanced_data_collection"

### Read-Only

- **items** (List of Object) List of Container Cluster Workload Profiles (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **assign_labels** (List of Object) Assigned labels container workload profile (see [below for nested schema](#nestedobjatt--items--assign_labels))
- **created_at** (String) Timestamp when this label group was first created
- **created_by** (Map of String) User who created this label group
- **description** (String) Description of the container workload profile
- **enforcement_mode** (String) Enforcement mode of container workload profiles to return
- **labels** (List of Object) Labels to assign to the workload that matches the namespace (see [below for nested schema](#nestedobjatt--items--labels))
- **linked** (Boolean) True if the namespace exists in the cluster and is reported by Kubelink
- **managed** (Boolean) If the namespace is managed or not
- **name** (String) A friendly name given to a profile if the namespace is not user-friendly
- **namespace** (String) Namespace name of the container workload profile
- **updated_at** (String) Timestamp when this label group was last updated
- **updated_by** (Map of String) User who last updated this label group
- **visibility_level** (String) Visibility Level of the container cluster workload profile

<a id="nestedobjatt--items--assign_labels"></a>
### Nested Schema for `items.assign_labels`

Read-Only:

- **href** (String) URI of the assigned label


<a id="nestedobjatt--items--labels"></a>
### Nested Schema for `items.labels`

Read-Only:

- **assignment** (List of Object) The label href to set. Single element list (see [below for nested schema](#nestedobjatt--items--labels--assignment))
- **key** (String) Key of the Label
- **restriction** (List of Object) The list of allowed label hrefs (see [below for nested schema](#nestedobjatt--items--labels--restriction))

<a id="nestedobjatt--items--labels--assignment"></a>
### Nested Schema for `items.labels.assignment`

Read-Only:

- **href** (String) URI of label
- **value** (String) Name of label


<a id="nestedobjatt--items--labels--restriction"></a>
### Nested Schema for `items.labels.restriction`

Read-Only:

- **href** (String) URI of label
- **value** (String) Name of label


