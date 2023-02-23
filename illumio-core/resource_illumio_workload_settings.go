// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

const WORKLOAD_SETTINGS_VEN_TYPE_DEFAULT = "server"

var validVENTypes = []string{"server", "endpoint"}

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
						"ven_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Default:          WORKLOAD_SETTINGS_VEN_TYPE_DEFAULT,
							Description:      `The VEN type that this property is applicable to. Must be "server" or "endpoint". An empty or missing value will default to "server" on the PCE`,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validVENTypes, true)),
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
						"ven_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Default:          WORKLOAD_SETTINGS_VEN_TYPE_DEFAULT,
							Description:      `The VEN type that this property is applicable to. Must be "server" or "endpoint". An empty or missing value will default to "server" on the PCE`,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validVENTypes, true)),
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
		Summary:  "Please use terraform import",
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

	for _, k := range []string{
		"workload_disconnected_timeout_seconds",
		"workload_goodbye_timeout_seconds",
	} {
		d.Set(k, extractWorkloadSettingsTimeout(data, k))
	}

	return diags
}

func extractWorkloadSettingsTimeout(data *gabs.Container, key string) []map[string]interface{} {
	if data.Exists(key) {
		d := data.S(key)
		m := []map[string]interface{}{}

		for _, v := range d.Children() {
			vm := extractMap(v, []string{"value", "ven_type"})
			if v.Exists("scope") {
				vm["scope"] = extractMapArray(v.S("scope"), []string{"href"})
			} else {
				vm["scope"] = []map[string]interface{}{}
			}
			m = append(m, vm)
		}

		return m
	}

	return nil
}

func resourceIllumioWorkloadSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	workloadSettings := &models.WorkloadSettings{
		WorkloadDisconnectedTimeoutSeconds: expandWorkloadSettingsTimeout(d, "workload_disconnected_timeout_seconds"),
		WorkloadGoodbyeTimeoutSeconds:      expandWorkloadSettingsTimeout(d, "workload_goodbye_timeout_seconds"),
	}

	_, err := illumioClient.Update(d.Id(), workloadSettings)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioWorkloadSettingsRead(ctx, d, m)
}

func expandWorkloadSettingsTimeout(d *schema.ResourceData, key string) []models.WorkloadSettingsTimeout {
	var timeoutSettings []models.WorkloadSettingsTimeout

	if items, ok := d.GetOk(key); ok {
		wts := items.(*schema.Set).List()

		for _, w := range wts {
			m := models.WorkloadSettingsTimeout{}
			wMap := w.(map[string]interface{})
			m.Value = wMap["value"].(int)

			if wMap["scope"].(*schema.Set).Len() > 0 {
				m.Scope = models.GetHrefs(wMap["scope"].(*schema.Set).List())
			} else {
				m.Scope = []models.Href{}
			}

			if venType, ok := wMap["ven_type"].(string); ok {
				m.VENType = &venType
			}

			timeoutSettings = append(timeoutSettings, m)
		}
	}

	return timeoutSettings
}

func resourceIllumioWorkloadSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "[illumio-core_workload_settings] Ignoring Delete Operation.",
	})

	return diags
}
