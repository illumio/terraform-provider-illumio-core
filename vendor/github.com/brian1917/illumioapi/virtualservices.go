package illumioapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

// A VirtualService represents a Virtual Service in the Illumio PCE
type VirtualService struct {
	ApplyTo               string              `json:"apply_to,omitempty"`
	CreatedAt             string              `json:"created_at,omitempty"`
	CreatedBy             *CreatedBy          `json:"created_by,omitempty"`
	DeletedAt             string              `json:"deleted_at,omitempty"`
	DeletedBy             *DeletedBy          `json:"deleted_by,omitempty"`
	Description           string              `json:"description,omitempty"`
	ExternalDataReference string              `json:"external_data_reference,omitempty"`
	ExternalDataSet       string              `json:"external_data_set,omitempty"`
	Href                  string              `json:"href,omitempty"`
	IPOverrides           []string            `json:"ip_overrides,omitempty"`
	Labels                []*Label            `json:"labels,omitempty"`
	Name                  string              `json:"name,omitempty"`
	PceFqdn               string              `json:"pce_fqdn,omitempty"`
	Service               *Service            `json:"service,omitempty"`
	ServiceAddresses      []*ServiceAddresses `json:"service_addresses,omitempty"`
	ServicePorts          []*ServicePort      `json:"service_ports,omitempty"`
	UpdateType            string              `json:"update_type,omitempty"`
	UpdatedAt             string              `json:"updated_at,omitempty"`
	UpdatedBy             *UpdatedBy          `json:"updated_by,omitempty"`
}

// ServiceAddresses are FQDNs for Virtual Services
type ServiceAddresses struct {
	IP          string   `json:"ip,omitempty"`
	Network     *Network `json:"network,omitempty"`
	Fqdn        string   `json:"fqdn,omitempty"`
	Description string   `json:"description,omitempty"`
}

// Network represents a network in the PCE
type Network struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetVirtualServices returns a slice of IP lists from the PCE. pStatus must be "draft" or "active".
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetVirtualServices(queryParameters map[string]string, pStatus string) (virtualServices []VirtualService, api APIResponse, err error) {
	// Validate pStatus
	pStatus = strings.ToLower(pStatus)
	if pStatus != "active" && pStatus != "draft" {
		return virtualServices, api, fmt.Errorf("invalid pStatus")
	}
	api, err = p.GetCollection("/sec_policy/"+pStatus+"/virtual_services", false, queryParameters, &virtualServices)
	if len(virtualServices) > 500 {
		virtualServices = nil
		api, err = p.GetCollection("/sec_policy/"+pStatus+"/virtual_services", true, queryParameters, &virtualServices)
	}
	return virtualServices, api, err

}

// GetVirtualServiceByName returns the virtual service based on name. A blank virtual service is return if no exact match.
func (p *PCE) GetVirtualServiceByName(name string, pStatus string) (VirtualService, APIResponse, error) {
	vsMatches, api, err := p.GetAllVirtualServices(map[string]string{"name": name}, pStatus)
	if err != nil {
		return VirtualService{}, api, err
	}
	for _, vs := range vsMatches {
		if vs.Name == name {
			return vs, api, nil
		}
	}
	return VirtualService{}, api, nil

}

// GetVirtualServiceByHref returns the virtualservice with a specific href
func (p *PCE) GetVirtualServiceByHref(href string) (virtualService VirtualService, api APIResponse, err error) {
	api, err = p.GetHref(href, &virtualService)
	return virtualService, api, err
}

// CreateVirtualService creates a new virtual service in the Illumio PCE.
func (p *PCE) CreateVirtualService(virtualService VirtualService) (createdVirtualService VirtualService, api APIResponse, err error) {
	virtualService.Sanitize()
	virtualService.Href = ""
	api, err = p.Post("sec_policy/draft/virtual_services", &virtualService, &createdVirtualService)
	return createdVirtualService, api, err
}

// UpdateVirtualService updates an existing virtual service in the PCE.
// The provided virtual service must include an Href.
// Properties that cannot be included in the PUT method will be ignored.
func (p *PCE) UpdateVirtualService(virtualService VirtualService) (APIResponse, error) {
	virtualService.Sanitize()
	return p.Put(&virtualService)
}

// BulkVS takes a bulk action on an array of workloads.
// Method must be create, update, or delete
func (p *PCE) BulkVS(virtualServices []VirtualService, method string) ([]APIResponse, error) {
	var apiResps []APIResponse
	var err error

	// Check on method
	method = strings.ToLower(method)
	if method != "create" && method != "update" && method != "delete" {
		return apiResps, errors.New("bulk vs error - method must be create, update, or delete")
	}

	// Sanitize update
	if method == "update" {
		sanitizedVSs := []VirtualService{}
		for _, vs := range virtualServices {
			vs.Sanitize()
			sanitizedVSs = append(sanitizedVSs, vs)
		}
		virtualServices = sanitizedVSs
	}

	// Build the API URL
	apiURL, err := url.Parse("https://" + p.cleanFQDN() + ":" + strconv.Itoa(p.Port) + "/api/v2/orgs/" + strconv.Itoa(p.Org) + "/sec_policy/draft/virtual_services/bulk_" + method)
	if err != nil {
		return apiResps, fmt.Errorf("bulk vs error - %s", err)
	}

	// If the method is delete, we can only send Hrefs
	if method == "delete" {
		hrefVSs := []VirtualService{}
		for _, vs := range virtualServices {
			hrefVSs = append(hrefVSs, VirtualService{Href: vs.Href})
		}
		// Re-assign virtual services to just the HREF
		virtualServices = hrefVSs
	}

	// Figure out how many API calls we need to make
	numAPICalls := int(math.Ceil(float64(len(virtualServices)) / 1000))

	// Build the array to be passed to the API
	apiArrays := [][]VirtualService{}
	for i := 0; i < numAPICalls; i++ {
		// Get 1,000 elements if this is not the last array
		if (i + 1) != numAPICalls {
			apiArrays = append(apiArrays, virtualServices[i*1000:(1+i)*1000])
			// Get the rest on the last array
		} else {
			apiArrays = append(apiArrays, virtualServices[i*1000:])
		}
	}

	// Call the API for each array
	for _, apiArray := range apiArrays {
		vsJSON, err := json.Marshal(apiArray)
		if err != nil {
			return apiResps, fmt.Errorf("bulk vs error - %s", err)
		}

		api, err := p.httpReq("PUT", apiURL.String(), vsJSON, false, true)
		api.ReqBody = string(vsJSON)

		apiResps = append(apiResps, api)

		if err != nil {
			return apiResps, fmt.Errorf("bulk vs error - %s", err)
		}

	}

	return apiResps, nil
}

// Sanitize removes fields for an update
func (vs *VirtualService) Sanitize() {
	// Remove fields for marshaling an update
	vs.CreatedAt = ""
	vs.CreatedBy = nil
	vs.DeletedAt = ""
	vs.DeletedBy = nil
	vs.UpdateType = ""
	vs.UpdatedAt = ""
	vs.UpdatedBy = nil
	vs.PceFqdn = ""
	// Make the service HREF only
	if vs.Service != nil {
		vs.Service = &Service{Href: vs.Service.Href}
	}
	// Adjust the labels and services to be HREF only
	hrefOnlyLabels := []*Label{}
	for _, l := range vs.Labels {
		hrefOnlyLabels = append(hrefOnlyLabels, &Label{Href: l.Href})
	}
	vs.Labels = hrefOnlyLabels
}

// SetActive changes the HREF of the Virtual Service Object to Active
func (vs *VirtualService) SetActive() VirtualService {
	vs.Href = strings.ReplaceAll(vs.Href, "draft", "active")
	return *vs
}
