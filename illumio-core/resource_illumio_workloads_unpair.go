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

func resourceIllumioWorkloadsUnpair() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioWorkloadsUnpairCreate,
		ReadContext:   resourceIllumioWorkloadsUnpairRead,
		UpdateContext: resourceIllumioWorkloadsUnpairUpdate,
		DeleteContext: resourceIllumioWorkloadsUnpairDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio Workloads Unpair",

		Schema: map[string]*schema.Schema{
			"workloads": {
				Type:     schema.TypeSet,
				Required: true,
				MinItems: 1,
				Set: func(i interface{}) int {
					m := i.(map[string]interface{})
					return hashcode(m["href"].(string))
				},
				Description: "List of Workloads to unpair. Max Items allowed: 1000",
				Elem:        hrefSchemaRequired("Workload", isWorkloadHref),
			},
			"ip_table_restore": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "default",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"saved", "default", "disable"}, false)),
				Description:      `The desired state of IP tables after the agent is uninstalled. Allowed values are "saved", "default" and "disable". Default value: "default"`,
			},
		},
	}
}

func resourceIllumioWorkloadsUnpairCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	orgID := pConfig.OrgID

	href := fmt.Sprintf("/orgs/%v/workloads/unpair", orgID)

	vup := expandIllumioWorkloadsUnpair(d, false, nil)

	responce, err := illumioClient.Update(href, vup)

	d.SetId(href)

	// Ignoring the read operation
	return handleUnpairAndUpgradeOperationErrors(err, responce, "unpair", "workloads")
}

func expandIllumioWorkloadsUnpair(d *schema.ResourceData, isUpdateCall bool, diags *diag.Diagnostics) *models.WorkloadsUnpair {
	hfs := []models.Href{}

	for _, v := range d.Get("workloads").(*schema.Set).List() {
		hfs = append(hfs, *getHrefObj(v))
	}

	return &models.WorkloadsUnpair{
		IPTableRestore: d.Get("ip_table_restore").(string),
		Hrefs:          hfs,
	}
}

func resourceIllumioWorkloadsUnpairRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceIllumioWorkloadsUnpairUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	vup := expandIllumioWorkloadsUnpair(d, true, &diags)

	responce, err := illumioClient.Update(d.Id(), vup)

	// Ignoring read operation
	return handleUnpairAndUpgradeOperationErrors(err, responce, "unpair", "workloads")
}

func resourceIllumioWorkloadsUnpairDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	log.Println("[WARN] [resource] [illumio-core_workloads_unpair] Ignoring delete operation. VEN's Unpair operation does not support delete operation")

	return diag.Diagnostics{}
}
