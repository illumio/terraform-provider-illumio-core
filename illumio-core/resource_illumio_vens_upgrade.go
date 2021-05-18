// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package illumiocore

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioVENsUpgrade() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioVENsUpgradeCreate,
		ReadContext:   resourceIllumioVENsUpgradeRead,
		UpdateContext: resourceIllumioVENsUpgradeUpdate,
		DeleteContext: resourceIllumioVENsUpgradeDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio VENs Upgrade",

		Schema: map[string]*schema.Schema{
			"vens": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Set: func(i interface{}) int {
					m := i.(map[string]interface{})
					return hashcode(m["href"].(string))
				},
				Description: "List of VENs to unpair. Max Items allowed: 25000",
				Elem:        hrefSchemaRequired("VEN", isVENHref),
			},
			"release": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The software release to upgrade to`,
			},
		},
	}
}

func resourceIllumioVENsUpgradeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	orgID := pConfig.OrgID

	href := fmt.Sprintf("/orgs/%v/vens/upgrade", orgID)

	vup := expandIllumioVENsUpgrade(d, false, nil)

	responce, err := illumioClient.Update(href, vup)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(href)

	// Ignoring the read operation
	return handleUnpairAndUpgradeOperationErrors(err, responce, "upgrade", "vens")

}

func expandIllumioVENsUpgrade(d *schema.ResourceData, isUpdateCall bool, diags *diag.Diagnostics) *models.VENsUpgrade {
	hfs := []models.Href{}

	for _, v := range d.Get("vens").(*schema.Set).List() {
		hfs = append(hfs, *getHrefObj(v))
	}

	return &models.VENsUpgrade{
		Release: d.Get("release").(string),
		Hrefs:   hfs,
	}
}

func resourceIllumioVENsUpgradeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceIllumioVENsUpgradeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	vup := expandIllumioVENsUpgrade(d, true, &diags)

	if diags != nil && len(diags) > 0 {
		return diags
	}

	responce, err := illumioClient.Update(d.Id(), vup)

	// Ignoring read operation
	return handleUnpairAndUpgradeOperationErrors(err, responce, "upgrade", "vens")

}

func resourceIllumioVENsUpgradeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	log.Println("[WARN] [resource] [illumio-core_vens_upgrade] Ignoring delete operation. VEN's upgrade operation does not support delete operation")

	return diag.Diagnostics{}
}
