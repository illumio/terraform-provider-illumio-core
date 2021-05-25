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
  "Description": "sring",
  "external_data_set": nul,
  "external_data_reference":null,
  "ip_ranges": [
    {
     "Description": "string",
      "from_ip": "string",
      "to_ip": "string",
      "exclusion": true
    }
  ],
  "fdns": [
    {
     "fqdn": "string",
      "Description": "sring"
    }
  ],
  "ceated_at": "2021-03-02T02:37:59Z",
  "updated_at": "2021-03-02T02:37:59Z",
  "deleted_at": "2021-03-02T02:37:59Z",
  "created_by": {
    "href": "strig"
  },
  "udated_by": {
    "href": "strig"
  },
  "dleted_by": {
    "href": "strig"
  }
}
*/

func datasourceIllumioIPList() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioIPListRead,
		SchemaVersion: version,
		Description:   "Represents Illumio IP List",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of the IP List",
				ValidateDiagFunc: isIPListHref,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the IP List",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the IP List",
			},
			"ip_ranges": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "IP addresses or ranges",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the IP Range",
						},
						"from_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address or a low end of IP range. Might be specified with CIDR notation",
						},
						"to_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "High end of an IP range",
						},
						"exclusion": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this IP address is an exclusion. Exclusions must be a strict subset of inclusive IP addresses",
						},
					},
				},
			},
			"fqdns": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of Fully Qualified Domain Names",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Fully Qualified Domain Name",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of FQDN",
						},
					},
				},
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The data source from which a resource originates",
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique identifier within the external data source",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IP List was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IP List was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IP List was deleted",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who created this IP List",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who last updated this IP List",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deleted_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who deleted this IP List",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func datasourceIllumioIPListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		"description",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"deleted_at",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("ip_ranges") {
		ip_ranges := data.S("ip_ranges")
		ip_rangeI := []map[string]interface{}{}

		for _, ip := range ip_ranges.Children() {
			ip_rangeI = append(ip_rangeI, extractMap(ip, []string{"description", "from_ip", "to_ip", "exclusion"}))
		}
		d.Set("ip_ranges", ip_rangeI)
	} else {
		d.Set("ip_ranges", nil)
	}

	if data.Exists("fqdns") {
		fqdns := data.S("fqdns")
		fqdnI := []map[string]interface{}{}

		for _, ip := range fqdns.Children() {
			fqdnI = append(fqdnI, extractMap(ip, []string{"fqdn", "description"}))
		}

		d.Set("fqdns", fqdnI)
	} else {
		d.Set("fqdns", nil)
	}

	return diagnostics
}
