// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type Vulnerability struct {
	ReferenceID string   "json:\"reference_id\""
	Score       int      "json:\"score\""
	CveIds      []string "json:\"cve_ids\""
	Description string   "json:\"description\""
	Name        string   "json:\"name\""
}

func (v *Vulnerability) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"reference_id": v.ReferenceID,
		"score":        v.Score,
		"cve_ids":      v.CveIds,
		"description":  v.Description,
		"name":         v.Name,
	}
}

type VulnerailityList struct {
	Values []Vulnerability
}

func (o *VulnerailityList) ToMap() (map[string]interface{}, error) {
	vls := []map[string]interface{}{}
	for _, v := range o.Values {
		vls = append(vls, v.ToMap())
	}

	return map[string]interface{}{"___items___": vls}, nil
}
