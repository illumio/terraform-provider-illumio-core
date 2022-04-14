# Service Bindings Example  

The `illumio-core_service_binding` resource must reference a virtual service that has already been provisioned and is in the `active` state. The virtual service configuration is provided in a separate subdirectory and referenced with a `local` backend data source.  

For convenience, `setup.sh` and `teardown.sh` scripts are included to create and remove the virtual service respectively.  

## Setup  

> **Note:** `setup.sh` requires `go` to be installed in order to build the `provision` binary

```sh
$ ./setup.sh
$ terraform init
$ terraform plan -out example-plan
$ terraform apply example-plan
```

## Teardown  

> **Note:** this will remove all created objects
```sh
$ ./teardown.sh
```
