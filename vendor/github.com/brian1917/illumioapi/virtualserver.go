package illumioapi

import (
	"fmt"
	"strings"
)

// VirtualServer represents a VirtualServer in the PCE
type VirtualServer struct {
	Href                    string                   `json:"href,omitempty"`
	CreatedAt               string                   `json:"created_at,omitempty"`
	UpdatedAt               string                   `json:"updated_at,omitempty"`
	DeletedAt               string                   `json:"deleted_at,omitempty"`
	CreatedBy               *CreatedBy               `json:"created_by,omitempty"`
	UpdatedBy               *UpdatedBy               `json:"updated_by,omitempty"`
	DeletedBy               *DeletedBy               `json:"deleted_by,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	Description             string                   `json:"description,omitempty"`
	DiscoveredVirtualServer *DiscoveredVirtualServer `json:"discovered_virtual_server,omitempty"`
	DvsName                 string                   `json:"dvs_name,omitempty"`
	DvsIdentifier           string                   `json:"dvs_identifier,omitempty"`
	Labels                  []*Label                 `json:"labels,omitempty"`
	Service                 *Service                 `json:"service,omitempty"`
	Providers               []interface{}            `json:"providers,omitempty"`
	Mode                    string                   `json:"mode,omitempty"`
}

// DiscoveredVirtualServer is part of a Virtual Server
type DiscoveredVirtualServer struct {
	Href string `json:"href"`
}

// GetVirtualServers returns a slice of IP lists from the PCE. pStatus must be "draft" or "active".
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetVirtualServers(queryParameters map[string]string, pStatus string) (virtualServers []VirtualServer, api APIResponse, err error) {
	// Validate pStatus
	pStatus = strings.ToLower(pStatus)
	if pStatus != "active" && pStatus != "draft" {
		return virtualServers, api, fmt.Errorf("invalid pStatus")
	}
	api, err = p.GetCollection("/sec_policy/"+pStatus+"/virtual_servers", false, queryParameters, &virtualServers)
	if len(virtualServers) > 500 {
		virtualServers = nil
		api, err = p.GetCollection("/sec_policy/"+pStatus+"/virtual_servers", true, queryParameters, &virtualServers)
	}
	return virtualServers, api, err
}
