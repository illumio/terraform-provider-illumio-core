// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioVulnerabilities() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioVulnerabilitiesCreate,
		ReadContext:   resourceIllumioVulnerabilitiesRead,
		UpdateContext: resourceIllumioVulnerabilitiesUpdate,
		DeleteContext: resourceIllumioVulnerabilitiesDelete,

		SchemaVersion: 1,
		Description:   "Manages Illumio Vulnerabilities",

		Schema: map[string]*schema.Schema{
			"vulnerability": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Description: "Collection of Vulnerabilites",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"reference_id": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotWhiteSpace),
							Description:      "reference id of vulnerability",
						},
						"score": {
							Type:             schema.TypeInt,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(0, 100)),
							Description:      "The normalized score of the vulnerability within the range of 0 to 100. CVSS Score can be used here with a 10x multiplier",
						},
						"cve_ids": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The cve_ids for the vulnerability",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "An arbitrary field to store some details of the vulnerability class",
						},
						"name": { // Name restrictions are not applied in API
							Type:        schema.TypeString,
							Required:    true,
							Description: "The title/name of the vulnerability",
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

func resourceIllumioVulnerabilitiesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	err := makeBatchedClientCalls(d, pConfig)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("/orgs/%v/vulnerabilities", illumioClient.OrgID))

	return resourceIllumioVulnerabilitiesRead(ctx, d, m)
}

func makeBatchedClientCalls(d *schema.ResourceData, pConfig Config) error {
	illumioClient := pConfig.IllumioClient
	href := fmt.Sprintf("/orgs/%v/vulnerabilities", illumioClient.OrgID)

	if v, ok := d.GetOk("vulnerability"); ok {
		batch := batchifyVulnerabilityList(v.([]interface{}))

		for _, j := range batch {
			_, _, err := pConfig.IllumioClient.Create(href, j)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func batchifyVulnerabilityList(vulnerabilities []interface{}) []*models.VulnerabilityList {
	listSize := len(vulnerabilities)

	result := []*models.VulnerabilityList{}

	for i := 0; i < listSize; i = i + 1000 {
		ub := i
		if listSize-i >= 1000 {
			ub += 1000
		} else {
			ub += listSize % 1000
		}

		result = append(result, expandIllumioVulnerabilityList(vulnerabilities[i:ub]))
	}

	return result
}

func expandIllumioVulnerabilityList(vulnerabilitylist []interface{}) *models.VulnerabilityList {
	list := []models.Vulnerability{}
	for _, vulnerability := range vulnerabilitylist {
		v := vulnerability.(map[string]interface{})
		list = append(list, models.Vulnerability{
			ReferenceID: v["reference_id"].(string),
			Score:       v["score"].(int),
			CveIds:      getStringList(v["cve_ids"].(*schema.Set).List()),
			Description: v["description"].(string),
			Name:        v["name"].(string),
		})
	}

	return &models.VulnerabilityList{
		Values: list,
	}
}

func resourceIllumioVulnerabilitiesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceIllumioVulnerabilitiesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)

	err := makeBatchedClientCalls(d, pConfig)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioVulnerabilitiesRead(ctx, d, m)
}

func resourceIllumioVulnerabilitiesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Ignoring Delete operation.",
			Detail:   "Delete operation is not supported for vulnerabilites resource.",
		},
	}
}
