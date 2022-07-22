package illumioapi

import (
	"fmt"
	"strings"
)

// IPRange repsents one of the IP ranges of an IP List.
type IPRange struct {
	Description string `json:"description,omitempty"`
	Exclusion   bool   `json:"exclusion,omitempty"`
	FromIP      string `json:"from_ip,omitempty"`
	ToIP        string `json:"to_ip,omitempty"`
}

// FQDN represents an FQDN in an IPList
type FQDN struct {
	FQDN string `json:"fqdn"`
}

// IPList represents an IP List in the Illumio PCE.
type IPList struct {
	CreatedAt             string      `json:"created_at,omitempty"`
	CreatedBy             *CreatedBy  `json:"created_by,omitempty"`
	DeletedAt             string      `json:"deleted_at,omitempty"`
	DeletedBy             *DeletedBy  `json:"deleted_by,omitempty"`
	Description           string      `json:"description,omitempty"`
	ExternalDataReference string      `json:"external_data_reference,omitempty"`
	ExternalDataSet       string      `json:"external_data_set,omitempty"`
	FQDNs                 *[]*FQDN    `json:"fqdns,omitempty"`
	Href                  string      `json:"href,omitempty"`
	IPRanges              *[]*IPRange `json:"ip_ranges,omitempty"`
	Name                  string      `json:"name,omitempty"`
	UpdatedAt             string      `json:"updated_at,omitempty"`
	UpdatedBy             *UpdatedBy  `json:"updated_by,omitempty"`
	Size                  int         `json:"size,omitempty"`
}

// GetIPLists returns a slice of IP lists from the PCE. pStatus must be "draft" or "active".
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetIPLists(queryParameters map[string]string, pStatus string) (ipLists []IPList, api APIResponse, err error) {
	// Validate pStatus
	pStatus = strings.ToLower(pStatus)
	if pStatus != "active" && pStatus != "draft" {
		return ipLists, api, fmt.Errorf("invalid pStatus")
	}
	api, err = p.GetCollection("/sec_policy/"+pStatus+"/ip_lists", false, queryParameters, &ipLists)
	if len(ipLists) > 500 {
		ipLists = nil
		api, err = p.GetCollection("/sec_policy/"+pStatus+"/ip_lists", true, queryParameters, &ipLists)
	}
	return ipLists, api, err
}

// GetIPListByName queries returns the IP List based on name. A blank IP List is return if no exact match.
func (p *PCE) GetIPListByName(name string, pStatus string) (IPList, APIResponse, error) {
	ipLists, api, err := p.GetIPLists(map[string]string{"name": name}, pStatus)
	if err != nil {
		return IPList{}, api, err
	}

	for _, ipl := range ipLists {
		if ipl.Name == name {
			return ipl, api, nil
		}
	}
	// If there is no match we are going to return an empty IP List
	return IPList{}, api, nil
}

// CreateIPList creates a new IP List in the PCE.
func (p *PCE) CreateIPList(ipList IPList) (createdIPL IPList, api APIResponse, err error) {
	api, err = p.Post("sec_policy/draft/ip_lists", &ipList, &createdIPL)
	return createdIPL, api, err
}

// UpdateIPList updates an existing IP List in the PCE.
// The provided IP List must include an Href.
// Properties that cannot be included in the PUT method will be ignored.
func (p *PCE) UpdateIPList(ipList IPList) (APIResponse, error) {
	ipList.CreatedAt = ""
	ipList.CreatedBy = nil
	ipList.DeletedAt = ""
	ipList.DeletedBy = nil
	ipList.UpdatedAt = ""
	ipList.UpdatedBy = nil

	api, err := p.Put(&ipList)
	return api, err
}
