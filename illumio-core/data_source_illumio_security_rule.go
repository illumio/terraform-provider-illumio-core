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
  "unscoped_consumers": true,
  "update_type": "string"
}
*/

func datasourceIllumioSecurityRule() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioSecurityRuleRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Security Rule",
		Schema:        securityRuleDatasourceSchema(true),
	}
}

func securityRuleDatasourceSchema(hrefRequired bool) map[string]*schema.Schema {
	m := map[string]*schema.Schema{
		"rule_set_href": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URI of the containing Rule Set",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Enabled flag. Determines whether this rule will be enabled in ruleset or not",
		},
		"description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Description of Security Rule",
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
		"ingress_services": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "Collection of Ingress Service",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"href": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "URI of service",
					},
					"proto": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Protocol number",
					},
					"port": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Port number used with protocol. Also, the starting port when specifying a range",
					},
					"to_port": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Upper end of port range",
					},
				},
			},
		},
		"resolve_labels_as": {
			Type:        schema.TypeList,
			Computed:    true,
			Description: "resolve label as for Security rule",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"providers": {
						Type:        schema.TypeList,
						Computed:    true,
						Description: "providers for resolve_labels_as",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
					"consumers": {
						Type:        schema.TypeList,
						Computed:    true,
						Description: "consumers for resolve_labels_as",
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
		"sec_connect": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Determines whether a secure connection is established",
		},
		"stateless": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Determines whether packet filtering is stateless for the rule",
		},
		"machine_auth": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Determines whether machine authentication is enabled",
		},
		"providers": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "providers for Security Rule",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"actors": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "All workloads provider filter",
					},
					"exclusion": {
						Type:        schema.TypeBool,
						Computed:    true,
						Description: "Boolean to specify whether or not the actor is an exclusion - only for labels and label groups. Requires PCE v22.5+",
					},
					"label": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Label provider filter",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"href": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Label URI",
								},
								"key": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Label key",
								},
								"value": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Label value",
								},
							},
						},
					},
					"label_group": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Label Group provider filter",
						Elem:        hrefSchemaComputed("Label Group", isLabelGroupHref),
					},
					"workload": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Workload provider filter",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"href": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Workload URI",
								},
								"name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Workload name",
								},
								"hostname": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Workload hostname",
								},
								"deleted": {
									Type:        schema.TypeBool,
									Computed:    true,
									Description: "Whether the workload has been deleted in the PCE",
								},
							},
						},
					},
					"virtual_service": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Virtual Service provider filter",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"href": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Virtual Service URI",
								},
								"name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Virtual Service name",
								},
							},
						},
					},
					"virtual_server": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Virtual Server provider filter",
						Elem:        hrefSchemaComputed("Virtual Server", isVirtualServiceHref),
					},
					"ip_list": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "IP List provider filter",
						Elem:        ipListDataSourceSchema(false),
					},
				},
			},
		},
		"consumers": {
			Type:        schema.TypeSet,
			Computed:    true,
			Description: "Consumers for Security Rule",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"actors": {
						Type:        schema.TypeString,
						Computed:    true,
						Description: "Workloads consumer filter",
					},
					"exclusion": {
						Type:        schema.TypeBool,
						Computed:    true,
						Description: "Boolean to specify whether or not the actor is an exclusion - only for labels and label groups. Requires PCE v22.5+",
					},
					"label": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Label consumer filter",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"href": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Label URI",
								},
								"key": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Label key",
								},
								"value": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Label value",
								},
							},
						},
					},
					"label_group": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Label Group consumer filter",
						Elem:        hrefSchemaComputed("Label Group", isLabelGroupHref),
					},
					"workload": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Workload consumer filter",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"href": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Workload URI",
								},
								"name": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Workload name",
								},
								"hostname": {
									Type:        schema.TypeString,
									Computed:    true,
									Description: "Workload hostname",
								},
								"deleted": {
									Type:        schema.TypeBool,
									Computed:    true,
									Description: "Whether the workload has been deleted in the PCE",
								},
							},
						},
					},
					"virtual_service": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "Virtual Service consumer filter",
						Elem:        hrefSchemaComputed("Virtual Service", isVirtualServiceHref),
					},
					"ip_list": {
						Type:        schema.TypeSet,
						Computed:    true,
						Description: "IP List consumer filter",
						Elem:        ipListDataSourceSchema(false),
					},
				},
			},
		},
		"unscoped_consumers": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Set the scope for rule consumers to All",
		},
		"update_type": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Type of update",
		},
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when this security rule was first created",
		},
		"updated_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when this security rule was last updated",
		},
		"deleted_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp when this security rule was deleted",
		},
		"created_by": {
			Type:        schema.TypeMap,
			Computed:    true,
			Description: "User who created this security rule",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"updated_by": {
			Type:        schema.TypeMap,
			Computed:    true,
			Description: "User who last updated this security rule",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"deleted_by": {
			Type:        schema.TypeMap,
			Computed:    true,
			Description: "User who deleted this security rule",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
	if hrefRequired {
		m["href"] = &schema.Schema{
			Type:             schema.TypeString,
			Required:         true,
			ValidateDiagFunc: isSecurityRuleHref,
			Description:      "URI of Security Rule",
		}
	} else {
		m["href"] = &schema.Schema{
			Type:        schema.TypeString,
			Computed:    true,
			Description: "URI of Security Rule",
		}
	}
	return m
}

func datasourceIllumioSecurityRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("href").(string)
	d.Set("rule_set_href", getParentHref(href))

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(href)

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

	for _, key := range rlKeys {
		if data.Exists(key) && data.S(key).Data() != nil {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	rlaKey := "resolve_labels_as"
	if data.Exists(rlaKey) {
		resLableAs := data.S(rlaKey)

		tm := make(map[string][]interface{})
		tm["providers"] = resLableAs.S("providers").Data().([]interface{})
		tm["consumers"] = resLableAs.S("consumers").Data().([]interface{})

		d.Set(rlaKey, []interface{}{tm})
	}

	isKey := "ingress_services"
	if data.Exists(isKey) {
		d.Set(isKey, extractSecurityRuleIngressService(data.S(isKey)))
	} else {
		d.Set(isKey, nil)
	}

	for _, key := range []string{"providers", "consumers"} {
		if data.Exists(key) {
			d.Set(key, extractRuleActors(data.S(key)))
		}
	}

	return diagnostics
}
