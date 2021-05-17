package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioVulnerabilities() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioVulnerabilitiesRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Vulnerabilities",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of vulnerabilities",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Vulnerability",
						},
						"score": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The normalized score of the vulnerability within the range of 0 to 100. CVSS Score can be used here with a 10x multiplier",
						},
						"cve_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The cve_ids for the vulnerability",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An arbitrary field to store some details of the vulnerability class",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The title/name of the vulnerability",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time (rfc3339 timestamp) at which this report was created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time (rfc3339 timestamp) at which this report was last updated",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The Href of the user who created this report",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The Href of the user who last updated this report",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of vulnerabilities to return. The integer should be a non-zero positive integer.",
			},
		},
	}
}

func dataSourceIllumioVulnerabilitiesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	paramKeys := []string{
		"max_results",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%v/vulnerabilities", orgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}

	for _, child := range data.Children() {
		m := map[string]interface{}{}

		for _, key := range []string{
			"href",
			"score",
			"cve_ids",
			"description",
			"name",
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
