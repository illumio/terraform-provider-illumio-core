package illumiocore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/go-cty/cty"
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

func datasourceIllumioServices() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioServicesRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Services",
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Long description of the servcie",
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "External data reference identifier",
			},
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "External data set identifier",
			},
			"max_results": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Maximum number of Services to return",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the servcie (does not need to be unique)",
			},
			"port": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Specify port or port range to filter results. The range is from -1 to 65535 (0 is not supported)",
				ValidateDiagFunc: servicePortValidation(),
			},
			"proto": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Protocol to filter on. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validServiceProtos, false)),
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of service hrefs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of service",
						},
					},
				},
			},
		},
	}
}

func dataSourceIllumioServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	paramKeys := []string{
		"description",
		"external_data_reference",
		"external_data_set",
		"max_results",
		"name",
		"port",
		"proto",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%d/sec_policy/draft/services", pConfig.OrgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	d.Set("items", extractHrefs(data))

	return diagnostics
}

func servicePortValidation() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		k, err := strconv.Atoi(v.(string))

		if err != nil {
			diags = append(diags, diag.Errorf("expected integer value, got: %v", v)...)
			return diags
		}

		if (1 > k || k > 65535) && k != -1 {
			diags = append(diags, diag.Errorf("expected to be in range 1-65535 or -1, got: %v", v)...)
		}

		return diags
	}
}
