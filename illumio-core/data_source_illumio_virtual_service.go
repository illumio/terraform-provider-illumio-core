package illumiocore

import (
	"context"
	"fmt"

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
			"virtual_service_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the virtual service",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the virtual service",
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
	orgID := pConfig.OrgID

	vsid := d.Get("virtual_service_id").(string)
	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%d/sec_policy/draft/virtual_services/%s", orgID, vsid), nil)
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
		"service_ports",
		"pce_fqdn",
		"service",
		"labels",
		"ip_overrides",
		"apply_to",
		"caps",
		"service_addresses",
	} {

		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		}
	}
	return diagnostics
}
