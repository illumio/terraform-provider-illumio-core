// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const ProviderName = "illumio-core"

var TestAccProvider *schema.Provider
var TestAccProviderFactories map[string]func() (*schema.Provider, error)

// make sure we're only configuring the provider once
var testAccProviderConfigure sync.Once

func init() {
	TestAccProvider = Provider()
	TestAccProviderFactories = map[string]func() (*schema.Provider, error){
		ProviderName: func() (*schema.Provider, error) { return Provider(), nil },
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	_ = Provider()
}

// testAccPreCheck To validate required environment values are available
func testAccPreCheck(t *testing.T) {
	testAccProviderConfigure.Do(func() {
		missing := []string{}
		for _, envKey := range []string{"ILLUMIO_PCE_HOST", "ILLUMIO_API_KEY_USERNAME", "ILLUMIO_API_KEY_SECRET"} {
			if v := os.Getenv(envKey); v == "" {
				missing = append(missing, envKey)
			}
		}
		if len(missing) > 0 {
			t.Fatalf("Required environment variables missing: %v", missing)
		}
	})
}
