package client

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Jeffail/gabs/v2"

	"github.com/illumio/terraform-provider-illumio-core/models"
)

// Errors for V2 API client
var (
	ErrPayloadCreation = errors.New("error in creating payload")
)

// Get Performs HTTP GET on endpoint with queryParams
//
// Illumio GET APIs return json response on success, available as *gabs.Container
func (c *V2) Get(endpoint string, queryParams *map[string]string) (*http.Response, *gabs.Container, error) {
	req, err := c.PrepareRequest(http.MethodGet, endpoint, nil, queryParams)
	if err != nil {
		return nil, nil, err
	}
	response, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	container, err := GetContainer(response)
	if err != nil {
		return response, nil, err
	}
	return response, container, nil
}

// AsyncGet Performs HTTP GET on endpoint with queryParams and polls until data is ready
//
// Illumio GET APIs return json response on success, available as *gabs.Container
func (c *V2) AsyncGet(endpoint string, queryParams *map[string]string) (*http.Response, *gabs.Container, error) {
	req, err := c.PrepareRequest(http.MethodGet, endpoint, nil, queryParams)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Prefer", "respond-async")
	response, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	jobHref := response.Header.Get("Location")
	// initial wait
	wait, _ := strconv.Atoi(response.Header.Get("Retry-After"))
	time.Sleep(time.Duration(wait) * time.Second)
	resultHref, err := pollJob(c, jobHref)
	if err != nil {
		return nil, nil, err
	}
	resp, container, err := c.Get(resultHref, nil)
	if err != nil {
		return resp, nil, err
	}
	return response, container, nil
}

func pollJob(c *V2, href string) (string, error) {
	jobStatus := ""
	waittime := c.backoffTime
	resultHref := ""
	for jobStatus != "done" && jobStatus != "failed" {
		resp, data, err := c.Get(href, nil)
		if err != nil {
			return resultHref, err
		}
		if v := data.S("status").Data().(string); v == "done" || v == "failed" {
			jobStatus = data.S("status").Data().(string)
		} else {
			if v := resp.Header.Get("Retry-After"); v != "" {
				waittime, _ = strconv.Atoi(v)
				fmt.Println("Wait Time: ", waittime)
			}
			time.Sleep(time.Duration(waittime) * time.Second)
		}
		if jobStatus == "done" {
			return data.S("result", "href").Data().(string), nil
		} else if jobStatus == "failed" {
			return "", errors.New("async get job failed")
		}
	}
	return resultHref, nil
}

// Create Performs HTTP POST on endpoint with model
//
// Illumio POST APIs return json response on success, available as *gabs.Container
func (c *V2) Create(endpoint string, model models.Model) (*http.Response, *gabs.Container, error) {
	jsonPayload, err := c.PrepareModel(model)
	if err != nil {
		return nil, nil, err
	}

	log.Println("[DEBUG] CREATE Payload: ", jsonPayload.String())

	req, err := c.PrepareRequest(http.MethodPost, endpoint, jsonPayload, nil)
	if err != nil {
		return nil, nil, err
	}

	response, err := c.Do(req)
	if err != nil {
		return response, nil, err
	}

	container, err := GetContainer(response)
	if err != nil {
		return response, nil, err
	}
	return response, container, nil
}

// Update Performs HTTP UPDATE on endpoint
//
// Illumio UPDATE APIs does not return any response
func (c *V2) Update(endpoint string, model models.Model) (*http.Response, error) {
	jsonPayload, err := c.PrepareModel(model)
	if err != nil {
		return nil, err
	}

	log.Println("[DEBUG] UPDATE Payload: ", jsonPayload.String())

	req, err := c.PrepareRequest(http.MethodPut, endpoint, jsonPayload, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// Delete Performs HTTP DELETE on endpoint
//
// Illumio DELETE APIs does not return any response
func (c *V2) Delete(endpoint string) (*http.Response, error) {
	req, err := c.PrepareRequest(http.MethodDelete, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// GetContainer parses HTTP Response and returns *gabs.Container
func GetContainer(resp *http.Response) (*gabs.Container, error) {
	defer resp.Body.Close()
	// Make sure response type is application/json
	if contentType := resp.Header.Get(headerContentType); contentType != mimeTypeApplicationJSON {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			bodyBytes = []byte{}
		}
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		return nil, fmt.Errorf("expected HTTP header 'Content-Type': 'application/json', got '%s'", contentType)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if len(bodyBytes) > 0 {
		return gabs.ParseJSON(bodyBytes)
	}
	return nil, errors.New("no body found")
}

// PrepareModel Creates container from model
func (c *V2) PrepareModel(model models.Model) (*gabs.Container, error) {
	modelMap, err := model.ToMap()
	if err != nil {
		return nil, err
	}

	payload := gabs.New()

	// Checking if modelMap is an array of map
	if v, ok := modelMap["___items___"]; len(modelMap) == 1 && ok {
		payload.Set(v)
	} else {
		for key, value := range modelMap {
			_, err = payload.Set(value, key)
			if err != nil {
				return nil, ErrPayloadCreation
			}
		}
	}

	return payload, nil
}
