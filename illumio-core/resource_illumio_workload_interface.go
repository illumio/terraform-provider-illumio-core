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

func resourceIllumioWorkloadInterface() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioWorkloadInterfaceCreate,
		ReadContext:   resourceIllumioWorkloadInterfaceRead,
		UpdateContext: resourceIllumioWorkloadInterfaceUpdate,
		DeleteContext: resourceIllumioWorkloadInterfaceDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio Workload Interface",

		Schema: map[string]*schema.Schema{
			"workload_href": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URI of Workload",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the Workload Interface",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the Workload Interface",
				ValidateDiagFunc: nameValidation,
			},
			"link_state": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Link State for Workload Interface. Allowed values are \"up\", \"down\", and \"unknown\" ",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(ValidWorkloadInterfaceLinkStateValues, false),
				),
			},
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IP Address to assign to this interface. The address should in the IPv4 or IPv6 format.",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.IsIPAddress,
				),
			},
			"cidr_block": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "CIDR BLOCK of the Workload Interface.",
			},
			"default_gateway_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IP Address of the default gateway. The Default Gateaway Address should in the IPv4 or IPv6 format.",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.IsIPAddress,
				),
			},
			"network": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Network for the Workload Interface.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"loopback": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Loopback for Workload Interface",
			},
			"network_detection_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network Detection Mode for Workload Interface",
			},
			"friendly_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "User-friendly name for Workload Interface",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
		},
	}
}

func resourceIllumioWorkloadInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	// orgID := pConfig.OrgID

	WorkloadInterface := &models.WorkloadInterface{
		Name:                  d.Get("name").(string),
		LinkState:             d.Get("link_state").(string),
		Address:               d.Get("address").(string),
		CidrBlock:             d.Get("cidr_block").(int),
		DefaultGatewayAddress: d.Get("default_gateway_address").(string),
		FriendlyName:          d.Get("friendly_name").(string),
	}

	wHref := d.Get("workload_href").(string)
	_, data, err := illumioClient.Create(fmt.Sprintf("%v/interfaces", wHref), WorkloadInterface)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.S("href").Data().(string))
	return resourceIllumioWorkloadInterfaceRead(ctx, d, m)
}

func resourceIllumioWorkloadInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
	for _, key := range []string{
		"href",
		"name",
		"link_state",
		"address",
		"cidr_block",
		"default_gateway_address",
		"network",
		"loopback",
		"network_detection_mode",
		"friendly_name",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	return diagnostics
}

func resourceIllumioWorkloadInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Detail:   "Cannot Update the Workload Interface Resource.",
		Summary:  "Ignoring the Update...",
	})

	return diags
}

func resourceIllumioWorkloadInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()
	_, err := illumioClient.Delete(href)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diagnostics
}
