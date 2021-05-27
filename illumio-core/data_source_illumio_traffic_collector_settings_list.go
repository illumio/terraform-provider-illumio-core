// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample API responce
[
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
]
*/

func datasourceIllumioTrafficCollectorSettingsList() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIllumioTrafficCollectorSettingsListRead,

		SchemaVersion: version,
		Description:   "Represents List of Illumio Traffic Collector Settings ",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of Traffic Collector Setting hrefs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Traffic Collector Setting",
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
				},
			},
		},
	}
}

func datasourceIllumioTrafficCollectorSettingsListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := fmt.Sprintf("/orgs/%v/settings/traffic_collector", pConfig.OrgID)

	// Does not support Async Call
	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(href)

	tcsKeys := []string{
		"href",
		"action",
	}

	tcss := []map[string]interface{}{}

	for _, trafficCS := range data.Children() {
		tcs := extractMap(trafficCS, tcsKeys)

		key := "transmission"
		if trafficCS.Exists(key) {
			switch trafficCS.S(key).Data().(string) {
			case "B":
				tcs[key] = "broadcast"
			case "M":
				tcs[key] = "multicast"
			}
		} else {
			tcs[key] = nil
		}

		key = "target"
		if trafficCS.Exists(key) {
			targetKeys := []string{
				"dst_port",
				"proto",
				"dst_ip",
			}
			tcs[key] = []interface{}{
				extractMap(trafficCS.S(key), targetKeys),
			}
		} else {
			tcs[key] = nil
		}

		tcss = append(tcss, tcs)
	}

	d.Set("items", tcss)

	return diags
}
