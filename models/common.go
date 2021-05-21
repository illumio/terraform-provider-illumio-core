// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

// Href - Represents href object for Illumio Resource
type Href struct {
	Href string `json:"href"`
}

// GetHrefs - Returns objects of Href from [{"href": "..."}, ...]
func GetHrefs(arr []interface{}) []Href {
	var hrefs []Href
	for _, elem := range arr {
		hrefs = append(hrefs, Href{
			Href: elem.(map[string]interface{})["href"].(string),
		})
	}
	return hrefs
}

// GetHrefMaps - Returns list of map with href as key
func GetHrefMaps(hrefs []Href) []map[string]string {
	l := []map[string]string{}
	for _, elem := range hrefs {
		l = append(l, map[string]string{
			"href": elem.Href,
		})
	}
	return l
}

// ToMap - Returns map for Href model
func (a *Href) ToMap() (map[string]interface{}, error) {
	hrefAttrMap := make(map[string]interface{})
	hrefAttrMap["href"] = a.Href
	return hrefAttrMap, nil
}

type ScopeObj struct {
	Label      *Href `json:"label"`
	LabelGroup *Href `json:"label_group"`
}

type Scope []ScopeObj
type Scopes []Scope

func (s Scopes) ToList() [][]map[string]map[string]interface{} {
	scps := [][]map[string]map[string]interface{}{}
	for _, scope := range s {
		scp := []map[string]map[string]interface{}{}
		for _, val := range scope {
			o := make(map[string]map[string]interface{})
			if val.Label != nil && val.Label.Href != "" {
				o["label"], _ = val.Label.ToMap()
			}
			if val.LabelGroup != nil && val.LabelGroup.Href != "" {
				o["label_group"], _ = val.LabelGroup.ToMap()
			}
			scp = append(scp, o)
		}
		scps = append(scps, scp)
	}
	return scps
}
