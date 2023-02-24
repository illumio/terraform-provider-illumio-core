// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type WorkloadSettingsTimeout struct {
	Scope   []Href  `json:"scope"`
	Value   int     `json:"value"`
	VENType *string `json:"ven_type"`
}

type WorkloadSettings struct {
	WorkloadDisconnectedTimeoutSeconds []WorkloadSettingsTimeout `json:"workload_disconnected_timeout_seconds,omitempty"`
	WorkloadGoodbyeTimeoutSeconds      []WorkloadSettingsTimeout `json:"workload_goodbye_timeout_seconds,omitempty"`
}

func (ws *WorkloadSettings) ToMap() (map[string]interface{}, error) {
	return toMap(ws)
}
