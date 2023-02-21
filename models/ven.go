// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type VEN struct {
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	TargetPCEFqdn string `json:"target_pce_fqdn,omitempty"`
	Status        string `json:"status"`
}

func (ven *VEN) ToMap() (map[string]interface{}, error) {
	return toMap(ven)
}
