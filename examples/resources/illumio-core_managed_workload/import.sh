# Managed workload objects cannot be created through Terraform, and must
# instead be imported from the Illumio PCE. The imported ID must
# match the HREF of the remote object.

terraform import illumio-core_managed_workload.example "/orgs/1/workloads/aabbccdd-eeff-0011-2233-445566778899"