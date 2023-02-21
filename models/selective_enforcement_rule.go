// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type SelectiveEnforcementRulesScope struct {
	Label      *Href `json:"label"`
	LabelGroup *Href `json:"label_group"`
}

type SelectiveEnforcementRules struct {
	Name             string                            `json:"name"`
	Scope            []*SelectiveEnforcementRulesScope `json:"scope"`
	EnforcedServices []Href                            `json:"enforced_services"`
}

// ToMap - Returns map for SelectiveEnforcementRules model
func (ser *SelectiveEnforcementRules) ToMap() (map[string]interface{}, error) {
	return toMap(ser)
}
