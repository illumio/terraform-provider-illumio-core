package illumiocore

import (
	"context"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Sample
/*
{
  "href": "string",
  "created_at": "2021-03-02T02:37:59Z",
  "updated_at": "2021-03-02T02:37:59Z",
  "deleted_at": "2021-03-02T02:37:59Z",
  "created_by": {
    "href": "string"
  },
  "updated_by": {
    "href": "string"
  },
  "deleted_by": {
    "href": "string"
  },
  "update_type": "string",
  "name": "string",
  "description": "string",
  "external_data_set": null,
  "external_data_reference": null,
  "enabled": true,
  "scopes": [
    [
      {
        "label": {
          "href": "string"
        },
        "label_group": {
          "href": "string"
        }
      }
    ]
  ],
  "rules": [
    {
      "href": "string",
      "enabled": true,
      "description": "string",
      "external_data_set": null,
      "external_data_reference": null,
      "ingress_services": [
        {
          "href": "string"
        }
      ],
      "resolve_labels_as": {
        "providers": [
          "workloads"
        ],
        "consumers": [
          "workloads"
        ]
      },
      "sec_connect": true,
      "stateless": true,
      "machine_auth": true,
      "providers": [
        {
          "actors": "ams",
          "label": {
            "href": "string"
          },
          "label_group": {
            "href": "string"
          },
          "workload": {
            "href": "string"
          },
          "virtual_service": {
            "href": "string"
          },
          "virtual_server": {
            "href": "string"
          },
          "ip_list": {
            "href": "string"
          }
        }
      ],
      "consumers": [
        {
          "actors": "ams",
          "label": {
            "href": "string"
          },
          "label_group": {
            "href": "string"
          },
          "workload": {
            "href": "string"
          },
          "virtual_service": {
            "href": "string"
          },
          "ip_list": {
            "href": "string"
          }
        }
      ],
      "consuming_security_principals": [
        {
          "href": "string"
        }
      ],
      "unscoped_consumers": true,
      "update_type": "string"
    }
  ],
  "ip_tables_rules": [
    {
      "href": "string",
      "enabled": true,
      "description": "string",
      "statements": [
        {
          "table_name": "nat",
          "chain_name": "PREROUTING",
          "parameters": "string"
        }
      ],
      "actors": [
        {
          "actors": "string",
          "label": {
            "href": "string"
          },
          "label_group": {
            "href": "string"
          },
          "workload": {
            "href": "string"
          }
        }
      ],
      "ip_version": "4"
    }
  ]
}
*/

func datasourceIllumioRuleSet() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioRuleSetRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Rule Set",
		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of Rule Set",
				ValidateDiagFunc: isRuleSetHref,
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this rule set was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this rule set was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this rule set was deleted",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who originally created this resource",
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
				Description: "Name of Rule Set",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of Rule Set",
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
				Description: "Enabled flag. Determines whether the Rule Set is enabled or not",
			},
			"scopes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "scopes for Rule Set",
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
							Description: "User who originally created this IP Table Rule",
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
		}
	}

	if data.Exists("rules") {
		rls := []map[string]interface{}{}

		rlKeys := []string{
			"href",
			"enabled",
			"description",
			"external_data_set",
			"external_data_reference",
			"sec_connect",
			"stateless",
			"machine_auth",
			"unscoped_consumers",
			"update_type",
			"created_at",
			"updated_at",
			"deleted_at",
			"created_by",
			"updated_by",
			"deleted_by",
		}

		for _, ruleData := range data.S("rules").Children() {
			rl := map[string]interface{}{}
			for _, key := range rlKeys {
				if ruleData.Exists(key) && ruleData.S(key).Data() != nil {
					rl[key] = ruleData.S(key).Data()
				} else {
					rl[key] = nil
				}
			}

			rlaKey := "resolve_labels_as"
			if ruleData.Exists(rlaKey) {
				resLableAs := ruleData.S(rlaKey)

				tm := make(map[string][]interface{})
				tm["providers"] = resLableAs.S("providers").Data().([]interface{})
				tm["consumers"] = resLableAs.S("consumers").Data().([]interface{})

				rl[rlaKey] = []interface{}{tm}
			}

			isKey := "ingress_services"
			if ruleData.Exists(isKey) {
				isKeys := []string{
					"href",
					"proto",
					"port",
					"to_port",
				}

				rl[isKey] = extractMapArray(ruleData.S(isKey), isKeys)
			}

			providersKey := "providers"
			if ruleData.Exists(providersKey) {
				rl[providersKey] = extractDatasourceActors(ruleData.S(providersKey))
			}

			consumerKey := "consumers"
			if ruleData.Exists(consumerKey) {
				rl[consumerKey] = extractDatasourceActors(ruleData.S(consumerKey))
			}

			rls = append(rls, rl)
		}

		d.Set("rules", rls)
	}

	key := "scopes"
	if data.Exists(key) {
		scps := []interface{}{}
		for _, scope := range data.S(key).Children() {
			scopeKeys := []string{"label", "label_group"}
			scps = append(scps, extractMapArray(scope, scopeKeys))
		}

		d.Set(key, scps)
	}

	key = "ip_table_rules"
	if data.Exists(key) {
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
		for _, iptRule := range data.S(key).Children() {

			iptr := extractMap(iptRule, iptrKeys)

			if iptRule.Exists(statKey) {
				iptr[statKey] = extractMapArray(iptRule.S(statKey), statKeys)
			}

			if iptRule.Exists(actorsKey) {
				iptr[actorsKey] = extractDatasourceActors(iptRule.S(actorsKey))
			}

			iptrs = append(iptrs, iptr)
		}

		d.Set(key, iptrs)
	}

	return diagnostics
}

func extractDatasourceActors(data *gabs.Container) []map[string]interface{} {
	actors := []map[string]interface{}{}

	validRuleActors := []string{
		"label",
		"label_group",
		"workload",
		"virtual_service",
		"virtual_server",
		"ip_list",
	}

	for _, actorArray := range data.Children() {

		actor := map[string]interface{}{}
		for k, v := range actorArray.ChildrenMap() {
			if k == "actors" {
				actor[k] = v.Data().(string)
			} else if contains(validRuleActors, k) {
				actor[k] = v.Data().(map[string]interface{})
			}
		}
		actors = append(actors, actor)
	}

	return actors
}
