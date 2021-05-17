package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceIllumioContainerClusterWorkloadProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioContainerClusterWorkloadProfilesRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Container Cluster Workload Profiles",

		Schema: map[string]*schema.Schema{
			"container_cluster_href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of the Container Cluster",
				ValidateDiagFunc: isContainerClusterHref,
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Container Cluster Workload Profiles",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Container Cluster Workload Profile",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A friendly name given to a profile if the namespace is not user friendly",
						},
						"namespace": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Namespace name of the container workload profile",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the container workload profile",
						},
						"assign_labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Assigned labels container workload profile",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of the assigned label",
									},
								},
							},
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Labels to assign to the workload that matches the namespace",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Key of the label",
									},
									"assignment": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The label href to set. Single element list",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"href": {
													Type:        schema.TypeString,
													Computed:    true,
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
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of allowed label hrefs.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"href": {
													Type:        schema.TypeString,
													Computed:    true,
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
							Computed:    true,
							Description: "Enforcement mode of container workload profiles to return.",
						},
						"visibility_level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Visibility Level of the container cluster workload profile.",
						},
						"managed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If the namespace is managed or not.",
						},
						"linked": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "True if the namespace exists in the cluster and is reported by kubelink.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this label group was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this label group was last updated",
						},
						"created_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who originally created this label group",
						},
						"updated_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who last updated this label group",
						},
					},
				},
			},
			"assign_labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of label URIs, encoded as a JSON string",
			},
			"enforcement_mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(ValidWorkloadEnforcementModeValues, false),
				),
				Description: "Filter by enforcement mode. Allowed values for enforcement modes are \"idle\",\"visibility_only\",\"full\", and \"selective\".",
			},
			"linked": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Filter by linked container workload profiles.",
			},
			"managed": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Filter by managed state",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of enforcement boundaries to return. The integer should be a non-zero positive integer.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name string to match.Supports partial matches.",
			},
			"namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Namespace string to match.Supports partial matches.",
			},
			"visibility_level": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(validVisibilityLevels, false),
				),
				Description: `Filter by visibility level. Allowed values are "flow_full_detail", "flow_summary", "flow_drops", "flow_off" and "enhanced_data_collection"`,
			},
		},
	}
}

func dataSourceIllumioContainerClusterWorkloadProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	paramKeys := []string{
		"assign_labels",
		"enforcement_mode",
		"linked",
		"managed",
		"max_results",
		"name",
		"namespace",
		"visibility_level",
	}

	params := resourceDataToMap(d, paramKeys)
	_, data, err := illumioClient.AsyncGet(d.Get("container_cluster_href").(string)+"/container_workload_profiles", &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}

	keys := []string{
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
	}

	for _, child := range data.Children() {
		m := extractMap(child, keys)

		key := "assign_labels"
		m[key] = extractDataSourceAttrs(child, key, []string{"href"})

		if child.Exists("labels") {
			labelKeys := []string{
				"key",
				"assignment",
				"restriction",
			}
			labels := child.S("labels")
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
			m["labels"] = labelI
		} else {
			m["labels"] = nil
		}

		dataMap = append(dataMap, m)

	}

	d.Set("items", dataMap)

	return diagnostics
}
