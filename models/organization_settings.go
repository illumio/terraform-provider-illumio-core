// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type OrganizationSettings struct {
	AuditEventRetentionSeconds int    `json:"audit_event_retention_seconds,omitempty"`
	AuditEventMinSeverity      string `json:"audit_event_min_severity,omitempty"`
	Format                     string `json:"format,omitempty"`
}

func (os *OrganizationSettings) ToMap() (map[string]interface{}, error) {
	return toMap(os)
}
