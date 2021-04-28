package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioSelectiveEnforcementRule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioSelectiveEnforcementRuleCreate,
		ReadContext:   resourceIllumioSelectiveEnforcementRuleRead,
		UpdateContext: resourceIllumioSelectiveEnforcementRuleUpdate,
		DeleteContext: resourceIllumioSelectiveEnforcementRuleDelete,
		Description:   "Manages Illumio Selective Enforcement Rule",
		SchemaVersion: version,

		Schema: map[string]*schema.Schema{
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
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who originally created this rule set",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this rule set",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the selective enforcement rule",
				ValidateDiagFunc: nameValidation,
			},
			"scope": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Scope of Selective Enforcement Rule",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        hrefSchemaRequired("Label", isLabelHref),
							Description: "Href of Label",
						},
						"label_group": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
							Description: "Href of Label Group",
						},
					},
				},
			},
			"enforced_services": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Collection of services that are enforced",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "URI of enforced service.",
						},
					},
				},
			},
		},
	}
}

func resourceIllumioSelectiveEnforcementRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID
	var diags diag.Diagnostics

	SER := &models.SelectiveEnforcementRules{
		Name: d.Get("name").(string),
	}

	scope, errs := expandIllumioSelectiveEnforcementRuleScope(d.Get("scope").([]interface{})[0])
	diags = append(diags, errs...)
	SER.Scope = scope

	if items, ok := d.GetOk("enforced_services"); ok {
		SER.EnforcedServices = models.GetHrefs(items.(*schema.Set).List())
	}

	if diags.HasError() {
		return diags
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/selective_enforcement_rules", orgID), SER)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref(pConfig.OrgID, "selective_enforcement_rules", data.S("href").Data().(string))
	d.SetId(data.S("href").Data().(string))

	return resourceIllumioSelectiveEnforcementRuleRead(ctx, d, m)
}

func expandIllumioSelectiveEnforcementRuleScope(scope interface{}) ([]*models.SelectiveEnforcementRulesScope, diag.Diagnostics) {
	var diags diag.Diagnostics

	sp := []*models.SelectiveEnforcementRulesScope{}

	scopeObj := scope.(map[string]interface{})

	labels := scopeObj["label"].(*schema.Set).List()
	labelGroups := scopeObj["label_group"].(*schema.Set).List()

	if len(labels)+len(labelGroups) > 4 {
		diags = append(diags, diag.Diagnostic{
			Severity:      diag.Error,
			Summary:       "at most 4 blocks of label/label_group are allowed inside scope",
			AttributePath: cty.Path{cty.GetAttrStep{Name: "scope"}}},
		)
	} else {

		for _, label := range labels {
			s := &models.SelectiveEnforcementRulesScope{
				Label: getHrefObj(label),
			}
			sp = append(sp, s)
		}

		for _, labelGroup := range labelGroups {
			s := &models.SelectiveEnforcementRulesScope{
				LabelGroup: getHrefObj(labelGroup),
			}
			sp = append(sp, s)
		}
	}

	return sp, diags
}

func resourceIllumioSelectiveEnforcementRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	SERHref := d.Id()

	_, data, err := illumioClient.Get(SERHref, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, key := range []string{"href", "created_at", "updated_at", "created_by", "updated_by", "name", "enforced_services"} {

		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	labels := []map[string]interface{}{}
	labelGroups := []map[string]interface{}{}

	for _, data := range data.S("scope").Children() {
		for k, v := range data.ChildrenMap() {
			if k == "label" {
				labels = append(labels, v.Data().(map[string]interface{}))
			} else if k == "label_group" {
				labelGroups = append(labelGroups, v.Data().(map[string]interface{}))
			}
		}
	}

	sc := map[string]interface{}{}
	sc["label"] = labels
	sc["label_group"] = labelGroups

	d.Set("scope", []interface{}{sc})

	return diagnostics

}

func resourceIllumioSelectiveEnforcementRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	var diags diag.Diagnostics

	SER := &models.SelectiveEnforcementRules{}

	if d.HasChange("name") {
		SER.Name = d.Get("name").(string)
	}

	scope, errs := expandIllumioSelectiveEnforcementRuleScope(d.Get("scope").([]interface{})[0])
	diags = append(diags, errs...)
	SER.Scope = scope

	SER.EnforcedServices = models.GetHrefs(d.Get("enforced_services").(*schema.Set).List())

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), SER)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "selective_enforcement_rules", d.Id())

	return resourceIllumioSelectiveEnforcementRuleRead(ctx, d, m)
}

func resourceIllumioSelectiveEnforcementRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, err := illumioClient.Delete(href)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref(pConfig.OrgID, "selective_enforcement_rules", href)
	d.SetId("")
	return diagnostics
}
