// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioRuleSet() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceIllumioRuleSetRead,
		CreateContext: resourceIllumioRuleSetCreate,
		UpdateContext: resourceIllumioRuleSetUpdate,
		DeleteContext: resourceIllumioRuleSetDelete,
		SchemaVersion: 1,
		Description:   "Manages Illumio Ruleset",
		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of Ruleset",
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of update",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this ruleset was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this ruleset was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this ruleset was deleted",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this ruleset",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this ruleset",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this ruleset",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: nameValidation,
				Description:      "Name of Ruleset. Valid name should be between 1 to 255 characters",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of Ruleset",
			},
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "External data set identifier",
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "External data reference identifier",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enabled flag. Determines whether the Ruleset is enabled or not. Default value: true",
			},
			"scopes": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "Ruleset label scopes. At most 3 blocks of label/label_group can be specified inside each scope block",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclusion": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Boolean to specify whether or not the scope is an exclusion",
							Default:     false,
						},
						"label": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Href of Label",
							Elem:        labelOptionalKeyValue(true),
						},
						"label_group": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Href of Label Group",
							Elem:        labelGroupOptionalKeyValue(true),
						},
					},
				},
			},
			"ip_tables_rules": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of IP Tables Rules",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of the Ip Tables Rules",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Enabled flag. Determines whether this IP Tables Rule is enabled or not",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of the IP Tables Rules",
						},
						"statements": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "statements for this IP Tables Rule",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"table_name": {
										Type:             schema.TypeString,
										Required:         true,
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"nat", "mangle", "filter"}, false)),
										Description:      "Name of the table. Allowed values are \"nat\", \"mangle\" and \"filter\"",
									},
									"chain_name": {
										Type:             schema.TypeString,
										Required:         true,
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"PREROUTING", "INPUT", "OUTPUT"}, false)),
										Description:      "Chain name for statement. Allowed values are \"PREROUTING\", \"INPUT\" and \"OUTPUT\"",
									},
									"parameters": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Parameters of statements",
									},
								},
							},
						},
						"actors": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "actors for IP Table Rule",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"actors": {
										Type:             schema.TypeString,
										Optional:         true,
										ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"ams"}, false)),
										Description:      "Set this if rule actors are all workloads. Allowed value is \"ams\"",
									},
									"label": {
										Type:        schema.TypeSet,
										Optional:    true,
										MaxItems:    1,
										Description: "Href of Label",
										Elem:        labelOptionalKeyValue(true),
									},
									"label_group": {
										Type:        schema.TypeSet,
										Optional:    true,
										MaxItems:    1,
										Description: "Href of Label Group",
										Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
									},
									"workload": {
										Type:        schema.TypeSet,
										Optional:    true,
										MaxItems:    1,
										Description: "Href of Workload",
										Elem:        hrefSchemaRequired("Workload", isWorkloadHref),
									},
								},
							},
						},
						"ip_version": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"4", "6"}, false)),
							Description:      "IP version for the rules to be applied to. Allowed values are \"4\" and \"6\"",
						},
						"update_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of update for IP Table Rule",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this IP Table Rule was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this IP Table Rule was last updated",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this IP Table Rule was deleted",
						},
						"created_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who created this IP Table Rule",
						},
						"updated_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who last updated this IP Table Rule",
						},
						"deleted_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who deleted this IP Table Rule",
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioRuleSetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	ruleSet, diags := expandIllumioRuleSet(d)

	if diags.HasError() {
		return *diags
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/rule_sets", orgID), ruleSet)
	if err != nil {
		return diag.Errorf(err.Error())
	}

	pConfig.StoreHref("rule_sets", data.S("href").Data().(string))

	d.SetId(data.S("href").Data().(string))

	return resourceIllumioRuleSetRead(ctx, d, m)
}

func expandIllumioRuleSet(d *schema.ResourceData) (*models.RuleSet, *diag.Diagnostics) {
	var diags diag.Diagnostics

	ruleSet := &models.RuleSet{
		Name:                  PtrTo(d.Get("name").(string)),
		Description:           PtrTo(d.Get("description").(string)),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		Enabled:               PtrTo(d.Get("enabled").(bool)),
	}

	scopes, errs := expandIllumioRuleSetScopes(d.Get("scopes").([]interface{}))
	diags = append(diags, *errs...)
	ruleSet.Scopes = scopes

	ipTableRules, errs := expandIllumioRuleSetIPTablesRules(d.Get("ip_tables_rules").(*schema.Set).List())
	diags = append(diags, *errs...)
	ruleSet.IPTablesRules = &ipTableRules

	return ruleSet, &diags
}

func expandIllumioRuleSetScopes(scopes []interface{}) ([][]*models.RuleSetScope, *diag.Diagnostics) {
	var diags diag.Diagnostics

	sps := [][]*models.RuleSetScope{}

	for _, scope := range scopes {

		sp := []*models.RuleSetScope{}

		if scope == nil {
			sps = append(sps, sp)
			continue
		}

		scopeObj := scope.(map[string]interface{})

		exclusion := scopeObj["exclusion"].(bool)
		labels := scopeObj["label"].(*schema.Set).List()
		labelGroups := scopeObj["label_group"].(*schema.Set).List()

		if len(labels)+len(labelGroups) > 3 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "[illumio-core_rule_set] At most 3 blocks of label/label_group are allowed inside scope",
			})
		} else {
			for _, label := range labels {
				s := &models.RuleSetScope{
					Exclusion: &exclusion,
					Label:     expandLabelOptionalKeyValue(label, false),
				}
				sp = append(sp, s)
			}

			for _, labelGroup := range labelGroups {
				s := &models.RuleSetScope{
					Exclusion:  &exclusion,
					LabelGroup: expandLabelGroupOptionalKeyValue(labelGroup, false),
				}
				sp = append(sp, s)
			}
		}

		sps = append(sps, sp)
	}
	return sps, &diags
}

func expandIllumioRuleSetIPTablesRules(ipTableRules []interface{}) ([]*models.RuleSetIPTablesRule, *diag.Diagnostics) {
	var diags diag.Diagnostics

	iptrs := []*models.RuleSetIPTablesRule{}
	for _, ipTableRule := range ipTableRules {
		i := ipTableRule.(map[string]interface{})

		iptr := &models.RuleSetIPTablesRule{}

		if v, ok := i["enabled"]; ok {
			iptr.Enabled = PtrTo(v.(bool))
		}

		if v, ok := i["description"]; ok {
			iptr.Description = PtrTo(v.(string))
		}

		if v, ok := i["statements"]; ok {
			statements, errs := expandIllumioRuleSetIPTablesRuleStatements(v.(*schema.Set).List())
			if errs.HasError() {
				diags = append(diags, *errs...)
			} else {
				iptr.Statements = &statements
			}
		}

		if v, ok := i["actors"]; ok {
			actors, errs := expandIllumioRuleSetIPTablesRuleActors(v.(*schema.Set).List())
			if errs.HasError() {
				diags = append(diags, *errs...)
			} else {
				iptr.Actors = &actors
			}
		}

		if v, ok := i["ip_version"]; ok {
			iptr.IPVersion = v.(string)
		}

		iptrs = append(iptrs, iptr)
	}

	return iptrs, &diags
}

func expandIllumioRuleSetIPTablesRuleStatements(statements []interface{}) ([]*models.RuleSetIPTablesRulesStatement, *diag.Diagnostics) {
	var diags diag.Diagnostics
	stats := []*models.RuleSetIPTablesRulesStatement{}

	for _, statement := range statements {
		s := statement.(map[string]interface{})
		stat := &models.RuleSetIPTablesRulesStatement{
			TableName:  s["table_name"].(string),
			ChainName:  s["chain_name"].(string),
			Parameters: s["parameters"].(string),
		}

		stats = append(stats, stat)
	}

	return stats, &diags
}

func expandIllumioRuleSetIPTablesRuleActors(actors []interface{}) ([]*models.RuleSetIPTablesRulesActor, *diag.Diagnostics) {
	var diags diag.Diagnostics
	acts := []*models.RuleSetIPTablesRulesActor{}

	for _, actor := range actors {
		a := actor.(map[string]interface{})
		act := &models.RuleSetIPTablesRulesActor{
			Actors:     a["actors"].(string),
			Label:      expandLabelOptionalKeyValue(a["label"], false),
			LabelGroup: getHrefObj(a["label_group"]),
			Workload:   getHrefObj(a["workload"]),
		}

		acts = append(acts, act)
	}

	return acts, &diags
}

func resourceIllumioRuleSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	_, data, err := illumioClient.Get(d.Id(), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, key := range []string{
		"href",
		"name",
		"description",
		"enabled",
		"external_data_set",
		"external_data_reference",
		"update_type",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("scopes") {
		d.Set("scopes", extractResourceScopes(data.S("scopes")))
	} else {
		d.Set("scopes", nil)
	}

	if data.Exists("ip_tables_rules") {
		d.Set("ip_tables_rules", extractResourceRuleSetIPTablesRules(data.S("ip_tables_rules")))
	} else {
		d.Set("ip_tables_rules", nil)
	}

	return diagnostics
}

func extractResourceRuleSetIPTablesRules(data *gabs.Container) []map[string]interface{} {
	ms := []map[string]interface{}{}

	iptrKeys := []string{
		"ip_version",
		"update_type",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"href",
		"enabled",
		"description",
	}

	statKeys := []string{
		"table_name",
		"chain_name",
		"parameters",
	}

	for _, ipTableRuleData := range data.Children() {
		m := extractMap(ipTableRuleData, iptrKeys)
		m["actors"] = extractRuleActors(ipTableRuleData.S("actors"))
		m["statements"] = extractMapArray(ipTableRuleData.S("statements"), statKeys)

		ms = append(ms, m)
	}
	return ms
}

func resourceIllumioRuleSetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	var diags diag.Diagnostics

	ruleSet := &models.RuleSet{
		Name:                  PtrTo(d.Get("name").(string)),
		Description:           PtrTo(d.Get("description").(string)),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		Enabled:               PtrTo(d.Get("enabled").(bool)),
		IPTablesRules:         nil,
	}

	scopes, errs := expandIllumioRuleSetScopes(d.Get("scopes").([]interface{}))
	diags = append(diags, *errs...)
	ruleSet.Scopes = scopes

	if d.HasChange("ip_tables_rules") {
		ipTableRules, errs := expandIllumioRuleSetIPTablesRules(d.Get("ip_tables_rules").(*schema.Set).List())
		ruleSet.IPTablesRules = &ipTableRules
		diags = append(diags, *errs...)
	}
	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), ruleSet)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref("rule_sets", d.Id())

	return resourceIllumioRuleSetRead(ctx, d, m)
}

func resourceIllumioRuleSetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, err := illumioClient.Delete(href)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref("rule_sets", href)

	d.SetId("")
	return diagnostics
}
