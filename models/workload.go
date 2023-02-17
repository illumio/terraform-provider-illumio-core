// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type Workload struct {
	Name                                  string `json:"name,omitempty"`
	Description                           string `json:"description,omitempty"`
	ExternalDataSet                       string `json:"external_data_set,omitempty"`
	ExternalDataReference                 string `json:"external_data_reference,omitempty"`
	Hostname                              string `json:"hostname,omitempty"`
	ServicePrincipalName                  string `json:"service_principal_name,omitempty"`
	PublicIP                              string `json:"public_ip,omitempty"`
	AgentToPceCertificateAuthenticationID string `json:"agent_to_pce_certificate_authentication_id,omitempty"`
	DistinguishedName                     string `json:"distinguished_name,omitempty"`
	ServiceProvider                       string `json:"service_provider,omitempty"`
	DataCenter                            string `json:"data_center,omitempty"`
	DataCenterZone                        string `json:"data_center_zone,omitempty"`
	OsID                                  string `json:"os_id,omitempty"`
	OsDetail                              string `json:"os_detail,omitempty"`
	EnforcementMode                       string `json:"enforcement_mode,omitempty"`

	// boolean false is omitted with `omitempty` when marshalling; use a pointer to fix this behaviour
	Online *bool `json:"online,omitempty"`

	// don't omitempty for lists - an empty array should remove all objects from the workload
	IgnoredInterfaceNames []string                `json:"ignored_interface_names"`
	Labels                []LabelOptionalKeyValue `json:"labels"`
	Interfaces            []WorkloadInterface     `json:"interfaces"`
}

// ToMap - Returns map for Workload model
func (w *Workload) ToMap() (map[string]interface{}, error) {
	return toMap(w)
}
