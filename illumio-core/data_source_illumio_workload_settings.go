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
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workload Disconnected Timeout Seconds for Workload Settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type:        schema.TypeList,
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
					},
				},
			},
			"workload_goodbye_timeout_seconds": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workload Goodbye Timeout Seconds for Workload Settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type:        schema.TypeList,
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

	if data.Exists("workload_disconnected_timeout_seconds") {
		wdtsS := data.S("workload_disconnected_timeout_seconds")
		wdtsI := []map[string]interface{}{}

		for _, wdts := range wdtsS.Children() {
			wdtsMap := extractMap(wdts, []string{"scope", "value"})
			if wdts.Exists("scope") {
				wdtsMap["scope"] = extractMapArray(wdts.S("scope"), []string{"href"})
			} else {
				wdtsMap["scope"] = nil
			}
			wdtsI = append(wdtsI, wdtsMap)
		}

		d.Set("workload_disconnected_timeout_seconds", wdtsI)
	} else {
		d.Set("workload_disconnected_timeout_seconds", nil)
	}

	if data.Exists("workload_goodbye_timeout_seconds") {
		wgtsS := data.S("workload_goodbye_timeout_seconds")
		wgtsI := []map[string]interface{}{}

		for _, wgts := range wgtsS.Children() {
			wgtsMap := extractMap(wgts, []string{"scope", "value"})
			if wgts.Exists("scope") {
				wgtsMap["scope"] = extractMapArray(wgts.S("scope"), []string{"href"})
			} else {
				wgtsMap["scope"] = nil
			}
			wgtsI = append(wgtsI, wgtsMap)
		}

		d.Set("workload_goodbye_timeout_seconds", wgtsI)
	} else {
		d.Set("workload_goodbye_timeout_seconds", nil)
	}

	return diagnostics
}
