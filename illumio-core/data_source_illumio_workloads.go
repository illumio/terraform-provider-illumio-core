// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func datasourceIllumioWorkloads() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioWorkloadsRead,
		SchemaVersion: 1,
		Description:   "Represents Illumio Workloads",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of workloads",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of workload",
						},
						"deleted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This indicates that the workload has been deleted",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Workload",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the Workload",
						},
						"hostname": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hostname of this workload",
						},
						"service_principal_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Kerberos Service Principal Name (SPN)",
						},
						"external_data_set": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The data source from which a resource originates",
						},
						"agent_to_pce_certificate_authentication_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PKI Certificate identifier to be used by the PCE for authenticating the VEN",
						},
						"distinguished_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "X.509 Subject distinguished name",
						},
						"enforcement_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Enforcement mode of workload(s) to return",
						},
						"visibility_level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Visibility Level of workload(s) to return",
						},
						"public_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public IP address of the server",
						},
						"external_data_reference": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique identifier within the external data source",
						},
						"interfaces": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A unique identifier within the external data source",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the Interface",
									},
									"loopback": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Loopback for Workload Interface",
									},
									"link_state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Link State of the Interface",
									},
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Address of the Interface",
									},
									"cidr_block": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "CIDR BLOCK of the Interface. The number of bits in the subnet /24 is 255.255.255.0",
									},
									"default_gateway_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Default Gateway Address of the Interface",
									},
									"network": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Default Gateway Address of the Interface",
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
										Description: "Friendly name of the Interface",
									},
								},
							},
						},
						"ignored_interface_names": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Ignored Interface Names for Workload",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"service_provider": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service provider for Workload",
						},
						"data_center": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data center for Workload",
						},
						"data_center_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data center Zone for Workload",
						},
						"os_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OS identifier for Workload",
						},
						"os_detail": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Additional OS details - just displayed to end-user",
						},
						"online": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines if this workload is online",
						},
						"ike_authentication_certificate": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "IKE authentication certificate for certificate-based Secure Connect and Machine Auth",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of label URIs",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of Label",
									},
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Workload Label key",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Workload Label value",
									},
								},
							},
						},
						"services": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Service report for Workload",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uptime_seconds": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "How long since the last reboot of this box - used as a timestamp for this",
									},
									"created_at": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Timestamp when this service was first created",
									},
									"open_service_ports": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A list of open ports",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocol": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Transport protocol for open service ports",
												},
												"address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "	The local address this service is bound to",
												},
												"port": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "	The local port this service is bound to",
												},
												"process_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The process name (including the full path)",
												},
												"user": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The user account that the process is running under",
												},
												"package": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The RPM/DEB package that the program is part of",
												},
												"win_service_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Name of the Windows service",
												},
											},
										},
									},
								},
							},
						},
						"vulnerabilities_summary": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Vulnerabilities summary associated with the workload",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"num_vulnerabilities": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of vulnerabilities associated with the workload",
									},
									"vulnerable_port_exposure": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The aggregated vulnerability port exposure score of the workload across all the vulnerable ports",
									},
									"vulnerable_port_wide_exposure": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "High end of an IP range",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"any": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "The boolean value representing if at least one port is exposed to internet (any rule) on the workload",
												},
												"ip_list": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "The boolean value representing if at least one port is exposed to ip_list(s) on the workload",
												},
											},
										},
									},
									"vulnerability_exposure_score": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The aggregated vulnerability exposure score of the workload across all the vulnerable ports",
									},
									"vulnerability_score": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The aggregated vulnerability score of the workload across all the vulnerable ports",
									},
									"max_vulnerability_score": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum of all the vulnerability scores associated with the detected_vulnerabilities on the workload",
									},
								},
							},
						},
						"detected_vulnerabilities": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address of the host where the vulnerability is found",
									},
									"port_wide_exposure": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "High end of an IP range",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"any": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "The boolean value representing if at least one port is exposed to internet (any rule) on the workload",
												},
												"ip_list": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "The boolean value representing if at least one port is exposed to ip_list(s) on the workload",
												},
											},
										},
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port that is associated with the vulnerability",
									},
									"proto": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The protocol that is associated with the vulnerability",
									},
									"port_exposure": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The exposure of the port based on the current policy",
									},
									"workload": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "URI of Workload",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"href": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URI of the workload to which this vulnerability belongs to",
												},
											},
										},
									},
									"vulnerability": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Vulnerability",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"href": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URI of the vulnerability class to which this vulnerability belongs to",
												},
												"score": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The normalized score of the vulnerability within the range of 0 to 100",
												},
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The title/name of the vulnerability",
												},
											},
										},
									},
									"vulnerability_report": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Vulnerability Report for Workload",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"href": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The URI of the report to which this vulnerability belongs to",
												},
											},
										},
									},
								},
							},
							Description: "Detected Vulnerabilities",
						},
						"containers_inherit_host_policy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This workload will apply the policy it receives both to itself and the containers hosted by it",
						},
						"firewall_coexistence": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Firewall coexistence mode",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"illumio_primary": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Illumio is the primary firewall if set to true",
									},
								},
							},
						},
						"selectively_enforced_services": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Selectively Enforced Services for Workload",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of service",
									},
									"proto": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Protocol number",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Port number used with protocol. Also, the starting port when specifying a range",
									},
									"to_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Upper end of port range",
									},
								},
							},
						},
						"container_cluster": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Container Cluster for Workload",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of container cluster",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of container cluster",
									},
								},
							},
						},
						"blocked_connection_action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Blocked Connection Action for Workload",
						},
						"ven": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "VENS for Workload",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"caps": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "User permissions for the object",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Workload was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Workload was last updated",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Workload was deleted",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who created this Workload",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who last updated this Workload",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"deleted_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who deleted this Workload",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"agent_active_pce_fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "FQDN of the PCE",
			},
			"container_clusters": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of container cluster URIs, encoded as a JSON string",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of workload(s) to return. Supports partial matches",
			},
			"enforcement_mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(ValidWorkloadEnforcementModeValues, false),
				),
				Description: `Enforcement mode of workload(s) to return. Allowed values are "idle", "visibility_only", "full" and "selective"`,
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
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hostname of workload(s) to return. Supports partial matches",
			},
			"include_deleted": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Include deleted workloads",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address of workload(s) to return. Supports partial matches",
			},
			"labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of lists of label URIs, encoded as a JSON string",
			},
			"last_heartbeat_on_gte": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Greater than or equal to value for last heartbeat on timestamp",
			},
			"last_heartbeat_on_lte": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Less than or equal to value for last heartbeat on timestamp",
			},
			"log_traffic": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Whether we want to log traffic events from this workload",
			},
			"managed": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Return managed or unmanaged workloads using this filter",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of workloads to return. The integer should be a non-zero positive integer",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of workload(s) to return. Supports partial matches",
			},
			"online": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Return online/offline workloads using this filter",
			},
			"os_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Operating System of workload(s) to return. Supports partial matches",
			},
			"policy_health": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"active", "warning", "error", "suspended"}, false),
				),
				Description: `Policy of health of workload(s) to return. Allowed values are "active", "warning", "error" and "suspended"`,
			},
			"security_policy_sync_state": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"staged"}, false),
				),
				Description: `Advanced search option for workload based on policy sync state. Allowed value: "staged"`,
			},
			"security_policy_update_mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"static", "adaptive"}, false),
				),
				Description: `Advanced search option for workload based on security policy update mode. Allowed values are "static" and "adaptive"`,
			},
			"ven": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URI of VEN to filter by",
			},
			"visibility_level": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(validVisibilityLevels, false),
				),
				Description: `Filter by visibility level. Allowed values are "flow_full_detail", "flow_summary", "flow_drops", "flow_off" and "enhanced_data_collection"`,
			},
			"vulnerability_summary_vulnerability_exposure_score_gte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Greater than or equal to value for vulnerability_exposure_score",
			},
			"vulnerability_summary_vulnerability_exposure_score_lte": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Less than or equal to value for vulnerability_exposure_score",
			},
			"match_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      PARTIAL_MATCH,
				ValidateFunc: validation.StringInSlice([]string{PARTIAL_MATCH, EXACT_MATCH}, true),
				Description:  `Indicates whether to return all partially-matching names or only exact matches. Allowed values are "partial" and "exact". Default value: "partial"`,
			},
		},
	}
}

func dataSourceIllumioWorkloadsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	paramKeys := []string{
		"container_clusters",
		"description",
		"enforcement_mode",
		"external_data_reference",
		"external_data_set",
		"hostname",
		"include_deleted",
		"ip_address",
		"labels",
		"log_traffic",
		"managed",
		"max_results",
		"name",
		"online",
		"os_id",
		"policy_health",
		"security_policy_sync_state",
		"security_policy_update_mode",
		"ven",
		"visibility_level",
	}

	params := resourceDataToMap(d, paramKeys)

	if value, ok := d.GetOk("agent_active_pce_fqdn"); ok {
		params["agent.active_pce_fqdn"] = value.(string)
	}
	if value, ok := d.GetOk("last_heartbeat_at_gte"); ok {
		params["last_heartbeat_at[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("last_heartbeat_at_lte"); ok {
		params["last_heartbeat_at[lte]"] = value.(string)
	}
	if value, ok := d.GetOk("vulnerability_summary_vulnerability_exposure_score_gte"); ok {
		params["vulnerability_summary.vulnerability_exposure_score[gte]"] = value.(string)
	}
	if value, ok := d.GetOk("vulnerability_summary_vulnerability_exposure_score_lte"); ok {
		params["vulnerability_summary.vulnerability_exposure_score[lte]"] = value.(string)
	}

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/workloads", illumioClient.OrgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}

	for _, child := range data.Children() {
		// if exact matching is enabled, skip the object if it's a partial match
		if d.Get("match_type").(string) == EXACT_MATCH {
			if !isExactMatch("name", d, child) {
				continue
			}
		}

		m := map[string]interface{}{}

		for _, key := range []string{
			"href",
			"deleted",
			"name",
			"description",
			"hostname",
			"service_principal_name",
			"agent_to_pce_certificate_authentication_id",
			"distinguished_name",
			"enforcement_mode",
			"visibility_level",
			"public_ip",
			"ignored_interface_names",
			"service_provider",
			"data_center",
			"data_center_zone",
			"os_id", "os_detail",
			"online",
			"containers_inherit_host_policy",
			"blocked_connection_action",
			"ven",
			"caps",
			"external_data_set",
			"external_data_reference",
			"created_at",
			"updated_at",
			"created_by",
			"updated_by",
			"deleted_by",
			"deleted_at",
			"ike_authentication_certificate",
		} {
			if child.Exists(key) {
				m[key] = child.S(key).Data()
			} else {
				m[key] = nil
			}
		}

		key := "interfaces"
		if child.Exists(key) {
			m[key] = extractMapArray(child.S(key), []string{
				"name",
				"loopback",
				"link_state",
				"address",
				"cidr_block",
				"default_gateway_address",
				"network",
				"network_detection_mode",
				"friendly_name",
			})
		} else {
			m[key] = nil
		}

		key = "labels"
		if child.Exists(key) {
			m[key] = extractMapArray(child.S(key), []string{
				"href",
				"key",
				"value",
			})
		} else {
			m[key] = nil
		}

		key = "services"
		if child.Exists(key) {
			services := child.S(key)
			srs := []map[string]interface{}{}

			for _, service := range services.Children() {
				sr := extractMap(service, []string{
					"uptime_seconds",
					"created_at",
					"open_service_ports",
				})
				if service.Exists("open_service_ports") {
					sr["open_service_ports"] = extractMapArray(service.S("open_service_ports"), []string{
						"protocol",
						"address",
						"port",
						"process_name",
						"user",
						"package",
						"win_service_name",
					})
				} else {
					sr["open_service_ports"] = nil
				}

				srs = append(srs, sr)

				m[key] = srs
			}
		} else {
			m[key] = nil
		}

		key = "vulnerabilities_summary"
		if child.Exists(key) {
			vss := child.S(key)
			srs := []map[string]interface{}{}

			for _, vs := range vss.Children() {
				sr := extractMap(vs, []string{
					"num_vulnerabilities",
					"vulnerable_port_exposure",
					"vulnerable_port_wide_exposure",
					"vulnerability_exposure_score",
					"vulnerability_score",
					"max_vulnerability_score",
				})
				if vs.Exists("vulnerable_port_wide_exposure") {
					sr["vulnerable_port_wide_exposure"] = extractMapArray(vs.S("vulnerable_port_wide_exposure"), []string{
						"any",
						"ip_list",
					})
				} else {
					sr["vulnerable_port_wide_exposure"] = nil
				}

				srs = append(srs, sr)

				m[key] = srs
			}
		} else {
			m[key] = nil
		}

		key = "detected_vulnerabilities"
		if child.Exists(key) {
			vss := child.S(key)
			srs := []map[string]interface{}{}

			for _, vs := range vss.Children() {
				sr := extractMap(vs, []string{
					"ip_address",
					"port",
					"proto",
					"port_exposure",
					"workload",
					"vulnerability",
					"port_wide_exposure",
					"vulnerability",
					"vulnerability_report",
				})
				if vs.Exists("port_wide_exposure") {
					sr["port_wide_exposure"] = extractMapArray(vs.S("port_wide_exposure"), []string{
						"any",
						"ip_list",
					})
				} else {
					sr["port_wide_exposure"] = nil
				}

				if vs.Exists("vulnerability") {
					sr["vulnerability"] = extractMapArray(vs.S("vulnerability"), []string{
						"href",
						"score",
						"name",
					})
				} else {
					sr["vulnerability"] = nil
				}

				srs = append(srs, sr)

				m[key] = srs
			}
		} else {
			m[key] = nil
		}

		if data.Exists("firewall_coexistence") {
			m["firewall_coexistence"] = extractMapArray(data.S("firewall_coexistence"), []string{
				"illumio_primary",
			})
		} else {
			m["firewall_coexistence"] = nil
		}

		if data.Exists("selectively_enforced_services") {
			m["selectively_enforced_services"] = extractMapArray(data.S("selectively_enforced_services"), []string{
				"href",
				"proto",
				"port",
				"to_port",
			})
		} else {
			m["selectively_enforced_services"] = nil
		}

		if data.Exists("container_cluster") {
			m["container_cluster"] = extractMapArray(data.S("container_cluster"), []string{
				"href",
				"name",
			})
		} else {
			m["container_cluster"] = nil
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
