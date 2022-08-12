// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

import (
	"encoding/json"
)

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
	Name                  string `json:"name,omitempty"`
	LinkState             string `json:"link_state,omitempty"`
	Address               string `json:"address,omitempty"`
	CidrBlock             int    `json:"cidr_block,omitempty"`
	DefaultGatewayAddress string `json:"default_gateway_address,omitempty"`
	FriendlyName          string `json:"friendly_name,omitempty"`
}

func (w *WorkloadInterface) ToMap() (map[string]interface{}, error) {
	encodedInterface, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(encodedInterface), &result)

	return result, nil
}
