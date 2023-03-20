// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

import (
	"encoding/json"
)

// Model Interface
type Model interface {
	ToMap() (map[string]interface{}, error)
}

// Href - Represents href object for Illumio Resource
type Href struct {
	Href string `json:"href,omitempty"`
}

// GetHrefs - Returns HrefObjects from [{"href": "..."}, ...]
func GetHrefs(arr []interface{}) []Href {
	hrefs := make([]Href, 0)
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

func toMap(o interface{}) (map[string]interface{}, error) {
	encodedObject, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(encodedObject), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// ToMap - Returns map for Href model
func (a *Href) ToMap() (map[string]interface{}, error) {
	return map[string]interface{}{"href": a.Href}, nil
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
