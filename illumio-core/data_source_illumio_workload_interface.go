// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample
{
 "href": "string",
  "name": "string",
  "link_state": "string",
  "address": "string",
  "cidr_block": 0,
  "default_gateway_address": "string",
  "network": {
    "href": "string"
  },
  "network_detection_mode": "string",
  "friendly_name": "string"
}
*/

func datasourceIllumioWorkloadInterface() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioWorkloadInterfaceRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Workload Interface",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isWorkloadInterfaceHref,
				Description:      "URI of the Workload Interface",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Workload Interface",
			},
			"link_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Link State for Workload Interface",
			},
			"address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP Address to assign to this interface",
			},
			"cidr_block": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of bits in the subnet /24 is 255.255.255.0",
			},
			"default_gateway_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IP Address of the default gateway",
			},
			"network": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Network for the Workload Interface. ",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"loopback": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Loopback for Workload Interface",
			},
			"network_detection_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network Detection Mode for Workload Interface",
			},
			"friendly_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User-friendly name for Workload Interface",
			},
		},
	}
}

func datasourceIllumioWorkloadInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	// orgID := pConfig.OrgID
	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
	for _, key := range []string{
		"href",
		"name",
		"link_state",
		"address",
		"cidr_block",
		"default_gateway_address",
		"network",
		"loopback",
		"network_detection_mode",
		"friendly_name",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}
	return diagnostics
}
