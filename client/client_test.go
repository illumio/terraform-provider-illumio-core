// Copyright 2021 Illumio, Inc. All Rights Reserved.

package client

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
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

func TestExpectedHTTPErrors(t *testing.T) {
	endpoint := "/health"

	tests := []struct {
		httpStatus    int
		expectedError string
	}{
		{http.StatusBadRequest, "failed: status code: 400"},
		{http.StatusUnauthorized, "unauthorized"},
		{http.StatusForbidden, "forbidden"},
		{http.StatusTooManyRequests, "max retries exceeded"},
		{http.StatusBadGateway, "server-error"},
		{http.StatusServiceUnavailable, "server-error"},
		{http.StatusInternalServerError, "server-error"},
		{http.StatusPermanentRedirect, "failed: status code: 308"},
	}

	for _, tt := range tests {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(tt.httpStatus)
		}))
		defer server.Close()

		apiClient, err := NewV2(
			server.URL, 1, "api_key", "api_secret",
			30, rate.NewLimiter(rate.Limit(float64(125)/float64(60)), 1),
			10, 0, false, "", "", "",
		)
		if err != nil {
			t.Fatalf("V2.New() => failed to create client: %v", err)
		}
		_, _, err = apiClient.Get(endpoint, nil)

		if err == nil {
			t.Errorf("V2.Get(%q, nil) => returned error was nil, want %q", endpoint, tt.expectedError)
		}

		if !strings.Contains(err.Error(), tt.expectedError) {
			t.Errorf("V2.Get(%q, nil) => error was %q, want %q", endpoint, err.Error(), tt.expectedError)
		}
	}
}

func TestDelete(t *testing.T) {
	href := "/orgs/1/labels/1"

	tests := []struct {
		httpStatus      int
		responseContent string
	}{
		{http.StatusNoContent, ""},
		{http.StatusNotAcceptable, `{"token": "label_already_deleted", "message": "Cannot delete a label that has already been deleted"}`},
	}

	for _, tt := range tests {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(tt.responseContent))
			w.WriteHeader(tt.httpStatus)
		}))
		defer server.Close()

		apiClient, err := NewV2(
			server.URL, 1, "api_key", "api_secret",
			30, rate.NewLimiter(rate.Limit(float64(125)/float64(60)), 1),
			10, 0, false, "", "", "",
		)
		if err != nil {
			t.Fatalf("V2.New() => failed to create client: %v", err)
		}
		_, err = apiClient.Delete(href)

		if err != nil {
			t.Errorf("V2.Delete(%q) => returned unexpected error %v", href, err)
		}
	}
}
