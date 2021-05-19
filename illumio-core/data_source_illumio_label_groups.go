// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/* Sample of API response
[
	{
		"name": "string",
		"description": null,
		"key": "string",
		"labels": [
			{
			"href": "string",
			"key": "string",
			"value": "string"
			}
		],
		"sub_groups": [
			{
			"href": "string",
			"name": "string"
			}
		],
		"usage": {
			"label_group": true,
			"ruleset": true,
			"rule": true,
			"static_policy_scopes": true,
			"containers_inherit_host_policy_scopes": true,
			"blocked_connection_reject_scope": true
		},
		"external_data_set": null,
		"external_data_reference": null,
		"update_type": null,
		"created_at": "2021-03-02T02:37:59Z",
		"updated_at": "2021-03-02T02:37:59Z",
		"deleted_at": null,
		"created_by": {
			"href": "string"
		},
		"updated_by": {
			"href": "string"
		},
		"deleted_by": {
			"href": "string"
		}
	}
]
*/

func datasourceIllumioLabelGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioLabelGroupsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Label Groups",

		Schema: map[string]*schema.Schema{
			"pversion": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "draft",
				ValidateDiagFunc: isValidPversion(),
				Description:      `pversion of the security policy. Allowed values are "draft", "active" and numbers greater than 0. Default value: "draft"`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The long description of the label group",
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "A unique identifier within the external data source",
			},
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "The data source from which a resource originates",
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(validLabelKeys, false),
				),
				Description: `Key in key-value pair of contained labels or label groups. Allowed values for key are "role", "loc", "app" and "env".`,
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of Labels to return. The integer should be a non-zero positive integer.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of Label Group(s) to return. Supports partial matches",
			},
			"usage": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Include label usage flags as well",
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of label group hrefs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of label group",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Label Group",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The long description of the Label Group",
						},
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Key in key-value pair of contained labels or Label Groups",
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Contained labels",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of label",
									},
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Label Key same as label group key",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Label Value in key-value pair",
									},
								},
							},
						},
						"sub_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Contained Label Groups",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of Label Group",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "name of sub Label Group",
									},
								},
							},
						},
						"external_data_set": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External Data set Identifier",
						},
						"external_data_reference": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External Data reference identifier",
						},
						"update_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of Update",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Label Group was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Label Group was last updated",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Label Group was last deleted",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who originally created this Label Group",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who last updated this Label Group",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"deleted_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who deleted this Label Group",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func datasourceIllumioLabelGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID
	pversion := d.Get("pversion").(string)

	paramKeys := []string{
		"description",
		"external_data_reference",
		"external_data_set",
		"key",
		"max_results",
		"name",
		"usage",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%d/sec_policy/%v/label_groups", orgID, pversion), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}
	keys := []string{
		"href",
		"name",
		"description",
		"key",
		"external_data_set",
		"external_data_reference",
		"update_type",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
	}

	for _, child := range data.Children() {
		m := extractMap(child, keys)

		if child.Exists("labels") {
			labels := child.S("labels")
			labelI := []map[string]interface{}{}

			for _, l := range labels.Children() {
				labelI = append(labelI, extractMap(l, []string{"href", "key", "value"}))
			}

			m["labels"] = labelI
		} else {
			m["labels"] = nil
		}

		if child.Exists("sub_groups") {
			sub_groups := child.S("sub_groups")
			sub_groupI := []map[string]interface{}{}

			for _, sg := range sub_groups.Children() {
				sub_groupI = append(sub_groupI, extractMap(sg, []string{"href", "name"}))
			}

			m["sub_groups"] = sub_groupI
		} else {
			m["sub_groups"] = nil
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
