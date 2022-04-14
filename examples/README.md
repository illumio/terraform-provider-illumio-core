# Illumio Terraform Provider Examples  

This directory contains basic usage examples for each resource and data source type in the Illumio Terraform provider. Each example is self-contained and can be applied to the PCE without modification.  

## Policy Workflows  

The `policy_workflows` folder contains examples that define more complete use-cases. They may be helpful to understand how Illumio objects interact and how complex configuration can be created and managed with Terraform.  

## Running the Examples  

> **Note:** some examples reuse names or other unique attributes for PCE resources. Make sure to destroy any previously applied examples before applying another  

To run the examples, start by cloning the repository and use the `.env.example` (*nix, Mac) or `env.example.bat` (Windows) files as templates to configure the necessary Terraform variables:  

```sh
$ git clone https://github.com/illumio/terraform-provider-illumio-core
$ cd terraform-provider-illumio-core/examples/
$ cp .env.example .env
# update .env with your PCE connection details
$ source .env
$ cd enforcement_boundary/
```

You can then run each example without setting the provider configuration. If you choose not to use the example env files, you'll be prompted to enter the PCE connection information when you run `terraform plan`.  

```sh
$ terraform init
$ terraform plan -out example-plan
$ terraform apply example-plan
```

> **Note:** some examples may require additional setup or customization. For any that do, see the included READMEs for details  

## Remove Example Objects  

Run  

```
$ terraform destroy
```

to remove all objects created by an example. If you've provisioned any of the policy objects, you may need to follow this with  

```sh
$ provision
$ terraform destroy
```

to clear out any dependent objects.  

If you find any issues with the examples, please [file a bug report](https://github.com/illumio/terraform-provider-illumio-core/issues/new/choose) or feel free to [contribute a fix](../.github/CONTRIBUTING.md)!  
