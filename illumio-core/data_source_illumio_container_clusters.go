// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioContainerClusters() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioContainerClustersRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Container Clusters",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Container Clusters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Container Cluster",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Cluster",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the Cluster",
						},
						"nodes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Nodes of the Cluster",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pod_subnet": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Pod Subnet of the node",
									},
								},
							},
						},
						"container_runtime": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Container Runtime used in this Cluster",
						},
						"manager_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Manager for this Cluster (and version)",
						},
						"last_connected": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time the Cluster last connected to",
						},
						"online": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the Cluster is online or not",
						},
						"errors": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Errors for Cluster",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audit_event": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Audit Event of Error",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"duplicate_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Duplicate IDs of Error",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"error_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Error Type of Error",
									},
								},
							},
						},
						"kubelink_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kubelink software version string for Cluster",
						},
						"pce_fqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PCE FQDN for this container cluster. Used in Supercluster only",
						},
						"caps": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Permission types",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of container clusters to return. The integer should be a non-zero positive integer",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the container cluster(s) to return. Supports partial matches",
			},
		},
	}
}

func datasourceIllumioContainerClustersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	paramKeys := []string{
		"max_results",
		"name",
	}

	params := resourceDataToMap(d, paramKeys)
	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%d/container_clusters", orgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}

	keys := []string{
		"href",
		"name",
		"description",
		"container_runtime",
		"manager_type",
		"last_connected",
		"online",
		"kubelink_version",
		"pce_fqdn",
		"caps",
	}

	for _, child := range data.Children() {
		m := extractMap(child, keys)

		for key, value := range map[string][]string{
			"nodes":  {"pod_subnet"},
			"errors": {"audit_event", "duplicate_ids", "error_type"},
		} {
			m[key] = extractDataSourceAttrs(child, key, value)
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
