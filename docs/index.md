---
layout: "illumio-core"
page_title: "Provider: Illumio Core"
sidebar_current: "docs-illumio-core-index"
description: |-
  The Illumio Core provider is used to interact with the resources provided by Illumio Core APIs.
---


Overview
--------------------------------------------------
Terraform provider Illumio Core is a Terraform plugin which can be used to manage the Illumio resources on the Illumio PCE platform with leveraging advantages of Terraform. 
Illumio Core Terraform provider lets users represent the infrastructure as a code and provides a way to enforce state on the infrastructure managed by the Terraform provider. 
Customers can use this provider to integrate the Terraform configuration with the DevOps pipeline to manage the Illumio Resource in a more flexible, consistent and reliable way.

Illumio Core Provider
------------
The Illumio Core provider is used to interact with the resources provided by Illumio Core APIs.
The provider needs to be configured with the pce host and api key details before it can be used.

Example Usage
------------

```hcl

#configure provider with your illumio host and api key details.
provider "illumio-core" {
    pce_host              = "https://pce.my-company.com:8443"
    api_username          = "api_xxxxxx"
    api_secret            = "big-secret"
    org_id                = 1
    request_timeout       = 30
    backoff_time          = 10
    max_retries           = 3
}

resource "illumio-core_label" "env_dev" {
  key     = "env"
  value   = "dev"
}

resource "illumio-core_label_group" "env_lg" {
  key           = "env"
  name          = "Dev Group"
  description   = "Label group for dev environments"
  labels {
    href = illumio-core_label.env_dev.href
  }
}
```

```hcl
# Configure provider with proxy (authentication)
provider "illumio-core" {
    # ... other configuration parameters
    proxy_url = "http:10.0.1.111:3128"
    proxy_creds = "root:password"
}
```

Some of the attributes can be specified via environment variables. Refer to schema for attributes which can be configured via environment variables.


Provisioning
------------
Currently terraform does not support post-processing of resources. To provision changes, provision command can be used.

To run provision, clone the provider repo and follow the commands.

```bash
cd cmd/provision
go build -o provision  # provision.exe for windows
```

Move the provision binary to the root dir of your tf module.
To use provision command, The required environment variables must be set  (`ILLUMIO_API_KEY_SECRET`, `ILLUMIO_API_KEY_USERNAME` and `ILLUMIO_PCE_HOST`).
Note that same environment variables can be used to configure provider.


Now provision command can be used with terraform apply

```bash
terraform apply && provision
```

### Managing Versioned and Non-versioned Resources Together

**Non-Versioned Resource**: Resource which does not require provisioning 

While managing versioned and non-versioned resources together, if you want to destroy a non-versioned resource which is already linked with the versioned resource then you must unlink/delete versioned resource and provision it first. You can perform the following steps for the same:
  -  Unlink the non-versioned resources from versioned resources OR destroy the versioned resources and then provision it using provisioning binary
  - Perform the deletion of non-versioned resources

For example, to delete the label `env_dev` which is referred in `env_lg`

```hcl
# label is non-versioned resource
resource "illumio-core_label" "env_dev" {
  key     = "env"
  value   = "dev"
}

# label_group is versioned resource
resource "illumio-core_label_group" "env_lg" {
  key           = "env"
  name          = "Dev Group"
  description   = "Label group for dev environments"
  labels {
    href = illumio-core_label.env_dev.href
  }
}
```

1. To delete label `env_dev`, we first need to either delete `env_lg` OR unlink `env_dev` from `env_lg`. We chose to unlink here. But we can not delete `env_dev` until we provision `env_lg`.

```hcl
# label is non-versioned resource
resource "illumio-core_label" "env_dev" {
  key     = "env"
  value   = "dev"
}

# label_group is versioned resource
resource "illumio-core_label_group" "env_lg" {
  key           = "env"
  name          = "Dev Group"
  description   = "Label group for dev environments"
  # labels {
  #   href = illumio-core_label.env_dev.href
  # }
}
```

2. Run `terraform apply && provision` on above configuration.

```hcl

# Removed `env_dev`

# label_group is versioned resource
resource "illumio-core_label_group" "env_lg" {
  key           = "env"
  name          = "Dev Group"
  description   = "Label group for dev environments"
}
```

3. Remove `env_dev` and run `terraform apply`. 

 **To identify such dependency, refer to the below list of versioned resources which can have reference to non-versioned resources.**
- label_group: label
- rule_set: label, workload
- virtual_service: label
- firewall_settings: label
- enforcement_boundary: label


## Schema

### Required

- **api_secret** (String, Sensitive) Secret of API Key. This can also be set by environment variable `ILLUMIO_API_KEY_SECRET`
- **api_username** (String) Username of API Key. This can also be set by environment variable `ILLUMIO_API_KEY_USERNAME`
- **pce_host** (String) Host URL of Illumio PCE. This can also be set by environment variable `ILLUMIO_PCE_HOST`

### Optional

- **backoff_time** (Number) Backoff Time (in seconds) on getting 429 (Too Many Requests). Default value: 10. Note: A default rate limit of 125 requests/min is already in place. A jitter of 1-5 seconds will be added to backoff time to randomize backoff.
- **max_retries** (Number) Maximum retries for an API request. Default value: 3
- **org_id** (Number) ID of the Organization. Default value: 1
- **request_timeout** (Number) Timeout for HTTP requests. Default value: 30
- **proxy_url** (String) Proxy Server URL with port number. This can also be set by environment variable `ILLUMIO_PROXY_URL`
- **proxy_creds** (String) Proxy credential in format `username:password`. This can also be set by environment variable `ILLUMIO_PROXY_CREDENTIALS`
- **ca_file** (String) The path to CA certificate file (PEM). In case, certificate is based on legacy CN instead of ASN, set env. variable `GODEBUG=x509ignoreCN=0`. This can also be set by environment variable `ILLUMIO_CA_FILE`
- **insecure** (String) Allow insecure TLS. Only `yes` will mark it insecure. This can also be set by environment variable `ILLUMIO_ALLOW_INSECURE_TLS`
