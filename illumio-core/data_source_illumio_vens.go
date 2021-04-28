package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceIllumioVENs() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioVENsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio VENs",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of VEN hrefs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of VEN",
						},
					},
				},
			},
			"activation_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"pairing_key", "kerberos", "certificate"}, false),
				),
				Description: `The method in which the VEN was activated. Allowed values are "pairing_key", "kerberos" and "certificate"`,
			},
			"active_pce_fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "FQDN of the PCE",
			},
			"condition": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"agent.upgrade_time_out", "agent.missing_heartbeats_after_upgrade", "agent.clone_detected", "agent.missed_heartbeats"}, false),
				),
				Description: `A specific error condition to filter by. Allowed values are "agent.upgrade_time_out", "agent.missing_heartbeats_after_upgrade", "agent.clone_detected" and "agent.missed_heartbeats"`,
			},
			"container_clusters": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Array of container cluster URIs, encoded as a JSON string",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of VEN(s) to return. Supports partial matches.",
			},
			"disconnected_before": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Return VENs that have been disconnected since the given time",
			},
			"health": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"healthy", "unhealthy", "error", "warning"}, false),
				),
				Description: `The overall health (condition) of the VEN. Allowed values are  "healthy", "unhealthy", "error" and "warning"`,
			},
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hostname of VEN(s) to return. Supports partial matches",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address of VEN(s) to return. Supports partial matches",
			},
			"labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "2D Array of label URIs, encoded as a JSON string",
			},
			"last_goodbye_at_gte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Greater than or equal to value for last goodbye at timestamp",
			},
			"last_goodbye_at_lte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Greater than or equal to value for last goodbye at timestamp",
			},
			"last_heartbeat_at_gte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Greater than or equal to value for last heartbeat timestamp",
			},
			"last_heartbeat_at_lte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Less than or equal to value for last heartbeat timestamp",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Maximum number of VENs to return.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of VEN(s) to return. Supports partial matches",
			},
			"os": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Operating System of VEN(s) to return. Supports partial matches.",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"active", "suspended", "stopped", "uninstalled"}, false),
				),
				Description: `The current status of the VEN. Allowed values are "active", "suspended", "stopped" and "uninstalled"`,
			},
			"upgrade_pending": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Only return VENs with/without a pending upgrade",
			},
			"version_gte": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Greater than or equal to value for version",
			},
			"version_lte": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Less than or equal to value for version",
			},
		},
	}
}

func dataSourceIllumioVENsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	paramKeys := []string{
		"activation_type",
		"active_pce_fqdn",
		"condition",
		"container_clusters",
		"description",
		"disconnected_before",
		"health",
		"hostname",
		"ip_address",
		"labels",
		"max_results",
		"name",
		"os",
		"status",
		"upgrade_pending",
	}

	params := resourceDataToMap(d, paramKeys)

	if value, ok := d.GetOk("version_gte"); ok {
		params["version[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("version_lte"); ok {
		params["version[lte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_goodbye_at_gte"); ok {
		params["last_goodbye_at[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_goodbye_at_lte"); ok {
		params["last_goodbye_at[lte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_heartbeat_at_gte"); ok {
		params["last_heartbeat_at[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_heartbeat_at_lte"); ok {
		params["last_heartbeat_at[lte]"] = value.(string)
	}

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%v/vens", pConfig.OrgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	d.Set("items", extractHrefs(data))

	return diagnostics
}
