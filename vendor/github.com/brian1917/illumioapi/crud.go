package illumioapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// GetHref returns the Illumio object with a specific href
func (p *PCE) GetHref(href string, response interface{}) (APIResponse, error) {
	var api APIResponse

	// Build the API URL
	apiURL, err := url.Parse("https://" + p.cleanFQDN() + ":" + strconv.Itoa(p.Port) + "/api/v2" + href)
	if err != nil {
		return api, err
	}

	// Call the API
	api, err = p.httpReq("GET", apiURL.String(), nil, false, true)
	if err != nil {
		return api, err
	}

	err = json.Unmarshal([]byte(api.RespBody), &response)
	return api, err
}

// SyncGet returns a collection of Illumio objects
func (p *PCE) GetCollection(endpoint string, async bool, queryParameters map[string]string, response interface{}) (APIResponse, error) {

	// Build the API URL
	url, err := url.Parse("https://" + p.cleanFQDN() + ":" + strconv.Itoa(p.Port) + "/api/v2/orgs/" + strconv.Itoa(p.Org) + "/" + endpoint)
	if err != nil {
		return APIResponse{}, err
	}

	// Set the query parameters
	for key, value := range queryParameters {
		q := url.Query()
		q.Set(key, value)
		url.RawQuery = q.Encode()
	}

	// Call the API
	api, err := p.httpReq("GET", url.String(), nil, async, true)
	if err != nil {
		return api, fmt.Errorf("get %s - %s", endpoint, err)
	}

	// Unmarshal response to struct and return
	json.Unmarshal([]byte(api.RespBody), &response)

	return api, nil

}

// Post sends a POST request to the PCE
func (p *PCE) Post(endpoint string, object, createdObject interface{}) (api APIResponse, err error) {

	// Build the API URL
	apiURL, err := url.Parse("https://" + p.cleanFQDN() + ":" + strconv.Itoa(p.Port) + "/api/v2/orgs/" + strconv.Itoa(p.Org) + "/" + endpoint)
	if err != nil {
		return api, err
	}

	// Create payload
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return api, err
	}

	// Call the API
	api, err = p.httpReq("POST", apiURL.String(), jsonBytes, false, true)
	if err != nil {
		return api, err
	}

	// Add the request body to the API object
	api.ReqBody = string(jsonBytes)

	// Unmarshal new label
	json.Unmarshal([]byte(api.RespBody), &createdObject)

	return api, nil
}

// Put sends a PUT request to the PCE.
// The object must include an Href field.
func (p *PCE) Put(object interface{}) (api APIResponse, err error) {

	// Build the API URL
	href := reflect.ValueOf(object).Elem().FieldByName("Href").Interface().(string)
	apiURL, err := url.Parse("https://" + p.cleanFQDN() + ":" + strconv.Itoa(p.Port) + "/api/v2" + href)
	if err != nil {
		return api, fmt.Errorf("update label - %s", err)
	}

	// Drop the HREF from the object
	reflect.ValueOf(object).Elem().FieldByName("Href").SetString("")

	// Create the payload
	jsonBytes, err := json.Marshal(object)
	if err != nil {
		return api, err
	}

	// Call the API
	api, err = p.httpReq("PUT", apiURL.String(), jsonBytes, false, true)
	if err != nil {
		return api, err
	}
	// Add the request body to the API object
	api.ReqBody = string(jsonBytes)

	return api, nil
}

// DeleteHref deletes an existing object in the PCE based on its href.
func (p *PCE) DeleteHref(href string) (APIResponse, error) {
	var api APIResponse

	// Build the API URL
	apiURL, err := url.Parse("https://" + p.cleanFQDN() + ":" + strconv.Itoa(p.Port) + "/api/v2" + href)
	if err != nil {
		return api, fmt.Errorf("delete href - %s", err)
	}

	// Call the API
	api, err = p.httpReq("DELETE", apiURL.String(), nil, false, true)
	if err != nil {
		return api, fmt.Errorf("delete href - %s", err)
	}

	return api, nil
}
