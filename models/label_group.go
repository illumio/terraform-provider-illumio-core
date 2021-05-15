package models

// LabelGroup represents label group resource
type LabelGroup struct {
	Name                  string `json:"name"`
	Description           string `json:"description"`
	Key                   string `json:"key"`
	Labels                []Href `json:"labels"`
	SubGroups             []Href `json:"sub_groups"`
	ExternalDataSet       string `json:"external_data_set"`
	ExternalDataReference string `json:"external_data_reference"`
}

// ToMap - Returns map for LabelGroup model
func (l *LabelGroup) ToMap() (map[string]interface{}, error) {
	lgAttrMap := make(map[string]interface{})
	if l.Key != "" {
		// for PUT request we cannot set key attribute
		lgAttrMap["key"] = l.Key
	}

	if l.Name != "" {
		lgAttrMap["name"] = l.Name
	}

	lgAttrMap["labels"] = GetHrefMaps(l.Labels)
	lgAttrMap["sub_groups"] = GetHrefMaps(l.SubGroups)

	lgAttrMap["description"] = l.Description

	lgAttrMap["external_data_reference"] = nil
	if l.ExternalDataReference != "" {
		lgAttrMap["external_data_reference"] = l.ExternalDataReference
	}

	lgAttrMap["external_data_set"] = nil
	if l.ExternalDataSet != "" {
		lgAttrMap["external_data_set"] = l.ExternalDataSet
	}

	return lgAttrMap, nil
}
