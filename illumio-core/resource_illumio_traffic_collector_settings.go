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

var (
	tcsTransmissionValidValues = []string{
		"broadcast",
		"multicast",
	}
	tcsActionValidValues = []string{"drop", "aggregate"}
	protoTCSValidValues  = []int{6, 17, 1, 58}
)

func resourceIllumioTrafficCollectorSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioTrafficCollectorSettingsCreate,
		ReadContext:   resourceIllumioTrafficCollectorSettingsRead,
		UpdateContext: resourceIllumioTrafficCollectorSettingsUpdate,
		DeleteContext: resourceIllumioTrafficCollectorSettingsDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio Traffic Collector Settings",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of traffic collecter settings",
			},
			"transmission": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `transmission type. Allowed values are "broadcast" and "multicast"`,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(tcsTransmissionValidValues, false),
				),
			},
			"action": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `action for target traffic. Allowed values are "drop" or "aggregate"`,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(tcsActionValidValues, false),
				),
			},
			"target": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: `target for traffic collector settings. Required if value of action is "drop"`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     -1,
							Description: "destination port for target. Allowed range is -1 to 65535. Default value: -1",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.Any(validation.IntBetween(-1, 65535)),
							),
						},
						"proto": {
							Type:     schema.TypeInt,
							Required: true,
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.IntInSlice(protoTCSValidValues),
							),
							Description: "protocol for target. Allowed values are 6 (TCP), 17 (UDP), 1 (ICMP) and 58 (ICMPv6)",
						},
						"dst_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "0.0.0.0/0",
							Description: `single IP address or CIDR. Default value: "0.0.0.0/0"`,
							ValidateDiagFunc: validation.ToDiagFunc(validation.Any(
								validation.IsIPAddress,
								validation.IsCIDR,
							)),
						},
					},
				},
			},
		},
	}
}

func resourceIllumioTrafficCollectorSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	orgID := pConfig.OrgID

	tcs := expandIllumioTrafficCollectorSettings(d)

	if tcs.Action == "drop" && tcs.Target == nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "[illumio-core_traffic_collector_settings] Expected target block",
				Detail:   "target block must be specified if action is \"drop\"",
			},
		}
	} else if tcs.Action == "aggregate" && tcs.Target != nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "[illumio-core_traffic_collector_settings] target block not allowed",
				Detail:   "target block is not allowed when action is \"aggregate\"",
			},
		}
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%v/settings/traffic_collector", orgID), tcs)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.S("href").Data().(string))

	return resourceIllumioTrafficCollectorSettingsRead(ctx, d, m)
}

func expandIllumioTrafficCollectorSettings(d *schema.ResourceData) *models.TrafficCollectorSettings {
	return &models.TrafficCollectorSettings{
		Transmission: d.Get("transmission").(string),
		Action:       d.Get("action").(string),
		Target:       expandIllumioTrafficCollectorSettingsTarget(d.Get("target")),
	}
}

func expandIllumioTrafficCollectorSettingsTarget(v interface{}) *models.TrafficCollectorSettingsTarget {
	if v != nil && len(v.([]interface{})) > 0 {
		t := v.([]interface{})[0].(map[string]interface{})
		return &models.TrafficCollectorSettingsTarget{
			DstPort: t["dst_port"].(int),
			Proto:   t["proto"].(int),
			DstIP:   t["dst_ip"].(string),
		}
	} else {
		return nil
	}
}

func resourceIllumioTrafficCollectorSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	setIllumioTrafficCollectorSettingState(d, data)

	return diags
}

func setIllumioTrafficCollectorSettingState(d *schema.ResourceData, data *gabs.Container) {
	for _, k := range []string{
		"href",
		"action",
	} {
		if data.Exists(k) {
			d.Set(k, data.S(k).Data())
		} else {
			d.Set(k, nil)
		}
	}

	key := "transmission"
	if data.Exists(key) {
		switch data.S(key).Data().(string) {
		case "B":
			d.Set(key, "broadcast")
		case "M":
			d.Set(key, "multicast")
		}
	} else {
		d.Set(key, nil)
	}

	key = "target"
	if data.Exists(key) {
		targetKeys := []string{
			"dst_port",
			"proto",
			"dst_ip",
		}
		d.Set(key, []interface{}{
			extractMap(data.S(key), targetKeys),
		})
	} else {
		d.Set(key, nil)
	}
}

func resourceIllumioTrafficCollectorSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	if d.HasChange("action") {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "[illumio-core_traffic_collector_settings] Unexpectade update: action attribute is not updatable",
			},
		}
	}

	tcs := &models.TrafficCollectorSettings{
		Transmission: d.Get("transmission").(string),
	}

	if v, ok := d.GetOk("target"); ok {
		tcs.Target = expandIllumioTrafficCollectorSettingsTarget(v)
	}

	action := d.Get("action").(string)
	if action == "aggregate" && tcs.Target != nil {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "[illumio-core_traffic_collector_settings] target block not allowed",
				Detail:   "target block is not allowed when action is \"aggregate\"",
			},
		}
	}

	_, err := illumioClient.Update(d.Id(), tcs)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioTrafficCollectorSettingsRead(ctx, d, m)
}

func resourceIllumioTrafficCollectorSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	_, err := illumioClient.Delete(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}
