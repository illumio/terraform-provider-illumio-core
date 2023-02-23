// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioWorkloadInterfaces() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioWorkloadInterfacesRead,
		SchemaVersion: 1,
		Description:   "Represents Illumio Workload Interfaces",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of Workload Interfaces",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Workload Interface",
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
							Description: "The IP Address to assign to this interface				",
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
				},
			},
			"workload_href": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URI of Workload",
			},
		},
	}
}

func dataSourceIllumioWorkloadInterfacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	wHref := d.Get("workload_href").(string)

	// Does not support Async call
	_, data, err := illumioClient.Get(fmt.Sprintf("%v/interfaces", wHref), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(wHref)))

	dataMap := []map[string]interface{}{}

	for _, child := range data.Children() {
		m := map[string]interface{}{}

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
			if child.Exists(key) {
				m[key] = child.S(key).Data()
			} else {
				m[key] = nil
			}
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
