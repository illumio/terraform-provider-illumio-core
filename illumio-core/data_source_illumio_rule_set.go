// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioRuleSet() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioRuleSetRead,
		SchemaVersion: 1,
		Description:   "Represents Illumio Ruleset",
		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of Ruleset",
				ValidateDiagFunc: isRuleSetHref,
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this ruleset was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this ruleset was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this ruleset was deleted",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who created this resource",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who last updated this resource",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deleted_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who deleted this resource",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of update",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of Ruleset",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of Ruleset",
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
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enabled flag. Determines whether the Ruleset is enabled or not",
			},
			"scopes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "scopes for Ruleset",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclusion": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Boolean to specify whether or not the scope is an exclusion",
						},
						"label": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Label scope",
							Elem:        labelOptionalKeyValue(false),
						},
						"label_group": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Label Group scope",
							Elem:        labelGroupOptionalKeyValue(false),
						},
					},
				},
			},
			"rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of Security Rules",
				Elem: &schema.Resource{
					Schema: securityRuleDatasourceSchema(false),
				},
			},
			"ip_tables_rules": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of IP Tables Rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of the Ip Tables Rule",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enabled flag. Determines whether this IP Tables Rule is enabled or not",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the Ip Tables Rule",
						},
						"statements": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "statements for in this IP Tables Rule",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"table_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the table",
									},
									"chain_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Chain name for statement",
									},
									"parameters": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameters of statement",
									},
								},
							},
						},
						"actors": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "actors for IP Table Rule",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"actors": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "actors for IP table Rule actors",
									},
									"label": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Href of Label",
										Elem:        labelOptionalKeyValue(false),
									},
									"label_group": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Label Group actor",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"workload": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Workload actor",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"ip_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP version for the rules to be applied to",
						},
						"update_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of update for IP Table Rule",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this IP Table Rule was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this IP Table Rule was last updated",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this IP Table Rule was deleted",
						},
						"created_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who created this IP Table Rule",
						},
						"updated_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who last updated this IP Table Rule",
						},
						"deleted_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who deleted this IP Table Rule",
						},
					},
				},
			},
		},
	}
}

func datasourceIllumioRuleSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(href)

	for _, key := range []string{
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
		"enabled",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("rules") {
		d.Set("rules", extractRules(data.S("rules")))
	}

	key := "scopes"
	if data.Exists(key) {
		d.Set(key, extractResourceScopes(data.S(key)))
	} else {
		d.Set(key, nil)
	}

	key = "ip_tables_rules"
	if data.Exists(key) {
		d.Set(key, extractResourceRuleSetIPTablesRules(data.S(key)))
	} else {
		d.Set(key, nil)
	}

	return diagnostics
}
