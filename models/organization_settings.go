// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

/* Sample
{
  "audit_event_retention_seconds": 86400,
  "audit_event_min_severity": "error",
  "format": "JSON"
}
*/

type OrganizationSettings struct {
	AuditEventRetentionSeconds int    `json:"audit_event_retention_seconds"`
	AuditEventMinSeverity      string `json:"audit_event_min_severity"`
	Format                     string `json:"format"`
}

func (os *OrganizationSettings) ToMap() (map[string]interface{}, error) {
	osAttrMap := make(map[string]interface{})

	osAttrMap["audit_event_retention_seconds"] = os.AuditEventRetentionSeconds

	osAttrMap["audit_event_min_severity"] = os.AuditEventMinSeverity

	osAttrMap["format"] = os.Format

	return osAttrMap, nil
}
