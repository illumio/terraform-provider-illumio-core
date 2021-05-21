// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

//Sample
/*
{
  "workload_disconnected_timeout_seconds": [
    {
      "scope": [
        {
          "href": "string"
        }
      ],
      "value": -1
    }
  ],
  "workload_goodbye_timeout_seconds": [
    {
      "scope": [
        {
          "href": "string"
        }
      ],
      "value": -1
    }
  ]
}
*/

type WorkloadSettingsTimeout struct {
	Scope []Href `json:"scope"`
	Value int    `json:"value"`
}

type WorkloadSettings struct {
	WorkloadDisconnectedTimeoutSeconds []WorkloadSettingsTimeout `json:"workload_disconnected_timeout_seconds"`
	WorkloadGoodbyeTimeoutSeconds      []WorkloadSettingsTimeout `json:"workload_goodbye_timeout_seconds"`
}

func (w *WorkloadSettings) ToMap() (map[string]interface{}, error) {
	workloadSettingsAttrMap := make(map[string]interface{})

	if len(w.WorkloadDisconnectedTimeoutSeconds) > 0 {
		wdtsMap := []map[string]interface{}{}
		for _, o := range w.WorkloadDisconnectedTimeoutSeconds {
			m := make(map[string]interface{})
			if len(o.Scope) > 0 {
				m["scope"] = GetHrefMaps(o.Scope)
			} else {
				m["scope"] = []map[string]string{}
			}
			m["value"] = o.Value

			wdtsMap = append(wdtsMap, m)
		}
		workloadSettingsAttrMap["workload_disconnected_timeout_seconds"] = wdtsMap
	}

	if len(w.WorkloadGoodbyeTimeoutSeconds) > 0 {
		wgtsMap := []map[string]interface{}{}
		for _, o := range w.WorkloadGoodbyeTimeoutSeconds {
			m := make(map[string]interface{})
			if len(o.Scope) > 0 {
				m["scope"] = GetHrefMaps(o.Scope)
			} else {
				m["scope"] = []map[string]string{}
			}

			m["value"] = o.Value

			wgtsMap = append(wgtsMap, m)
		}
		workloadSettingsAttrMap["workload_goodbye_timeout_seconds"] = wgtsMap
	}

	return workloadSettingsAttrMap, nil
}
