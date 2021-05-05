package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample of API response
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
*/

func datasourceIllumioLabel() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioLabelRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Label",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URI of this label",
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
	}
}

func dataSourceIllumioLabelRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}
	return diagnostics
}
