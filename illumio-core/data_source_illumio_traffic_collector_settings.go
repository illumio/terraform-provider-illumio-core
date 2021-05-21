// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample API responce
{
  "href": "string",
  "transmission": "string",
  "target": {
    "dst_port": 0,
    "proto": 0,
    "dst_ip": "string"
  },
  "action": "string"
}
*/

func datasourceIllumioTrafficCollectorSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIllumioTrafficCollectorSettingsRead,

		SchemaVersion: version,
		Description:   "Represents Illumio Traffic Collector Settings",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isTrafficCollectorSettingsHref,
				Description:      "URI of traffic collecter settings",
			},
			"transmission": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `transmission type`,
			},
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `action for target traffic`,
			},
			"target": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "target for traffic collector settings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "destination port for target",
						},
						"proto": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "protocol for target",
						},
						"dst_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `single IP address or CIDR`,
						},
					},
				},
			},
		},
	}
}

func datasourceIllumioTrafficCollectorSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(href)

	setIllumioTrafficCollectorSettingState(d, data)

	return diags
}
