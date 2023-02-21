// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type Service struct {
	Name                  string                 `json:"name"`
	Description           string                 `json:"description,omitempty"`
	ProcessName           string                 `json:"process_name,omitempty"`
	ServicePorts          []ServicePort          `json:"service_ports,omitempty"`
	WindowsServices       []WindowsService       `json:"windows_services"`
	WindowsEgressServices []WindowsEgressService `json:"windows_services"`
	ExternalDataSet       string                 `json:"external_data_set,omitempty"`
	ExternalDataReference string                 `json:"external_data_reference,omitempty"`
}

type ServicePort struct {
	Port     *int `json:"port,omitempty"`
	ToPort   *int `json:"to_port,omitempty"`
	Proto    int  `json:"proto"`
	ICMPType *int `json:"icmp_type,omitempty"`
	ICMPCode *int `json:"icmp_code,omitempty"`
}

type WindowsService struct {
	ServiceName string `json:"service_name,omitempty"`
	ProcessName string `json:"process_name,omitempty"`
	Port        *int   `json:"port,omitempty"`
	ToPort      *int   `json:"to_port,omitempty"`
	Proto       *int   `json:"proto,omitempty"`
	ICMPType    *int   `json:"icmp_type,omitempty"`
	ICMPCode    *int   `json:"icmp_code,omitempty"`
}

type WindowsEgressService struct {
	ServiceName string `json:"service_name,omitempty"`
	ProcessName string `json:"process_name,omitempty"`
}

// ToMap - Returns map for Service model
func (s *Service) ToMap() (map[string]interface{}, error) {
	return toMap(s)
}
