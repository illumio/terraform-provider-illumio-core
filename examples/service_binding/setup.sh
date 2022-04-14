#!/usr/bin/env bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

# build the provision binary
go build -o ${SCRIPT_DIR}/virtual_service/provision ${SCRIPT_DIR}/../../cmd/provision

cd ${SCRIPT_DIR}/virtual_service

terraform init
# create and provision the virtual service
terraform apply -auto-approve
./provision
