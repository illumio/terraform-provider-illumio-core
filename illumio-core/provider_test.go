package illumiocore

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]*schema.Provider
var testAccProviderFactories map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"illumio-core": testAccProvider,
	}
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"illumio-core": func() (*schema.Provider, error) { return Provider(), nil },
	}
}

func testAccProviderFactoriesInit(provider **schema.Provider, providerName string) map[string]func() (*schema.Provider, error) {
	var factories = make(map[string]func() (*schema.Provider, error))

	p := Provider()

	factories[providerName] = func() (*schema.Provider, error) {
		return p, nil
	}

	if provider != nil {
		*provider = p
	}

	return factories
}

func testAccProviderFactoriesInternal(provider **schema.Provider) map[string]func() (*schema.Provider, error) {
	return testAccProviderFactoriesInit(provider, "illumio-core")
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
	missing := []string{}
	for _, envKey := range []string{"ILLUMIO_PCE_HOST", "ILLUMIO_API_KEY_USERNAME", "ILLUMIO_API_KEY_SECRET"} {
		if v := os.Getenv(envKey); v == "" {
			missing = append(missing, envKey)
		}
	}
	if len(missing) > 0 {
		t.Fatalf("Required environment variables missing: %v", missing)
	}
}

func testAccCheckIllumioGeneralizeDestroy(providerInstance *schema.Provider, resourceType string, checkDeleted bool) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		pConfig := (*providerInstance).Meta().(Config)

		for _, rs := range s.RootModule().Resources {
			if rs.Type == resourceType {
				resp, cont, err := pConfig.IllumioClient.Get(rs.Primary.ID, nil)
				if checkDeleted {
					if !cont.S("deleted").Data().(bool) {
						return fmt.Errorf("%s still exists, respStatus: %v, ResourceID: %v", resourceType, resp.Status, rs.Primary.ID)
					}
				} else {
					if err == nil { // got successful response
						return fmt.Errorf("%s still exists, ResourceID: %v", resourceType, rs.Primary.ID)
					}
					if !strings.Contains(err.Error(), "not-found") {
						return fmt.Errorf("%s still exists, ResourceID: %v", resourceType, rs.Primary.ID)
					}
				}

			}
		}

		return nil
	}
}
