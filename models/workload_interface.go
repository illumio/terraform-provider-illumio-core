// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package models

//Sample
/*
{
  "name": "string",
  "link_state": "up",
  "address": "string",
  "cidr_block": 0,
  "default_gateway_address": "string",
  "friendly_name": "string"
}
*/

type WorkloadInterface struct {
	Name                  string `json:"name"`
	LinkState             string `json:"link_state"`
	Address               string `json:"address"`
	CidrBlock             int    `json:"cidr_block"`
	DefaultGatewayAddress string `json:"default_gateway_address"`
	FriendlyName          string `json:"friendly_name"`
}

func (w *WorkloadInterface) ToMap() (map[string]interface{}, error) {
	workloadInterfaceAttrMap := make(map[string]interface{})
	workloadInterfaceAttrMap["name"] = w.Name
	workloadInterfaceAttrMap["link_state"] = w.LinkState
	if w.Address != "" {
		workloadInterfaceAttrMap["address"] = w.Address
	}
	workloadInterfaceAttrMap["cidr_block"] = w.CidrBlock
	if w.DefaultGatewayAddress != "" {
		workloadInterfaceAttrMap["default_gateway_address"] = w.DefaultGatewayAddress
	}
	workloadInterfaceAttrMap["friendly_name"] = w.FriendlyName

	return workloadInterfaceAttrMap, nil
}
