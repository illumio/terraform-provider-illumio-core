## 1.0.0 (TBD)  

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

NOTES:  

* Add `interfaces` back into managed/unmanaged workload schema  

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
