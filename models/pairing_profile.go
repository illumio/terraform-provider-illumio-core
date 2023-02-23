// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type PairingProfile struct {
	Name                  string `json:"name,omitempty"`
	Description           string `json:"description,omitempty"`
	Enabled               bool   `json:"enabled"`
	AgentSoftwareRelease  string `json:"agent_software_release,omitempty"`
	AllowedUsesPerKey     *int   `json:"allowed_uses_per_key,omitempty"`
	KeyLifespan           *int   `json:"key_lifespan,omitempty"`
	Labels                []Href `json:"labels"`
	EnforcementMode       string `json:"enforcement_mode,omitempty"`
	EnforcementModeLock   *bool  `json:"enforcement_mode_lock,omitempty"`
	EnvLabelLock          *bool  `json:"env_label_lock,omitempty"`
	LocLabelLock          *bool  `json:"loc_label_lock,omitempty"`
	RoleLabelLock         *bool  `json:"role_label_lock,omitempty"`
	AppLabelLock          *bool  `json:"app_label_lock,omitempty"`
	LogTraffic            *bool  `json:"log_traffic,omitempty"`
	LogTrafficLock        *bool  `json:"log_traffic_lock,omitempty"`
	VisibilityLevel       string `json:"visibility_level,omitempty"`
	VisibilityLevelLock   *bool  `json:"visibility_level_lock,omitempty"`
	ExternalDataSet       string `json:"external_data_set,omitempty"`
	ExternalDataReference string `json:"external_data_reference,omitempty"`
}

// ToMap - Returns map for PairingProfile model
func (pp *PairingProfile) ToMap() (map[string]interface{}, error) {
	return toMap(pp)
}
