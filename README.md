<a href="https://terraform.io">
    <img src="https://raw.githubusercontent.com/hashicorp/terraform-website/master/public/img/logo-hashicorp.svg" alt="Terraform logo" title="Terraform" align="right" valign="center" height="75px" />
</a>

# Terraform Provider for Illumio Core  

The Terraform Illumio provider allows users to define HCL configuration to manage resources in the Illumio Policy Compute Engine (PCE).  

For more information about Illumio, please visit the [Illumio Website](https://www.illumio.com). Documentation about the Illumio Core product can be found on the [Illumio documentation portal](https://docs.illumio.com).  

The provider can be used to manage policy and objects within the Illumio Policy Compute Engine. Objects that can be managed in Terraform include, but are not limited to:

- Workloads
- Labels
- IP Lists
- Services
- Security Rules and Rulesets
- Enforcement Boundaries
- Pairing Profiles and Pairing Keys

See [the documentation](https://registry.terraform.io/providers/illumio/illumio-core/latest/docs) for a more comprehensive list.  

The following versions of the Illumio Core Policy Compute Engine are currently supported:  

- PCE v21.2
- PCE v21.5
- PCE v22.2
- SaaS PCEs (v22.5)
- PCE v23.1 (experimental)

## Getting Started  

- [Using the provider](https://registry.terraform.io/providers/illumio/illumio-core/latest/docs)
- [Provider development](DEVELOPMENT.md)
- [Usage examples](./examples/README.md)

## Contributing

For information on how you can contribute to the provider, please refer to the [contributor guidelines](.github/CONTRIBUTING.md).

If you believe you have found a security issue or vulnerability in the provider or in the Illumio Core product, please refer to the [Security document](.github/SECURITY.md) for steps on how to contact the Illumio security team. **Please do not file a public issue.**

## Support

The Illumio-Core Terraform Provider is released and distributed as open source software subject to the [LICENSE](LICENSE). 
Please read the entire [LICENSE](LICENSE) for additional information regarding the permissions and limitations.  

For bugs and feature requests, please open a Github Issue and label it appropriately. 
