// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample of API response
[
	{
		"href": "string",
		"name": "string",
		"description": "string",
		"enforcement_mode": "idle",
		"status": "string",
		"enabled": true,
		"total_use_count": 0,
		"allowed_uses_per_key": 1,
		"key_lifespan": 1,
		"last_pairing_at": "string",
		"created_at": "2021-03-02T02:37:59Z",
		"updated_at": "2021-03-02T02:37:59Z",
		"created_by": {
			"href": "string"
		},
		"updated_by": {
			"href": "string"
		},
		"is_default": true,
		"labels": [
			{
			"href": "string"
			}
		],
		"env_label_lock": true,
		"loc_label_lock": true,
		"role_label_lock": true,
		"app_label_lock": true,
		"enforcement_mode_lock": true,
		"log_traffic": true,
		"log_traffic_lock": true,
		"visibility_level": "string",
		"visibility_level_lock": true,
		"status_lock": true,
		"external_data_set": null,
		"external_data_reference": null,
		"agent_software_release": null
	}
]
*/

func datasourceIllumioPairingProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioPairingProfilesRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Pairing Profiles",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of pairing profiles",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of pairing profile",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The short friendly name of the pairing profile",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The long description of the pairing profile",
						},
						"enforcement_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Filter by mode",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State of VEN",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The enabled flag of the pairing profile",
						},
						"total_use_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of times the pairing profile has been used",
						},
						"allowed_uses_per_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The number of times the pairing profile can be use",
						},
						"key_lifespan": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Number of seconds pairing profile keys will be valid for",
						},
						"last_pairing_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this pairing profile was last used for pairing a workload",
						},
						"external_data_set": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The data source from which a resource originates",
						},
						"external_data_reference": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique identifier within the external data source",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this pairing profile was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this pairing profile was last updated",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who originally created this pairing profile",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who last updated this pairing profile",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"is_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag indicating this is default auto-created pairing profile",
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Assigned labels",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Label URI",
									},
								},
							},
						},
						"env_label_lock": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag that controls whether env label can be overridden from pairing script",
						},
						"loc_label_lock": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag that controls whether loc label can be overridden from pairing script",
						},
						"role_label_lock": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag that controls whether role label can be overridden from pairing script",
						},
						"app_label_lock": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag that controls whether app label can be overridden from pairing script",
						},
						"enforcement_mode_lock": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag that controls whether app label can be overridden from pairing script",
						},
						"log_traffic": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "State of VEN",
						},
						"log_traffic_lock": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag that controls whether log_traffic can be overridden from pairing script",
						},
						"visibility_level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Visibility level of the workload",
						},
						"visibility_level_lock": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag that controls whether visibility_level can be overridden from pairing script",
						},
						"status_lock": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag that controls whether status can be overridden from pairing script",
						},
						"agent_software_release": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Agent software release associated with this paring profile",
						},
						"caps": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "CAP",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"agent_software_release": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The agent software release for pairing profiles",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of Pairing Profile(s) to return. Supports partial matches",
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
			"labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of lists of label URIs, encoded as a JSON string",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of Pairing Profiles to return. The integer should be a non-zero positive integer.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of Pairing Profile(s) to return. Supports partial matches",
			},
		},
	}
}

func dataSourceIllumioPairingProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	orgID := pConfig.OrgID

	href := fmt.Sprintf("/orgs/%v/pairing_profiles", orgID)

	paramKeys := []string{
		"agent_software_release",
		"description",
		"external_data_reference",
		"external_data_set",
		"labels",
		"max_results",
		"name",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.Get(href, &params)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	ppKeys := []string{
		"href",
		"name",
		"description",
		"enforcement_mode",
		"status",
		"enabled",
		"total_use_count",
		"allowed_uses_per_key",
		"key_lifespan",
		"last_pairing_at",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"is_default",
		"env_label_lock",
		"loc_label_lock",
		"role_label_lock",
		"app_label_lock",
		"enforcement_mode_lock",
		"log_traffic",
		"log_traffic_lock",
		"visibility_level",
		"visibility_level_lock",
		"status_lock",
		"agent_software_release",
		"caps",
	}

	pps := []map[string]interface{}{}

	for _, child := range data.Children() {
		pp := extractMap(child, ppKeys)

		key := "labels"
		if child.Exists(key) {
			pp[key] = extractMapArray(child.S(key), []string{"href"})
		} else {
			pp[key] = nil
		}

		pps = append(pps, pp)
	}

	d.Set("items", pps)

	return diagnostics
}
