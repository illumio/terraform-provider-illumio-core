// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioOrganizationSettingsCreate,
		ReadContext:   resourceIllumioOrganizationSettingsRead,
		UpdateContext: resourceIllumioOrganizationSettingsUpdate,
		DeleteContext: resourceIllumioOrganizationSettingsDelete,

		SchemaVersion: 1,
		Description:   "Manages Illumio Organization Settings",

		Schema: map[string]*schema.Schema{
			"audit_event_retention_seconds": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The time in seconds an audit event is stored in the database. The value should be between 86400 and 17280000",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.IntBetween(86400, 17280000),
				),
			},
			"audit_event_min_severity": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Minimum severity level of audit event messages. Allowed values are \"error\", \"warning\", and \"informational\" ",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"error", "warning", "informational"}, false),
				),
			},
			"format": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The log format (JSON, CEF, LEEF), which applies to all remote Syslog destinations. Allowed values are \"JSON\", \"CEF\", and \"LEEF\" ",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"JSON", "CEF", "LEEF"}, false),
				),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioOrganizationSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	d.SetId(fmt.Sprintf("/orgs/%v/settings/events", illumioClient.OrgID))
	return resourceIllumioOrganizationSettingsUpdate(ctx, d, m)
}

func resourceIllumioOrganizationSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

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

	return diags
}

func resourceIllumioOrganizationSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	organizationSettings := &models.OrganizationSettings{
		AuditEventRetentionSeconds: d.Get("audit_event_retention_seconds").(int),
		AuditEventMinSeverity:      d.Get("audit_event_min_severity").(string),
		Format:                     d.Get("format").(string),
	}

	if d.HasChanges("audit_event_retention_seconds", "audit_event_min_severity", "format") {
		_, err := illumioClient.Update(d.Id(), organizationSettings)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceIllumioOrganizationSettingsRead(ctx, d, m)
}

func resourceIllumioOrganizationSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "[illumio-core_organization_settings] Ignoring Delete Operation.",
	})

	return diags
}
