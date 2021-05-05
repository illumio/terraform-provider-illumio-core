package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample of API response
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
*/

func datasourceIllumioLabelGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioLabelGroupRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Label Group",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URI of Label Group",
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
				Type:        schema.TypeSet,
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
				Type:        schema.TypeSet,
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
	}
}

func datasourceIllumioLabelGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
	// set computed/optional values from api response
	for _, key := range []string{
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
	} {

		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}

		if data.Exists("labels") {
			labels := data.S("labels")
			labelI := []map[string]interface{}{}

			for _, l := range labels.Children() {
				labelI = append(labelI, gabsToMap(l, []string{"href", "key", "value"}))
			}

			d.Set("labels", labelI)
		} else {
			d.Set("labels", nil)
		}

		if data.Exists("sub_groups") {
			sub_groups := data.S("sub_groups")
			sub_groupI := []map[string]interface{}{}

			for _, sg := range sub_groups.Children() {
				sub_groupI = append(sub_groupI, gabsToMap(sg, []string{"href", "name"}))
			}

			d.Set("sub_groups", sub_groupI)
		} else {
			d.Set("sub_groups", nil)
		}

	}
	return diagnostics
}
