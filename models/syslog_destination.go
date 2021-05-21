// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

import "fmt"

type SyslogDestination struct {
	PceScope           []string                             `json:"pce_scope"`
	Type               string                               `json:"type"`
	Description        string                               `json:"description"`
	AuditEventLogger   *SyslogDestinationAuditEventLogger   `json:"audit_event_logger"`
	TrafficEventLogger *SyslogDestinationTrafficEventLogger `json:"traffic_event_logger"`
	NodeStatusLogger   *SyslogDestinationNodeStatusLogger   `json:"node_status_logger"`
	RemoteSyslog       *SyslogDestinationRemoteSyslog       `json:"remote_syslog"`
}

func (o *SyslogDestination) ToMap() (map[string]interface{}, error) {
	if o == nil {
		return nil, fmt.Errorf("[Error] Syslog Destination: got nil in ToMap function")
	}

	m := map[string]interface{}{}

	if o.PceScope != nil {
		m["pce_scope"] = o.PceScope
	}

	if o.Type != "" {
		m["type"] = o.Type
	}

	m["description"] = o.Description

	if o.AuditEventLogger != nil {
		m["audit_event_logger"] = o.AuditEventLogger.ToMap()
	}

	if o.TrafficEventLogger != nil {
		m["traffic_event_logger"] = o.TrafficEventLogger.ToMap()
	}

	if o.NodeStatusLogger != nil {
		m["node_status_logger"] = o.NodeStatusLogger.ToMap()
	}

	if o.RemoteSyslog != nil {
		m["remote_syslog"] = o.RemoteSyslog.ToMap()
	}

	return m, nil
}

type SyslogDestinationAuditEventLogger struct {
	ConfigurationEventIncluded bool   `json:"configuration_event_included"`
	SystemEventIncluded        bool   `json:"system_event_included"`
	MinSeverity                string `json:"min_severity"`
}

func (o *SyslogDestinationAuditEventLogger) ToMap() map[string]interface{} {
	if o == nil {
		return nil
	}
	return map[string]interface{}{
		"configuration_event_included": o.ConfigurationEventIncluded,
		"system_event_included":        o.SystemEventIncluded,
		"min_severity":                 o.MinSeverity,
	}
}

type SyslogDestinationTrafficEventLogger struct {
	TrafficFlowAllowedEventIncluded            bool `json:"traffic_flow_allowed_event_included"`
	TrafficFlowPotentiallyBlockedEventIncluded bool `json:"traffic_flow_potentially_blocked_event_included"`
	TrafficFlowBlockedEventIncluded            bool `json:"traffic_flow_blocked_event_included"`
}

func (o *SyslogDestinationTrafficEventLogger) ToMap() map[string]interface{} {
	if o == nil {
		return nil
	}
	return map[string]interface{}{
		"traffic_flow_allowed_event_included":             o.TrafficFlowAllowedEventIncluded,
		"traffic_flow_potentially_blocked_event_included": o.TrafficFlowPotentiallyBlockedEventIncluded,
		"traffic_flow_blocked_event_included":             o.TrafficFlowBlockedEventIncluded,
	}
}

type SyslogDestinationNodeStatusLogger struct {
	NodeStatusIncluded bool `json:"node_status_included"`
}

func (o *SyslogDestinationNodeStatusLogger) ToMap() map[string]interface{} {
	if o == nil {
		return nil
	}
	return map[string]interface{}{
		"node_status_included": o.NodeStatusIncluded,
	}
}

type SyslogDestinationRemoteSyslog struct {
	Address       string `json:"address"`
	Port          int    `json:"port"`
	Protocol      int    `json:"protocol"`
	TLSEnabled    bool   `json:"tls_enabled"`
	TLSCaBundle   string `json:"tls_ca_bundle"`
	TLSVerifyCert bool   `json:"tls_verify_cert"`
}

func (o *SyslogDestinationRemoteSyslog) ToMap() map[string]interface{} {
	if o == nil {
		return nil
	}
	return map[string]interface{}{
		"address":         o.Address,
		"port":            o.Port,
		"protocol":        o.Protocol,
		"tls_enabled":     o.TLSEnabled,
		"tls_ca_bundle":   o.TLSCaBundle,
		"tls_verify_cert": o.TLSVerifyCert,
	}
}
