// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package models

type VEN struct {
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	TargetPCEFqdn *string `json:"target_pce_fqdn"`
	Status        string  `json:"status"`
}

func (o *VEN) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{
		"name":        o.Name,
		"description": o.Description,
		"status":      o.Status,
	}

	if o.TargetPCEFqdn != nil {
		m["target_pce_fqdn"] = o.TargetPCEFqdn
	}

	return m, nil
}
