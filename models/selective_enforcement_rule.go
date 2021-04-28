package models

/* Sample
{
  "name": "string",
  "scope": [
    {
      "label": {
        "href": "string"
      },
      "label_group": {
        "href": "string"
      }
    }
  ],
  "enforced_services": [
    {
      "href": "string"
    }
  ]
}
*/

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
func (s *SelectiveEnforcementRules) ToMap() (map[string]interface{}, error) {
	serM := make(map[string]interface{})

	if s.Name != "" {
		serM["name"] = s.Name
	}

	n := []map[string]map[string]interface{}{}
	for _, val := range s.Scope {
		o := make(map[string]map[string]interface{})
		if val.Label != nil && val.Label.Href != "" {
			o["label"], _ = val.Label.ToMap()
		}
		if val.LabelGroup != nil && val.LabelGroup.Href != "" {
			o["label_group"], _ = val.LabelGroup.ToMap()
		}
		n = append(n, o)
	}
	serM["scope"] = n

	serM["enforced_services"] = GetHrefMaps(s.EnforcedServices)

	return serM, nil
}
