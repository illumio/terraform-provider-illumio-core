// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioEnforcementBoundary() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioEnforcementBoundaryRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Enforcement Boundary",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of this Enforcement Boundary",
				ValidateDiagFunc: isEnforcementBoundaryHref,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Enforcement Boundary",
			},
			"ingress_services": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of Ingress Service. Only one of the {\"href\"} or {\"proto\", \"port\", \"to_port\"} parameter combination is allowed",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protocol number",
						},
						"port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port number used with protocol or starting port when specifying a range",
						},
						"to_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Upper end of port range",
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Ingress Service",
						},
					},
				},
			},
			"providers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "providers for Enforcement Boundary. Only one actor can be specified in one providers block",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actors": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "actors for providers",
						},
						"label": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Href of Label",
							Elem:        hrefSchemaRequired("Label", isLabelHref),
						},
						"label_group": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Href of Label Group",
							Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
						},
						"ip_list": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Href of IP List",
							Elem:        hrefSchemaRequired("IP List", isIPListHref),
						},
					},
				},
			},
			"consumers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Consumers for Enforcement Boundary. Only one actor can be specified in one consumer block",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actors": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "actors for consumers parameter. Allowed values are \"ams\" and \"container_host\"",
						},
						"label": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Href of Label",
							Elem:        hrefSchemaRequired("Label", isLabelHref),
						},
						"label_group": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Href of Label Group",
							Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
						},
						"ip_list": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Href of IP List",
							Elem:        hrefSchemaRequired("IP List", isIPListHref),
						},
					},
				},
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Enforcement Boundary was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Enforcement Boundary was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Enforcement Boundary was last deleted",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this Enforcement Boundary",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this Enforcement Boundary",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last deleted this Enforcement Boundary",
			},
			"caps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User permissions for the object",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func datasourceIllumioEnforcementBoundaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"caps",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("ingress_services") {
		ingServs := data.S("ingress_services").Data().([]interface{})
		iss := []map[string]interface{}{}

		for _, ingServ := range ingServs {
			is := ingServ.(map[string]interface{})

			for k, v := range ingServ.(map[string]interface{}) {
				if k == "href" {
					is[k] = v
				} else {
					is[k] = strconv.Itoa(int(v.(float64)))
				}

				iss = append(iss, is)
			}
		}

		d.Set("ingress_services", iss)
	} else {
		d.Set("ingress_services", nil)
	}

	d.Set("providers", extractEBActors(data.S("providers")))
	d.Set("consumers", extractEBActors(data.S("consumers")))

	return diagnostics
}
