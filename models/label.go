// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

// Label represents label resource
type Label struct {
	Key                   string `json:"key"`
	Value                 string `json:"value"`
	ExternalDataSet       string `json:"external_data_set,omitempty"`
	ExternalDataReference string `json:"external_data_reference,omitempty"`
}

type LabelOptionalKeyValue struct {
	Href  string `json:"href"`
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

// ToMap - Returns map for Label model
func (l *Label) ToMap() (map[string]interface{}, error) {
	return toMap(l)
}
