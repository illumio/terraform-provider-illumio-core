package models

/* Sample
{
	"href": "string",
	"deleted": true,
	"key": "string",
	"value": "string",
	"external_data_set": null,
	"external_data_reference": null,
	"created_at": "2020-08-19T21:34:26Z",
	"updated_at": "2020-08-19T21:34:26Z",
	"created_by": {
		"href": "string"
	},
	"updated_by": {
		"href": "string"
	}
}
*/

// Label represents label resource
type Label struct {
	Key                   string `json:"key"`
	Value                 string `json:"value"`
	ExternalDataSet       string `json:"external_data_set"`
	ExternalDataReference string `json:"external_data_reference"`
}

// ToMap - Returns map for Label model
func (l *Label) ToMap() (map[string]interface{}, error) {
	labelAttrMap := make(map[string]interface{})
	if l.Key != "" {
		// for PUT request we cannot set key attribute
		labelAttrMap["key"] = l.Key
	}
	if l.Value != "" {
		labelAttrMap["value"] = l.Value
	}

	labelAttrMap["external_data_set"] = nil
	if l.ExternalDataSet != "" {
		labelAttrMap["external_data_set"] = l.ExternalDataSet
	}

	labelAttrMap["external_data_reference"] = nil
	if l.ExternalDataReference != "" {
		labelAttrMap["external_data_reference"] = l.ExternalDataReference
	}

	return labelAttrMap, nil
}
