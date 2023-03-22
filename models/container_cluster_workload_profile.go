// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type ContainerClusterWorkloadProfileLabel struct {
	Key         *string `json:"key,omitempty"`
	Assignment  *Href   `json:"assignment,omitempty"`
	Restriction *[]Href `json:"restriction,omitempty"`
}

type ContainerClusterWorkloadProfile struct {
	Name            *string                                 `json:"name,omitempty"`
	Description     *string                                 `json:"description,omitempty"`
	AssignLabels    *[]Href                                 `json:"assign_labels,omitempty"`
	Labels          *[]ContainerClusterWorkloadProfileLabel `json:"labels,omitempty"`
	EnforcementMode *string                                 `json:"enforcement_mode,omitempty"`
	Managed         *bool                                   `json:"managed,omitempty"`
}

func (o *ContainerClusterWorkloadProfileLabel) HasConflicts() bool {
	var restrictionLen int
	if o.Restriction != nil {
		restrictionLen = len(*o.Restriction)
	}

	if (o.Assignment.Href != "") && (restrictionLen > 0) {
		return true
	} else if o.Assignment.Href != "" {
		return false
	} else if restrictionLen > 0 {
		return false
	}
	return true
}

func (ccwp *ContainerClusterWorkloadProfile) ToMap() (map[string]interface{}, error) {
	return toMap(ccwp)
}
