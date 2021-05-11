# Illumio Provider
Provider to maintain Illumio Resource with terraform.

# Development
- Clone this repo to `$GOPATH/src/github.com/illumio/terraform-provider-illumio-core`
- Install terraform and create path for plugin `%APPDATA%\terraform.d\plugins\illumio.com\labs\illumio-core\0.1\windows_amd64`
- Copy env.example.bat to env.bat and replace dummy values with credentials

## Commands to build and use plugin in terraform

- `go build -o terraform-provider-illumio-core.exe`
- `move terraform-provider-illumio-core.exe %APPDATA%\terraform.d\plugins\illumio.com\labs\illumio-core\0.1\windows_amd64`

## Build binaries for multiple platform

- Run `go get github.com/mitchellh/gox`. This will place gox executable in Go's binary directory. Make sure that is part of the PATH.
- Run `sh scripts/build.sh`. This will create `pkg` directory with binaries for configuration in `build.sh`
- Same can be done for `provision` executable
- Run `cd cmd/provision` and  `sh ../../scripts/build.sh`. This will create `cmd/provision/pkg` directory with binaries for configuration in `build.sh`

## To run terraform on any examples

- Run `env.bat`
- Run `cd examples/<resource|datasource>/`
- Run `terraform init`
- Run `terraform apply`