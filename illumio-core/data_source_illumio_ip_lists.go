// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceIllumioIPLists() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioIPListsRead,
		SchemaVersion: 1,
		Description:   "Represents Illumio IP Lists",

		Schema: map[string]*schema.Schema{
			"pversion": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "draft",
				ValidateDiagFunc: isValidPversion(),
				Description:      `pversion of the security policy. Allowed values are "draft", "active", and numbers greater than 0. Default value: "draft"`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of IP list(s) to return. Supports partial matches",
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A unique identifier within the external data source",
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The data source from which a resource originates",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP lists matching FQDN. Supports partial matches",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address matching IP list(s) to return. Supports partial matches",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of IP Lists to return. The integer should be a non-zero positive integer",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of IP list(s) to return. Supports partial matches",
			},
			"match_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      PARTIAL_MATCH,
				ValidateFunc: validation.StringInSlice([]string{PARTIAL_MATCH, EXACT_MATCH}, true),
				Description:  `Indicates whether to return all partially-matching names or only exact matches. Allowed values are "partial" and "exact". Default value: "partial"`,
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of IP Lists",
				Elem:        ipListDataSourceSchema(false),
			},
		},
	}
}

func ipListDataSourceSchema(hrefRequired bool) *schema.Resource {
	s := map[string]*schema.Schema{
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
	}

	hrefSchema := &schema.Schema{
		Type:        schema.TypeString,
		Description: "URI of this IP List",
	}

	if hrefRequired {
		hrefSchema.Required = true
	} else {
		hrefSchema.Computed = true
	}
	s["href"] = hrefSchema

	return &schema.Resource{Schema: s}
}

func datasourceIllumioIPListsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID
	pversion := d.Get("pversion").(string)

	paramKeys := []string{
		"description",
		"external_data_reference",
		"external_data_set",
		"fqdn",
		"ip_address",
		"max_results",
		"name",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/sec_policy/%v/ip_lists", orgID, pversion), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}
	keys := []string{
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
	}
	for _, child := range data.Children() {
		// if exact matching is enabled, skip the object if it's a partial match
		if d.Get("match_type").(string) == EXACT_MATCH {
			if !isExactMatch("name", d, child) {
				continue
			}
		}

		m := extractMap(child, keys)

		if child.Exists("ip_ranges") {
			ip_ranges := child.S("ip_ranges")
			ip_rangeI := []map[string]interface{}{}

			for _, ip := range ip_ranges.Children() {
				ip_rangeI = append(ip_rangeI, extractMap(ip, []string{"description", "from_ip", "to_ip", "exclusion"}))
			}

			m["ip_ranges"] = ip_rangeI
		} else {
			m["ip_ranges"] = nil
		}

		if child.Exists("fqdns") {
			fqdns := child.S("fqdns")
			fqdnI := []map[string]interface{}{}

			for _, ip := range fqdns.Children() {
				fqdnI = append(fqdnI, extractMap(ip, []string{"fqdn", "description"}))
			}

			m["fqdns"] = fqdnI
		} else {
			m["fqdns"] = nil
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
