package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioTrafficCollectorSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIllumioTrafficCollectorSettingsRead,

		SchemaVersion: version,
		Description:   "Represents Illumio Traffic Collector Settings",

		Schema: map[string]*schema.Schema{
			"traffic_collector_setting_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Traffic Collector Settings ID",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of traffic collecter settings",
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
							Description: `single ip address or CIDR`,
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

	orgID := pConfig.OrgID
	refID := d.Get("traffic_collector_setting_id").(string)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/settings/traffic_collector/%v", orgID, refID), nil)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(refID)

	setIllumioTrafficCollectorSettingState(d, data)

	return diags
}
