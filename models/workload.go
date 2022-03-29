// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

import (
	"encoding/json"
)

/* Sample
{
  "name": "string",
  "description": "string",
  "external_data_set": null,
  "external_data_reference": null,
  "hostname": "string",
  "service_principal_name": null,
  "public_ip": "string",
  "interfaces": [
    {
      "name": "string",
      "link_state": "up",
      "address": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
      "cidr_block": 0,
      "default_gateway_address": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
      "friendly_name": "string"
    }
  ],
  "service_provider": "string",
  "data_center": "string",
  "data_center_zone": "string",
  "os_id": "string",
  "os_detail": "string",
  "online": true,
  "labels": [
    {
      "href": "string"
    }
  ],
  "agent": {
    "config": {
      "mode": "idle",
      "log_traffic": true
    }
  },
  "enforcement_mode": "idle"
}
*/

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
	Online                                bool   `json:"online,omitempty"`
	Labels                                []Href `json:"labels,omitempty"`
	EnforcementMode                       string `json:"enforcement_mode,omitempty"`

	/* Following code is commented to prevent the race condition
	 * between Workload and Workload Interface Resources. Preserved for future use.
	 * Bug#15
	 */
	// Interfaces                            []WorkloadInterface `json:"interfaces"`
}

// ToMap - Returns map for Workload model
func (w *Workload) ToMap() (map[string]interface{}, error) {
	encodedWorkload, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(encodedWorkload), &result)

	return result, nil
}
