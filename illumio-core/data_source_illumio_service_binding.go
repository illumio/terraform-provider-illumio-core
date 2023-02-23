// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioServiceBinding() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioServiceBindingRead,
		SchemaVersion: 1,
		Description:   "Represents Illumio Service Binding",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isServiceBindingHref,
				Description:      "URI of the Service Binding",
			},
			"bound_service": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Bound service href",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"virtual_service": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Virtual service href",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workload URI",
						},
					},
				},
			},
			"workload": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Workload Object for Service Bindings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workload URI",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workload Name",
						},
						"hostname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workload Hostname",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether the workload is deleted",
						},
					},
				},
			},
			"container_workload": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Container Workload href",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workload URI",
						},
					},
				},
			},
			"port_overrides": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Port Overrides for Service Bindings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port Number in the original service which to override (integer 0-65535). Starting port when specifying a range",
						},
						"proto": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Transport protocol in the original service which to override",
						},
						"new_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Overriding port number (or starting point when specifying a range)",
						},
						"new_to_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Overriding port range ending port",
						},
					},
				},
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External Data Set identifier",
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External Data reference identifier",
			},
		},
	}
}

func datasourceIllumioServiceBindingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		"bound_service",
		"external_data_set",
		"external_data_reference",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	key := "virtual_service"
	if data.Exists(key) {
		d.Set(key, []interface{}{extractMap(data.S(key), []string{"href"})})
	} else {
		d.Set(key, nil)
	}

	key = "container_workload"
	if data.Exists(key) {
		d.Set(key, []interface{}{extractMap(data.S(key), []string{"href"})})
	} else {
		d.Set(key, nil)
	}

	key = "port_overrides"
	if data.Exists(key) {
		poS := data.S(key)
		poI := []map[string]interface{}{}

		for _, po := range poS.Children() {
			poI = append(poI, extractMap(po, []string{"port", "proto", "new_port", "new_to_port"}))
		}

		d.Set(key, poI)
	} else {
		d.Set(key, nil)
	}

	key = "workload"
	if data.Exists(key) {
		d.Set(key, []interface{}{extractMap(data.S(key), []string{"href", "name", "hostname", "deleted"})})
	} else {
		d.Set(key, nil)
	}

	return diagnostics
}
