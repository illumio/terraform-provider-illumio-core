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

func resourceIllumioServiceBinding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioServiceBindingCreate,
		ReadContext:   resourceIllumioServiceBindingRead,
		UpdateContext: resourceIllumioServiceBindingUpdate,
		DeleteContext: resourceIllumioServiceBindingDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio Service Binding",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the Service Binding",
			},
			"bound_service": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Bound service href",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"virtual_service": {
				Type:        schema.TypeSet,
				Required:    true,
				MaxItems:    1,
				Description: "Virtual service href",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Virtual Service URI",
						},
					},
				},
			},
			"workload": {
				Type:         schema.TypeSet,
				Optional:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"workload", "container_workload"},
				Description:  "Workload Object for Service Bindings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Workload URI",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workload Name",
						},
						"hostname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Workload Hostname",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether the workload is deleted",
						},
					},
				},
			},
			"container_workload": {
				Type:        schema.TypeSet,
				Optional:    true,
				MaxItems:    1,
				Description: "Container Workload href",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Container Workload URI",
						},
					},
				},
			},
			"port_overrides": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Port Overrides for Service Bindings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:             schema.TypeInt,
							Optional:         true,
							Description:      "Port Number in the original service which to override (integer 0-65535). Starting port when specifying a range",
							ValidateDiagFunc: portNumberValidation,
						},
						"proto": {
							Type:             schema.TypeInt,
							Optional:         true,
							Description:      "Transport protocol in the original service which to override. Allowed values are 6 (TCP) and 17 (UDP)",
							ValidateDiagFunc: validation.ToDiagFunc(validation.IntInSlice([]int{6, 17})),
						},
						"new_port": {
							Type:             schema.TypeInt,
							Required:         true,
							Description:      "Overriding port number (or starting point when specifying a range). Allowed range is 0 - 65535",
							ValidateDiagFunc: portNumberValidation,
						},
						"new_to_port": {
							Type:             schema.TypeInt,
							Optional:         true,
							Description:      "Overriding port range ending port. Allowed range is 0 - 65535",
							ValidateDiagFunc: portNumberValidation,
						},
					},
				},
			},
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "External Data Set identifier",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "External Data reference identifier",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioServiceBindingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	ServiceBinding := &models.ServiceBinding{
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		PortOverrides:         []models.ServiceBindingPortOverrides{},
	}

	item := d.Get("virtual_service").(*schema.Set).List()
	ServiceBinding.VirtualService.Href = item[0].(map[string]interface{})["href"].(string)

	item = d.Get("workload").(*schema.Set).List()
	if len(item) > 0 {
		ServiceBinding.Workload = &models.Href{
			Href: item[0].(map[string]interface{})["href"].(string),
		}
	}

	item = d.Get("container_workload").(*schema.Set).List()
	if len(item) > 0 {
		ServiceBinding.ContainerWorkload = &models.Href{
			Href: item[0].(map[string]interface{})["href"].(string),
		}
	}

	if items, ok := d.GetOk("port_overrides"); ok {
		pos := items.(*schema.Set).List()
		posModel := []models.ServiceBindingPortOverrides{}
		for _, po := range pos {
			poI := models.ServiceBindingPortOverrides{}
			poMap := po.(map[string]interface{})
			poI.Port = poMap["port"].(*int)
			poI.NewPort = poMap["new_port"].(int)
			poI.NewToPort = poMap["new_to_port"].(*int)
			poI.Proto = poMap["proto"].(*int)
			posModel = append(posModel, poI)
		}
		ServiceBinding.PortOverrides = posModel
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/service_bindings", orgID), ServiceBinding)
	if err != nil {
		return diag.FromErr(err)
	}
	status := data.Children()[0].S("status").Data()
	if data.Children()[0].Exists("href") {
		d.SetId(data.Children()[0].S("href").Data().(string))
	} else {
		var diags diag.Diagnostics

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Detail:   "[illumio-core_service_binding] Error occured while creating the Service Binding",
			Summary:  fmt.Sprintln("Error Status: ", status),
		})

		return diags
	}

	return resourceIllumioServiceBindingRead(ctx, d, m)
}

func resourceIllumioServiceBindingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		"bound_service",
		"external_data_set",
		"external_data_reference",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("port_overrides") {
		poS := data.S("port_overrides")
		poI := []map[string]interface{}{}

		for _, po := range poS.Children() {
			poI = append(poI, extractMap(po, []string{"port", "proto", "new_port", "new_to_port"}))
		}

		d.Set("port_overrides", poI)
	} else {
		d.Set("port_overrides", nil)
	}

	if data.Exists("workload") {
		d.Set("workload", []interface{}{extractMap(data.S("workload"), []string{"href", "name", "hostname", "deleted"})})
	} else {
		d.Set("workload", nil)
	}

	for _, x := range []string{"virtual_service", "container_workload"} {
		if data.Exists(x) {
			d.Set(x, []interface{}{extractMap(data.S(x), []string{"href"})})
		} else {
			d.Set(x, nil)
		}
	}

	return diagnostics
}

func resourceIllumioServiceBindingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Detail:   "Cannot Update the Service Binding Resource.",
		Summary:  "Ignoring the Update...",
	})

	return diags
}

func resourceIllumioServiceBindingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
