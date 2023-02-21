// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type WorkloadSettingsTimeout struct {
	Scope []Href `json:"scope"`
	Value int    `json:"value"`
}

type WorkloadSettings struct {
	WorkloadDisconnectedTimeoutSeconds []WorkloadSettingsTimeout `json:"workload_disconnected_timeout_seconds"`
	WorkloadGoodbyeTimeoutSeconds      []WorkloadSettingsTimeout `json:"workload_goodbye_timeout_seconds"`
}

func (ws *WorkloadSettings) ToMap() (map[string]interface{}, error) {
	return toMap(ws)
}
