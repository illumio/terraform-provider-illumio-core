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
		"href": string
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

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%d/labels", orgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	d.Set("items", extractHrefs(data))

	return diagnostics
}
