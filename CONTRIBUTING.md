# Contributing Guide

## Sign the CLA

Before you can contribute, you will need to sign the [Contributor License Agreement](CLA.md).
## GitHub workflow

Github workflow followed by the project:

- Fork the repo.
- Create a new branch on the fork.
- Push the branch to your fork.
- Submit a pull request!

## Open a Pull Request

For pull requests please follow the standard [github pull request](https://help.github.com/articles/about-pull-requests/) process.

## Bug Reporting

Please report any bug you found at opensource@illumio.com.
You can also open an issue in the issue tracker.

Before reporting any bugs, please do a quick search if any existing bug is already reported or not. If it is already reported, please add a comment on the same issue. It helps in tracking issues.

While reporting a fresh bug, provide a minimal example to reproduce the bug. Include `.tf` files (**REMOVE ANY SECRETS**). Also include `crash.log` in case of panic.


## Issue Assignment

Once a bug/issue is raised and acknowledged as an issue by repo maintainers.
Anyone can work on the issue. Before start working on an issue, make sure to assign the issue to yourself or mention it in the issue. Also, update the progress on the issue regularly.

## Testing

While submitting a new resource/datasource, please make sure you add acceptance tests for them. While updating an existing resource/datasource, update test steps to test the changes done. Also, feel free to add unit tests.

You can refer to the terraform [testing guideline](https://www.terraform.io/docs/extend/testing/index.html) to test resource/datasource.


## Documentation

Documentation is an important aspect of the project. Resource and Datasource changes should reflect in their respective document files.
Also, make sure to add an entry for change in [CHANGELOG](./CHANGELOG.md)

## Development

For Development, refer the [DEVELOPMENT GUIDE](./DEVELOPMENT.md)
