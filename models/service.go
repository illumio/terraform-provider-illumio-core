// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package models

/* Sample
{
  "name": "string",
  "description": "string",
  "process_name": "string",
  "service_ports": [
    {
      "port": 0,
      "to_port": 0,
      "proto": 0,
      "icmp_type": 0,
      "icmp_code": 0
    }
  ],
  "windows_services": [
    {
      "service_name": "string",
      "process_name": "string",
      "port": 0,
      "to_port": 0,
      "proto": 0,
      "icmp_type": 0,
      "icmp_code": 0
    }
  ],
  "external_data_set": "string",
  "external_data_reference": "string"
}
*/

type Service struct {
	Name                  string                   `json:"name"`
	Description           string                   `json:"description"`
	ProcessName           string                   `json:"process_name"`
	ServicePorts          []map[string]interface{} `json:"service_ports"`
	WindowsServices       []map[string]interface{} `json:"windows_services"`
	ExternalDataSet       string                   `json:"external_data_set"`
	ExternalDataReference string                   `json:"external_data_reference"`
}

// ToMap - Returns map for Service model
func (s *Service) ToMap() (map[string]interface{}, error) {
	sAttrMap := make(map[string]interface{})

	if s.Name != "" {
		sAttrMap["name"] = s.Name
	}

	sAttrMap["external_data_set"] = nil
	if s.ExternalDataSet != "" {
		sAttrMap["external_data_set"] = s.ExternalDataSet
	}

	sAttrMap["external_data_reference"] = nil
	if s.ExternalDataReference != "" {
		sAttrMap["external_data_reference"] = s.ExternalDataReference
	}

	sAttrMap["description"] = s.Description

	if s.ProcessName != "" {
		sAttrMap["process_name"] = s.ProcessName
	}

	if s.ServicePorts != nil {
		sAttrMap["service_ports"] = s.ServicePorts
	}

	if s.WindowsServices != nil {
		sAttrMap["windows_services"] = s.WindowsServices
	}

	return sAttrMap, nil
}
