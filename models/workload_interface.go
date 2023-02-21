// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type WorkloadInterface struct {
	Name                  string `json:"name,omitempty"`
	LinkState             string `json:"link_state,omitempty"`
	Address               string `json:"address,omitempty"`
	CidrBlock             int    `json:"cidr_block,omitempty"`
	DefaultGatewayAddress string `json:"default_gateway_address,omitempty"`
	FriendlyName          string `json:"friendly_name,omitempty"`
}

func (wi *WorkloadInterface) ToMap() (map[string]interface{}, error) {
	return toMap(wi)
}
