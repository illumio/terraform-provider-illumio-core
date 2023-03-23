// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type SyslogDestinationAuditEventLogger struct {
	ConfigurationEventIncluded *bool   `json:"configuration_event_included,omitempty"`
	SystemEventIncluded        *bool   `json:"system_event_included,omitempty"`
	MinSeverity                *string `json:"min_severity,omitempty"`
}

type SyslogDestinationTrafficEventLogger struct {
	TrafficFlowAllowedEventIncluded            *bool `json:"traffic_flow_allowed_event_included,omitempty"`
	TrafficFlowPotentiallyBlockedEventIncluded *bool `json:"traffic_flow_potentially_blocked_event_included,omitempty"`
	TrafficFlowBlockedEventIncluded            *bool `json:"traffic_flow_blocked_event_included,omitempty"`
}

type SyslogDestinationNodeStatusLogger struct {
	NodeStatusIncluded *bool `json:"node_status_included,omitempy"`
}

type SyslogDestinationRemoteSyslog struct {
	Address       *string `json:"address,omitempty"`
	Port          *int    `json:"port,omitempty"`
	Protocol      *int    `json:"protocol,omitempty"`
	TLSEnabled    *bool   `json:"tls_enabled,omitempty"`
	TLSCaBundle   *string `json:"tls_ca_bundle,omitempty"`
	TLSVerifyCert *bool   `json:"tls_verify_cert,omitempty"`
}

type SyslogDestination struct {
	PceScope           *[]string                            `json:"pce_scope,omitempty"`
	Type               *string                              `json:"type,omitempty"`
	Description        *string                              `json:"description,omitempty"`
	AuditEventLogger   *SyslogDestinationAuditEventLogger   `json:"audit_event_logger,omitempty"`
	TrafficEventLogger *SyslogDestinationTrafficEventLogger `json:"traffic_event_logger,omitempty"`
	NodeStatusLogger   *SyslogDestinationNodeStatusLogger   `json:"node_status_logger,omitempty"`
	RemoteSyslog       *SyslogDestinationRemoteSyslog       `json:"remote_syslog,omitempty"`
}

func (sd *SyslogDestination) ToMap() (map[string]interface{}, error) {
	return toMap(sd)
}
