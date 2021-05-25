// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample API Responce
{
  "audit_event_retention_seconds": 0,
  "audit_event_min_severity": "error",
  "format": "string"
}
*/

func datasourceIllumioOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioOrganizationSettingsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Organization Settings",

		Schema: map[string]*schema.Schema{
			"audit_event_retention_seconds": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The time in seconds an audit event is stored in the database",
			},
			"audit_event_min_severity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Minimum severity level of audit event messages",
			},
			"format": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The log format (JSON, CEF, LEEF), which applies to all remote Syslog destinations",
			},
		},
	}
}

func datasourceIllumioOrganizationSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/settings/events", orgID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("/orgs/%v/settings/events", orgID))

	for _, key := range []string{
		"audit_event_retention_seconds",
		"audit_event_min_severity",
		"format",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	return diagnostics
}
