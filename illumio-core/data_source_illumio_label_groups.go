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

func datasourceIllumioLabelGroups() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioLabelGroupsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Label Groups",

		Schema: map[string]*schema.Schema{
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
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Maximum number of Labels to return",
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

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%d/sec_policy/draft/label_groups", orgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	d.Set("items", extractHrefs(data))

	return diagnostics
}
