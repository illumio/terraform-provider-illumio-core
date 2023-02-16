// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceIllumioLabelTypes() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioLabelTypesRead,
		SchemaVersion: version,
		Description:   "Represents a list of Illumio Label Types. Requires PCE version 22.5.0 or higher",

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name on which to filter. Supports partial matches",
			},
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
					validation.StringLenBetween(1, LABEL_KEY_LENGTH_MAX),
				),
				Description: `Label type key to filter by. The value must be a string between 1 and 64 characters long`,
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of Labels to return. The integer should be a non-zero positive integer",
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of label types",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of this label type",
						},
						"key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Key in key-value pair. The value must be a string between 1 and 64 characters long`,
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the label type",
						},
						"display_info": {
							// XXX: surely there must be a better way to represent
							//      subobjects in the schema
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Object containing UI display information for the label type",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"initial": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "1-2 initial characters for use in the UI display. Defaults to the first letter of the label type's display_name",
									},
									"icon": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Icon for use in the UI display",
									},
									"background_color": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Background color in hexadecimal for UI display",
									},
									"foreground_color": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Foreground color in hexadecimal for UI display",
									},
									"sort_ordinal": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Optional user provided sort order for label type",
									},
									"display_name_plural": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Optional pluralized form of the display name for the label type",
									},
								},
							},
						},
						"usage": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Usage of the label type",
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"external_data_set": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The data source from which a resource originates",
						},
						"external_data_reference": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique identifier within the external data source",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Flag to indicate whether the label type has been deleted or not",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this label type was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this label type was last updated",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this label type was deleted",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who created this label type",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who last updated this label type",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"deleted_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who deleted this label type",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"caps": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Object permissions",
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

func dataSourceIllumioLabelTypesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	paramKeys := []string{
		"display_name",
		"external_data_set",
		"external_data_reference",
		"include_deleted",
		"key",
		"max_results",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%d/label_dimensions", orgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}
	keys := []string{
		"href",
		"key",
		"display_name",
		"usage",
		"external_data_set",
		"external_data_reference",
		"deleted",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"deleted_at",
		"deleted_by",
		"caps",
	}

	for _, child := range data.Children() {
		m := extractMap(child, keys)

		if child.Exists("display_info") {
			m["display_info"] = []interface{}{
				extractMap(
					child.S("display_info"),
					[]string{
						"initial",
						"icon",
						"background_color",
						"foreground_color",
						"sort_ordinal",
						"display_name_plural",
					},
				),
			}
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
