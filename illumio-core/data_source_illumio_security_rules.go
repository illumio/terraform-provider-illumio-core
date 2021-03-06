// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample
[
	{
		"href": "string",
		"enabled": true,
		"description": "string",
		"external_data_set": null,
		"external_data_reference": null,
		"ingress_services": [
			{
			"href": "string"
			}
		],
		"resolve_labels_as": {
			"providers": [
			"workloads"
			],
			"consumers": [
			"workloads"
			]
		},
		"sec_connect": true,
		"stateless": true,
		"machine_auth": true,
		"providers": [
			{
			"actors": "ams",
			"label": {
				"href": "string"
			},
			"label_group": {
				"href": "string"
			},
			"workload": {
				"href": "string"
			},
			"virtual_service": {
				"href": "string"
			},
			"virtual_server": {
				"href": "string"
			},
			"ip_list": {
				"href": "string"
			}
			}
		],
		"consumers": [
			{
			"actors": "ams",
			"label": {
				"href": "string"
			},
			"label_group": {
				"href": "string"
			},
			"workload": {
				"href": "string"
			},
			"virtual_service": {
				"href": "string"
			},
			"ip_list": {
				"href": "string"
			}
			}
		],
		"unscoped_consumers": true,
		"update_type": "string"
	}
]
*/

func datasourceIllumioSecurityRules() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioSecurityRulesRead,
		SchemaVersion: version,
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

	srs := []map[string]interface{}{}
	for _, rule := range data.Children() {
		sr := extractMap(rule, rlKeys)

		rlaKey := "resolve_labels_as"
		if rule.Exists(rlaKey) {
			sr[rlaKey] = extractSecurityRuleResolveLabelAs(rule.S(rlaKey))
		}

		isKey := "ingress_services"
		if rule.Exists(isKey) {
			isKeys := []string{
				"href",
				"proto",
				"port",
				"to_port",
			}

			sr[isKey] = extractMapArray(rule.S(isKey), isKeys)
		}

		providersKey := "providers"
		if rule.Exists(providersKey) {
			sr[providersKey] = extractDatasourceActors(rule.S(providersKey))
		}

		consumerKey := "consumers"
		if rule.Exists(consumerKey) {
			sr[consumerKey] = extractDatasourceActors(rule.S(consumerKey))
		}

		srs = append(srs, sr)
	}

	d.Set("items", srs)

	return diagnostics
}
