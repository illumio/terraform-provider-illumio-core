// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioSecurityRules() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioSecurityRulesRead,
		SchemaVersion: 1,
		Description:   "Represents Illumio Security Rules",

		Schema: map[string]*schema.Schema{
			"rule_set_href": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URI of ruleset",
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of Security Rule hrefs",
				Elem: &schema.Resource{
					Schema: securityRuleDatasourceSchema(false),
				},
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
		},
	}
}

func datasourceIllumioSecurityRulesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	ruleSetHref := d.Get("rule_set_href").(string)

	paramKeys := []string{
		"rule_set_href",
		"external_data_reference",
		"external_data_set",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.Get(fmt.Sprintf("%v/sec_rules", ruleSetHref), &params)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	d.Set("items", extractRules(data))

	return diagnostics
}

func extractRules(data *gabs.Container) []map[string]any {
	rlKeys := []string{
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
	}

	srs := []map[string]any{}

	for _, rule := range data.Children() {
		sr := extractMap(rule, rlKeys)

		key := "resolve_labels_as"
		if rule.Exists(key) {
			sr[key] = extractSecurityRuleResolveLabelAs(rule.S(key))
		}

		key = "ingress_services"
		if rule.Exists(key) {
			sr[key] = extractSecurityRuleIngressService(rule.S(key))
		} else {
			sr[key] = nil
		}

		key = "use_workload_subnets"
		if rule.Exists(key) {
			sr[key] = getStringList(rule.S(key).Data().([]any))
		}

		key = "providers"
		if rule.Exists(key) {
			sr[key] = extractRuleActors(rule.S(key))
		}

		key = "consumers"
		if rule.Exists(key) {
			sr[key] = extractRuleActors(rule.S(key))
		}

		srs = append(srs, sr)
	}

	return srs
}
