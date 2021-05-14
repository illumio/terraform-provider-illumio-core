package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample of API respones
{
  "href": "string",
  "name": "string",
  "description": "string",
  "description_url": "string",
  "process_name": "string",
  "service_ports": [
    {
      "port": 0,
      "to_port": 0,
      "proto": 0,
      "icmp_type": 0,
      "icmp_code": 0
    }
  ],
  "windows_services": [
    {
      "service_name": "string",
      "process_name": "string",
      "port": 0,
      "to_port": 0,
      "proto": 0,
      "icmp_type": 0,
      "icmp_code": 0
    }
  ],
  "external_data_set": "string",
  "external_data_reference": "string",
  "created_at": "1970-01-01T00:00:00.000Z",
  "updated_at": "1970-01-01T00:00:00.000Z",
  "deleted_at": "1970-01-01T00:00:00.000Z",
  "created_by": {
    "href": "string"
  },
  "updated_by": {
    "href": "string"
  },
  "deleted_by": {
    "href": "string"
  },
  "update_type": "string"
}
*/

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
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port Number ( the starting port when specifying a range)",
						},
						"to_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "High end of port range",
						},
						"proto": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Transport protocol",
						},
						"icmp_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ICMP Type",
						},
						"icmp_code": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ICMP Code",
						},
					},
				},
				Description: "Service ports of Illumio Service",
			},
			"windows_services": {
				Type:     schema.TypeList,
				Computed: true,
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
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port Number, also the starting port when specifying a range",
						},
						"to_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "High end of port range",
						},
						"proto": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Transport protocol",
						},
						"icmp_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ICMP Type",
						},
						"icmp_code": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ICMP Code",
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
				Description: "User who originally created this Service",
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

	if data.Exists("service_ports") {
		sps := data.S("service_ports")
		spI := []map[string]interface{}{}

		for _, sp := range sps.Children() {
			spI = append(spI, extractMap(sp, []string{"port", "to_port", "proto", "icmp_type", "icmp_code"}))
		}
		d.Set("service_ports", spI)
	} else {
		d.Set("service_ports", nil)
	}

	if data.Exists("windows_services") {
		wss := data.S("windows_services")
		wsI := []map[string]interface{}{}

		for _, ws := range wss.Children() {
			wsI = append(wsI, extractMap(ws, []string{"port", "to_port", "proto", "icmp_type", "icmp_code", "service_name", "process_name"}))
		}
		d.Set("windows_services", wsI)
	} else {
		d.Set("windows_services", nil)
	}

	return diagnostics
}
