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
		"href": "string",
		"deleted": true,
		"key": "string",
		"value": "string",
		"external_data_set": null,
		"external_data_reference": null,
		"created_at": "2020-08-19T21:34:26Z",
		"updated_at": "2020-08-19T21:34:26Z",
		"created_by": {
			"href": "string"
		},
		"updated_by": {
			"href": "string"
		}
	}
]
*/

func datasourceIllumioLabels() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioLabelsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Labels",

		Schema: map[string]*schema.Schema{
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "The data source from which a resource originates",
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "A unique identifier within the external data source",
			},
			"include_deleted": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Include deleted labels",
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(validLabelKeys, false),
				),
				Description: `Key in key-value pair. Allowed values for key are "role", "loc", "app" and "env"`,
			},
			"max_results": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Maximum number of Labels to return",
			},
			"usage": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Include label usage flags as well",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Value on which to filter. Supports partial matches",
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of label hrefs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of label",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag to indicate whether deleted or not",
						},
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Key in key-value pair",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value in key-value pair",
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
							Description: "Timestamp when this label was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this label was last updated",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who originally created this label",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who last updated this label",
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

func dataSourceIllumioLabelsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	paramKeys := []string{
		"external_data_set",
		"external_data_reference",
		"include_deleted",
		"key",
		"max_results",
		"usage",
		"value",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%d/labels", orgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}

	for _, child := range data.Children() {
		m := map[string]interface{}{}

		for _, key := range []string{
			"href",
			"deleted",
			"key",
			"value",
			"external_data_set",
			"external_data_reference",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
		} {
			if child.Exists(key) {
				m[key] = child.S(key).Data()
			} else {
				m[key] = nil
			}
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
