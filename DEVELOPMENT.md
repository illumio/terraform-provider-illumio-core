# Development Environment Setup

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) (0.13, 0.14) (to run acceptance tests)
- [Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)

## Quick Start

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](#requirements) before proceeding).

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e `$HOME/development/terraform-providers/`).

Clone repository to: `$HOME/development/terraform-providers/`

```sh
$ mkdir -p $HOME/development/terraform-providers/; cd $HOME/development/terraform-providers/
$ git clone git@github.com:terraform-providers/terraform-provider-illumio-core
...
```

Enter the provider directory and run `make tools`. This will install the needed tools for the provider.

```sh
$ make tools
```

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-illumio-core
...
```

### Build binaries for multiple platform

- Run `go get github.com/mitchellh/gox`. This will place gox executable in Go's binary directory. Make sure that is part of the PATH.
- Run `sh scripts/build.sh`. This will create `pkg` directory with binaries for configuration in `build.sh`
- Same can be done for `provision` executable
- Run `cd cmd/provision` and  `sh ../../scripts/build.sh`. This will create `cmd/provision/pkg` directory with binaries for configuration in `build.sh`

## Using the Provider

Move the generated binary from the build step to the [plugin directory](https://www.terraform.io/docs/cli/config/config-file.html#implied-local-mirror-directories)/illumio.com/labs/illumio-core/`<version>`/`<os>_<arch>`. Examples for `<os>_<arch>` are `windows_amd64`, `linux_arm`, `darwin_amd64`, etc. be After placing it into your plugins directory, run `terraform init` to initialize it.

*Note:* Make sure `ILLUMIO_PCE_HOST`, `ILLUMIO_API_KEY_USERNAME` and `ILLUMIO_API_KEY_SECRET` variables are set.

Example
```hcl
terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
  request_timeout = 30
  org_id          = 1
}

resource "illumio-core_container_cluster" "example" {
    name = "contianer cluster name"
    description = "contianer cluster desc"
}
```

## Testing the Provider

In order to test the provider, you can run `make test`.

*Note:* Make sure `ILLUMIO_PCE_HOST`, `ILLUMIO_API_KEY_USERNAME` and `ILLUMIO_API_KEY_SECRET` variables are set. For windows, copy env.example.bat to env.bat and replace dummy values with credentials and then execute the bat file with command promt.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run. Please read [Running and Writing Acceptance Tests](contributing/running-and-writing-acceptance-tests.md) in the contribution guidelines for more information on usage.

```sh
$ make testacc
```

## Debugging and Troubleshooting

- Set environment variable `TF_LOG` to one of the log levels `TRACE`, `DEBUG`, `INFO`, `WARN` or `ERROR`
- Set environment variable `TF_LOG_PATH` to write logs in a file. e.g. `TF_LOG_PATH=tf.log`

For more details visit - [Terraform Debugging](https://www.terraform.io/docs/internals/debugging.html)

## Documentation

Once done with changes/development of any resource/datasource, document the changes.

- Update parameter/description as per the resource changes.
- If a resource is designed to behave in a specific way that might be strange to the end-user, that should be documented.
- Add an example usage section and add different examples about how to use it.
- Document any environment variable on which resource/datasource is dependent in the example usage section or in the argument description itself.

## JSON Format TF Configuration Files
- The user can also use the JSON format TF configuration files similar to HCL configuration files. 
- Please refer to the github link for conversion of HCL to JSON format TF configuration files and vice versa. Note - In the output JSON file generated after conversion of the corresponding HCL file using the above tool, please convert the  terraform.required_providers[0].illumio-core (highlighted in bold in the below JSON example) from List to Map.
- Below is an example of the HCL and the corresponding JSON file.

### HCL (.tf file)

```hcl
terraform {
  required_providers {
    illumio-core = {
      version = "0.1"
      source  = "illumio.com/labs/illumio-core"
    }
  }
}

provider "illumio-core" {
  request_timeout = 30
  org_id          = 1
}

resource "illumio-core_rule_set" "name" {
  name = "example-json-hcl"
  scopes {
    label {
      href = "/orgs/1/labels/69"
    }
    label {
      href = "/orgs/1/labels/1"
    }
    label_group {
      href = "/orgs/1/sec_policy/draft/label_groups/64126bda-0f9d-47fc-846b-0f9adbe290d6"
    }
  }
}
```

### JSON (.tf.json file)

```JSON
{
    "resource": {
        "illumio-core_rule_set": {
            "name-json": {
                "name": "example-hcl-json",
                "scopes": [
                    {
                        "label": [
                            {
                                "href": "/orgs/1/labels/69"
                            },
                            {
                                "href": "/orgs/1/labels/1"
                            }
                        ],
                        "label_group": [
                            {
                                "href": "/orgs/1/sec_policy/draft/label_groups/64126bda-0f9d-47fc-846b-0f9adbe290d6"
                            }
                        ]
                    }
                ]
            }
        }
    },
    "provider": [
        {
            "illumio-core": [
                {
                    "org_id": 1,
                    "request_timeout": 30
                }
            ]
        }
    ],
    "terraform": [
        {
            "required_providers": [
                {
                    "illumio-core": {
                        "source": "illumio.com/labs/illumio-core",
                        "version": "0.1"
                    }
                }
            ]
        }
    ]
}
```