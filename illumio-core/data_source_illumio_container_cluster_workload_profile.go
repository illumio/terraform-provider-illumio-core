// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Sample

/*
{
  "href": "string",
  "name": null,
  "namespace": null,
  "description": "string",
  "assign_labels": [
    {
      "href": "string"
    }
  ],
  "labels": [
    {
      "key": "string",
      "assignment": {
        "href": "string",
        "value": "string"
      }
    }
  ],
  "enforcement_mode": "idle",
  "managed": true,
  "linked": true,
  "created_at": "2021-03-02T02:37:59Z",
  "created_by": {
    "href": "string"
  },
  "updated_by": {
    "href": "string"
  },
  "updated_at": "2021-03-02T02:37:59Z"
}
*/

func datasourceIllumioContainerClusterWorkloadProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioContainerClusterWorkloadProfileRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Container Cluster Workload Profile",

		Schema: map[string]*schema.Schema{
			"container_cluster_href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of Container Cluster",
			},
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of the container workload profile",
				ValidateDiagFunc: isContainerClusterWorkloadProfileHref,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A friendly name given to a profile if the namespace is not user-friendly",
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
							Description: "The list of allowed label hrefs",
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
				Description: "Enforcement mode of container workload profiles to return",
			},
			"visibility_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Visibility Level of the container cluster workload profile",
			},
			"managed": {
				Type:        schema.TypeBool,
				Computed:    true,
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
				Description: "User who created this label group",
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
	}
}

func datasourceIllumioContainerClusterWorkloadProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	// orgID := pConfig.OrgID
	href := d.Get("href").(string)
	d.Set("container_cluster_href", getParentHref(href))

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))

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
