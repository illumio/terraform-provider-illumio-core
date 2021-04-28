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
			"rule_set_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Numerical ID of rule set",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of Rule Set",
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
				Description: "Enabled flag. Determines wheter the Rule Set is enabled or not",
			},
			"scopes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "scopes for Rule Set",
				Elem: &schema.Schema{
					Type:     schema.TypeList,
					MinItems: 1,
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
					Schema: securityRuleDatasourceBaseSchemaMap(),
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
										Description: "Href of Worklaod",
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
	orgID := pConfig.OrgID

	secrutiyRuleID := d.Get("rule_set_id").(int)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/sec_policy/draft/rule_sets/%v", orgID, secrutiyRuleID), nil)
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
		"enabled",
		"ip_tables_rules",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		}
	}

	if data.Exists("rules") {
		rls := []map[string]interface{}{}

		for _, data := range data.S("rules").Children() {
			r := map[string]interface{}{}
			for _, key := range []string{
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
				"ingress_services",
				"consumers",
				"providers",
			} {
				if data.Exists(key) {
					r[key] = data.S(key).Data()
				}
			}

			if data.Exists("resolve_labels_as") {
				rs := data.S("resolve_labels_as")

				tm := make(map[string][]interface{})
				tm["providers"] = rs.S("providers").Data().([]interface{})
				tm["consumers"] = rs.S("consumers").Data().([]interface{})

				r["resolve_labels_as"] = []interface{}{tm}
			}

			rls = append(rls, r)
		}

		d.Set("rules", rls)
	}

	if data.Exists("scopes") {
		scps := []interface{}{}
		for _, v := range data.S("scopes").Children() {
			scps = append(scps, v.Data())
		}

		d.Set("scopes", scps)
	}

	return diagnostics
}
