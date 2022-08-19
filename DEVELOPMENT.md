# Development Environment Setup

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) (0.13+) (to run acceptance tests)
- [Go](https://golang.org/doc/install) 1.18+ (to build the provider plugin)

## Quick Start

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](#requirements) before proceeding).

> This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume you're using a directory in your home directory outside of the standard GOPATH  

Clone the repository, then run the following `make` commands to install the necessary tools and build a local version of the provider:  

```sh
$ git clone git@github.com:terraform-providers/terraform-provider-illumio-core
$ make tools
$ make build
```

`make build` compile the provider and put the binary under `$GOPATH/bin/terraform-provider-illumio-core`.

### Build binaries for multiple platform

- Run `go get github.com/mitchellh/gox`. This will place gox executable in Go's binary directory. Make sure that is part of the PATH.
- Run `sh scripts/build.sh`. This will create `pkg` directory with binaries for configuration in `build.sh`
- Same can be done for `provision` executable
- Run `cd cmd/provision` and  `sh ../../scripts/build.sh`. This will create `cmd/provision/pkg` directory with binaries for configuration in `build.sh`

## Using the Provider

It's recommended to use local [Developer Overrides](https://www.terraform.io/cli/config/config-file#development-overrides-for-provider-developers) while working on the provider. If using the build steps above, the following `.terraformrc` or `terraform.rc` configuration will set up overrides for the Illumio provider:  

```hcl
provider_installation {
  dev_overrides {
    "illumio/illumio-core" = "/path/to/GOPATH/bin"
  }

  direct {}
}
```

Move the generated binary from the build step to the [plugin directory](https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories)/illumio/illumio-core/`<version>`/`<os>_<arch>`. Examples for `<os>_<arch>` are `windows_amd64`, `linux_arm`, `darwin_amd64`, etc. After placing it into your plugins directory, run `terraform init` to initialize it.  

> **Note:** PCE connection parameters for the provider can also be set with the `ILLUMIO_PCE_HOST`, `ILLUMIO_API_KEY_USERNAME` and `ILLUMIO_API_KEY_SECRET` environment variables. Optionally, the `ILLUMIO_PCE_ORG_ID` variable can be set to specify a non-default Organization ID. See the [Provider documentation](https://registry.terraform.io/providers/illumio/illumio-core/latest/docs) for details.  

Example:

```hcl
terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  pce_host      = "https://pce.my-company.com:8443"
  api_username  = "api_xxxxxx"
  api_secret    = "xxxxxxxxxx"
  org_id        = 10
}

resource "illumio-core_container_cluster" "example" {
    name        = "Container cluster name"
    description = "Container cluster desc"
}
```

## Testing the Provider

Unit tests can be run for the provider with `make test`.  

> **Note:** the `ILLUMIO_PCE_HOST`, `ILLUMIO_API_KEY_USERNAME` and `ILLUMIO_API_KEY_SECRET` variables must be set for tests to run. Organization IDs other than the default (1) can be set with `ILLUMIO_PCE_ORG_ID`.  

```sh
$ make test
```

In order to run the full suite of acceptance tests, run `make testacc`.  

> **Note:** Acceptance tests create real resources in the PCE, and may cost money to run.  

```sh
$ make testacc
```

To get code coverage, set `-coverprofile=cover.out` while running `go test` command.
To analyze code coverage, we can use the standard go tool [cover](https://golang.org/cmd/cover/).

To check coverage per function - `go tool cover -func=cover.out`
To check code lines covered - `go tool cover -html=cover.out`

*Note: Current code coverage artifacts available at [here](.code-coverage/)*

## Debugging and Troubleshooting  

- Set environment variable `TF_LOG` to one of the log levels `TRACE`, `DEBUG`, `INFO`, `WARN` or `ERROR`
- Set environment variable `TF_LOG_PATH` to write logs in a file. e.g. `TF_LOG_PATH=tf.log`

For more details, see the [Terraform Debugging](https://www.terraform.io/docs/internals/debugging.html) documentation.  

## Documentation  

Once done with changes/development of any resource/datasource, document the changes.  

- Update parameter/description as per the resource changes.
- If a resource is designed to behave in a specific way that might be strange to the end-user, that should be documented.
- Add an example usage section and add different examples about how to use it.
- Document any environment variable on which resource/datasource is dependent in the example usage section or in the argument description itself.
