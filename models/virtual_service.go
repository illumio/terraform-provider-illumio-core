// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type ServiceAdd struct {
	IP          string `json:"ip"`
	Network     *Href  `json:"network"`
	Port        *int   `json:"port"`
	Fqdn        string `json:"fqdn"`
	Description string `json:"description"`
}

type VirtualService struct {
	Name                  string        `json:"name"`
	Description           string        `json:"description,omitempty"`
	Labels                []Href        `json:"labels"`
	ServicePorts          []ServicePort `json:"service_ports"`
	Service               Href          `json:"service"`
	IPOverrides           []string      `json:"ip_overrides"`
	ApplyTo               string        `json:"apply_to,omitempty"`
	ServiceAddresses      []ServiceAdd  `json:"service_addresses"`
	ExternalDataSet       string        `json:"external_data_set,omitempty"`
	ExternalDataReference string        `json:"external_data_reference,omitempty"`
}

// ToMap - Returns map for VirtualService model
func (vs *VirtualService) ToMap() (map[string]interface{}, error) {
	return toMap(vs)
}
