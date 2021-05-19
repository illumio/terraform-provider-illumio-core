// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

import (
	"strconv"
)

/* Sample
{
  "href": "string",
  "name": "string",
  "description": "string",
  "mode": "idle",
  "enforcement_mode": "idle",
  "status": "string",
  "enabled": true,
  "total_use_count": 0,
  "allowed_uses_per_key": 1,
  "key_lifespan": 1,
  "last_pairing_at": "string",
  "created_at": "2021-03-02T02:37:59Z",
  "updated_at": "2021-03-02T02:37:59Z",
  "created_by": {
    "href": "string"
  },
  "updated_by": {
    "href": "string"
  },
  "is_default": true,
  "labels": [
    {
      "href": "string"
    }
  ],
  "env_label_lock": true,
  "loc_label_lock": true,
  "role_label_lock": true,
  "app_label_lock": true,
  "mode_lock": true,
  "enforcement_mode_lock": true,
  "log_traffic": true,
  "log_traffic_lock": true,
  "visibility_level": "string",
  "visibility_level_lock": true,
  "status_lock": true,
  "external_data_set": null,
  "external_data_reference": null,
  "agent_software_release": null
}
*/

type PairingProfile struct {
	Name                  string `json:"name"`
	Description           string `json:"description"`
	EnforcementMode       string `json:"enforcement_mode"`
	EnforcementModeLock   bool   `json:"enforcement_mode_lock"`
	Enabled               bool   `json:"enabled"`
	AllowedUsesPerKey     string `json:"allowed_uses_per_key"`
	KeyLifespan           string `json:"key_lifespan"`
	Labels                []Href `json:"labels"`
	EnvLabelLock          bool   `json:"env_label_lock"`
	LocLabelLock          bool   `json:"loc_label_lock"`
	RoleLabelLock         bool   `json:"role_label_lock"`
	AppLabelLock          bool   `json:"app_label_lock"`
	LogTraffic            bool   `json:"log_traffic"`
	LogTrafficLock        bool   `json:"log_traffic_lock"`
	VisibilityLevel       string `json:"visibility_level"`
	VisibilityLevelLock   bool   `json:"visibility_level_lock"`
	ExternalDataSet       string `json:"external_data_set"`
	ExternalDataReference string `json:"external_data_reference"`
	AgentSoftwareRelease  string `json:"agent_software_release"`
}

// ToMap - Returns map for PairingProfile model
func (pp *PairingProfile) ToMap() (map[string]interface{}, error) {
	ppAttrMap := make(map[string]interface{})

	if pp.Name != "" {
		ppAttrMap["name"] = pp.Name
	}

	if pp.EnforcementMode != "" {
		ppAttrMap["enforcement_mode"] = pp.EnforcementMode
	}

	if pp.AllowedUsesPerKey != "" {
		if pp.AllowedUsesPerKey != "unlimited" {
			ppAttrMap["allowed_uses_per_key"], _ = strconv.Atoi(pp.AllowedUsesPerKey)
		} else {
			ppAttrMap["allowed_uses_per_key"] = "unlimited"
		}
	}

	if pp.KeyLifespan != "" {
		if pp.KeyLifespan != "unlimited" {
			ppAttrMap["key_lifespan"], _ = strconv.Atoi(pp.KeyLifespan)
		} else {
			ppAttrMap["key_lifespan"] = "unlimited"
		}
	}

	if pp.VisibilityLevel != "" {
		ppAttrMap["visibility_level"] = pp.VisibilityLevel
	}
	if pp.AgentSoftwareRelease != "" {
		ppAttrMap["agent_software_release"] = pp.AgentSoftwareRelease
	}

	ppAttrMap["labels"] = GetHrefMaps(pp.Labels)

	ppAttrMap["description"] = pp.Description

	ppAttrMap["external_data_set"] = nil
	if pp.ExternalDataSet != "" {
		ppAttrMap["external_data_set"] = pp.ExternalDataSet
	}

	ppAttrMap["external_data_reference"] = nil
	if pp.ExternalDataReference != "" {
		ppAttrMap["external_data_reference"] = pp.ExternalDataReference
	}

	ppAttrMap["enabled"] = pp.Enabled
	ppAttrMap["enforcement_mode_lock"] = pp.EnforcementModeLock
	ppAttrMap["env_label_lock"] = pp.EnvLabelLock
	ppAttrMap["loc_label_lock"] = pp.LocLabelLock
	ppAttrMap["role_label_lock"] = pp.RoleLabelLock
	ppAttrMap["app_label_lock"] = pp.AppLabelLock
	ppAttrMap["log_traffic"] = pp.LogTraffic
	ppAttrMap["log_traffic_lock"] = pp.LogTrafficLock
	ppAttrMap["visibility_level_lock"] = pp.VisibilityLevelLock

	return ppAttrMap, nil
}
