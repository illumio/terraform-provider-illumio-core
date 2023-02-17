// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

import "reflect"

type SecurityRuleResolveLabelAs struct {
	Providers []string `json:"providers"`
	Consumers []string `json:"consumers"`
}

type SecurityRuleProvider struct {
	Actors         string `json:"actors,omitempty"`
	Label          *Href  `json:"label,omitempty"`
	LabelGroup     *Href  `json:"label_group,omitempty"`
	Workload       *Href  `json:"workload,omitempty"`
	VirtualService *Href  `json:"virtual_service,omitempty"`
	VirtualServer  *Href  `json:"virtual_server,omitempty"`
	IPList         *Href  `json:"ip_list,omitempty"`
}

type SecurityRuleConsumer struct {
	Actors         string `json:"actors,omitempty"`
	Label          *Href  `json:"label,omitempty"`
	LabelGroup     *Href  `json:"label_group,omitempty"`
	Workload       *Href  `json:"workload,omitempty"`
	VirtualService *Href  `json:"virtual_service,omitempty"`
	IPList         *Href  `json:"ip_list,omitempty"`
}

type IngressService struct {
	Href   string `json:"href,omitempty"`
	Port   *int   `json:"port,omitempty"`
	ToPort *int   `json:"to_port,omitempty"`
	Proto  *int   `json:"proto,omitempty"`
}

type SecurityRule struct {
	Enabled               bool                        `json:"enabled"`
	Description           string                      `json:"description,omitempty"`
	IngressServices       []IngressService            `json:"ingress_services"`
	ResolveLabelsAs       *SecurityRuleResolveLabelAs `json:"resolve_labels_as"`
	SecConnect            *bool                       `json:"sec_connect,omitempty"`
	Stateless             *bool                       `json:"stateless,omitempty"`
	MachineAuth           *bool                       `json:"machine_auth,omitempty"`
	Providers             []*SecurityRuleProvider     `json:"providers"`
	Consumers             []*SecurityRuleConsumer     `json:"consumers"`
	UnscopedConsumers     *bool                       `json:"unscoped_consumers,omitempty"`
	ExternalDataSet       string                      `json:"external_data_set,omitempty"`
	ExternalDataReference string                      `json:"external_data_reference,omitempty"`
}

// Checks if Security Rule ResolveLabelAs.providers list containes only "virtual_services"
func (r *SecurityRuleResolveLabelAs) ProviderIsVirtualService() bool {
	if r == nil {
		return false
	}

	if len(r.Providers) == 1 && r.Providers[0] == "virtual_services" {
		return true
	}

	return false
}

// Checks if conflicting values are set in SecurityRule
func (s *SecurityRule) HasConflicts() bool {
	if s == nil {
		return false
	}

	if (*s.SecConnect && *s.Stateless) || (*s.SecConnect && *s.MachineAuth) || (*s.Stateless && *s.MachineAuth) {
		return true
	}
	return false
}

// HasOneActor validates provider and consumer structs to
// ensure only one actor value was provided
func HasOneActor(o interface{}) bool {
	if o == nil {
		return false
	}

	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return false
	}

	var count int

	for i := 0; i < v.NumField(); i++ {
		if f := v.Field(i); f.IsValid() {
			switch f.Kind() {
			case reflect.String:
				if f.String() != "" {
					count++
				}
			case reflect.Pointer:
				if !f.IsNil() {
					count++
				}
			}
		}
	}

	return count == 1
}

// ToMap - Returns map for SecurityRule model
func (sr *SecurityRule) ToMap() (map[string]interface{}, error) {
	return toMap(sr)
}
