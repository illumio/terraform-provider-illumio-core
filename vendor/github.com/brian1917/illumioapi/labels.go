package illumioapi

import (
	"errors"
	"strings"
)

// A Label represents an Illumio Label.
type Label struct {
	CreatedAt             string      `json:"created_at,omitempty"`
	CreatedBy             *CreatedBy  `json:"created_by,omitempty"`
	Deleted               bool        `json:"deleted,omitempty"`
	ExternalDataReference string      `json:"external_data_reference,omitempty"`
	ExternalDataSet       string      `json:"external_data_set,omitempty"`
	Href                  string      `json:"href,omitempty"`
	Key                   string      `json:"key,omitempty"`
	UpdatedAt             string      `json:"updated_at,omitempty"`
	UpdatedBy             *UpdatedBy  `json:"updated_by,omitempty"`
	Value                 string      `json:"value,omitempty"`
	LabelUsage            *LabelUsage `json:"usage,omitempty"`
}

// CreatedBy represents the CreatedBy property of an object
type CreatedBy struct {
	Href string `json:"href"`
}

// UpdatedBy represents the UpdatedBy property of an object
type UpdatedBy struct {
	Href string `json:"href"`
}

type LabelUsage struct {
	VirtualServer                     bool `json:"virtual_server"`
	LabelGroup                        bool `json:"label_group"`
	Ruleset                           bool `json:"ruleset"`
	StaticPolicyScopes                bool `json:"static_policy_scopes"`
	PairingProfile                    bool `json:"pairing_profile"`
	Permission                        bool `json:"permission"`
	Workload                          bool `json:"workload"`
	ContainerWorkload                 bool `json:"container_workload"`
	FirewallCoexistenceScope          bool `json:"firewall_coexistence_scope"`
	ContainersInheritHostPolicyScopes bool `json:"containers_inherit_host_policy_scopes"`
	ContainerWorkloadProfile          bool `json:"container_workload_profile"`
	BlockedConnectionRejectScope      bool `json:"blocked_connection_reject_scope"`
	EnforcementBoundary               bool `json:"enforcement_boundary"`
	LoopbackInterfacesInPolicyScopes  bool `json:"loopback_interfaces_in_policy_scopes"`
	VirtualService                    bool `json:"virtual_service"`
}

// GetLabels returns a slice of labels from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetLabels(queryParameters map[string]string) (labels []Label, api APIResponse, err error) {
	api, err = p.GetCollection("labels", false, queryParameters, &labels)
	if len(labels) >= 500 {
		labels = nil
		api, err = p.GetCollection("labels", true, queryParameters, &labels)
	}
	return labels, api, err
}

// GetLabelByKeyValue finds a label based on the key and value. A blank label is return if no exact match.
func (p *PCE) GetLabelByKeyValue(key, value string) (Label, APIResponse, error) {
	labels, api, err := p.GetLabels(map[string]string{"key": key, "value": value})
	for _, label := range labels {
		if label.Value == value {
			return label, api, err
		}
	}
	return Label{}, api, nil
}

// GetLabelbyHref returns a label based on the provided HREF.
func (p *PCE) GetLabelByHref(href string) (Label, APIResponse, error) {
	var label Label
	api, err := p.GetHref(href, &label)
	return label, api, err
}

// CreateLabel creates a new Label in the PCE.
func (p *PCE) CreateLabel(label Label) (createdLabel Label, api APIResponse, err error) {
	// Check to make sure the label key is valid
	label.Key = strings.ToLower(label.Key)
	if label.Key != "app" && label.Key != "env" && label.Key != "role" && label.Key != "loc" {
		return Label{}, APIResponse{}, errors.New("label key is not app, env, role, or loc")
	}
	api, err = p.Post("labels", &label, &createdLabel)
	return createdLabel, api, err
}

// UpdateLabel updates an existing label in the PCE.
// The provided label must include an Href.
// Properties that cannot be included in the PUT method will be ignored.
func (p *PCE) UpdateLabel(label Label) (APIResponse, error) {
	// Create a new label with just the fields that should be updated and the href
	l := Label{
		Href:                  label.Href,
		Value:                 label.Value,
		ExternalDataReference: label.ExternalDataReference,
		ExternalDataSet:       label.ExternalDataSet,
	}
	api, err := p.Put(&l)
	return api, err
}

// LabelsToRuleStructure takes a slice of labels and returns a slice of slices for how the labels would be organized as read by the PCE rule processing.
// For example {"A-ERP", "A-CRM", "E-PROD"} will return [{"A-ERP, E-PROD"}. {"A-CRM", "E-PROD"}]
func LabelsToRuleStructure(labels []Label) ([][]Label, error) {
	// Create 4 slices: roleLabels, appLabels, envLabels, locLabels and put each label in the correct one
	var roleLabels, appLabels, envLabels, locLabels []Label
	for _, l := range labels {
		switch l.Key {
		case "role":
			roleLabels = append(roleLabels, l)
		case "app":
			appLabels = append(appLabels, l)
		case "env":
			envLabels = append(envLabels, l)
		case "loc":
			locLabels = append(locLabels, l)
		default:
			return nil, errors.New("label key is not role, app, env, or loc")
		}
	}
	// If any of the label slices are empty, put a filler that we will ignore in with blank key and value
	targets := []*[]Label{&roleLabels, &appLabels, &envLabels, &locLabels}
	for _, t := range targets {
		if len(*t) == 0 {
			*t = append(*t, Label{Key: "", Value: ""})
		}
	}
	// Produce an array for every combination that is needed.
	var results [][]Label
	for _, r := range roleLabels {
		for _, a := range appLabels {
			for _, e := range envLabels {
				for _, l := range locLabels {
					n := []Label{}
					if r.Value != "" {
						n = append(n, r)
					}
					if a.Value != "" {
						n = append(n, a)
					}
					if e.Value != "" {
						n = append(n, e)
					}
					if l.Value != "" {
						n = append(n, l)
					}
					results = append(results, n)
				}
			}
		}
	}

	return results, nil

}
