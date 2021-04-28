# Illumio Provider
Provider to maintain Illumio Resource with terraform.

# Development
- Clone this repo to `$GOPATH/src/bitbucket.org/crestdatasys/terraform-provider-illumio`
- Clone illumio-go-client to `$GOPATH/src/github.com/illumio/terraform-provider-illumio-core`
- Install terraform and create path for plugin `%APPDATA%\terraform.d\plugins\illumio.com\labs\illumio\0.1\windows_amd64`
- Copy env.example.bat to env.bat and replace dummy values with credentials

## Commands to build and use plugin in terraform

- `go build -o terraform-provider-illumio.exe`
- `move terraform-provider-illumio.exe %APPDATA%\terraform.d\plugins\illumio.com\labs\illumio\0.1\windows_amd64`


## To run terraform on any examples

- Run `env.bat`
- Run `cd examples/<resource|datasource>/`
- Run `terraform init`
- Run `terraform apply`

github.com/illumio/terraform-provider-illumio-core