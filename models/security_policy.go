// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

// SecurityPolicyChangeSubset represents change subset of security policy resource
type SecurityPolicyChangeSubset struct {
	LabelGroups      []Href `json:"label_groups,omitempty"`
	Services         []Href `json:"services,omitempty"`
	RuleSets         []Href `json:"rule_sets,omitempty"`
	IPLists          []Href `json:"ip_lists,omitempty"`
	VirtualServices  []Href `json:"virtual_services,omitempty"`
	EBoundaries      []Href `json:"enforcement_boundaries,omitempty"`
	FirewallSettings []Href `json:"firewall_settings,omitempty"`
}

func (cs *SecurityPolicyChangeSubset) AppendHref(rtype, href string) {
	switch rtype {
	case "label_groups":
		cs.LabelGroups = append(cs.LabelGroups, Href{Href: href})
	case "services":
		cs.Services = append(cs.Services, Href{Href: href})
	case "rule_sets":
		cs.RuleSets = append(cs.RuleSets, Href{Href: href})
	case "ip_lists":
		cs.IPLists = append(cs.IPLists, Href{Href: href})
	case "virtual_services":
		cs.VirtualServices = append(cs.VirtualServices, Href{Href: href})
	case "enforcement_boundaries":
		cs.EBoundaries = append(cs.EBoundaries, Href{Href: href})
	case "firewall_settings":
		cs.FirewallSettings = append(cs.FirewallSettings, Href{Href: href})
	}
}

func (cs *SecurityPolicyChangeSubset) Size() int {
	return len(cs.LabelGroups) + len(cs.Services) + len(cs.RuleSets) +
		len(cs.IPLists) + len(cs.VirtualServices) + len(cs.EBoundaries) +
		len(cs.FirewallSettings)
}

// SecurityPolicy represents security policy resource
type SecurityPolicy struct {
	UpdateDesc   string                     `json:"update_description"`
	ChangeSubset SecurityPolicyChangeSubset `json:"change_subset"`
}

// ToMap - Returns map for SecurityPolicy model
func (sp *SecurityPolicy) ToMap() (map[string]interface{}, error) {
	return toMap(sp)
}
