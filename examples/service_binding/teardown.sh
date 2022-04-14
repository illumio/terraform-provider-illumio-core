#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
cd ${SCRIPT_DIR}

# tear down the workload and service binding first
terraform destroy -auto-approve

cd ${SCRIPT_DIR}/virtual_service

# provision removal of the virtual service
terraform destroy -auto-approve
./provision
# a second call to destroy is required to get rid of the
# labels due to the references in the virtual service
terraform destroy -auto-approve

# remove the provision binary
rm provision
