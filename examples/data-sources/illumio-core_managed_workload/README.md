# Managed Workload Example  

`illumio-core_managed_workload` resources cannot be created. For this example to work, you will need to have paired at least one workload. The import statement below gets the first managed workload returned from a GET call to the PCE `/workloads` endpoint:  

> **Note:** the command below requires `curl` and `jq` to be installed on your local system  

```sh
$ terraform init
$ terraform import illumio-core_managed_workload.example "$(curl -s -u "$ILLUMIO_API_KEY_USERNAME:$ILLUMIO_API_KEY_SECRET" "$ILLUMIO_PCE_HOST/api/v2/orgs/$ILLUMIO_PCE_ORG_ID/workloads?managed=true" -H "Accept: application/json" | jq -r '.[0].href')"
$ terraform plan -out example-plan
$ terraform apply example-plan
```

> **Note:** if you want to align the values for the managed workload resource with the values in the PCE, you will need to update the HCL fields manually to match.  
