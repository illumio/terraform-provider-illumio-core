package illumioapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// APIResponse contains the information from the response of the API
type APIResponse struct {
	RespBody   string
	StatusCode int
	Header     http.Header
	Request    *http.Request
	ReqBody    string
	Warnings   []string
}

// Unexported struct for handling the asyncResults
type asyncResults struct {
	Href        string `json:"href"`
	JobType     string `json:"job_type"`
	Description string `json:"description"`
	Result      struct {
		Href string `json:"href"`
	} `json:"result"`
	Status       string `json:"status"`
	RequestedAt  string `json:"requested_at"`
	TerminatedAt string `json:"terminated_at"`
	RequestedBy  struct {
		Href string `json:"href"`
	} `json:"requested_by"`
}

func (p *PCE) httpSetup(action, apiURL string, body []byte, async bool, headers [][2]string) (APIResponse, error) {
	var asyncResults asyncResults

	// Get the base URL
	u, err := url.Parse(apiURL)
	if err != nil {
		return APIResponse{}, err
	}
	baseURL := "https://" + u.Host + "/api/v2"

	// Create body
	httpBody := bytes.NewBuffer(body)

	// Create HTTP client and request
	client := &http.Client{}
	if p.DisableTLSChecking {
		client.Transport = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	}

	req, err := http.NewRequest(action, apiURL, httpBody)
	if err != nil {
		return APIResponse{}, err
	}

	// Set basic authentication and headers
	req.SetBasicAuth(p.User, p.Key)
	for _, h := range headers {
		req.Header.Set(h[0], h[1])
	}
	if async {
		req.Header.Set("Prefer", "respond-async")
	}

	// Make HTTP Request
	resp, err := client.Do(req)
	if err != nil {
		return APIResponse{}, err
	}

	// Process Async requests
	if async {
		for asyncResults.Status != "done" {
			asyncResults, err = p.asyncPoll(baseURL, resp)
			if err != nil {
				return APIResponse{}, err
			}
		}

		finalReq, err := http.NewRequest("GET", baseURL+asyncResults.Result.Href, httpBody)
		if err != nil {
			return APIResponse{}, err
		}

		// Set basic authentication and headers
		finalReq.SetBasicAuth(p.User, p.Key)
		finalReq.Header.Set("Content-Type", "application/json")

		// Make HTTP Request
		resp, err = client.Do(finalReq)
		if err != nil {
			return APIResponse{}, err
		}
	}

	// Process response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return APIResponse{}, err
	}

	// Put relevant response info into struct
	var response APIResponse
	response.RespBody = string(data[:])
	response.StatusCode = resp.StatusCode
	response.Header = resp.Header
	response.Request = resp.Request

	// Check for a 200 response code
	if strconv.Itoa(resp.StatusCode)[0:1] != "2" {
		return response, errors.New("http status code of " + strconv.Itoa(response.StatusCode))
	}

	// Return data and nil error
	return response, nil
}

// asyncPoll is used in async requests to check when the data is ready
func (p *PCE) asyncPoll(baseURL string, origResp *http.Response) (asyncResults asyncResults, err error) {

	// Create HTTP client and request
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	pollReq, err := http.NewRequest("GET", baseURL+origResp.Header.Get("Location"), nil)
	if err != nil {
		return asyncResults, err
	}

	// Set basic authentication and headers
	pollReq.SetBasicAuth(p.User, p.Key)
	pollReq.Header.Set("Content-Type", "application/json")

	// Wait for recommended time from Retry-After
	wait, err := strconv.Atoi(origResp.Header.Get("Retry-After"))
	if err != nil {
		return asyncResults, err
	}
	duration := time.Duration(wait) * time.Second
	time.Sleep(duration)

	// Check if the data is ready
	pollResp, err := client.Do(pollReq)
	if err != nil {
		return asyncResults, err
	}

	// Process Response
	data, err := ioutil.ReadAll(pollResp.Body)
	if err != nil {
		return asyncResults, err
	}

	// Put relevant response info into struct
	json.Unmarshal(data[:], &asyncResults)

	return asyncResults, err
}

// httpReq makes an API call to the PCE with sepcified options
// httpAction must be GET, POST, PUT, or DELETE.
// apiURL is the full endpoint being called.
// PUT and POST methods should have a body that is JSON run through the json.marshal function so it's a []byte.
// async parameter should be set to true for any GET requests returning > 500 items.
func (p *PCE) httpReq(action, apiURL string, body []byte, async bool, jsonContentType bool) (APIResponse, error) {
	// Set headers based on jsonContentType
	headers := [][2]string{{"Content-Type", "application/json"}}
	if !jsonContentType {
		headers = nil
	}

	// Make initial http call
	api, err := p.httpSetup(action, apiURL, body, async, headers)
	retry := 0

	// If the status code is 429, try 3 times
	for api.StatusCode == 429 {
		// If we have already tried 3 times, exit
		if retry > 2 {
			return api, errors.New("received three 429 errors with 30 second pauses between attempts")
		}
		// Increment the retry counter and sleep for 30 seconds
		retry++
		time.Sleep(30 * time.Second)
		// Retry the API call
		api, err = p.httpSetup(action, apiURL, body, async, headers)
	}
	// Return once response code isn't 429
	return api, err
}

// cleanFQDN cleans up the provided PCE FQDN in case of common errors
func (p *PCE) cleanFQDN() string {
	// Remove trailing slash if included
	p.FQDN = strings.TrimSuffix(p.FQDN, "/")
	// Remove HTTPS if included
	p.FQDN = strings.TrimPrefix(p.FQDN, "https://")
	return p.FQDN
}
