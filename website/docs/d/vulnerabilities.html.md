---
layout: "illumio-core"
page_title: "illumio-core_vulnerabilities Data Source - terraform-provider-illumio-core"
sidebar_current: "docs-illumio-core-data-source-vulnerabilities"
subcategory: ""
description: |-
  Represents Illumio Vulnerabilities
---

# illumio-core_vulnerabilities (Data Source)

Represents Illumio Vulnerabilities


Example Usage
------------

```hcl
data "illumio-core_vulnerabilities" "example" {
  max_results = "5"
}
```


## Schema

### Optional

- **max_results** (String) Maximum number of vulnerabilities to return. The integer should be a non-zero positive integer. 

### Read-Only

- **items** (List of Object) list of vulnerabilities (see [below for nested schema](#nestedatt--items))

<a id="nestedatt--items"></a>
### Nested Schema for `items`

Read-Only:

- **created_at** (String) The time (rfc3339 timestamp) at which this report was created
- **created_by** (Map of String) The Href of the user who created this report
- **cve_ids** (Set of String) The cve_ids for the vulnerability
- **description** (String) An arbitrary field to store some details of the vulnerability class
- **name** (String) The title/name of the vulnerability
- **score** (Number) The normalized score of the vulnerability within the range of 0 to 100. CVSS Score can be used here with a 10x multiplier
- **updated_at** (String) The time (rfc3339 timestamp) at which this report was last updated
- **updated_by** (Map of String) The Href of the user who last updated this report


