// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioService() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioServiceRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Service",
		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of Service",
				ValidateDiagFunc: isServiceHref,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The short friendly name of the service",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Long Description of Service",
			},
			"description_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description URL Read-only to prevent XSS attacks",
			},
			"process_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The process name",
			},
			"service_ports": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port Number ( the starting port when specifying a range)",
						},
						"to_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "High end of port range",
						},
						"proto": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Transport protocol",
						},
						"icmp_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ICMP Type",
						},
						"icmp_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ICMP Code",
						},
					},
				},
				Description: "Service ports of Illumio Service",
			},
			"windows_services": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "windows_services for services",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of Windows Service",
						},
						"process_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of running process",
						},
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port Number, also the starting port when specifying a range",
						},
						"to_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "High end of port range",
						},
						"proto": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Transport protocol",
						},
						"icmp_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ICMP Type",
						},
						"icmp_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ICMP Code",
						},
					},
				},
			},
			"windows_egress_services": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Windows Egress services",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of Windows Service",
						},
						"process_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of running process",
						},
					},
				},
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External data set identifier",
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External data reference identifier",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Service was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Service was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Service was deleted",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who created this Service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who last updated this Service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deleted_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who deleted this Service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of update",
			},
		},
	}
}

func dataSourceIllumioServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))

	for _, key := range []string{
		"href",
		"name",
		"description",
		"description_url",
		"process_name",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"update_type",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	key := "service_ports"
	if data.Exists(key) {
		d.Set(key, extractServicePorts(data))
	} else {
		d.Set(key, nil)
	}

	key = "windows_services"
	if data.Exists(key) {
		d.Set(key, extractWindowsServices(data))
	} else {
		d.Set(key, nil)
	}

	key = "windows_egress_services"
	if data.Exists(key) {
		d.Set(key, extractWindowsEgressServices(data))
	} else {
		d.Set(key, nil)
	}

	return diagnostics
}
