// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioLabelType() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioLabelTypeRead,
		SchemaVersion: version,
		Description:   "Represents an Illumio Label Type. Requires PCE version 22.5.0 or higher",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of this label type",
				ValidateDiagFunc: isLabelTypeHref,
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
	}
}

func dataSourceIllumioLabelTypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
	for _, key := range []string{
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
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("display_info") {
		d.Set("display_info", []interface{}{
			extractMap(
				data.S("display_info"),
				[]string{
					"initial",
					"icon",
					"background_color",
					"foreground_color",
					"sort_ordinal",
					"display_name_plural",
				},
			),
		})
	} else {
		d.Set("display_info", nil)
	}

	return diagnostics
}
