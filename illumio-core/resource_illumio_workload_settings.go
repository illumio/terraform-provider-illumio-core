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

func resourceIllumioWorkloadSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioWorkloadSettingsCreate,
		ReadContext:   resourceIllumioWorkloadSettingsRead,
		UpdateContext: resourceIllumioWorkloadSettingsUpdate,
		DeleteContext: resourceIllumioWorkloadSettingsDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio Workload Settings",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the Workload Settings",
			},
			"workload_disconnected_timeout_seconds": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Workload Disconnected Timeout Seconds for Workload Settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Assigned labels for Workload Disconnected Timeout Seconds",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Label URI",
									},
								},
							},
						},
						"value": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Property value associated with the scope. Allowed range is 300 - 2147483647 or -1",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.Any(validation.IntBetween(300, 2147483647), validation.IntInSlice([]int{-1})),
							),
						},
					},
				},
			},
			"workload_goodbye_timeout_seconds": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Workload Goodbye Timeout Seconds for Workload Settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Assigned labels for Workload Goodbye Timeout Seconds",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Label URI",
									},
								},
							},
						},
						"value": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Property value associated with the scope. Allowed range is 300 - 2147483647 or -1",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.Any(validation.IntBetween(300, 2147483647), validation.IntInSlice([]int{-1})),
							),
						},
					},
				},
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioWorkloadSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Detail:   "[illumio-core_workload_settings] Cannot use create operation.",
		Summary:  "Please use terraform import...",
	})

	return diags
}

func resourceIllumioWorkloadSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := fmt.Sprintf("/orgs/%v/settings/workloads", illumioClient.OrgID)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(href)
	d.Set("href", href)

	if data.Exists("workload_disconnected_timeout_seconds") {
		wdtsS := data.S("workload_disconnected_timeout_seconds")
		wdtsI := []map[string]interface{}{}

		for _, wdts := range wdtsS.Children() {
			wdtsMap := extractMap(wdts, []string{"scope", "value"})
			if wdts.Exists("scope") {
				wdtsMap["scope"] = extractMapArray(wdts.S("scope"), []string{"href"})
			} else {
				wdtsMap["scope"] = []map[string]interface{}{}
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
				wgtsMap["scope"] = []map[string]interface{}{}
			}
			wgtsI = append(wgtsI, wgtsMap)
		}

		d.Set("workload_goodbye_timeout_seconds", wgtsI)
	} else {
		d.Set("workload_goodbye_timeout_seconds", nil)
	}

	return diags
}

func resourceIllumioWorkloadSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	WorkloadSettings := &models.WorkloadSettings{}

	for _, x := range []string{"workload_disconnected_timeout_seconds", "workload_goodbye_timeout_seconds"} {

		if items, ok := d.GetOk(x); ok {
			wdts := items.(*schema.Set).List()
			wdtsModel := []models.WorkloadSettingsTimeout{}

			for _, w := range wdts {
				wdtsI := models.WorkloadSettingsTimeout{}
				wdtsMap := w.(map[string]interface{})
				wdtsI.Value = wdtsMap["value"].(int)
				if wdtsMap["scope"].(*schema.Set).Len() > 0 {
					wdtsI.Scope = models.GetHrefs(wdtsMap["scope"].(*schema.Set).List())
				} else {
					wdtsI.Scope = nil
				}
				wdtsModel = append(wdtsModel, wdtsI)
			}
			if x == "workload_disconnected_timeout_seconds" {
				WorkloadSettings.WorkloadDisconnectedTimeoutSeconds = wdtsModel
			} else {
				WorkloadSettings.WorkloadGoodbyeTimeoutSeconds = wdtsModel
			}
		}
	}

	_, err := illumioClient.Update(d.Id(), WorkloadSettings)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioWorkloadSettingsRead(ctx, d, m)
}

func resourceIllumioWorkloadSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "[illumio-core_workload_settings] Ignoring Delete Operation...",
	})

	return diags
}
