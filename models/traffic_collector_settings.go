// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type TrafficCollectorSettings struct {
	Transmission string                          `json:"transmission"`
	Target       *TrafficCollectorSettingsTarget `json:"target"`
	Action       string                          `json:"action"`
}

type TrafficCollectorSettingsTarget struct {
	DstPort int    `json:"dst_port"`
	Proto   int    `json:"proto"`
	DstIP   string `json:"dst_ip"`
}

func (o *TrafficCollectorSettings) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{}

	if o.Transmission != "" {
		m["transmission"] = o.Transmission
	}

	if o.Action != "" {
		m["action"] = o.Action
	}

	if o.Target != nil {
		m["target"] = o.Target.ToMAP()
	}

	return m, nil
}

func (o *TrafficCollectorSettingsTarget) ToMAP() map[string]interface{} {
	return map[string]interface{}{
		"dst_port": o.DstPort,
		"proto":    o.Proto,
		"dst_ip":   o.DstIP,
	}
}
