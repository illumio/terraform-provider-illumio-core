// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample
{
	"items": [
		{
			"href": "string",
			"pce_scope": [
			"string"
			],
			"type": "string",
			"description": "string",
			"audit_event_logger": {
			"configuration_event_included",: true,
			"system_event_included",: true,
			"min_severity",: "error"
			},
			"traffic_event_logger": {
			"traffic_flow_allowed_event_included": true,
			"traffic_flow_potentially_blocked_event_included": true,
			"traffic_flow_blocked_event_included": true
			},
			"node_status_logger": {
			"node_status_included": true
			},
			"remote_syslog": {
			"address": "string",
			"port": 0,
			"protocol": 0,
			"tls_enabled": true,
			"tls_ca_bundle": "string",
			"tls_verify_cert": true
			}
		}
	]
}
*/

func datasourceIllumioServiceBindings() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioServiceBindingsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Service Bindings",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of Service Bindings",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Service Binding",
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
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Virtual service href",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"workload": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Workload Object for Service Bindings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
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
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Container Workload href",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"port_overrides": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Port Overrides for Service Bindings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Port Number in the original service which to override (integer 0-65535). Starting port when specifying a range.",
									},
									"proto": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Transport protocol in the original service which to override",
									},
									"new_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Overriding port number (or starting point when specifying a range)",
									},
									"new_to_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Overriding port range ending port",
									},
								},
							},
						},
						"external_data_set": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External Data Set identifier",
						},
						"external_data_reference": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External Data reference identifier",
						},
					},
				},
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A unique identifier within the external data source",
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The data source from which a resource originates",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of virtual service bindings to return. The integer should be a non-zero positive integer.",
			},
			"virtual_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Virtual service URI",
			},
			"workload": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Workload URI",
			},
		},
	}
}

func dataSourceIllumioServiceBindingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	paramKeys := []string{
		"external_data_reference",
		"external_data_set",
		"virtual_service",
		"workload",
		"max_results",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/service_bindings", orgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}

	for _, child := range data.Children() {
		m := map[string]interface{}{}

		for _, key := range []string{
			"href",
			"bound_service",
			"virtual_service",
			"external_data_set",
			"external_data_reference",
		} {
			if child.Exists(key) {
				m[key] = child.S(key).Data()
			} else {
				m[key] = nil
			}
		}

		if child.Exists("container_workload") {
			m["container_workload"] = map[string]string{"href": child.S("container_workload").S("href").Data().(string)}
		} else {
			m["container_workload"] = nil
		}

		if child.Exists("port_overrides") {
			poS := child.S("port_overrides")
			poI := []map[string]interface{}{}

			for _, po := range poS.Children() {
				poI = append(poI, extractMap(po, []string{"port", "proto", "new_port", "new_to_port"}))
			}

			m["port_overrides"] = poI
		} else {
			m["port_overrides"] = nil
		}

		if child.Exists("workload") {
			m["workload"] = []interface{}{extractMap(child.S("workload"), []string{"href", "name", "hostname", "deleted"})}
		} else {
			m["workload"] = nil
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
