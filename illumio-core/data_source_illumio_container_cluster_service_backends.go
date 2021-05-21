// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioContainerClusterServiceBackends() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioContainerClusterServiceBackendsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Container Cluster Service Backend",

		Schema: map[string]*schema.Schema{
			"container_cluster_href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "URI of Container Cluster",
				ValidateDiagFunc: isContainerClusterHref,
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of Container Cluster Service Backends",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Container Cluster Backend",
						},
						"kind": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type (or kind) of Container Cluster Backend",
						},
						"namespace": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The namespace of the Container Cluster Backend",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time (rfc339 timestamp) at which the Container Cluster Backend was last updated",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time (rfc339 timestamp) in which the Container Cluster Backend was created",
						},
						"virtual_service": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Associated virtual service. Single element list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URI to the associated virtual service",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of virtual service",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func datasourceIllumioContainerClusterServiceBackendsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("container_cluster_href").(string) + "/service_backends"

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(href)

	sbsKeys := []string{
		"name",
		"kind",
		"namespace",
		"updated_at",
		"created_at",
	}

	sbs := []map[string]interface{}{}
	for _, child := range data.Children() {
		sb := extractMap(child, sbsKeys)

		if child.Exists("virtual_service") {
			sb["virtual_service"] = []interface{}{
				extractMap(child.S("virtual_service"), []string{"href", "name"}),
			}
		}
		sbs = append(sbs, sb)
	}
	d.Set("items", sbs)

	return diagnostics
}
