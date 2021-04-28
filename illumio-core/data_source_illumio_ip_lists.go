package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample of API response
[
	{
		"href": string
	}
]
*/

func datasourceIllumioIPLists() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioIPListsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio IP Lists",

		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of IP list(s) to return. Supports partial matches",
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A unique identifier within the external data source",
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The data source from which a resource originates",
			},
			"fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP lists matching FQDN. Supports partial matches",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address matching IP list(s) to return. Supports partial matches",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Maximum number of IP Lists to return.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of IP list(s) to return. Supports partial matches",
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of IP List hrefs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of this IP List",
						},
					},
				},
			},
		},
	}
}

func datasourceIllumioIPListsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	paramKeys := []string{
		"description",
		"external_data_reference",
		"external_data_set",
		"fqdn",
		"ip_address",
		"max_results",
		"name",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%v/sec_policy/draft/ip_lists", orgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	d.Set("items", extractHrefs(data))

	return diagnostics
}
