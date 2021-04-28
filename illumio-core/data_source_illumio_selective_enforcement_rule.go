package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample of API response
{
  "href": "string",
  "created_at": "2021-03-02T02:37:59Z",
  "updated_at": "2021-03-02T02:37:59Z",
  "created_by": {
    "href": "string"
  },
  "updated_by": {
    "href": "string"
  },
  "name": "string",
  "scope": [
    {
      "label": {
        "href": "string"
      },
      "label_group": {
        "href": "string"
      }
    }
  ],
  "enforced_services": [
    {
      "href": "string"
    }
  ]
}
*/

func datasourceIllumioSelectiveEnforcementRule() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioSelectiveEnforcementRuleRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Selective Enforcement Rule",

		Schema: map[string]*schema.Schema{
			"ser_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "selective enforcement rule id",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the selective enforcement rule",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this rule set was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this rule set was last updated",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who originally created this rule set",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who last updated this rule set",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the selective enforcement rule",
			},
			"scope": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Scope of Selective Enforcement Rule",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Label URI",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"label_group": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Label Group URI",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"enforced_services": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of services that are enforced",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Href Link",
						},
					},
				},
			},
		},
	}
}

func datasourceIllumioSelectiveEnforcementRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	serID := d.Get("ser_id").(int)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%d/sec_policy/draft/selective_enforcement_rules/%d", orgID, serID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
	for _, key := range []string{"href", "created_at", "updated_at", "created_by", "updated_by", "name", "scope", "enforced_services"} {

		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		}
	}
	return diagnostics
}
