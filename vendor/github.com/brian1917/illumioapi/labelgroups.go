package illumioapi

import (
	"fmt"
	"strings"
)

// LabelGroup represents a Label Group in the Illumio PCE
type LabelGroup struct {
	Description           string       `json:"description,omitempty"`
	ExternalDataReference string       `json:"external_data_reference,omitempty"`
	ExternalDataSet       string       `json:"external_data_set,omitempty"`
	Href                  string       `json:"href,omitempty"`
	Key                   string       `json:"key,omitempty"`
	Labels                []*Label     `json:"labels,omitempty"`
	Name                  string       `json:"name,omitempty"`
	SubGroups             []*SubGroups `json:"sub_groups,omitempty"`
	Usage                 *Usage       `json:"usage,omitempty"`
}

// SubGroups represent SubGroups for Label Groups
type SubGroups struct {
	Href string `json:"href"`
	Name string `json:"name,omitempty"`
}

// Usage covers how a LabelGroup is used in the PCE
type Usage struct {
	LabelGroup         bool `json:"label_group"`
	Rule               bool `json:"rule"`
	Ruleset            bool `json:"ruleset"`
	StaticPolicyScopes bool `json:"static_policy_scopes,omitempty"`
}

// GetLabelGroups returns a slice of label groups from the PCE. pStatus must be "draft" or "active"
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetLabelGroups(queryParameters map[string]string, pStatus string) (labelGroups []LabelGroup, api APIResponse, err error) {

	// Validate pStatus
	pStatus = strings.ToLower(pStatus)
	if pStatus != "active" && pStatus != "draft" {
		return labelGroups, api, fmt.Errorf("invalid pStatus")
	}
	api, err = p.GetCollection("/sec_policy/"+pStatus+"/label_groups", false, queryParameters, &labelGroups)
	if len(labelGroups) > 500 {
		labelGroups = nil
		api, err = p.GetCollection("/sec_policy/"+pStatus+"/label_groups", true, queryParameters, &labelGroups)
	}
	return labelGroups, api, err
}

// CreateLabelGroup creates a new label group in the PCE.
func (p *PCE) CreateLabelGroup(labelGroup LabelGroup) (createdLabelGroup LabelGroup, api APIResponse, err error) {
	api, err = p.Post("sec_policy/draft/label_groups", &labelGroup, &createdLabelGroup)
	return createdLabelGroup, api, err
}

// UpdateLabelGroup updates an existing label group in the PCE.
// The provided label group must include an Href.
// Properties that cannot be included in the PUT method will be ignored.
func (p *PCE) UpdateLabelGroup(labelGroup LabelGroup) (APIResponse, error) {
	labelGroup.Usage = nil
	labelGroup.Key = ""

	api, err := p.Put(&labelGroup)
	return api, err
}

// ExpandLabelGroup returns a string of label hrefs in a label group
// Every subgroup (and nested subgroup) is expanded
func (p *PCE) ExpandLabelGroup(href string) (labelHrefs []string) {

	// Get the labels from the original label group
	a, _ := p.expandLabelGroup(href)
	labelHrefs = append(labelHrefs, a...)

	// Iterate through the subgroups of the original label group
	for _, sg := range p.LabelGroups[href].SubGroups {
		// Get the labels in that subgroup and the additional subgroups
		l, moreSGs := p.expandLabelGroup(sg.Href)
		// Append the labels
		labelHrefs = append(labelHrefs, l...)
		// While there are more subgroups, continue expanding them
		for len(moreSGs) > 0 {
			for _, newSG := range moreSGs {
				l, moreSGs = p.expandLabelGroup(newSG)
				// Append the labels
				labelHrefs = append(labelHrefs, l...)
			}
		}
	}

	// De-dupe and return
	labelGroupMap := make(map[string]bool)
	for _, l := range labelHrefs {
		labelGroupMap[l] = true
	}
	labelHrefs = []string{}
	for l := range labelGroupMap {
		labelHrefs = append(labelHrefs, l)
	}
	return labelHrefs
}

func (p *PCE) expandLabelGroup(href string) (labelHrefs []string, moreSGs []string) {
	for _, l := range p.LabelGroups[href].Labels {
		labelHrefs = append(labelHrefs, l.Href)
	}
	for _, sg := range p.LabelGroups[href].SubGroups {
		moreSGs = append(moreSGs, sg.Href)
	}
	return labelHrefs, moreSGs
}
