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

		SchemaVersion: version,
		Description:   "Manages Illumio Organization Settings",

		Schema: map[string]*schema.Schema{
			"audit_event_retention_seconds": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The time in seconds an audit event is stored in the database. The value should be in between 86400 and 17280000",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.IntBetween(86400, 17280000),
				),
			},
			"audit_event_min_severity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Minimum severity level of audit event messages. Allowed values : \"error\", \"warning\", and \"informational\" ",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"error", "warning", "informational"}, false),
				),
			},
			"format": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The log format (JSON, CEF, LEEF), which applies to all remote syslog destinations. Allowed values : \"JSON\", \"CEF\", and \"LEEF\" ",
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
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Detail:   "[illumio-core_organization_settings] Cannot use create operation.",
		Summary:  "Please use terraform import...",
	})

	return diags
}

func resourceIllumioOrganizationSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
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

	return diags
}

func resourceIllumioOrganizationSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	OrganizationSettings := &models.OrganizationSettings{}

	OrganizationSettings.AuditEventRetentionSeconds = d.Get("audit_event_retention_seconds").(int)

	OrganizationSettings.AuditEventMinSeverity = d.Get("audit_event_min_severity").(string)

	OrganizationSettings.Format = d.Get("format").(string)

	_, err := illumioClient.Update(d.Id(), OrganizationSettings)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioOrganizationSettingsRead(ctx, d, m)
}

func resourceIllumioOrganizationSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "[illumio-core_organization_settings] Ignoring Delete Operation...",
	})

	return diags
}
