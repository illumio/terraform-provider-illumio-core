// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type IPRange struct {
	Description string `json:"description,omitempty"`
	FromIP      string `json:"from_ip"`
	ToIP        string `json:"to_ip,omitempty"`
	Exclusion   *bool  `json:"exclusion,omitempty"`
}

type FQDN struct {
	FQDN        string `json:"fqdn"`
	Description string `json:"description,omitempty"`
}

type IPList struct {
	Name                  string    `json:"name"`
	Description           string    `json:"description,omitempty"`
	ExternalDataSet       string    `json:"external_data_set,omitempty"`
	ExternalDataReference string    `json:"external_data_reference,omitempty"`
	IPRanges              []IPRange `json:"ip_ranges"`
	FQDNs                 []FQDN    `json:"fqdns"`
}

// ToMap - Returns map for IP List model
func (ipl *IPList) ToMap() (map[string]interface{}, error) {
	return toMap(ipl)
}
