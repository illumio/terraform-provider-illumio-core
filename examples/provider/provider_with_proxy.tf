terraform {
  required_providers {
    illumio-core = {
      source  = "illumio/illumio-core"
    }
  }
}

provider "illumio-core" {
  # PCE connection config
  pce_host        = "https://pce.my-company.com:8443"
  api_username    = "api_xxxxxx"
  api_secret      = "xxxxxxxxxx"
  org_id          = 1

  # HTTP request config
  request_timeout = 60
  backoff_time    = 5
  max_retries     = 5

  # TLS config
  ca_file         = "/path/to/devtest-selfsign.pem"
  insecure        = yes

  # Proxy config
  proxy_url       = "http://10.0.1.111:3128"
  proxy_creds     = "username:password"
}
