// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package models

/*
{
  "update_description": "string",
  "change_subset": {
    "label_groups": [
      {
        "href": "string"
      }
    ],
    "services": [
      {
        "href": "string"
      }
    ],
    "rule_sets": [
      {
        "href": "string"
      }
    ],
    "ip_lists": [
      {
        "href": "string"
      }
    ],
    "virtual_services": [
      {
        "href": "string"
      }
    ],
    "firewall_settings": [
      {
        "href": "string"
      }
    ],
    "secure_connect_gateways": [
      {
        "href": "string"
      }
    ],
    "virtual_servers": [
      {
        "href": "string"
      }
    ],
    "enforcement_boundaries": [
      {
        "href": "string"
      }
    ]
  }
}
*/

// SecurityPolicyChangeSubset represents change subset of security policy resource
type SecurityPolicyChangeSubset struct {
	LabelGroups      []Href `json:"label_groups"`
	Services         []Href `json:"services"`
	RuleSets         []Href `json:"rule_sets"`
	IPLists          []Href `json:"ip_lists"`
	VirtualServices  []Href `json:"virtual_services"`
	EBoundaries      []Href `json:"enforcement_boundaries"`
	FirewallSettings []Href `json:"firewall_settings"`
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

// SecurityPolicy represents security policy resource
type SecurityPolicy struct {
	UpdateDesc   string                     `json:"update_description"`
	ChangeSubset SecurityPolicyChangeSubset `json:"change_subset"`
}

// ToMap - Returns map for SecurityPolicy model
func (s *SecurityPolicy) ToMap() (map[string]interface{}, error) {
	spAttrMap := make(map[string]interface{})
	spAttrMap["update_description"] = s.UpdateDesc

	changeSubset := make(map[string]interface{})
	if v := GetHrefMaps(s.ChangeSubset.LabelGroups); len(v) > 0 {
		changeSubset["label_groups"] = v
	}
	if v := GetHrefMaps(s.ChangeSubset.Services); len(v) > 0 {
		changeSubset["services"] = v
	}
	if v := GetHrefMaps(s.ChangeSubset.RuleSets); len(v) > 0 {
		changeSubset["rule_sets"] = v
	}
	if v := GetHrefMaps(s.ChangeSubset.IPLists); len(v) > 0 {
		changeSubset["ip_lists"] = v
	}
	if v := GetHrefMaps(s.ChangeSubset.VirtualServices); len(v) > 0 {
		changeSubset["virtual_services"] = v
	}
	if v := GetHrefMaps(s.ChangeSubset.EBoundaries); len(v) > 0 {
		changeSubset["enforcement_boundaries"] = v
	}
	if v := GetHrefMaps(s.ChangeSubset.FirewallSettings); len(v) > 0 {
		changeSubset["firewall_settings"] = v
	}
	spAttrMap["change_subset"] = changeSubset

	return spAttrMap, nil
}
