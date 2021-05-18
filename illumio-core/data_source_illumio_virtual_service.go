// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Sample
/*
{
  "name": "string",
  "description": "string",
  "external_data_set": null,
  "external_data_reference": null,
  "labels": [
    {
      "href": "string"
    }
  ],
  "service_ports": [
    {
      "port": 0,
      "to_port": 0,
      "proto": 0
    }
  ],
  "service": {},
  "apply_to": "host_only",
  "ip_overrides": [
    "string"
  ],
  "service_addresses": [
    {}
  ]
}
*/

func datasourceIllumioVirtualService() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioVirtualServiceRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Virtual Services",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isVirtualServiceHref,
				Description:      "URI of the virtual service",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the virtual service",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The long description of this virtual service",
			},
			"pce_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PCE FQDN for this container cluster. Used in Supercluster only",
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The data source from which a resource originates",
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unque identifier within the external data source",
			},
			"labels": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Assigned labels",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of label",
						},
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Key in key-value pair",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value in key-value pair",
						},
					},
				},
			},
			"service_ports": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Service Ports",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port Number. Also starting port when specifying port range",
						},
						"to_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specify port or port range to filter results.",
						},
						"proto": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Transport Protocol",
						},
					},
				},
			},
			"service": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "URI of associated service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of service",
						},
					},
				},
			},
			"apply_to": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Firewall rule target for workloads bound to this virtual service: host_only or internal_bridge_network",
			},
			"ip_overrides": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of IPs or CIDRs as IP overrides",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"service_addresses": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Serivce Addresses",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "FQDN to assign to the virtual service",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description for given fqdn",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port Number. Also the starting port when specifying a range",
						},
						"ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address to assign to the virtual service",
						},
						"network": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Network URI for this IP address",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"caps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "CAPS",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update Type",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Virtual Service was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Virtual Service was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Virtual Service was deleted",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who originally created this Virtual Service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who last updated this Virtual Service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this Virtual Service",
			},
		},
	}
}

func dataSourceIllumioVirtualServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"update_type",
		"name",
		"description",
		"external_data_set",
		"external_data_reference",
		"pce_fqdn",
		"labels",
		"ip_overrides",
		"apply_to",
		"caps",
	} {

		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	key := "service"
	if data.Exists(key) {
		l := []map[string]string{}
		l = append(l, map[string]string{"href": data.S(key, "href").Data().(string)})
		d.Set(key, l)
	} else {
		d.Set(key, nil)
	}

	key = "service_addresses"
	if data.Exists(key) {
		l := []map[string]interface{}{}
		for _, child := range data.S(key).Children() {
			val := map[string]interface{}{}

			if v := child.S("fqdn").Data(); v != nil {
				val["fqdn"] = v.(string)
			}
			if v := child.S("description").Data(); v != nil {
				val["description"] = v.(string)
			}
			if v := child.S("port").Data(); v != nil {
				val["port"] = v
			}
			if v := child.S("ip").Data(); v != nil {
				val["ip"] = v.(string)
			}
			if v := child.S("network").Data(); v != nil {
				val["network"] = v
			}
			l = append(l, val)
		}
		d.Set(key, l)
	} else {
		d.Set(key, nil)
	}

	key = "service_ports"
	if data.Exists(key) {
		sps := []map[string]interface{}{}

		for _, serPort := range data.S(key).Children() {
			sp := map[string]interface{}{}

			for k, v := range serPort.ChildrenMap() {
				if k == "proto" {
					sp[k] = v.Data()
				} else if k == "port" || k == "to_port" {
					sp[k] = v.Data()
				}

				sps = append(sps, sp)
			}
		}

		d.Set(key, sps)
	} else {
		d.Set(key, nil)
	}

	return diagnostics
}
