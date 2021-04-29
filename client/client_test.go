package client

import (
	"fmt"
	"log"
	"os"
	"testing"

	"golang.org/x/time/rate"
)

var (
	apiHost      = os.Getenv("TEST_API_HOST")
	apiUsername  = os.Getenv("TEST_API_USERNAME")
	apiKeySecret = os.Getenv("TEST_API_KEY_SECRET")
)

func init() {
	if apiHost == "" {
		log.Fatal("Illumio Host for  tests is missing, please set in TEST_API_HOST env var")
	}
	if apiUsername == "" || apiKeySecret == "" {
		log.Fatal("Credentials for client tests is missing, " +
			"please set in [TEST_API_USERNAME, TEST_API_KEY_SECRET] env var")
	}
}

func GetTestClient() (*V2, error) {
	return NewV2(
		apiHost,
		apiUsername,
		apiKeySecret,
		30,
		rate.NewLimiter(rate.Limit(float64(125)/float64(60)), 1), // 125 requests per 60 seconds
		10,
		3,
		false,
		"",
		"",
	)
}

func TestClient(t *testing.T) {
	_, err := GetTestClient()
	if err != nil {
		t.Error("Error in creating client")
	}
}

func TestAsyncGet(t *testing.T) {
	testClient, err := GetTestClient()
	if err != nil {
		t.Error("Error in creating client")
		return
	}
	_, data, err := testClient.AsyncGet("/orgs/1/labels", nil)
	fmt.Printf("AsyncGet - Data Length: %v\n", len(data.Children()))
	if err != nil {
		t.Error("Error in async get request", err)
	}
}

func TestGet(t *testing.T) {
	testClient, err := GetTestClient()
	if err != nil {
		t.Error("Error in creating client")
		return
	}
	_, _, err = testClient.Get("/health", nil)
	if err != nil {
		t.Error("Error in fetching health")
	}
}
