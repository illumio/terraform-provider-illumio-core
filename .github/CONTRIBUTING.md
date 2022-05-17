# Contributing Guide  

## GitHub workflow  

Github workflow followed by the project:  

- Fork the repo
- Create a new branch on the fork
- Push the branch to your fork
- Submit a pull request!

## Open a Pull Request  

For pull requests please follow the standard [github pull request](https://help.github.com/articles/about-pull-requests/) process.  

You will be prompted to sign the Illumio Contributor License Agreement when you submit your PR. This step is mandatory for all contributors, but only needs to be completed once per person.  

## Bug Reporting  

> If you find a bug or issue that you believe to be a security vulnerability, please see the [SECURITY](SECURITY.md) document for instructions on reporting. **Do not file a public issue.**

Please report any bugs you find as GitHub issues. If the bug needs an urgent fix, you can send an email to the Illumio [Integrations team](mailto:app-integrations@illumio.com).  

Before reporting any bugs, please do a quick search to see if it has already been reported. If so, please add a comment on the existing issue rather than creating a new one.  

While reporting a bug, please provide a minimal example to reproduce the issue. Include `.tf` files, **making sure to remove any secrets**. If applicable, include the `crash.log` file as well.  

## Testing  

When submitting a new resource or datasource, please follow the current convention of including acceptance tests that set up, verify, and tear down the target resource. When making changes to existing resources or datasources, update the corresponding tests and add any unit tests you deem necessary to ensure the changes are working as expected and have not introduced regressions.  

Refer to the terraform [testing guideline](https://www.terraform.io/docs/extend/testing/index.html) for instructions on testing resources and datasources.  

## Documentation  

Documentation is an important aspect of the project. Changes to resources or datasources should be reflected in their respective docs files. Make sure to update the [CHANGELOG](../CHANGELOG.md) describing your changes. This project follows the [CHANGELOG specification](https://www.terraform.io/plugin/sdkv2/best-practices/versioning#changelog-specification) recommended by Hashicorp.  

## Development  

For Development, refer to the [DEVELOPMENT GUIDE](../DEVELOPMENT.md)  
