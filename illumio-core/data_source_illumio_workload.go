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
  "deleted": true,
  "delete_type": "string",
  "name": "string",
  "description": "string",
  "hostname": "string",
  "service_principal_name": "string",
  "agent_to_pce_certificate_authentication_id": null,
  "distinguished_name": "string",
  "public_ip": "string",
  "external_data_set": null,
  "external_data_reference": null,
  "interfaces": {
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
  },
  "service_provider": "string",
  "data_center": "string",
  "data_center_zone": "string",
  "os_id": "string",
  "os_detail": "string",
  "online": true,
  "firewall_coexistence": null,
  "containers_inherit_host_policy": true,
  "blocked_connection_action": "drop",
  "labels": [
    {
      "href": "string"
    }
  ],
  "services": {
    "uptime_seconds": 0,
    "created_at": "2021-03-02T02:37:59Z",
    "open_service_ports": [
      {
        "protocol": 0,
        "address": "string",
        "port": 0,
        "process_name": "string",
        "user": "string",
        "package": "string",
        "win_service_name": "string"
      }
    ]
  },
  "vulnerabilities_summary": {
    "num_vulnerabilities": 0,
    "vulnerable_port_exposure": null,
    "vulnerable_port_wide_exposure": {
      "any": null,
      "ip_list": null
    },
    "vulnerability_exposure_score": null,
    "vulnerability_score": 0,
    "max_vulnerability_score": 0
  },
  "detected_vulnerabilities": [
    {
      "ip_address": "string",
      "port": 0,
      "proto": 0,
      "port_exposure": null,
      "port_wide_exposure": {
        "any": null,
        "ip_list": null
      },
      "workload": {
        "href": "string"
      },
      "vulnerability": {
        "href": "string",
        "score": 0,
        "name": "string"
      },
      "vulnerability_report": {
        "href": "string"
      }
    }
  ],
  "agent": {
    "config": {
      "mode": "string",
      "log_traffic": true,
      "security_policy_update_mode": "string"
    },
    "href": "string",
    "secure_connect": {
      "matching_issuer_name": "string"
    },
    "status": {
      "uid": "string",
      "last_heartbeat_on": null,
      "uptime_seconds": null,
      "agent_version": "string",
      "managed_since": "2021-03-02T02:37:59Z",
      "fw_config_current": true,
      "firewall_rule_count": 0,
      "security_policy_refresh_at": "2021-03-02T02:37:59Z",
      "security_policy_applied_at": "2021-03-02T02:37:59Z",
      "security_policy_received_at": "2021-03-02T02:37:59Z",
      "agent_health_errors": {
        "errors": [
          "string"
        ],
        "warnings": [
          "string"
        ]
      },
      "agent_health": [
        {
          "type": "string",
          "severity": "string",
          "audit_event": "string"
        }
      ],
      "security_policy_sync_state": "string"
    },
    "active_pce_fqdn": "string",
    "target_pce_fqdn": "string",
    "type": "string"
  },
  "ven": {
    "href": "string",
    "hostname": "string",
    "name": "string",
    "status": "string"
  },
  "enforcement_mode": "idle",
  "selectively_enforced_services": [
    {
      "href": "string"
    }
  ],
  "created_at": "2021-03-02T02:37:59Z",
  "updated_at": "2021-03-02T02:37:59Z",
  "deleted_at": "2021-03-02T02:37:59Z",
  "created_by": {
    "href": "string"
  },
  "updated_by": {
    "href": "string"
  },
  "deleted_by": {
    "href": "string"
  },
  "container_cluster": {
    "href": "string",
    "name": "string"
  }
}
*/

func datasourceIllumioWorkload() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioWorkloadRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Workload",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isWorkloadHref,
				Description:      "URI of the Workload",
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
							Description: "CIDR BLOCK of the Interface. The number of bits in the subnet /24 is 255.255.255.0.",
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
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Description: "IKE authentication certificate for certificate-based Secure Connect and Machine Auth.",
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
										Type:     schema.TypeString,
										Computed: true,
										Description: "	The local address this service is bound to",
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
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
							Description: "The aggregated vulnerability exposure score of the workload across all the vulnerable ports.",
						},
						"vulnerability_score": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The aggregated vulnerability score of the workload across all the vulnerable ports.",
						},
						"max_vulnerability_score": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum of all the vulnerability scores associated with the detected_vulnerabilities on the workload.",
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
							Description: "The exposure of the port based on the current policy.",
						},
						"workload": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "URI of Workload.",
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
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"caps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "CAPS for Workload",
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this Workload",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this Workload",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this Workload",
			},
		},
	}
}

func dataSourceIllumioWorkloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	// orgID := pConfig.OrgID
	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
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
		"labels",
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
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	key := "interfaces"
	if data.Exists(key) {
		d.Set("interfaces", extractMapArray(data.S(key), []string{
			"name",
			"loopback",
			"link_state",
			"address",
			"cidr_block",
			"default_gateway_address",
			"network",
			"network_detection_mode",
			"friendly_name",
		}))
	} else {
		d.Set(key, nil)
	}

	key = "services"
	if data.Exists(key) {
		services := data.S(key)
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
		}
		d.Set(key, srs)
	} else {
		d.Set(key, nil)
	}

	key = "vulnerabilities_summary"
	if data.Exists(key) {
		vss := data.S(key)
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

			d.Set(key, srs)
		}
	} else {
		d.Set(key, nil)
	}

	key = "detected_vulnerabilities"
	if data.Exists(key) {
		vss := data.S(key)
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

			d.Set(key, srs)
		}
	} else {
		d.Set(key, nil)
	}

	if data.Exists("firewall_coexistence") {
		d.Set("firewall_coexistence", extractMapArray(data.S("firewall_coexistence"), []string{
			"illumio_primary",
		}))
	}

	if data.Exists("selectively_enforced_services") {
		d.Set("selectively_enforced_services", extractMapArray(data.S("selectively_enforced_services"), []string{
			"href",
			"proto",
			"port",
			"to_port",
		}))
	}

	if data.Exists("container_cluster") {
		d.Set("container_cluster", extractMapArray(data.S("container_cluster"), []string{
			"href",
			"name",
		}))
	}

	return diagnostics
}
