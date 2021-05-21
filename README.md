<a href="https://terraform.io">
    <img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider for Illumio-Core

The Terraform Illumio-Core provider is a plugin for Terraform that one can use with Terraform to work with Illumio Core. Using this provider, Illumio security policies can be created and applied to your workloads.. 

For more information about Illumio, please visit https://www.illumio.com

Documentation about Illumio Core can be found at [Illumio Docs portal](https://docs.illumio.com)

The following versions of Illumio Core are supported:
- Illumio Core 21.2

In case of a security finding in the Terraform Illumio-Core Provider, please inform security@illumio.com

## Getting Started

- [Using the provider](docs/index.md)
- [Develop, Build, Test, and Debug the provider](./DEVELOPMENT.md)

The primary use-case for the Illumio Core provider is managing the following resources:
- labels
- iplists
- pairing profiles and pairing keys
- workloads 
- rules and rulesets using labels, and iplists
This use-case is verified and validated with Acceptance tests.  Over time, other use-cases will be developed. 

If you have a specific policy use-case for this provider, please contact opensource@illumio.com  

### Security Policy Provisioning  

When Illumio policy is created or updated, the changes are staged as `draft`.
Illumio policy needs to be provisioned to make the policy changes `active`.

This provisioning is only done after all the changes are complete.
To execute provisioning, a separate tool `provision` is provided.

This `provision` tool should be executed after running `terraform apply`. 
For example: `terrafrom apply && provision`

Similarly, after running `terraform destroy`, the changes are made active using provision. 

## Contributing

To contribute, please refer the [contributor guidelines](./CONTRIBUTING.md)

## Support

The Illumio-Core Terraform Provider is released and distributed as open source software subject to the [LICENSE](LICENSE). 
Please read the entire [LICENSE](LICENSE) for additional information regarding the permissions and limitations. 

For bugs and feature requests, please open a Github Issue and label it appropriately. 
