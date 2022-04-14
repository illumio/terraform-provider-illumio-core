// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioVENsUnpair() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioVENsUnpairCreate,
		ReadContext:   resourceIllumioVENsUnpairRead,
		UpdateContext: resourceIllumioVENsUnpairUpdate,
		DeleteContext: resourceIllumioVENsUnpairDelete,

		SchemaVersion:      version,
		Description:        "Manages Illumio VENs Unpair",
		DeprecationMessage: "DEPRECATED in v0.2.0. Will be removed in v1.0.0. Use resource/unmanaged_workload and resource/managed_workload instead.",

		Schema: map[string]*schema.Schema{
			"vens": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Set: func(i interface{}) int {
					m := i.(map[string]interface{})
					return hashcode(m["href"].(string))
				},
				Description: "List of VENs to unpair. Max Items allowed: 1000",
				Elem:        hrefSchemaRequired("VEN", isVENHref),
			},
			"firewall_restore": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "default",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"saved", "default", "disable"}, false)),
				Description:      `The strategy to use to restore the firewall state after the VEN is uninstalled. Allowed values are "saved", "default" and "disable". Default value: "default"`,
			},
		},
	}
}

func resourceIllumioVENsUnpairCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	orgID := illumioClient.OrgID

	href := fmt.Sprintf("/orgs/%v/vens/unpair", orgID)

	vup := expandIllumioVENsUnpair(d, false, nil)

	responce, err := illumioClient.Update(href, vup)

	d.SetId(href)

	// Ignoring the read operation
	return handleUnpairAndUpgradeOperationErrors(err, responce, "unpair", "vens")
}

func expandIllumioVENsUnpair(d *schema.ResourceData, isUpdateCall bool, diags *diag.Diagnostics) *models.VENsUnpair {
	hfs := []models.Href{}

	for _, v := range d.Get("vens").(*schema.Set).List() {
		hfs = append(hfs, *getHrefObj(v))
	}

	return &models.VENsUnpair{
		FirewallRestore: d.Get("firewall_restore").(string),
		Hrefs:           hfs,
	}
}

func resourceIllumioVENsUnpairRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceIllumioVENsUnpairUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	vup := expandIllumioVENsUnpair(d, true, &diags)

	responce, err := illumioClient.Update(d.Id(), vup)

	// Ignoring read operation
	return handleUnpairAndUpgradeOperationErrors(err, responce, "unpair", "vens")
}

func resourceIllumioVENsUnpairDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	log.Println("[WARN] [resource] [illumio-core_vens_unpair] Ignoring delete operation. VEN's Unpair operation does not support delete operation")
	return diag.Diagnostics{}
}
