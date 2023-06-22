## 1.1.3 (Jun 21, 2023)

BUG FIXES:

* update client to not pass an error back if a DELETE returns a 406 `already-deleted` response

ENHANCEMENTS:

* set `ForceNew` on the deleted flag for the following resources and add customizers to ensure a remote deletion prompts Terraform to recreate the resource
    * `illumio-core_label`
    * `illumio-core_label_type`
    * `illumio-core_unmanaged_workload`

## 1.1.2 (Apr 5, 2023)

BUG FIXES:

* Fix `illumio-core_rule_set` resource to allow more than 3 label/label group refs in each scope

## 1.1.1 (Mar 23, 2023)

BUG FIXES:

* Fix resource updates broken by model changes
* Fix bugs and missing fields in several resources
* Fix consumer/provider exclusions not being applied for `illumio-core_security_rule` resources

ENHANCEMENTS:

* The following resources now allow create operations, and will "adopt" the remote object rather than returning an error:
    * illumio-core_firewall_settings
    * illumio-core_organization_settings
    * illumio-core_workload_settings

## 1.1.0 (Feb 24, 2023)

BUG FIXES:

* Allow arbitrary string values for label keys to work with label changes introduced in PCE v22.5

**NOTE:** this change, while backwards compatible, removes validation guarantees for older PCE versions. Versions prior to 22.5 still restrict key values to `role`, `app`, `env`, and `loc`.

* Several resources/data sources have been fixed so that import state mirrors the resource schema

NEW FEATURES:

* **New resource:** `label_type`

As of PCE v22.5, new label types can be configured (up to a default hard limit of 20 types). The `illumio-core_label_type` resource has been added to manage these types through the provider.

`illumio-core_label_type` and `illumio-core_label_types` data sources have also been added.

ENHANCEMENTS:

* Broad refactor to simplify object conversions

SCHEMA UPDATES:

* `resource.illumio-core_virtual_service` v2 - The `network_href` attribute has been changed to match the API schema and data source. It is now a nested object with an HREF attribute:

```
resource "illumio-core_virtual_service" "example" {
    ...

    network {
        href = "/orgs/1/networks/..."
    }
}
```

This change was necessary to align the import state with virtual service resource definitions.

* added `ven_type` nested fields to settings blocks in `illumio-core_workload_settings`

* added `use_workload_subnets` field to `illumio-core_security_rule`

## 1.0.3 (Dec 16, 2022)

BUG FIXES:

* Fix missing label sets in `unmanaged_workload` and `managed_workload` resource state

## 1.0.2 (Oct 4, 2022)

BUG FIXES:

* Fix issue where Label SubGroups whose parent group is not in the Terraform state will fail to provision
    * Add `/member_of` call to Label Group resource delete function to add parent group HREFs to provision list

ENHANCEMENTS:

* Refactor `provision` binary

## 1.0.1 (Sep 9, 2022)

BUG FIXES:

* Remove valid port slice constraint for service creation in favour of bounded int validation (-1 <= N <= 255) to support any IANA proto number

## 1.0.0 (Aug 19, 2022)  

BREAKING CHANGES:  

* **resource/workload_interface has been REMOVED**  
* **data_source/workload_interface has been REMOVED**  

The PCE `/workloads/:uuid/interfaces/:name` endpoint is incompatible with Terraform in its current implementation.
The API uses the interface name as a unique identifier, but interfaces were later changed to support multiple addresses (IPv4 and IPv6) without a change to the underlying representation expecting names and addresses to be 1-1. Since Terraform expects a unique relationship for resources to remote objects, it's not possible to define an interface with both an IPv4 and IPv6 address using HCL.  

The `/workloads` API, on the other hand, *does* return multiple interface objects with the same name, and a `/workloads` POST can be used to create an interface with both IPv4 and IPv6 addresses (though they still need to be defined as distinct objects!).  

As such, and because interfaces have no interaction with objects outside of the workload they belong to, the workload interface object is removed and the schema for both managed and unmanaged workloads have been updated to include an `interfaces` set.  

* **resource/vens_unpair has been REMOVED**  

VEN unpairing has been moved to the delete call for managed_workload resources to provide a workflow that is more consistent with Terraform's resource lifecycle. As such, the dedicated unpair resources are no longer needed.  

* **resource/workloads_unpair has been REMOVED**  

The PCE `/workloads/unpair` API has been deprecated in favour of `/vens/unpair` for a while, and is made doubly redundant by the addition of the managed_workload resource.  

* **resource/vens_upgrade has been REMOVED**  

vens_upgrade was another example of a utility resource that didn't conform to the Terraform resource lifecycle. Terraform already has the `null_resource` type and `local-exec` provisioner to handle these special cases if needed. An example of how to perform an upgrade using `local-exec` has been added to the `examples/ven_upgrade` folder.  

* **resource/workload has been REMOVED**

Workloads were split into the distinct managed_workload/unmanaged_workload resources in version 0.2.0. The workload resource simply duplicated unmanaged workload functionality, and is no longer necessary.

ENHANCEMENTS:

* Add `interfaces` back into managed/unmanaged workload schema  
* Add import to several objects
  * `resource/container_cluster_workload_profile`
  * `resource/enforcement_boundary`
  * `resource/ip_list`
  * `resource/pairing_profile`
  * `resource/rule_set`
  * `resource/security_rule`
  * `resource/service_binding`
  * `resource/syslog_destination`
  * `resource/traffic_collector_settings`
  * `resource/virtual_service`
  * `resource/vulnerability`
  * `resource/vulnerability_report`
* Fix docs typos

NOTES:  

* Bump pinned go version to 1.18
* Add arm64 to gox build targets

BUG FIXES:

* Extract parent HREF for child objects (rules/container workload profiles) on read to set for imports

## 0.2.2 (May 19, 2022)

NOTES:

* Remove reflective mapping of workload objects in favour of splitting managed/unmanaged model mapping

BUG FIXES:

* Fix `ignored_interface_names` parameter to be editable for `managed_workload` resources

## 0.2.1 (May 17, 2022)

NOTES:

* Clean up commented code from previous change

BUG FIXES:

* Fix bug causing managed workload update to fail in some cases

## 0.2.0 (May 10, 2022)

NOTES:

* resource/workload: DEPRECATED - use `resource/unmanaged_workload`. Will be removed in v1.0.0
* resource/workloads_unpair: DEPRECATED. Will be removed in v1.0.0
* resource/vens_unpair: DEPRECATED. Will be removed in v1.0.0
* Users should instead use the `destroy` lifecycle operation on imported `managed_workload` resources to unpair their VENs
* resource/vens_upgrade: DEPRECATED. Will be removed in v1.0.0

FEATURES:

* **New resource:** `managed_workload`
* **New resource:** `unmanaged_workload`
* resource/workload: Split managed/unmanaged workloads into separate resources
* Update workload model logic to accommodate both payloads

ENHANCEMENTS:

* Restructure and rewrite tests to be PCE-agnostic
* Improve documentation
* Improve HCL examples and remove duplicate JSON examples

BUG FIXES:

* Remove auto-provision from `resource/virtual_service` as the provision calls if the virtual service references other draft objects

## 0.1.1 (Oct 28, 2021)

ENHANCEMENTS:

* Minor bug fixes 
* Update examples and documentation

BUG FIXES:

* Fix race conditions for nested dependent objects (rule_set/rule, workload/workload_interface) [GH-15]
* Fix state inconsistency after `terraform apply` on `resource/pairing_key` [GH-17]

## 0.1.0 (May 31, 2021)

* Initial Release
