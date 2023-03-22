// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

		err := TestAccProvider.Configure(context.Background(), terraform.NewResourceConfigRaw(nil))
		if err != nil {
			t.Fatal(err)
		}
	})
}

func testAccPreCheckSaaSPCE(t *testing.T) {
	conf := TestAccProvider.Meta().(Config)
	illumioClient := conf.IllumioClient
	settingsEndpoint := fmt.Sprintf("/orgs/%d/settings/events", illumioClient.OrgID)

	_, _, err := illumioClient.Get(settingsEndpoint, nil)
	if err != nil {
		t.Skipf("skipping acceptance test: %s", err)
	}
}

func skipIfPCEVersionBelow(v string) func() (bool, error) {
	return func() (bool, error) {
		checkVersion, err := version.NewVersion(v)
		if err != nil {
			return false, err
		}

		conf := TestAccProvider.Meta().(Config)
		illumioClient := conf.IllumioClient

		_, data, err := illumioClient.Get("/product_version", nil)
		if err != nil {
			return false, err
		}

		pceVersion, err := version.NewVersion(data.S("version").Data().(string))
		if err != nil {
			return false, err
		}

		return pceVersion.LessThan(checkVersion), nil
	}
}
