package illumiocore

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/illumio/terraform-provider-illumio-core/client"
	"golang.org/x/time/rate"
)

// TestMain parses test flags and invokes sweepers
func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func sharedClient(r string) (*client.V2, error) {
	pceHost := os.Getenv("ILLUMIO_PCE_HOST")
	if pceHost == "" {
		return nil, fmt.Errorf("ILLUMIO_PCE_HOST environment variable must be set.")
	}

	orgID, err := strconv.Atoi(os.Getenv("ILLUMIO_PCE_ORG_ID"))
	if err != nil {
		orgID = 1
	}

	apiKeyUsername := os.Getenv("ILLUMIO_API_KEY_USERNAME")
	if apiKeyUsername == "" {
		return nil, fmt.Errorf("ILLUMIO_API_KEY_USERNAME environment variable must be set.")
	}

	apiKeySecret := os.Getenv("ILLUMIO_API_KEY_SECRET")
	if apiKeySecret == "" {
		return nil, fmt.Errorf("ILLUMIO_API_KEY_SECRET environment variable must be set.")
	}

	insecure := false
	if os.Getenv("ILLUMIO_ALLOW_INSECURE_TLS") == "yes" {
		insecure = true
	}

	return client.NewV2(
		pceHost,
		orgID,
		apiKeyUsername,
		apiKeySecret,
		30,
		rate.NewLimiter(rate.Limit(float64(125)/float64(60)), 1), // limits API calls 125/min
		10,
		3,
		insecure,
		os.Getenv("ILLUMIO_CA_FILE"),
		os.Getenv("ILLUMIO_PROXY_URL"),
		os.Getenv("ILLUMIO_PROXY_CREDENTIALS"),
	)
}

func sweep(objectType, matchKey, prefix, endpoint string) resource.SweeperFunc {
	return func(r string) error {
		illumioClient, err := sharedClient(r)
		if err != nil {
			return fmt.Errorf("Error creating Illumio client for sweepers: %s", err)
		}

		endpoint = fmt.Sprintf(endpoint, illumioClient.OrgID)
		_, data, err := illumioClient.Get(endpoint, &map[string]string{
			matchKey: prefix,
		})

		if err != nil {
			return fmt.Errorf("Error fetching objects for %s sweeper: %s", objectType, err)
		}

		// XXX: currently this will not remove objects that have been provisioned
		for _, o := range data.Children() {
			href := o.S("href").Data().(string)
			_, err := illumioClient.Delete(href)
			if err != nil {
				fmt.Printf("Failed to sweep %s with HREF: %s", objectType, href)
			}
		}

		return nil
	}
}
