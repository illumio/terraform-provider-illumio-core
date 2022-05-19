## 0.2.2 (Unreleased)

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
