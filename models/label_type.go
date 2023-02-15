// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

import (
	"encoding/json"
)

/* Sample
{
  "href": "/orgs/1179745/label_dimensions/313b80fa-97c0-404d-9570-46ba3ab6c171",
  "key": "os",
  "display_name": "OS",
  "created_at": "2023-02-15T19:18:09.608Z",
  "updated_at": "2023-02-15T19:18:09.608Z",
  "display_info": {
    "icon": "container-workload",
    "initial": "OS",
    "background_color": "#818286",
    "foreground_color": "#ffffff",
    "display_name_plural": "OS"
  },
  "deleted": false,
  "deleted_at": null,
  "usage": {
    "labels": false,
    "label_groups": false
  },
  "caps": [
    "write",
    "delete"
  ],
  "created_by": {
    "href": "/users/4503599627370772"
  },
  "updated_by": {
    "href": "/users/4503599627370772"
  },
  "deleted_by": null
}
*/

type LabelTypeDisplayInfo struct {
	Icon              string `json:"icon,omitempty"`
	Initial           string `json:"initial,omitempty"`
	BackgroundColor   string `json:"background_color,omitempty"`
	ForegroundColor   string `json:"foreground_color,omitempty"`
	SortOrdinal       *int   `json:"sort_ordinal,omitempty"`
	DisplayNamePlural string `json:"display_name_plural,omitempty"`
}

// Label represents label resource
type LabelType struct {
	Key                   string `json:"key"`
	DisplayName           string `json:"display_name"`
	*LabelTypeDisplayInfo `json:"display_info,omitempty"`
	ExternalDataSet       string `json:"external_data_set,omitempty"`
	ExternalDataReference string `json:"external_data_reference,omitempty"`
}

// ToMap - Returns map for LabelType model
func (lt *LabelType) ToMap() (map[string]interface{}, error) {
	encodedLabelType, err := json.Marshal(lt)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(encodedLabelType), &result)
	return result, err
}
