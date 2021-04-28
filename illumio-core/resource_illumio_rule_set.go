package illumiocore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

var (
	RuleSetIPTableRuleSupportedParams = []string{
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
		"statements",
	}
)

func resourceIllumioRuleSet() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceIllumioRuleSetRead,
		CreateContext: resourceIllumioRuleSetCreate,
		UpdateContext: resourceIllumioRuleSetUpdate,
		DeleteContext: resourceIllumioRuleSetDelete,
		SchemaVersion: version,
		Description:   "Manages Illumio Rule Set",
		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of Rule Set",
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of update",
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
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this rule set was deleted",
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
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this rule set",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: nameValidation,
				Description:      "Name of Rule Set. Valid name should be in between 1 to 255 characters",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of Rule Set",
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
				Description: "Enabled flag. Determines wheter the Rule Set is enabled or not. Default value: true",
			},
			"scope": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "scopes for Rule Set. At most 3 blocks of label/label_group can be specified inside each scope block",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"label": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Href of Label",
							Elem:        hrefSchemaRequired("Label", isLabelHref),
						},
						"label_group": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Href of Label Group",
							Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
						},
					},
				},
			},
			"rule": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Security Rules",
				Elem: &schema.Resource{
					Schema: securityRuleResourceBaseSchemaMap(),
				},
			},
			"ip_tables_rule": {
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
										Description:      "Set this if rule actors are all workloads. Allowed value: \"ams\"",
									},
									"label": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "Href of Label",
										Elem:        hrefSchemaRequired("Label", isLabelHref),
									},
									"label_group": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "Href of Label Group",
										Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
									},
									"workload": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "Href of Worklaod",
										Elem:        hrefSchemaRequired("Workload", isWorklaodHref),
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
							Description: "User who originally created this IP Table Rule",
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
	}
}

func resourceIllumioRuleSetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	ruleSet, diags := expandIllumioRuleSet(d)

	if diags.HasError() {
		return *diags
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/rule_sets", orgID), ruleSet)
	if err != nil {
		return diag.Errorf(err.Error())
	}

	pConfig.StoreHref(pConfig.OrgID, "rule_sets", data.S("href").Data().(string))

	d.SetId(data.S("href").Data().(string))

	return resourceIllumioRuleSetRead(ctx, d, m)
}

func expandIllumioRuleSet(d *schema.ResourceData) (*models.RuleSet, *diag.Diagnostics) {
	var diags diag.Diagnostics

	ruleSet := &models.RuleSet{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		Enabled:               d.Get("enabled").(bool),
	}

	scopes, errs := expandIllumioRuleSetScopes(d.Get("scope").([]interface{}))
	diags = append(diags, *errs...)
	ruleSet.Scopes = scopes

	rules, errs := expandIllumioRuleSetSecurityRules(d.Get("rule").(*schema.Set).List())
	diags = append(diags, *errs...)
	ruleSet.Rules = rules

	ipTableRules, errs := expandIllumioRuleSetIPTablesRules(d.Get("ip_tables_rule").(*schema.Set).List())
	diags = append(diags, *errs...)
	ruleSet.IPTablesRules = ipTableRules

	return ruleSet, &diags
}

func expandIllumioRuleSetScopes(scopes []interface{}) ([][]*models.RuleSetScope, *diag.Diagnostics) {
	var diags diag.Diagnostics

	sps := [][]*models.RuleSetScope{}

	for i, scope := range scopes {
		sp := []*models.RuleSetScope{}

		scopeObj := scope.(map[string]interface{})

		labels := scopeObj["label"].(*schema.Set).List()
		labelGroups := scopeObj["label_group"].(*schema.Set).List()

		if len(labels)+len(labelGroups) > 3 {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "at most 3 blocks of label/label_group are allowed inside scope",
				AttributePath: cty.Path{cty.GetAttrStep{Name: "scope"}, cty.IndexStep{Key: cty.NumberIntVal(int64(i))}},
			})
		} else {

			for _, label := range labels {
				s := &models.RuleSetScope{
					Label: getHrefObj(label),
				}
				sp = append(sp, s)
			}

			for _, labelGroup := range labelGroups {
				s := &models.RuleSetScope{
					LabelGroup: getHrefObj(labelGroup),
				}
				sp = append(sp, s)
			}
		}

		sps = append(sps, sp)
	}
	return sps, &diags
}

func expandIllumioRuleSetSecurityRules(rules []interface{}) ([]*models.SecurityRule, *diag.Diagnostics) {
	var diags diag.Diagnostics
	rls := []*models.SecurityRule{}

	for _, rule := range rules {
		r := rule.(map[string]interface{})
		rl := &models.SecurityRule{
			Enabled:               r["enabled"].(bool),
			Description:           r["description"].(string),
			ExternalDataSet:       r["external_data_set"].(string),
			ExternalDataReference: r["external_data_reference"].(string),
			SecConnect:            r["sec_connect"].(bool),
			Stateless:             r["stateless"].(bool),
			MachineAuth:           r["machine_auth"].(bool),
			UnscopedConsumers:     r["unscoped_consumers"].(bool),
		}

		if rl.HasConflicts() {
			diags = append(diags, diag.Errorf("Only one of [\"sec_connect\", \"machine_auth\", \"stateless\"] can be set to true")...)
		}

		resLabelAs, errs := expandIllumioSecurityRuleResolveLabelsAs(r["resolve_labels_as"].([]interface{})[0])
		diags = append(diags, errs...)
		rl.ResolveLabelsAs = resLabelAs

		ingServs, errs := expandIllumioSecurityRuleIngressService(
			r["ingress_service"].(*schema.Set).List(),
			rl.ResolveLabelsAs.ProviderIsVirtualService(),
		)
		diags = append(diags, errs...)
		rl.IngressServices = ingServs

		povs, errs := expandIllumioSecurityRuleProviders(r["illumio_provider"].(*schema.Set).List())
		diags = append(diags, errs...)
		rl.Providers = povs

		cons, errs := expandIllumioSecurityRuleConsumers(r["consumer"].(*schema.Set).List())
		diags = append(diags, errs...)
		rl.Consumers = cons

		rls = append(rls, rl)
	}

	return rls, &diags
}

func expandIllumioRuleSetIPTablesRules(ipTableRules []interface{}) ([]*models.RuleSetIPTablesRule, *diag.Diagnostics) {
	var diags diag.Diagnostics

	iptrs := []*models.RuleSetIPTablesRule{}
	for _, ipTableRule := range ipTableRules {
		i := ipTableRule.(map[string]interface{})

		iptr := &models.RuleSetIPTablesRule{}

		if v, ok := i["enabled"]; ok {
			iptr.Enabled = v.(bool)
		}

		if v, ok := i["description"]; ok {
			iptr.Description = v.(string)
		}

		if v, ok := i["statements"]; ok {
			statements, errs := expandIllumioRuleSetIPTablesRuleStatements(v.(*schema.Set).List())
			if errs.HasError() {
				diags = append(diags, *errs...)
			} else {
				iptr.Statements = statements
			}
		}

		if v, ok := i["actors"]; ok {
			actors, errs := expandIllumioRuleSetIPTablesRuleActors(v.(*schema.Set).List())
			if errs.HasError() {
				diags = append(diags, *errs...)
			} else {
				iptr.Actors = actors
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
			Label:      getHrefObj(a["label"]),
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
		d.Set("scope", extractResourceScopes(data.S("scopes")))
	} else {
		d.Set("scope", nil)
	}

	if data.Exists("rules") {
		d.Set("rule", resourceIllumioRuleSetReadSecurityRules(data.S("rules")))
	} else {
		d.Set("rule", nil)
	}

	if data.Exists("ip_tables_rules") {
		d.Set("ip_tables_rule", resourceIllumioRuleSetReadIPTablesRules(data.S("ip_tables_rules")))
	} else {
		d.Set("ip_tables_rule", nil)
	}

	return diagnostics
}

func resourceIllumioRuleSetReadIPTablesRules(data *gabs.Container) []map[string]interface{} {
	ms := []map[string]interface{}{}
	for _, data := range data.Children() {
		m := map[string]interface{}{}
		for k, v := range data.ChildrenMap() {
			if k == "actors" {
				m[k] = getRuleActors(v)
			} else if contains(RuleSetIPTableRuleSupportedParams, k) {
				m[k] = v.Data()
			}
		}
		ms = append(ms, m)
	}
	return ms
}

func resourceIllumioRuleSetReadSecurityRules(data *gabs.Container) []map[string]interface{} {
	ms := []map[string]interface{}{}
	for _, data := range data.Children() {
		m := map[string]interface{}{}

		for _, key := range []string{
			"href",
			"enabled",
			"description",
			"external_data_set",
			"external_data_reference",
			"sec_connect",
			"stateless",
			"machine_auth",
			"unscoped_consumers",
			"update_type",
			"created_at",
			"updated_at",
			"deleted_at",
			"created_by",
			"updated_by",
			"deleted_by",
		} {
			if data.Exists(key) {
				m[key] = data.S(key).Data()
			} else {
				m[key] = nil
			}
		}

		obj := map[string]interface{}{}
		for k, v := range data.S("resolve_labels_as").Data().(map[string]interface{}) {
			obj[k] = getStringList(v)
		}
		resLabAs := []map[string]interface{}{obj}

		// Req param, Will be present in JSON responce
		m["resolve_labels_as"] = resLabAs

		// Req param, Will be present in JSON responce
		m["illumio_provider"] = getRuleActors(data.S("providers"))

		// Req param, Will be present in JSON responce
		m["consumer"] = getRuleActors(data.S("consumers"))

		ingServs := data.S("ingress_services").Data().([]interface{})

		iss := []map[string]interface{}{}

		for _, ingServ := range ingServs {
			is := ingServ.(map[string]interface{})

			for k, v := range ingServ.(map[string]interface{}) {
				if k == "href" {
					is[k] = v
				} else {
					is[k] = strconv.Itoa(int(v.(float64)))
				}

				iss = append(iss, is)
			}
		}

		// Req param, Will be present in JSON responce
		m["ingress_service"] = iss

		ms = append(ms, m)
	}

	return ms
}

func resourceIllumioRuleSetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	var diags diag.Diagnostics

	ruleSet := &models.RuleSet{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		Enabled:               d.Get("enabled").(bool),
		Rules:                 nil,
		IPTablesRules:         nil,
	}

	scopes, errs := expandIllumioRuleSetScopes(d.Get("scope").([]interface{}))
	diags = append(diags, *errs...)
	ruleSet.Scopes = scopes

	if d.HasChange("rule") {
		rules, errs := expandIllumioRuleSetSecurityRules(d.Get("rule").(*schema.Set).List())
		ruleSet.Rules = rules
		diags = append(diags, *errs...)
	}

	if d.HasChange("ip_tables_rule") {
		ipTableRules, errs := expandIllumioRuleSetIPTablesRules(d.Get("ip_tables_rule").(*schema.Set).List())
		ruleSet.IPTablesRules = ipTableRules
		diags = append(diags, *errs...)
	}
	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), ruleSet)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref(pConfig.OrgID, "rule_sets", d.Id())

	return resourceIllumioRuleSetRead(ctx, d, m)
}

func resourceIllumioRuleSetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, err := illumioClient.Delete(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "rule_sets", href)

	d.SetId("")
	return diagnostics
}
