// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioWorkloadSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioWorkloadSettingsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Workload Settings",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the Workload Settings",
			},
			"workload_disconnected_timeout_seconds": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Workload Disconnected Timeout Seconds for Workload Settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Assigned labels for Workload Disconnected Timeout Seconds",
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
						"value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Property value associated with the scope",
						},
						"ven_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The VEN type that this property is applicable to. Must be "server" or "endpoint". An empty or missing value will default to "server" on the PCE`,
						},
					},
				},
			},
			"workload_goodbye_timeout_seconds": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Workload Goodbye Timeout Seconds for Workload Settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Assigned labels for Workload Goodbye Timeout Seconds",
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
						"value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Property value associated with the scope",
						},
						"ven_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `The VEN type that this property is applicable to. Must be "server" or "endpoint". An empty or missing value will default to "server" on the PCE`,
						},
					},
				},
			},
		},
	}
}

func datasourceIllumioWorkloadSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/settings/workloads", orgID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))

	d.Set("href", data.S("href").Data().(string))

	for _, k := range []string{
		"workload_disconnected_timeout_seconds",
		"workload_goodbye_timeout_seconds",
	} {
		d.Set(k, extractWorkloadSettingsTimeout(data, k))
	}

	return diagnostics
}
