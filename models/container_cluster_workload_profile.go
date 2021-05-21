// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

// Sample
/*
{
  "name": null,
  "description": "string",
  "assign_labels": [
    {
      "href": "string"
    }
  ],
  "labels": [
    {
      "key": "string",
      "assignment": {
        "href": "string"
      }
    }
  ],
  "enforcement_mode": "idle",
  "managed": true
}
*/

type ContainerClusterWorkloadProfileLabel struct {
	Key         string `json:"key"`
	Assignment  Href   `json:"assignment"`
	Restriction []Href `json:"restriction"`
}

func (o *ContainerClusterWorkloadProfileLabel) HasConflicts() bool {
	if (o.Assignment.Href != "") && (len(o.Restriction) > 0) {
		return true
	} else if o.Assignment.Href != "" {
		return false
	} else if len(o.Restriction) > 0 {
		return false
	}
	return true
}

type ContainerClusterWorkloadProfile struct {
	Name            string                                 `json:"name"`
	Description     string                                 `json:"description"`
	AssignLabels    []Href                                 `json:"assign_labels"`
	Labels          []ContainerClusterWorkloadProfileLabel `json:"labels"`
	EnforcementMode string                                 `json:"enforcement_mode"`
	Managed         bool                                   `json:"managed"`
}

func (ccwp *ContainerClusterWorkloadProfile) ToMap() (map[string]interface{}, error) {
	ccWorkloadProfileMap := make(map[string]interface{})

	if ccwp.Name != "" {
		ccWorkloadProfileMap["name"] = ccwp.Name
	}
	ccWorkloadProfileMap["description"] = ccwp.Description

	if ccwp.EnforcementMode != "" {
		ccWorkloadProfileMap["enforcement_mode"] = ccwp.EnforcementMode
	}

	ccWorkloadProfileMap["managed"] = ccwp.Managed

	if ccwp.AssignLabels != nil {
		ccWorkloadProfileMap["assign_labels"] = GetHrefMaps(ccwp.AssignLabels)
	}

	if ccwp.Labels != nil {
		labelMaps := []map[string]interface{}{}
		for _, l := range ccwp.Labels {
			labelIMap := map[string]interface{}{}
			labelIMap["key"] = l.Key

			if l.Assignment.Href != "" {
				labelIMap["assignment"], _ = l.Assignment.ToMap()
			} else {
				labelIMap["restriction"] = GetHrefMaps(l.Restriction)
			}
			labelMaps = append(labelMaps, labelIMap)
		}
		ccWorkloadProfileMap["labels"] = labelMaps
	}

	return ccWorkloadProfileMap, nil
}
