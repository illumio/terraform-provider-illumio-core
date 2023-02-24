// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioRuleSets() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioRuleSetsRead,
		SchemaVersion: 1,
		Description:   "Represents Illumio Rulesets",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of Rulesets",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Ruleset",
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
							Elem: &schema.Schema{
								Type: schema.TypeList,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"label": {
											Type:        schema.TypeMap,
											Computed:    true,
											Description: "Href of Label",
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
										"label_group": {
											Type:        schema.TypeMap,
											Computed:    true,
											Description: "Href of Label Group",
											Elem: &schema.Schema{
												Type: schema.TypeString,
											},
										},
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
										Description: "statements for IP Tables Rule",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"table_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "table name of statement",
												},
												"chain_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "chain name of statement",
												},
												"parameters": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "parameters of statement",
												},
											},
										},
									},
									"actors": {
										Type:        schema.TypeList,
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
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "Href of Label",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"label_group": {
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "Href of Label Group",
													Elem:        &schema.Schema{Type: schema.TypeString},
												},
												"workload": {
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "Href of Workload",
													Elem:        &schema.Schema{Type: schema.TypeString},
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
				},
			},
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
				Description: "Description of Ruleset(s) to return. Supports partial matches",
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
			"labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of lists of label URIs, encoded as a JSON string",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of Rulesets to return. The integer should be a non-zero positive integer",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of Ruleset(s) to return. Supports partial matches",
			},
		},
	}
}

func datasourceIllumioRuleSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	pversion := d.Get("pversion").(string)

	href := fmt.Sprintf("/orgs/%v/sec_policy/%v/rule_sets", illumioClient.OrgID, pversion)

	paramKeys := []string{
		"description",
		"external_data_reference",
		"external_data_set",
		"labels",
		"max_results",
		"name",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.Get(href, &params)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	rsKeys := []string{
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
		"enabled",
	}

	rss := []map[string]interface{}{}

	for _, ruleSet := range data.Children() {
		rs := extractMap(ruleSet, rsKeys)

		key := "rules"
		if ruleSet.Exists(key) {
			rs[key] = extractRules(data.S("rules"))
		}

		key = "scopes"
		if ruleSet.Exists(key) {
			scps := []interface{}{}
			scopeKeys := []string{"label", "label_group"}
			for _, scope := range ruleSet.S(key).Children() {
				scps = append(scps, extractMapArray(scope, scopeKeys))
			}

			rs[key] = scps
		}

		key = "ip_table_rules"
		if ruleSet.Exists(key) {
			iptrKeys := []string{
				"href",
				"enabled",
				"description",
				"ip_version",
				"update_type",
				"created_at",
				"updated_at",
				"deleted_at",
				"created_by",
				"updated_by",
				"deleted_by",
			}

			statKey := "statements"
			statKeys := []string{
				"table_name",
				"chain_name",
				"parameters",
			}

			actorsKey := "actors"

			iptrs := []map[string]interface{}{}
			for _, iptRule := range ruleSet.S(key).Children() {

				iptr := extractMap(iptRule, iptrKeys)

				if iptRule.Exists(statKey) {
					iptr[statKey] = extractMapArray(iptRule.S(statKey), statKeys)
				}

				if iptRule.Exists(actorsKey) {
					iptr[actorsKey] = extractRuleActors(iptRule.S(actorsKey))
				}

				iptrs = append(iptrs, iptr)
			}

			rs[key] = iptrs
		}

		rss = append(rss, rs)
	}

	d.Set("items", rss)

	return diagnostics
}
