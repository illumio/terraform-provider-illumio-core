// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type ServiceAddress struct {
	IP          string `json:"ip,omitempty"`
	Network     *Href  `json:"network,omitempty"`
	Port        *int   `json:"port,omitempty"`
	Fqdn        string `json:"fqdn,omitempty"`
	Description string `json:"description,omitempty"`
}

type VirtualService struct {
	Name                  string           `json:"name"`
	Description           string           `json:"description,omitempty"`
	Labels                []Href           `json:"labels"`
	ServicePorts          []ServicePort    `json:"service_ports,omitempty"`
	Service               *Href            `json:"service,omitempty"`
	IPOverrides           []string         `json:"ip_overrides"`
	ApplyTo               string           `json:"apply_to,omitempty"`
	ServiceAddresses      []ServiceAddress `json:"service_addresses"`
	ExternalDataSet       string           `json:"external_data_set,omitempty"`
	ExternalDataReference string           `json:"external_data_reference,omitempty"`
}

// ToMap - Returns map for VirtualService model
func (vs *VirtualService) ToMap() (map[string]interface{}, error) {
	return toMap(vs)
}
