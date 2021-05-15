package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/* Sample
{
	"items" : [
		{
		"href": "string",
		"name": "string",
		"description": "string",
		"hostname": "string",
		"uid": "string",
		"os_id": "string",
		"os_detail": "string",
		"os_platform": "string",
		"version": "string",
		"status": "string",
		"activation_type": "string",
		"active_pce_fqdn": "string",
		"target_pce_fqdn": "string",
		"labels": [
			{
			"href": "string",
			"key": "string",
			"value": "string"
			}
		],
		"interfaces": [
			{
			"name": "string",
			"link_state": "string",
			"address": "string",
			"cidr_block": 0,
			"default_gateway_address": "string",
			"network": {
				"href": "string"
			},
			"network_detection_mode": "string",
			"friendly_name": "string"
			}
		],
		"workloads": [
			{
			"href": "string",
			"name": "string",
			"hostname": "string",
			"os_id": "string",
			"os_detail": "string",
			"labels": [
				{
				"href": "string",
				"key": "string",
				"value": "string"
				}
			],
			"public_ip": "string",
			"interfaces": [
				{
				"name": "string",
				"link_state": "string",
				"address": "string",
				"cidr_block": 0,
				"default_gateway_address": "string",
				"network": {
					"href": "string"
				},
				"network_detection_mode": "string",
				"friendly_name": "string"
				}
			],
			"security_policy_applied_at": "2021-03-02T02:37:59Z",
			"security_policy_received_at": "2021-03-02T02:37:59Z",
			"mode": "idle",
			"enforcement_mode": "idle",
			"visibility_level": "string",
			"online": true
			}
		],
		"container_cluster": {
			"href": "string",
			"name": "string"
		},
		"secure_connect": {
			"matching_issuer_name": "string"
		},
		"last_heartbeat_at": null,
		"last_goodbye_at": "2021-03-02T02:37:59Z",
		"created_at": "2021-03-02T02:37:59Z",
		"created_by": {
			"href": "string"
		},
		"updated_at": "2021-03-02T02:37:59Z",
		"updated_by": {
			"href": "string"
		},
		"conditions": [
			{
			"first_reported_timestamp": "2021-03-02T02:37:59Z",
			"latest_event": {
				"notification_type": "string",
				"severity": "error",
				"href": "string",
				"info": {},
				"timestamp": "2021-03-02T02:37:59Z"
			}
			}
		],
		"caps": [
			"string"
		]
		}
	]
}
*/

func datasourceIllumioVENs() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioVENsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio VENs",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of VENs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of VEN",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Friendly name for the VEN",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the VEN",
						},
						"hostname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hostname of the host managed by the VEN",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of the host managed by the VEN",
						},
						"os_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OS identifier of the host managed by the VEN",
						},
						"os_detail": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Additional OS details from the host managed by the VEN",
						},
						"os_platform": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OS platform of the host managed by the VEN",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Software version of the VEN",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the VEN",
						},
						"activation_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The method by which the VEN was activated",
						},
						"active_pce_fqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The FQDN of the PCE that the VEN last connected to",
						},
						"target_pce_fqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The FQDN of the PCE that the VEN will use for future connections",
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Labels assigned to the host managed by the VEN",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Label URI",
									},
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Key of the label",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of the label",
									},
								},
							},
						},
						"interfaces": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Network interfaces of the host managed by the VEN",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Interface URI",
									},
									"loopback": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Loopback for Workload Interface",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Interface name",
									},
									"link_state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link State",
									},
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP Address to assign to this interface",
									},
									"cidr_block": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of bits in the subnet",
									},
									"default_gateway_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP Address of the default gateway",
									},
									"network": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Network that the interface belongs to",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"network_detection_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network Detection Mode",
									},
									"friendly_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "User-friendly name for interface",
									},
								},
							},
						},
						"workloads": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "collection of Workloads",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of the Workload",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The short friendly name of the workload",
									},
									"hostname": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The hostname of this workload",
									},
									"os_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "OS identifier for the workload",
									},
									"os_detail": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Additional OS details",
									},
									"labels": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Labels assigned to the host managed by the VEN",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"href": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Label URI",
												},
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Key of the label",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Value of the label",
												},
											},
										},
									},
									"public_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The public IP address of the server",
									},
									"interfaces": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Network interfaces of the workload",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Name of Interface",
												},
												"loopback": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Loopback for Workload Interface",
												},
												"link_state": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Link State for the workload Interface.",
												},
												"address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The IP Address to assign to this interface.",
												},
												"cidr_block": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "CIDR BLOCK of the Interface.",
												},
												"default_gateway_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Default Gateaway Address of the Interface",
												},
												"network": {
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "Network of the Interface",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"network_detection_mode": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Network Detection Mode of the Interface",
												},
												"friendly_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "User-friendly name for interface",
												},
											},
										},
									},
									"security_policy_applied_at": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Last reported time when policy was applied to the workload (UTC)",
									},
									"security_policy_received_at": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Last reported time when policy was received by the workload (UTC)",
									},
									"mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Policy enforcement mode",
									},
									"enforcement_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Policy enforcement mode",
									},
									"visibility_level": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Visibility level of the workload",
									},
									"online": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "If this workload is online and present in policy",
									},
								},
							},
						},
						"container_cluster": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "container_cluster details for ven. Single element list",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URI of the container cluster managed by this VEN",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the container cluster managed by this VEN, only present in expanded representations",
									},
								},
							},
						},
						"secure_connect": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "secure_connect details for vens",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"matching_issuer_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Issuer name match criteria for certificate used during establishing secure connections",
									},
								},
							},
						},
						"unpair_allowed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "",
						},
						"last_heartbeat_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last time (rfc3339 timestamp) a heartbeat was received from this VEN",
						},
						"last_goodbye_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time (rfc3339 timestamp) of the last goodbye from the VEN",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time (rfc3339 timestamp) at which this VEN was created",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The href of the user who created this VEN",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time (rfc3339 timestamp) at which this VEN was last updated",
						},
						"updated_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The href of the user who last updated this VEN",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"conditions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"first_reported_timestamp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timestamp of the first event that reported this condition",
									},
									"latest_event": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The latest notification event that was generated for the corresponding condition. Single element list",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"notification_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The notification_type of the event",
												},
												"severity": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Severity of the condition, same as the event",
												},
												"href": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The href of the event",
												},
												"info": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The information from the notification event that was generated by the condition",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"agent": {
																Type:        schema.TypeMap,
																Computed:    true,
																Description: "",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"timestamp": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "RFC 3339 timestamp at which this event was created",
												},
											},
										},
									},
								},
							},
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
			"activation_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"pairing_key", "kerberos", "certificate"}, false),
				),
				Description: `The method in which the VEN was activated. Allowed values are "pairing_key", "kerberos" and "certificate"`,
			},
			"active_pce_fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "FQDN of the PCE",
			},
			"condition": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"agent.upgrade_time_out", "agent.missing_heartbeats_after_upgrade", "agent.clone_detected", "agent.missed_heartbeats"}, false),
				),
				Description: `A specific error condition to filter by. Allowed values are "agent.upgrade_time_out", "agent.missing_heartbeats_after_upgrade", "agent.clone_detected" and "agent.missed_heartbeats"`,
			},
			"container_clusters": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Array of container cluster URIs, encoded as a JSON string",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of VEN(s) to return. Supports partial matches.",
			},
			"disconnected_before": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Return VENs that have been disconnected since the given time",
			},
			"health": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"healthy", "unhealthy", "error", "warning"}, false),
				),
				Description: `The overall health (condition) of the VEN. Allowed values are  "healthy", "unhealthy", "error" and "warning"`,
			},
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hostname of VEN(s) to return. Supports partial matches",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address of VEN(s) to return. Supports partial matches",
			},
			"labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "2D Array of label URIs, encoded as a JSON string",
			},
			"last_goodbye_at_gte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Greater than or equal to value for last goodbye at timestamp",
			},
			"last_goodbye_at_lte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Greater than or equal to value for last goodbye at timestamp",
			},
			"last_heartbeat_at_gte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Greater than or equal to value for last heartbeat timestamp",
			},
			"last_heartbeat_at_lte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Less than or equal to value for last heartbeat timestamp",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of VENs to return. The integer should be a non-zero positive integer.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of VEN(s) to return. Supports partial matches",
			},
			"os": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Operating System of VEN(s) to return. Supports partial matches.",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"active", "suspended", "stopped", "uninstalled"}, false),
				),
				Description: `The current status of the VEN. Allowed values are "active", "suspended", "stopped" and "uninstalled"`,
			},
			"upgrade_pending": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Only return VENs with/without a pending upgrade",
			},
			"version_gte": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Greater than or equal to value for version",
			},
			"version_lte": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Less than or equal to value for version",
			},
		},
	}
}

func dataSourceIllumioVENsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	paramKeys := []string{
		"activation_type",
		"active_pce_fqdn",
		"condition",
		"container_clusters",
		"description",
		"disconnected_before",
		"health",
		"hostname",
		"ip_address",
		"labels",
		"max_results",
		"name",
		"os",
		"status",
		"upgrade_pending",
	}

	params := resourceDataToMap(d, paramKeys)

	if value, ok := d.GetOk("version_gte"); ok {
		params["version[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("version_lte"); ok {
		params["version[lte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_goodbye_at_gte"); ok {
		params["last_goodbye_at[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_goodbye_at_lte"); ok {
		params["last_goodbye_at[lte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_heartbeat_at_gte"); ok {
		params["last_heartbeat_at[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_heartbeat_at_lte"); ok {
		params["last_heartbeat_at[lte]"] = value.(string)
	}

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%v/vens", pConfig.OrgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}
	keys := []string{
		"href",
		"name",
		"description",
		"hostname",
		"uid",
		"os_id",
		"os_detail",
		"os_platform",
		"version",
		"status",
		"activation_type",
		"active_pce_fqdn",
		"target_pce_fqdn",
		"labels",
		"unpair_allowed",
		"last_heartbeat_at",
		"last_goodbye_at",
		"created_at",
		"created_by",
		"updated_at",
		"updated_by",
		"caps",
	}

	for _, child := range data.Children() {
		m := extractMap(child, keys)

		intfKeys := []string{
			"href",
			"name",
			"link_state",
			"address",
			"loopback",
			"cidr_block",
			"default_gateway_address",
			"network",
			"network_detection_mode",
			"friendly_name",
		}

		key := "interfaces"
		if child.Exists(key) && child.S(key).Data() != nil {
			m["interfaces"] = extractMapArray(child.S(key), intfKeys)
		} else {
			m["interfaces"] = nil
		}

		key = "workloads"
		if child.Exists(key) {
			workloadKeys := []string{
				"href",
				"name",
				"hostname",
				"os_id",
				"os_detail",
				"labels",
				"public_ip",
				"security_policy_applied_at",
				"security_policy_received_at",
				"mode",
				"enforcement_mode",
				"visibility_level",
				"online",
			}

			workloads := child.S(key)

			wrs := []map[string]interface{}{}

			for _, workload := range workloads.Children() {
				wr := extractMap(workload, workloadKeys)
				if workload.Exists("interfaces") {
					wr["interfaces"] = extractMapArray(workload.S("interface"), intfKeys)
				} else {
					wr["interfaces"] = nil
				}

				wrs = append(wrs, wr)
			}

			m[key] = wrs
		} else {
			m[key] = nil
		}

		key = "container_cluster"
		if child.Exists(key) && child.S(key).Data() != nil {
			ccKeys := []string{
				"href",
				"name",
			}

			m[key] = []interface{}{extractMap(child.S(key), ccKeys)}
		} else {
			m[key] = nil
		}

		key = "secure_connect"
		if child.Exists(key) && child.S(key).Data() != nil {
			ccKeys := []string{
				"matching_issuer_name",
			}

			m[key] = []interface{}{extractMap(child.S(key), ccKeys)}

		} else {
			m[key] = nil
		}

		key = "conditions"
		if child.Exists(key) && child.S(key).Data() != nil {
			cnds := []map[string]interface{}{}
			for _, condition := range child.S(key).Children() {
				cnd := map[string]interface{}{}

				for k, v := range condition.ChildrenMap() {
					switch k {
					case "first_reported_timestamp":
						cnd[k] = v.Data()

					case "latest_event":
						eventKeys := []string{
							"notification_type",
							"severity",
							"href",
							"timestamp",
						}

						if v.Data() != nil {
							t := extractMap(v, eventKeys)

							if v.Exists("info", "agent") {
								info := make(map[string]interface{})
								info["agent"] = extractMap(v.S("info", "agent"), []string{"name", "hostname", "href"})

								t["info"] = []interface{}{info}
							}

							cnd[k] = []interface{}{t}
						} else {
							cnd[k] = nil
						}
					}
				}
				cnds = append(cnds, cnd)
			}

			m[key] = cnds
		} else {
			m[key] = nil
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
