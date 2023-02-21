// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type ServiceBindingPortOverrides struct {
	Port      *int `json:"port,omitempty"`
	Proto     *int `json:"proto,omitempty"`
	NewPort   int  `json:"new_port"`
	NewToPort *int `json:"new_to_port,omitempty"`
}

type ServiceBinding struct {
	VirtualService        Href                          `json:"virtual_service"`
	Workload              *Href                         `json:"workload,omitempty"`
	PortOverrides         []ServiceBindingPortOverrides `json:"port_overrides,omitempty"`
	ExternalDataReference string                        `json:"external_data_reference,omitempty"`
	ExternalDataSet       string                        `json:"external_data_set,omitempty"`
	ContainerWorkload     *Href                         `json:"container_workload,omitempty"`
}

func (sb *ServiceBinding) ToMap() (map[string]interface{}, error) {
	return toMap(sb)
}
