// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
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
		SchemaVersion: 1,
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
				Deprecated:  `[illumio-core_container_cluster_workload_profile] assign_labels is deprecated: Use "labels" instead`,
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
							Description: "Key of the Label. The value must be a string between 1 and 64 characters long",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.StringLenBetween(1, LABEL_KEY_LENGTH_MAX),
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
				Description: "Enforcement mode of container workload profiles to return. Allowed values for enforcement modes are \"idle\", \"visibility_only\", \"full\", and \"selective\". Default value: \"idle\"",
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
		CustomizeDiff: customdiff.Sequence(
			customizeAssignedLabels(),
		),
	}
}

// XXX: The /container_cluster_workload_settings schema
// defines both assign_labels (now deprecated) and labels
// fields to represent labels assigned to container workloads
// in a cluster namespace. The two fields are mutually exclusive.
//
// When a profile is created, the PCE will create the other
// entry mirroring the one provided, so given
//
// {"assign_labels": [{"href": "..."}]}
//
// it will return a state with both assign_labels and
//
// {"labels": [{"key": "role", "assignment": {"href": ..., "value": "..."}}]}
//
// This requires the fields to be Optional + Computed as the state
// for one or the other parameter may be set by the PCE. The issue
// then is that empty Computed fields default to the value of the
// remote - adding and then removing assign_labels of labels should
// strip the assigned labels from the profile, but instead defaults
// to whatever value was set previously.
//
// To work around this, we need to check the raw HCL against the
// state when updating and explicitly clear the labels fields.
//
// NB: it seems unintuitive that a CustomizeDiffFunc is needed
// in addition to the logic in the update function to explicitly
// pass an empty list, but even when the read returns an empty
// list after a successful update, the fields retain their
// previously Computed values in state. This may be due to some
// misconfiguration, but without both I was unable to get the
// desired functionality. -- Duncan
func customizeAssignedLabels() schema.CustomizeDiffFunc {
	return func(ctx context.Context, d *schema.ResourceDiff, m any) error {
		if assignedLabelsRemoved(d.GetRawConfig(), d.GetRawState()) {
			// clear the assign_labels and labels values
			d.SetNewComputed("assign_labels")
			d.SetNewComputed("labels")
		}

		return nil
	}
}

func resourceIllumioContainerClusterWorkloadProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	var isAssignLabelsSet bool
	var assignLabels *[]models.Href
	var labels *[]models.ContainerClusterWorkloadProfileLabel

	if _, ok := d.GetOk("assign_labels"); ok {
		assignLabels = PtrTo(models.GetHrefs(d.Get("assign_labels").(*schema.Set).List()))
		isAssignLabelsSet = true
	}

	if _, ok := d.GetOk("labels"); ok {
		if isAssignLabelsSet {
			return append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "[illumio-core_container_cluster_workload_profile] ExactlyOneOf : {\"assign_labels\", \"labels\"} in the block. Please provide one of them",
			})
		}

		labels, diags = expandCCWPLabels(d, diags)

		if diags.HasError() {
			return diags
		}
	}

	ccwp := &models.ContainerClusterWorkloadProfile{
		Name:            PtrTo(d.Get("name").(string)),
		Description:     PtrTo(d.Get("description").(string)),
		EnforcementMode: PtrTo(d.Get("enforcement_mode").(string)),
		Managed:         PtrTo(d.Get("managed").(bool)),
		AssignLabels:    assignLabels,
		Labels:          labels,
	}

	cchref := d.Get("container_cluster_href").(string)
	_, data, err := illumioClient.Create(fmt.Sprintf("%v/container_workload_profiles", cchref), ccwp)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
	return resourceIllumioContainerClusterWorkloadProfileRead(ctx, d, m)
}

func expandCCWPLabels(d *schema.ResourceData, diags diag.Diagnostics) (*[]models.ContainerClusterWorkloadProfileLabel, diag.Diagnostics) {
	labels := d.Get("labels").(*schema.Set).List()
	labelsModel := []models.ContainerClusterWorkloadProfileLabel{}
	for _, label := range labels {
		labelsI := models.ContainerClusterWorkloadProfileLabel{}
		labelMap := label.(map[string]interface{})
		labelsI.Key = PtrTo(labelMap["key"].(string))

		assignmentList := labelMap["assignment"].(*schema.Set).List()
		if len(assignmentList) > 0 {
			t := models.Href{
				Href: assignmentList[0].(map[string]interface{})["href"].(string),
			}
			labelsI.Assignment = &t
		}

		if labelMap["restriction"].(*schema.Set).Len() > 0 {
			restriction := models.GetHrefs(labelMap["restriction"].(*schema.Set).List())
			labelsI.Restriction = &restriction
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

	return &labelsModel, diags
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

	var isAssignLabelsSet bool
	var assignLabels *[]models.Href
	var labels *[]models.ContainerClusterWorkloadProfileLabel

	if assignedLabelsRemoved(d.GetRawConfig(), d.GetRawState()) {
		// prefer to set labels as assign_labels is deprecated in
		// newer versions of the PCE.
		labels = &[]models.ContainerClusterWorkloadProfileLabel{}
	} else {
		if d.HasChange("assign_labels") {
			assignLabels = PtrTo(models.GetHrefs(d.Get("assign_labels").(*schema.Set).List()))
			isAssignLabelsSet = true
		}

		if d.HasChange("labels") {
			if isAssignLabelsSet {
				return append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "[illumio-core_container_cluster_workload_profile] ExactlyOneOf : {\"assign_labels\", \"labels\"} in the block. Please provide one of them",
				})
			}

			labels, diags = expandCCWPLabels(d, diags)

			if diags.HasError() {
				return diags
			}
		}
	}

	ccwp := &models.ContainerClusterWorkloadProfile{
		Name:            PtrTo(d.Get("name").(string)),
		Description:     PtrTo(d.Get("description").(string)),
		EnforcementMode: PtrTo(d.Get("enforcement_mode").(string)),
		Managed:         PtrTo(d.Get("managed").(bool)),
		AssignLabels:    assignLabels,
		Labels:          labels,
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

func assignedLabelsRemoved(conf cty.Value, state cty.Value) bool {
	confMap := conf.AsValueMap()

	if assignLabels, ok := confMap["assign_labels"]; ok {
		if labels, ok := confMap["labels"]; ok {
			if len(assignLabels.AsValueSlice()) == 0 && len(labels.AsValueSlice()) == 0 {
				for _, key := range []string{"assign_labels", "labels"} {
					if keyState, ok := state.AsValueMap()[key]; ok {
						if len(keyState.AsValueSlice()) > 0 {
							return true
						}
					}
				}
			}
		}
	}

	return false
}
