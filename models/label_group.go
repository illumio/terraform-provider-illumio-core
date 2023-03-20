// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

// LabelGroup represents label group resource
type LabelGroup struct {
	Name                  *string `json:"name,omitempty"`
	Description           *string `json:"description,omitempty"`
	Key                   *string `json:"key,omitempty"`
	Labels                *[]Href `json:"labels,omitempty"`
	SubGroups             *[]Href `json:"sub_groups,omitempty"`
	ExternalDataSet       string  `json:"external_data_set,omitempty"`
	ExternalDataReference string  `json:"external_data_reference,omitempty"`
}

type LabelGroupOptionalKeyValue struct {
	Href string  `json:"href"`
	Key  *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
}

// ToMap - Returns map for LabelGroup model
func (lg *LabelGroup) ToMap() (map[string]interface{}, error) {
	return toMap(lg)
}
