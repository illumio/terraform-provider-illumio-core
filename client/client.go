// Copyright 2021 Illumio, Inc. All Rights Reserved.

package client

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	mathrand "math/rand"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/Jeffail/gabs/v2"
	"golang.org/x/time/rate"
)

const (
	headerAccept            = "Accept"
	headerContentType       = "Content-Type"
	mimeTypeApplicationJSON = "application/json"
)

// V2 client for Illumio REST APIs
type V2 struct {
	HostURL      string
	OrgID        int
	apiUsername  string
	apiKeySecret string
	httpClient   *http.Client
	rateLimiter  *rate.Limiter
	backoffTime  int // in seconds // Timeout to wait for requests when server responds with 429
	maxRetries   int
}

// NewV2 Constructor for V2 Client
//
// defaultTimeout (in seconds)
// e.g. NewV2("https://pce.my-company.com:8443", 1, "api_xxxxxx", "big-secret", 30, rate.NewLimiter(rate.Limit(float64(125)/float64(60)), 1), 10, 3, false, "", "")
func NewV2(hostURL string, orgID int, apiUsername string, apiKeySecret string, defaultTimeout int, rateLimiter *rate.Limiter,
	waitTime int, maxRetries int, insecure bool, caFile string, proxyURL string, proxyCreds string) (*V2, error) {
	if !strings.HasPrefix(hostURL, "http") {
		return nil, errors.New("hostURL scheme must be 'http(s)'")
	}
	// hostURL should end with port number and not with trailing slash
	// if trailing slash present, remove it
	hostURL = strings.TrimSuffix(hostURL, "/")
	baseURL := fmt.Sprintf("%s/api/v2", hostURL)
	transport := http.DefaultTransport.(*http.Transport)
	if proxyURL != "" {
		pUrl, err := url.Parse(proxyURL)
		if err != nil {
			return nil, err
		}
		transport.Proxy = http.ProxyURL(pUrl)
		if proxyCreds != "" {
			basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(proxyCreds))
			transport.ProxyConnectHeader = http.Header{}
			transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)
		}

	}
	tlsConfig := &tls.Config{InsecureSkipVerify: insecure}
	if len(caFile) > 0 {
		caCert, err := ioutil.ReadFile(filepath.Clean(caFile))
		if err != nil {
			return nil, err
		}
		caCertPool, _ := x509.SystemCertPool()
		if caCertPool == nil {
			caCertPool = x509.NewCertPool()
		}

		if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
			return nil, fmt.Errorf("could not append CA file to CA Cert pool")
		}

		tlsConfig.RootCAs = caCertPool
	}
	transport.TLSClientConfig = tlsConfig
	httpClient := &http.Client{
		Timeout:   time.Second * time.Duration(defaultTimeout),
		Transport: transport,
	}

	return &V2{
		OrgID:        orgID,
		HostURL:      baseURL,
		apiUsername:  apiUsername,
		apiKeySecret: apiKeySecret,
		httpClient:   httpClient,
		rateLimiter:  rateLimiter,
		maxRetries:   maxRetries,
		backoffTime:  waitTime,
	}, nil
}

// Do function performs HTTP API Call
func (c *V2) Do(req *http.Request) (*http.Response, error) {
	log.Printf("[DEBUG] Beginning DO method %s", req.URL.String())
	var resp *http.Response
	retryCount := 0
	maxRetriesExceeded := false
	for { // retry in case of 429 - Too many requests
		err := c.rateLimiter.Wait(context.Background())
		if err != nil {
			return nil, err
		}
		resp, err = c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode == http.StatusTooManyRequests {
			if retryCount == c.maxRetries {
				maxRetriesExceeded = true
				break
			}
			log.Printf("[DEBUG] Retrying DO method %s", req.URL.String())
			retryCount++
			// Sleep for configured time and retry
			// jitter in 1-5 seconds
			jitter := 1
			// using strong crypto/rand
			n, err := rand.Int(rand.Reader, big.NewInt(5))
			if err != nil {
				// in case of error, fallback to weak math/rand
				jitter = mathrand.Intn(5) + 1
			} else {
				jitter = int(n.Int64()) + 1
			}
			time.Sleep(time.Duration(c.backoffTime+jitter) * time.Second)
		} else {
			// No indication of rate limit from server, we can proceed
			break
		}
	}
	if maxRetriesExceeded {
		return nil, fmt.Errorf("max retries exceeded for %v %v - Error: %v",
			req.Method, req.URL.String(), resp.Status)
	}

	return resp, checkForErrors(resp)
}

// PrepareRequest Creates *http.Request with required headers
func (c *V2) PrepareRequest(method string, endpoint string, body *gabs.Container, queryParams *map[string]string) (*http.Request, error) {
	urlString := fmt.Sprintf("%s%s", c.HostURL, endpoint)
	// validate url
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}
	var req *http.Request
	if method == http.MethodGet || method == http.MethodDelete {
		req, err = http.NewRequest(method, urlString, nil)
	} else if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		req, err = http.NewRequest(method, urlString, bytes.NewBuffer(body.Bytes()))
	} else {
		return nil, errors.New("invalid method")
	}
	if err != nil {
		return nil, err
	}
	// apiUsername, apiKeySecret works like basic auth
	req.SetBasicAuth(c.apiUsername, c.apiKeySecret)

	// content header handling
	if method == http.MethodGet {
		req.Header.Set(headerAccept, mimeTypeApplicationJSON)
	} else {
		req.Header.Set(headerContentType, mimeTypeApplicationJSON)
	}

	if queryParams != nil {
		q := req.URL.Query()
		for qk, qv := range *queryParams {
			q.Add(qk, qv)
		}
		req.URL.RawQuery = q.Encode()
	}
	return req, nil
}

// checkForErrors - Checks for http status code based errors
func checkForErrors(resp *http.Response) error {
	resourcePath := strings.Replace(resp.Request.URL.Path, "/api/v2", "", 1)
	method := resp.Request.Method

	switch resp.StatusCode {
	// success checks
	case http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent:
		return nil

	// client config/parameter/payload related error checks
	case http.StatusNotFound:
		return fmt.Errorf("not-found: %s", resourcePath)
	case http.StatusUnauthorized:
		return errors.New("unauthorized: please check your credentials")
	case http.StatusForbidden:
		return errors.New("forbidden: you do not have permission OR org_id is invalid")
	case http.StatusMethodNotAllowed:
		return fmt.Errorf("not-allowed: %s is not allowed on %s", method, resourcePath)
	case http.StatusNotAcceptable:
		container, err := GetContainer(resp)
		if err != nil {
			return fmt.Errorf("not-acceptable: %v", err)
		}

		return fmt.Errorf("not-acceptable: %v", formatNotAcceptableErrors(container))

	// server side errors
	case http.StatusInternalServerError, http.StatusServiceUnavailable, http.StatusBadGateway:
		return fmt.Errorf("server-error: %s", resp.Status)
	default:
		log.Printf("[DEBUG] Unable to identify HTTP error")
	}

	// fallback success check
	if 200 <= resp.StatusCode && resp.StatusCode < 300 {
		return nil
	}
	// Handling of unknown errors
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	log.Printf("[DEBUG] HTTP REQUEST FAILED [%s] %s %d - response: %s",
		method, resourcePath, resp.StatusCode, bodyString)
	return fmt.Errorf("failed: status code: %d - error: %v", resp.StatusCode, bodyString)
}

func formatNotAcceptableErrors(data *gabs.Container) string {
	var errString string

	if data.Exists("token") {
		errString += fmt.Sprintf("Token: %s: %s", data.S("token").Data(), data.S("message").Data())
	} else {
		for _, child := range data.Children() {
			errString += fmt.Sprintf("\nToken: %s\nMessage: %s\n", child.S("token").Data(), child.S("message").Data())
			for k, v := range child.ChildrenMap() {
				if k != "message" && k != "token" {
					errString += fmt.Sprintf("%s: %s\n", k, v.String())
				}
			}
		}
	}

	if errString == "" {
		return data.String()
	} else {
		return errString
	}
}
