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
		SchemaVersion: version,
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
				Description: "scopes for Ruleset. At most 3 blocks of label/label_group can be specified inside each scope block",
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
			// "rules": {
			// 	Type:        schema.TypeSet,
			// 	Optional:    true,
			// 	Description: "Collection of Security Rules",
			// 	Elem: &schema.Resource{
			// 		Schema: securityRuleResourceBaseSchemaMap(),
			// 	},
			// },
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
										Elem:        hrefSchemaRequired("Label", isLabelHref),
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

	scopes, errs := expandIllumioRuleSetScopes(d.Get("scopes").([]interface{}))
	diags = append(diags, *errs...)
	ruleSet.Scopes = scopes

	// rules, errs := expandIllumioRuleSetSecurityRules(d.Get("rules").(*schema.Set).List())
	// diags = append(diags, *errs...)
	// ruleSet.Rules = rules

	ipTableRules, errs := expandIllumioRuleSetIPTablesRules(d.Get("ip_tables_rules").(*schema.Set).List())
	diags = append(diags, *errs...)
	ruleSet.IPTablesRules = ipTableRules

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

// func expandIllumioRuleSetSecurityRules(rules []interface{}) ([]*models.SecurityRule, *diag.Diagnostics) {
// 	var diags diag.Diagnostics
// 	rls := []*models.SecurityRule{}

// 	for _, rule := range rules {
// 		r := rule.(map[string]interface{})
// 		rl := &models.SecurityRule{
// 			Enabled:               r["enabled"].(bool),
// 			Description:           r["description"].(string),
// 			ExternalDataSet:       r["external_data_set"].(string),
// 			ExternalDataReference: r["external_data_reference"].(string),
// 			SecConnect:            r["sec_connect"].(bool),
// 			Stateless:             r["stateless"].(bool),
// 			MachineAuth:           r["machine_auth"].(bool),
// 			UnscopedConsumers:     r["unscoped_consumers"].(bool),
// 		}

// 		if rl.HasConflicts() {
// 			diags = append(diags, diag.Errorf("[illumio-core_rule_set] Exactly one of [\"sec_connect\", \"machine_auth\", \"stateless\"] can be set to true inside rules")...)
// 		}

// 		resLabelAs, errs := expandIllumioRuleSetSRResolveLabelsAs(r["resolve_labels_as"].([]interface{})[0])
// 		diags = append(diags, errs...)
// 		rl.ResolveLabelsAs = resLabelAs

// 		ingServs, errs := expandIllumioRuleSetSRIngressService(
// 			r["ingress_services"].(*schema.Set).List(),
// 			rl.ResolveLabelsAs.ProviderIsVirtualService(),
// 		)
// 		diags = append(diags, errs...)
// 		rl.IngressServices = ingServs

// 		povs, errs := expandIllumioRuleSetSRProviders(r["providers"].(*schema.Set).List())
// 		diags = append(diags, errs...)
// 		rl.Providers = povs

// 		cons, errs := expandIllumioRuleSetSRConsumers(r["consumers"].(*schema.Set).List())
// 		diags = append(diags, errs...)
// 		rl.Consumers = cons

// 		rls = append(rls, rl)
// 	}

// 	return rls, &diags
// }

// func expandIllumioRuleSetSRResolveLabelsAs(o interface{}) (*models.SecurityRuleResolveLabelAs, diag.Diagnostics) {
// 	var diags diag.Diagnostics
// 	resLabelsAs := o.(map[string]interface{})

// 	rProvs := getStringList(resLabelsAs["providers"].(*schema.Set).List())
// 	rCons := getStringList(resLabelsAs["consumers"].(*schema.Set).List())

// 	if !validateList(rProvs, validSRResLabelAsValues) {
// 		diags = append(diags, diag.Errorf(`[illumio-core_rule_set] Invalid value for resolve_value_as.providers, allowed values are "workloads" and "virtual_services" inside rules`)...)
// 	}

// 	if !validateList(rCons, validSRResLabelAsValues) {
// 		diags = append(diags, diag.Errorf(`[illumio-core_rule_set] Invalid value for resolve_value_as.consumers, allowed values are "workloads" and "virtual_services" inside rules`)...)
// 	}

// 	v := &models.SecurityRuleResolveLabelAs{
// 		Providers: rProvs,
// 		Consumers: rCons,
// 	}

// 	return v, diags
// }

// func expandIllumioRuleSetSRIngressService(inServices []interface{}, setEmpty bool) ([]map[string]interface{}, diag.Diagnostics) {
// 	var diags diag.Diagnostics

// 	iss := []map[string]interface{}{}

// 	// Throw error if virtual_services is the only value set in resolve_label_as.provider and ingress_service's resource is non empty
// 	if setEmpty && len(inServices) > 0 {
// 		diags = append(diags, diag.Errorf("[illumio-core_rule_set] If the only value in the providers of resolve_label_as block is \"virtual_services\", then setting ingress_services is not allowed inside rules")...)
// 	}

// 	if !setEmpty {
// 		if len(inServices) == 0 {
// 			diags = append(diags, diag.Errorf("[illumio-core_rule_set] At least one ingress_service must be specified if providers of resolve_label_as block has \"workloads\" inside rules")...)
// 		}
// 		for _, service := range inServices {
// 			s := service.(map[string]interface{})

// 			m := make(map[string]interface{})

// 			if isRuleSetSRIngressServiceSchemaValid(s, &diags) {
// 				if s["href"].(string) != "" {
// 					m["href"] = s["href"].(string)
// 				}

// 				if v, ok := getInt(s["proto"]); ok {
// 					m["proto"] = v
// 					if vPort, ok := getInt(s["port"]); ok {
// 						m["port"] = vPort
// 						if vToPort, ok := getInt(s["to_port"]); ok {
// 							if vToPort <= vPort {
// 								diags = append(diags, diag.Errorf(" [illumio-core_rule_set] Value of to_port can't be less or equal to value of port inside ingress_services, inside rules")...)
// 							} else {
// 								m["to_port"] = vToPort
// 							}
// 						}
// 					}
// 				}
// 			}

// 			iss = append(iss, m)
// 		}
// 	}

// 	return iss, diags
// }

// Validates schema of the security_rule.ingress_service parameter.
//
// Verifes if required fileds are defined or not.
// func isRuleSetSRIngressServiceSchemaValid(s map[string]interface{}, diags *diag.Diagnostics) bool {
// 	hrefOk := s["href"].(string) != ""
// 	protoOk := s["proto"].(string) != ""
// 	portOk := s["port"].(string) != ""
// 	toPortOk := s["to_port"].(string) != ""

// 	switch {
// 	case !hrefOk && !protoOk:
// 		*diags = append(*diags, diag.Errorf("[illumio-core_rule_set] Atleast one of [href, proto] is required inside ingress_services, inside rules")...)

// 	case hrefOk && protoOk:
// 		*diags = append(*diags, diag.Errorf("[illumio-core_rule_set] Exactly one of [href, proto] is allowed inside ingress_services, inside rules")...)

// 	case hrefOk:
// 		if portOk || toPortOk { // If port or to_port are defined with href, return error
// 			*diags = append(*diags, diag.Errorf("[illumio-core_rule_set] port/proto is not allowed with href inside ingress_services, inside rules")...)
// 			return false
// 		}
// 		return true

// 	case protoOk:
// 		if !portOk && toPortOk { // If to_port is defined without port, return error
// 			*diags = append(*diags, diag.Errorf("[illumio-core_rule_set] port is required with to_port inside ingress_services, inside rules")...)
// 			return false
// 		}
// 		return true
// 	}

// 	return false
// }

// func expandIllumioRuleSetSRProviders(providers []interface{}) ([]*models.SecurityRuleProvider, diag.Diagnostics) {
// 	provs := []*models.SecurityRuleProvider{}

// 	for _, provider := range providers {
// 		p := provider.(map[string]interface{})
// 		prov := &models.SecurityRuleProvider{
// 			Actors:         p["actors"].(string),
// 			Label:          getHrefObj(p["label"]),
// 			LabelGroup:     getHrefObj(p["label_group"]),
// 			Workload:       getHrefObj(p["workload"]),
// 			VirtualService: getHrefObj(p["virtual_service"]),
// 			VirtualServer:  getHrefObj(p["virtual_server"]),
// 			IPList:         getHrefObj(p["ip_list"]),
// 		}
// 		if !prov.HasOneActor() {
// 			return nil, diag.Errorf("[illumio-core_rule_set] Provider block can have only one rule actor inside rules")
// 		}

// 		provs = append(provs, prov)
// 	}
// 	return provs, diag.Diagnostics{}
// }

// func expandIllumioRuleSetSRConsumers(consumers []interface{}) ([]*models.SecurityRuleConsumer, diag.Diagnostics) {
// 	cons := []*models.SecurityRuleConsumer{}

// 	for _, consumer := range consumers {
// 		p := consumer.(map[string]interface{})

// 		con := &models.SecurityRuleConsumer{
// 			Actors:         p["actors"].(string),
// 			Label:          getHrefObj(p["label"]),
// 			LabelGroup:     getHrefObj(p["label_group"]),
// 			Workload:       getHrefObj(p["workload"]),
// 			VirtualService: getHrefObj(p["virtual_service"]),
// 			IPList:         getHrefObj(p["ip_list"]),
// 		}

// 		if !con.HasOneActor() {
// 			return nil, diag.Errorf("[illumio-core_rule_set] Consumer block can have only one rule actor inside rules")
// 		}
// 		cons = append(cons, con)
// 	}

// 	return cons, diag.Diagnostics{}
// }

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
		d.Set("scopes", extractResourceScopes(data.S("scopes")))
	} else {
		d.Set("scopes", nil)
	}

	// if data.Exists("rules") {
	// 	d.Set("rules", extractResourceRuleSetSecurityRules(data.S("rules")))
	// } else {
	// 	d.Set("rules", nil)
	// }

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
		m["actors"] = extractResourceRuleActors(ipTableRuleData.S("actors"))
		m["statements"] = extractMapArray(ipTableRuleData.S("statements"), statKeys)

		ms = append(ms, m)
	}
	return ms
}

// func extractResourceRuleSetSecurityRules(data *gabs.Container) []map[string]interface{} {

// 	srKeys := []string{
// 		"href",
// 		"enabled",
// 		"description",
// 		"external_data_set",
// 		"external_data_reference",
// 		"sec_connect",
// 		"stateless",
// 		"machine_auth",
// 		"unscoped_consumers",
// 		"update_type",
// 		"created_at",
// 		"updated_at",
// 		"deleted_at",
// 		"created_by",
// 		"updated_by",
// 		"deleted_by",
// 	}

// 	srs := []map[string]interface{}{}
// 	for _, secRuleData := range data.Children() {
// 		sr := extractMap(secRuleData, srKeys)

// 		rlaKey := "resolve_labels_as"
// 		if secRuleData.Exists(rlaKey) {
// 			sr[rlaKey] = extractSecurityRuleResolveLabelAs(secRuleData.S(rlaKey))
// 		}

// 		prkey := "providers"
// 		if secRuleData.Exists(prkey) {
// 			sr[prkey] = extractResourceRuleActors(secRuleData.S(prkey))
// 		}

// 		cnKeys := "consumers"
// 		if secRuleData.Exists(cnKeys) {
// 			sr[cnKeys] = extractResourceRuleActors(secRuleData.S(cnKeys))
// 		}

// 		isKey := "ingress_services"
// 		if secRuleData.Exists(isKey) {
// 			sr[isKey] = extractResourceSecurityRuleIngressService(secRuleData.S(isKey))
// 		}

// 		srs = append(srs, sr)
// 	}

// 	return srs
// }

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
		// Rules:                 nil,
		IPTablesRules: nil,
	}

	scopes, errs := expandIllumioRuleSetScopes(d.Get("scopes").([]interface{}))
	diags = append(diags, *errs...)
	ruleSet.Scopes = scopes

	// if d.HasChange("rules") {
	// 	rules, errs := expandIllumioRuleSetSecurityRules(d.Get("rules").(*schema.Set).List())
	// 	ruleSet.Rules = rules
	// 	diags = append(diags, *errs...)
	// }

	if d.HasChange("ip_tables_rules") {
		ipTableRules, errs := expandIllumioRuleSetIPTablesRules(d.Get("ip_tables_rules").(*schema.Set).List())
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

	_, err := illumioClient.Delete(href)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "rule_sets", href)

	d.SetId("")
	return diagnostics
}
