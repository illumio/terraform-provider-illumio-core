package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Sample

/*
{
  "href": "string",
  "name": "string",
  "description": "string",
  "nodes": [
    {
      "pod_subnet": "string"
    }
  ],
  "container_runtime": "string",
  "manager_type": "string",
  "last_connected": "2021-03-02T02:37:59Z",
  "online": true,
  "errors": [
    {
      "audit_event": {
        "href": "string"
      },
      "duplicate_ids": [],
      "error_type": "string"
    }
  ],
  "kubelink_version": "string",
  "pce_fqdn": "string"
}
*/

func datasourceIllumioContainerCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext:   datasourceIllumioContainerClusterRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Container Cluster",

		Schema: map[string]*schema.Schema{
			"container_cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Numerical ID of Container Cluster",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the Cluster",
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
	}
}

func datasourceIllumioContainerClusterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID
	ccID := d.Get("container_cluster_id").(string)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/container_clusters/%v", orgID, ccID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))

	for _, key := range []string{
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
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("nodes") {
		nodeKeys := []string{"pod_subnet"}
		nodes := data.S("nodes")
		nodeI := []map[string]interface{}{}

		for _, node := range nodes.Children() {
			nodeI = append(nodeI, gabsToMap(node, nodeKeys))
		}

		d.Set("nodes", nodeI)
	} else {
		d.Set("nodes", nil)
	}

	if data.Exists("errors") {
		errorKeys := []string{
			"audit_event",
			"duplicate_ids",
			"error_type",
		}
		errors := data.S("errors")
		errorI := []map[string]interface{}{}

		for _, error := range errors.Children() {
			errorMap := gabsToMap(error, errorKeys)
			if error.Exists("audit_events") {
				errorMap["audit_events"] = gabsToMap(error.S("audit_events"), []string{"href"})
			} else {
				errorMap["audit_events"] = nil
			}
			errorI = append(errorI, errorMap)
		}
		d.Set("errors", errorI)
	} else {
		d.Set("errors", nil)
	}

	return diagnostics
}
