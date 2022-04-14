---
name: Bug report
about: Report an issue in the Illumio provider
---

<!-- Problem overview -->

## Expected Result

<!-- What you expected to happen -->

## Actual Result

<!-- What actually happened -->

## Steps to reproduce

<!-- Provide the simplest HCL and PCE configuration that reproduces the issue -->
```hcl
terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {}

...
```

## System Information

<!-- Information about your system and setup - Terraform, provider, and PCE versions, OS, plugins installed,
    anything you think could help identify or may contribute to the problem -->

**Are you able to help contribute a fix for this issue?** yes/no
