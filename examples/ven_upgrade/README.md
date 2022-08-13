# VEN Upgrade Example  

The `illumio-core_vens_upgrade` resource has been removed. This example shows how Terraform's built-in `local-exec` provisioner can be leveraged to perform an upgrade of all VENs in the PCE to a given version (set using the `ven_version` TF variable).  

> **Note:** the command below requires `curl` and `jq` to be installed on your local system

```sh
$ terraform init
$ terraform plan -out example-plan
$ terraform apply example-plan
```
