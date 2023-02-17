// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type LabelTypeDisplayInfo struct {
	Icon              string `json:"icon,omitempty"`
	Initial           string `json:"initial,omitempty"`
	BackgroundColor   string `json:"background_color,omitempty"`
	ForegroundColor   string `json:"foreground_color,omitempty"`
	SortOrdinal       *int   `json:"sort_ordinal,omitempty"`
	DisplayNamePlural string `json:"display_name_plural,omitempty"`
}

// Label represents label resource
type LabelType struct {
	Key                   string `json:"key"`
	DisplayName           string `json:"display_name"`
	*LabelTypeDisplayInfo `json:"display_info,omitempty"`
	ExternalDataSet       string `json:"external_data_set,omitempty"`
	ExternalDataReference string `json:"external_data_reference,omitempty"`
}

// ToMap - Returns map for LabelType model
func (lt *LabelType) ToMap() (map[string]interface{}, error) {
	return toMap(lt)
}
