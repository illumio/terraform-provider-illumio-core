package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/* Sample of API response
[
	{
		"href": string
	}
]
*/

func datasourceIllumioWorkloads() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioWorkloadsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Workload",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of workload hrefs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of workload",
						},
					},
				},
			},
			"agent_active_pce_fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "FQDN of the PCE",
			},
			"container_clusters": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of container cluster URIs, encoded as a JSON string",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of workload(s) to return. Supports partial matches",
			},
			"enforcement_mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(ValidWorkloadEnforcementModeValues, false),
				),
				Description: `Enforcement mode of workload(s) to return. Allowed values are "idle", "visibility_only", "full" and "selective"`,
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A unique identifier within the external data source",
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The data source from which a resource originates",
			},
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hostname of workload(s) to return. Supports partial matches",
			},
			"include_deleted": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Include deleted workloads",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address of workload(s) to return. Supports partial matches",
			},
			"labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of lists of label URIs, encoded as a JSON string",
			},
			"last_heartbeat_on_gte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Greater than or equal to value for last heartbeat on timestamp",
			},
			"last_heartbeat_on_lte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Less than or equal to value for last heartbeat on timestamp",
			},
			"log_traffic": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Whether we want to log traffic events from this workload",
			},
			"managed": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Return managed or unmanaged workloads using this filter",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Maximum number of workloads to return.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of workload(s) to return. Supports partial matches",
			},
			"online": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Return online/offline workloads using this filter",
			},
			"os_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Operating System of workload(s) to return. Supports partial matches",
			},
			"policy_health": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"active", "warning", "error", "suspended"}, false),
				),
				Description: `Policy of health of workload(s) to return. Allowed values are "active", "warning", "error" and "suspended"`,
			},
			"security_policy_sync_state": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"staged"}, false),
				),
				Description: `Advanced search option for workload based on policy sync state. Allowed value: "staged"`,
			},
			"security_policy_update_mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"static", "adaptive"}, false),
				),
				Description: `Advanced search option for workload based on security policy update mode. Allowed values are "static" and "adaptive"`,
			},
			"ven": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URI of VEN to filter by.",
			},
			"visibility_level": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(validVisibilityLevels, false),
				),
				Description: `Filter by visibility level. Allowed values are "flow_full_detail", "flow_summary", "flow_drops", "flow_off" and "enhanced_data_collection"`,
			},
			"vulnerability_summary_vulnerability_exposure_score_gte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Greater than or equal to value for vulnerability_exposure_score",
			},
			"vulnerability_summary_vulnerability_exposure_score_lte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Less than or equal to value for vulnerability_exposure_score",
			},
		},
	}
}

func dataSourceIllumioWorkloadsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	paramKeys := []string{
		"container_clusters",
		"description",
		"enforcement_mode",
		"external_data_reference",
		"external_data_set",
		"hostname",
		"include_deleted",
		"ip_address",
		"labels",
		"log_traffic",
		"managed",
		"max_results",
		"name",
		"online",
		"os_id",
		"policy_health",
		"security_policy_sync_state",
		"security_policy_update_mode",
		"ven",
		"visibility_level",
	}

	/*
		"last_heartbeat_on[gte]",
		"last_heartbeat_on[lte]",
		"vulnerability_summary.vulnerability_exposure_score[lte]",
		"vulnerability_summary.vulnerability_exposure_score[gte]",
		"agent.active_pce_fqdn",
	*/

	params := resourceDataToMap(d, paramKeys)

	if value, ok := d.GetOk("agent_active_pce_fqdn"); ok {
		params["agent.active_pce_fqdn"] = value.(string)
	}
	if value, ok := d.GetOk("last_heartbeat_at_gte"); ok {
		params["last_heartbeat_at[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_heartbeat_at_lte"); ok {
		params["last_heartbeat_at[lte]"] = value.(string)
	}
	if value, ok := d.GetOk("vulnerability_summary_vulnerability_exposure_score_gte"); ok {
		params["vulnerability_summary.vulnerability_exposure_score[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("vulnerability_summary_vulnerability_exposure_score_lte"); ok {
		params["vulnerability_summary.vulnerability_exposure_score[lte]"] = value.(string)
	}

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%v/workloads", pConfig.OrgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	d.Set("items", extractHrefs(data))

	return diagnostics
}
