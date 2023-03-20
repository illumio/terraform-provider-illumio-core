// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type EnforcementBoundaryProviderConsumer struct {
	Actors     string `json:"actors,omitempty"`
	Label      *Href  `json:"label,omitempty"`
	LabelGroup *Href  `json:"label_group,omitempty"`
	IPList     *Href  `json:"ip_list,omitempty"`
}

type EnforcementBoundary struct {
	Name            *string                                 `json:"name,omitempty"`
	Enabled         *bool                                   `json:"enabled,omitempty"`
	Providers       *[]*EnforcementBoundaryProviderConsumer `json:"providers,omitempty"`
	Consumers       *[]*EnforcementBoundaryProviderConsumer `json:"consumers,omitempty"`
	IngressServices *[]map[string]interface{}               `json:"ingress_services,omitempty"`
}

func (eb *EnforcementBoundary) ToMap() (map[string]interface{}, error) {
	return toMap(eb)
}
