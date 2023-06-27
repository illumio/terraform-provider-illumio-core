// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioContainerCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioContainerClusterRead,
		SchemaVersion: 1,
		Description:   "Represents Illumio Container Cluster",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of the Cluster",
				ValidateDiagFunc: isContainerClusterHref,
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
				Description: "User permissions for the object",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"container_cluster_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Convenience variable for the cluster UUID contained in the HREF",
			},
		},
	}
}

func datasourceIllumioContainerClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	resourceIllumioContainerClusterReadResult(d, data)

	return diagnostics
}
