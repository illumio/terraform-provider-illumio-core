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
    "selective_enforcement_rules": [
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
	SER              []Href `json:"selective_enforcement_rules"`
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
	case "selective_enforcement_rules":
		cs.SER = append(cs.SER, Href{Href: href})
	case "firewall_settings":
	cs.SER = append(cs.FirewallSettings, Href{Href: href})
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
	spAttrMap["change_subset"] = make(map[string]interface{})
	spAttrMap["change_subset"].(map[string]interface{})["label_groups"] = GetHrefMaps(s.ChangeSubset.LabelGroups)
	spAttrMap["change_subset"].(map[string]interface{})["services"] = GetHrefMaps(s.ChangeSubset.Services)
	spAttrMap["change_subset"].(map[string]interface{})["rule_sets"] = GetHrefMaps(s.ChangeSubset.RuleSets)
	spAttrMap["change_subset"].(map[string]interface{})["ip_lists"] = GetHrefMaps(s.ChangeSubset.IPLists)
	spAttrMap["change_subset"].(map[string]interface{})["virtual_services"] = GetHrefMaps(s.ChangeSubset.VirtualServices)
	// spAttrMap["change_subset"].(map[string]interface{})["selective_enforcement_rules"] = GetHrefMaps(s.ChangeSubset.SER)
	// TODO Confirm removal
	spAttrMap["change_subset"].(map[string]interface{})["firewall_settings"] = GetHrefMaps(s.ChangeSubset.SER)
	return spAttrMap, nil
}
