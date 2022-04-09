// Copyright 2021 Illumio, Inc. All Rights Reserved.

package client

import (
	"log"
	"os"
	"strconv"
	"testing"

	"golang.org/x/time/rate"
)

var (
	apiHost      = os.Getenv("ILLUMIO_PCE_HOST")
	orgID        int
	apiUsername  = os.Getenv("ILLUMIO_API_KEY_USERNAME")
	apiKeySecret = os.Getenv("ILLUMIO_API_KEY_SECRET")
)

func init() {
	if apiHost == "" {
		log.Fatal("Illumio Host for tests is missing, please set in ILLUMIO_PCE_HOST env var")
	} else if apiUsername == "" || apiKeySecret == "" {
		log.Fatal("Credentials for client tests is missing, " +
			"please set in [ILLUMIO_API_KEY_USERNAME, ILLUMIO_API_KEY_SECRET] env var")
	}

	var err error
	orgID, err = strconv.Atoi(os.Getenv("ILLUMIO_PCE_ORG_ID"))

	if err != nil {
		orgID = 1
	}
}

func GetTestClient() (*V2, error) {
	return NewV2(
		apiHost,
		orgID,
		apiUsername,
		apiKeySecret,
		30,
		rate.NewLimiter(rate.Limit(float64(125)/float64(60)), 1), // 125 requests per 60 seconds
		10,
		3,
		false,
		"",
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
