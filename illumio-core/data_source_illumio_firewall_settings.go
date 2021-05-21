// Copyright 2021 Illumio, Inc. All Rights Reserved.

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
  "static_policy_scopes": [
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
  "firewall_coexistence": [{"illumio_primary": true,
                           "scope": [{"href": "string"},
                                     {"href": "string"},
                                     {"href": "string"},
                                     {"href": "string"}],
                           "workload_mode": "illuminated"},
                          {"illumio_primary": true,
                           "scope": [{"href": "string"},
                                     {"href": "string"},
                                     {"href": "string"},
                                     {"href": "string"}]}],
  "containers_inherit_host_policy_scopes": [
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
  "blocked_connection_reject_scopes": [
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
  "update_type": "string"
}
*/

func datasourceIllumioFirewallSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioFirewallSettingsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Firewall Settings",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of Firewall Settings",
				ValidateDiagFunc: isFirewallSettingsHref,
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of Update",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when these firewall settings were first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when these firewall settings were last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when these firewall settings were deleted",
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
			"ike_authentication_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IKE authentication type to use for IPsec (SecureConnect and Machine Authentication)",
			},
			"static_policy_scopes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "scopes for static policy",
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
			"firewall_coexistence": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Firewall coexistence configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"illumio_primary": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether Illumio is primary firewall or not",
						},
						"scope": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Href of Label",
									},
								},
							},
						},
						"workload_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Match criteria to select workload(s)",
						},
					},
				},
			},
			"containers_inherit_host_policy_scopes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "scopes for container inherit host policy",
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
			"blocked_connection_reject_scopes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "scopes for reject connections",
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
			"loopback_interfaces_in_policy_scopes": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "scopes for loopback interfaces",
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
		},
	}
}

func datasourceIllumioFirewallSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	firewallSettingsKeys := []string{
		"href",
		"update_type",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"ike_authentication_type",
	}
	d.SetId(data.S("href").Data().(string))
	for _, key := range firewallSettingsKeys {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		}
	}
	for _, k := range []string{
		"static_policy_scopes",
		"containers_inherit_host_policy_scopes",
		"blocked_connection_reject_scopes",
		"loopback_interfaces_in_policy_scopes",
	} {
		d.Set(k, extractDatasourceScopes(data, k))
	}
	d.Set("firewall_coexistence", extractFirewallCoexistence(data))
	return diagnostics
}

func extractFirewallCoexistence(data *gabs.Container) []interface{} {
	key := "firewall_coexistence"
	if data.Exists(key) {
		l1 := []interface{}{}
		for _, child1 := range data.S(key).Children() {
			obj := map[string]interface{}{}
			k := "illumio_primary"
			if child1.Exists(k) {
				obj[k] = child1.S(k).Data().(bool)
			} else {
				obj[k] = nil
			}
			k = "workload_mode"
			if child1.Exists(k) {
				obj[k] = child1.S(k).Data().(string)
			} else {
				obj[k] = ""
			}
			k = "scope"
			if child1.Exists(k) {
				l2 := []map[string]string{}
				for _, child2 := range child1.S(k).Children() {
					l2 = append(l2, map[string]string{"href": child2.S("href").Data().(string)})
				}
				obj[k] = l2
			} else {
				obj[k] = nil
			}
			l1 = append(l1, obj)
		}
		return l1
	}
	return nil
}

func extractDatasourceScopes(data *gabs.Container, key string) []interface{} {
	if data.Exists(key) {
		l1 := []interface{}{}
		for _, child1 := range data.S(key).Children() {
			l2 := []map[string]interface{}{}
			for _, child2 := range child1.Children() {
				scopeObj := map[string]interface{}{}
				for _, k := range []string{"label", "label_group"} {
					if child2.Exists(k) {
						scopeObj[k] = map[string]string{"href": child2.S(k, "href").Data().(string)}
					} else {
						scopeObj[k] = nil
					}
				}
				l2 = append(l2, scopeObj)
			}
			l1 = append(l1, l2)
		}
		return l1
	}
	return nil
}
