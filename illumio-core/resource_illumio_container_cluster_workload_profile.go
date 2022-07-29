// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioContainerClusterWorkloadProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioContainerClusterWorkloadProfileCreate,
		ReadContext:   resourceIllumioContainerClusterWorkloadProfileRead,
		UpdateContext: resourceIllumioContainerClusterWorkloadProfileUpdate,
		DeleteContext: resourceIllumioContainerClusterWorkloadProfileDelete,
		SchemaVersion: version,
		Description:   "Manages Illumio Container Cluster",

		Schema: map[string]*schema.Schema{
			"container_cluster_href": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "URI of Container Cluster",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the container workload profile",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "A friendly name given to a profile if the namespace is not user-friendly. The name should be up to 255 characters",
				ValidateDiagFunc: nameValidation,
			},
			"namespace": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Namespace name of the container workload profile",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the container workload profile",
			},
			"assign_labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Assigned labels container workload profile",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "URI of the assigned label",
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Labels to assign to the workload that matches the namespace",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Key of the Label. Allowed values for key are \"role\", \"loc\", \"app\" and \"env\"",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.StringInSlice(validLabelKeys, false),
							),
						},
						"assignment": {
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    1,
							Description: "The label href to set. Single element list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "URI of label",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of label",
									},
								},
							},
						},
						"restriction": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "The list of allowed label hrefs",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "URI of label",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of label",
									},
								},
							},
						},
					},
				},
			},
			"enforcement_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "idle",
				Description: "Enforcement mode of container workload profiles to return. Allowed values for enforcement modes are \"idle\",\"visibility_only\",\"full\", and \"selective\". Default value: \"idle\" ",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(ValidWorkloadEnforcementModeValues, false),
				),
			},
			"visibility_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Visibility Level of the container cluster workload profile",
			},
			"managed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If the namespace is managed or not",
			},
			"linked": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if the namespace exists in the cluster and is reported by Kubelink",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this container workload profile was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this container workload profile was last updated",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this container workload profile",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this container workload profile",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioContainerClusterWorkloadProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	// orgID := pConfig.OrgID

	cchref := d.Get("container_cluster_href").(string)

	ccwp := &models.ContainerClusterWorkloadProfile{
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		EnforcementMode: d.Get("enforcement_mode").(string),
		Managed:         d.Get("managed").(bool),
	}

	var isAssignLabelsAvailable bool
	var isLabelsAvailable bool

	if items, ok := d.GetOk("assign_labels"); ok {
		ccwp.AssignLabels = models.GetHrefs(items.(*schema.Set).List())
		isAssignLabelsAvailable = true
	}

	if items, ok := d.GetOk("labels"); ok {
		labels := items.(*schema.Set).List()
		labelsModel := []models.ContainerClusterWorkloadProfileLabel{}
		for _, label := range labels {
			labelsI := models.ContainerClusterWorkloadProfileLabel{}
			labelMap := label.(map[string]interface{})
			labelsI.Key = labelMap["key"].(string)

			assList := labelMap["assignment"].(*schema.Set).List()
			if len(assList) > 0 {
				t := models.Href{
					Href: assList[0].(map[string]interface{})["href"].(string),
				}
				labelsI.Assignment = t
			}

			if labelMap["restriction"].(*schema.Set).Len() > 0 {
				labelsI.Restriction = models.GetHrefs(labelMap["restriction"].(*schema.Set).List())
			}

			if labelsI.HasConflicts() {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "[illumio-core_container_cluster_workload_profile] ExactlyOneOf : {\"assignment\", \"restriction\"} in the labels block",
					AttributePath: cty.Path{
						cty.GetAttrStep{
							Name: "labels",
						},
					},
				})
			}
			labelsModel = append(labelsModel, labelsI)
		}
		ccwp.Labels = labelsModel
		isLabelsAvailable = true
	}

	if isAssignLabelsAvailable && isLabelsAvailable {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "[illumio-core_container_cluster_workload_profile] ExactlyOneOf : {\"assign_labels\", \"labels\"} in the block. Please provide one of them",
		})
	}

	if diags.HasError() {
		return diags
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("%v/container_workload_profiles", cchref), ccwp)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
	isAssignLabelsAvailable = false
	isLabelsAvailable = false
	return resourceIllumioContainerClusterWorkloadProfileRead(ctx, d, m)
}

func resourceIllumioContainerClusterWorkloadProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// extract the parent HREF and set it
	href = data.S("href").Data().(string)
	d.Set("container_cluster_href", getParentHref(href))

	d.SetId(href)
	for _, key := range []string{
		"href",
		"name",
		"namespace",
		"description",
		"enforcement_mode",
		"managed",
		"linked",
		"visibility_level",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("assign_labels") {
		ALKeys := []string{"href"}
		ALs := data.S("assign_labels")
		ALI := []map[string]interface{}{}

		for _, AL := range ALs.Children() {
			ALI = append(ALI, extractMap(AL, ALKeys))
		}

		d.Set("assign_labels", ALI)
	} else {
		d.Set("assign_labels", nil)
	}

	if data.Exists("labels") {
		labelKeys := []string{
			"key",
			"assignment",
			"restriction",
		}
		labels := data.S("labels")
		labelI := []map[string]interface{}{}

		for _, label := range labels.Children() {
			labelMap := extractMap(label, labelKeys)
			if label.Exists("assignment") {
				labelMap["assignment"] = []interface{}{extractMap(label.S("assignment"), []string{"href", "value"})}
			} else {
				labelMap["assignment"] = nil
			}
			if label.Exists("restriction") {
				labelMap["restriction"] = extractMapArray(label.S("restriction"), []string{"href", "value"})
			} else {
				labelMap["restriction"] = nil
			}
			labelI = append(labelI, labelMap)
		}
		d.Set("labels", labelI)
	} else {
		d.Set("labels", nil)
	}

	return diagnostics
}

func resourceIllumioContainerClusterWorkloadProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	ccwp := &models.ContainerClusterWorkloadProfile{
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		EnforcementMode: d.Get("enforcement_mode").(string),
		Managed:         d.Get("managed").(bool),
	}

	if d.HasChange("assign_labels") {
		items := d.Get("assign_labels")
		ccwp.AssignLabels = models.GetHrefs(items.(*schema.Set).List())
	}

	if d.HasChange("labels") {
		items := d.Get("labels")
		labels := items.(*schema.Set).List()
		labelsModel := []models.ContainerClusterWorkloadProfileLabel{}
		for _, label := range labels {
			labelsI := models.ContainerClusterWorkloadProfileLabel{}
			labelMap := label.(map[string]interface{})
			labelsI.Key = labelMap["key"].(string)

			assList := labelMap["assignment"].(*schema.Set).List()

			if len(assList) > 0 {
				t := models.Href{
					Href: assList[0].(map[string]interface{})["href"].(string),
				}
				labelsI.Assignment = t
			}
			if labelMap["restriction"].(*schema.Set).Len() > 0 {
				labelsI.Restriction = models.GetHrefs(labelMap["restriction"].(*schema.Set).List())
			}

			if labelsI.HasConflicts() {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "[illumio-core_container_cluster_workload_profile] ExactlyOneOf : {\"assignment\", \"restriction\"} in the labels block",
					AttributePath: cty.Path{
						cty.GetAttrStep{
							Name: "labels",
						},
					},
				})
			}
			labelsModel = append(labelsModel, labelsI)
		}
		ccwp.Labels = labelsModel
	}

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), ccwp)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioContainerClusterWorkloadProfileRead(ctx, d, m)
}

func resourceIllumioContainerClusterWorkloadProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
