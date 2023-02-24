// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type ContainerClusterWorkloadProfileLabel struct {
	Key         string `json:"key"`
	Assignment  Href   `json:"assignment"`
	Restriction []Href `json:"restriction"`
}

type ContainerClusterWorkloadProfile struct {
	Name            string                                 `json:"name"`
	Description     string                                 `json:"description,omitempty"`
	AssignLabels    []Href                                 `json:"assign_labels,omitempty"`
	Labels          []ContainerClusterWorkloadProfileLabel `json:"labels,omitempty"`
	EnforcementMode string                                 `json:"enforcement_mode,omitempty"`
	Managed         *bool                                  `json:"managed,omitempty"`
}

func (o *ContainerClusterWorkloadProfileLabel) HasConflicts() bool {
	if (o.Assignment.Href != "") && (len(o.Restriction) > 0) {
		return true
	} else if o.Assignment.Href != "" {
		return false
	} else if len(o.Restriction) > 0 {
		return false
	}
	return true
}

func (ccwp *ContainerClusterWorkloadProfile) ToMap() (map[string]interface{}, error) {
	return toMap(ccwp)
}
