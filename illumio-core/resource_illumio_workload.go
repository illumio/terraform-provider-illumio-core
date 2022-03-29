// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

var (
	ValidWorkloadInterfaceLinkStateValues = []string{"up", "down", "unknown"}
	ValidWorkloadEnforcementModeValues    = []string{"idle", "visibility_only", "full", "selective"}
)

func resourceIllumioWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioWorkloadCreate,
		ReadContext:   resourceIllumioWorkloadRead,
		UpdateContext: resourceIllumioWorkloadUpdate,
		DeleteContext: resourceIllumioWorkloadDelete,
		Description:   "Manages Illumio Workload",

		SchemaVersion: version,

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the Workload",
			},
			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				AtLeastOneOf:     []string{"name", "hostname"},
				Description:      "Name of the Workload. The name should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"hostname": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The hostname of this workload. The hostname should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The long description of the workload",
			},
			"service_principal_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "The Kerberos Service Principal Name (SPN). The SPN should be between 1 to 255 characters",
				ValidateDiagFunc: nameValidation,
			},
			"public_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The public IP address of the server. The public IP should in the IPv4 or IPv6 format",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.IsIPAddress,
				),
			},
			/* Following code is commented to prevent the race condition
			 * between Workload and Workload Interface Resources. Preserved for future use.
			 * Bug#15
			 */
			// "interfaces": {
			// 	Type:        schema.TypeSet,
			// 	Computed:    true,
			// 	Description: "Workload network interfaces",
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"name": {
			// 				Type:        schema.TypeString,
			// 				Computed:    true,
			// 				Description: "Name of Interface. The name should be up to 255 characters",
			// 				// ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			// 			},
			// 			"link_state": {
			// 				Type:        schema.TypeString,
			// 				Computed:    true,
			// 				Description: "Link State for the workload Interface. Allowed values are \"up\", \"down\", and \"unknown\" ",
			// 				// ValidateDiagFunc: validation.ToDiagFunc(
			// 				// 	validation.StringInSlice(ValidWorkloadInterfaceLinkStateValues, false),
			// 				// ),
			// 			},
			// 			"address": {
			// 				Type:        schema.TypeString,
			// 				Computed:    true,
			// 				Description: "The Address to assign to this interface. The address should in the IPv4 or IPv6 format",
			// 				// ValidateDiagFunc: validation.ToDiagFunc(
			// 				// 	validation.IsIPAddress,
			// 				// ),
			// 			},
			// 			"cidr_block": {
			// 				Type:        schema.TypeInt,
			// 				Computed:    true,
			// 				Description: "CIDR BLOCK of the Interface",
			// 			},
			// 			"default_gateway_address": {
			// 				Type:        schema.TypeString,
			// 				Computed:    true,
			// 				Description: "Default Gateway Address of the Interface. The Default Gateway Address should in the IPv4 or IPv6 format",
			// 				// ValidateDiagFunc: validation.ToDiagFunc(
			// 				// 	validation.IsIPAddress,
			// 				// ),
			// 			},
			// 			"network": {
			// 				Type:        schema.TypeMap,
			// 				Computed:    true,
			// 				Description: "Href of Network of the Interface",
			// 				Elem: &schema.Schema{
			// 					Type: schema.TypeString,
			// 				},
			// 			},
			// 			"network_detection_mode": {
			// 				Type:        schema.TypeString,
			// 				Computed:    true,
			// 				Description: "Network Detection Mode of the Interface",
			// 			},
			// 			"friendly_name": {
			// 				Type:        schema.TypeString,
			// 				Computed:    true,
			// 				Description: "User-friendly name for interface. The name should be up to 255 characters",
			// 				// ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			// 			},
			// 			"loopback": {
			// 				Type:        schema.TypeBool,
			// 				Computed:    true,
			// 				Description: "Loopback for Workload Interface",
			// 			},
			// 		},
			// 	},
			// },
			"service_provider": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Service provider for Workload. The service_provider should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"data_center": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Data center for Workload. The data_center should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"data_center_zone": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Data center Zone for Workload. The data_center_zone should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"os_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "OS identifier for Workload. The os_id should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"os_detail": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Additional OS details - just displayed to end-user. The os_details should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"online": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if this workload is online. Default value: false",
				Default:     false,
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Assigned labels for workload",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "URI of label",
						},
					},
				},
			},
			"enforcement_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "visibility_only",
				Description: "Enforcement mode of workload(s) to return. Allowed values for enforcement modes are \"idle\",\"visibility_only\",\"full\", and \"selective\". Default value: \"visibility_only\" ",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(ValidWorkloadEnforcementModeValues, false),
				),
			},
			"deleted": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "This indicates that the workload has been deleted",
			},
			"agent_to_pce_certificate_authentication_id": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "PKI Certificate identifier to be used by the PCE for authenticating the VEN. The ID should be between 1 to 255 characters",
				ValidateDiagFunc: nameValidation,
			},
			"distinguished_name": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "X.509 Subject distinguished name. The name should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"visibility_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Visibility Level of workload(s) to return",
			},
			"ignored_interface_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Ignored Interface Names for Workload",
				Elem:        &schema.Schema{Type: schema.TypeString},
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
										Description: "The local address this service is bound to",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The local port this service is bound to",
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
										Required:    true,
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
							Description: "HThe aggregated vulnerability score of the workload across all the vulnerable ports",
						},
						"max_vulnerability_score": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum of all the vulnerability scores associated with the detected_vulnerabilities on the workload",
						},
					},
				},
			},
			"ike_authentication_certificate": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "IKE authentication certificate for certificate-based Secure Connect and Machine Auth",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The port that is associated with the vulnerability",
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
										Required:    true,
										Description: "The URI of the workload to which this vulnerability belongs to",
									},
								},
							},
						},
						"vulnerability": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Vulnerability for Workload",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The URI of the vulnerability class to which this vulnerability belongs to",
									},
									"score": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The normalized score of the vulnerability within the range of 0 to 100",
									},
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The title/name of the vulnerability",
									},
								},
							},
						},
						"vulnerability_report": {
							Type:        schema.TypeSet,
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
				Description: "Detected Vulnerabilities for Workload",
			},
			"containers_inherit_host_policy": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "This workload will apply the policy it receives both to itself and the containers hosted by it",
			},
			"firewall_coexistence": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Firewall coexistence mode for Workload",
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
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "The data source from which a resource originates",
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "A unique identifier within the external data source",
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
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this label group was last deleted",
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
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this label group",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioWorkloadCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	workload := &models.Workload{
		Name:                                  d.Get("name").(string),
		Description:                           d.Get("description").(string),
		ExternalDataSet:                       d.Get("external_data_set").(string),
		ExternalDataReference:                 d.Get("external_data_reference").(string),
		Hostname:                              d.Get("hostname").(string),
		ServicePrincipalName:                  d.Get("service_principal_name").(string),
		PublicIP:                              d.Get("public_ip").(string),
		ServiceProvider:                       d.Get("service_provider").(string),
		DataCenter:                            d.Get("data_center").(string),
		DataCenterZone:                        d.Get("data_center_zone").(string),
		OsID:                                  d.Get("os_id").(string),
		OsDetail:                              d.Get("os_detail").(string),
		Online:                                d.Get("online").(bool),
		EnforcementMode:                       d.Get("enforcement_mode").(string),
		AgentToPceCertificateAuthenticationID: d.Get("agent_to_pce_certificate_authentication_id").(string),
		DistinguishedName:                     d.Get("distinguished_name").(string),
	}
	if items, ok := d.GetOk("labels"); ok {
		workload.Labels = models.GetHrefs(items.(*schema.Set).List())
	}
	/* Following code is commented to prevent the race condition
	 * between Workload and Workload Interface Resources. Preserved for future use.
	 * Bug#15
	 */
	// if items, ok := d.GetOk("interfaces"); ok {
	// 	workload.Interfaces = expandIllumioWorkloadInterface(items.(*schema.Set).List())
	// }
	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/workloads", orgID), workload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.S("href").Data().(string))
	return resourceIllumioWorkloadRead(ctx, d, m)
}

func resourceIllumioWorkloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

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
		"os_id",
		"os_detail",
		"online",
		"labels",
		"containers_inherit_host_policy",
		"blocked_connection_action",
		"ven",
		"caps",
		"ike_authentication_certificate",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"deleted_at",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	/* Following code is commented to prevent the race condition
	 * between Workload and Workload Interface Resources. Preserved for future use.
	 * Bug#15
	 */
	// key := "interfaces"
	// if data.Exists(key) {
	// 	d.Set("interfaces", extractMapArray(data.S(key), []string{
	// 		"name",
	// 		"loopback",
	// 		"link_state",
	// 		"address",
	// 		"cidr_block",
	// 		"default_gateway_address",
	// 		"network",
	// 		"network_detection_mode",
	// 		"friendly_name",
	// 	}))
	// } else {
	// 	d.Set(key, nil)
	// }

	key := "services"
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
		}
		d.Set(key, srs)
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
		}
		d.Set(key, srs)
	} else {
		d.Set(key, nil)
	}

	if data.Exists("firewall_coexistence") {
		d.Set("firewall_coexistence", extractMapArray(data.S("firewall_coexistence"), []string{
			"illumio_primary",
		}))
	} else {
		d.Set("firewall_coexistence", nil)
	}

	if data.Exists("selectively_enforced_services") {
		d.Set("selectively_enforced_services", extractMapArray(data.S("selectively_enforced_services"), []string{
			"href",
			"proto",
			"port",
			"to_port",
		}))
	} else {
		d.Set("selectively_enforced_services", nil)
	}

	if data.Exists("container_cluster") {
		d.Set("container_cluster", extractMapArray(data.S("container_cluster"), []string{
			"href",
			"name",
		}))
	} else {
		d.Set("container_cluster", nil)
	}

	return diagnostics
}

func resourceIllumioWorkloadUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	var diags diag.Diagnostics

	workload := &models.Workload{}
	populateFromResourceData(workload, d)

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), workload)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioWorkloadRead(ctx, d, m)
}

func resourceIllumioWorkloadDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	_, err := illumioClient.Delete(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diagnostics
}

func populateFromResourceData(w *models.Workload, d *schema.ResourceData) {
	val := reflect.TypeOf(*w)

	// iterate over the Workload struct fields and
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := field.Type.Kind()
		jsonTag := field.Tag.Get("json")
		parts := strings.Split(jsonTag, ",")
		jsonFieldName := parts[0]

		if d.HasChange(jsonFieldName) {
			if fieldType == reflect.Slice || fieldType == reflect.Array {
				continue
			}
			resourceValue := reflect.ValueOf(d.Get(parts[0]))
			reflect.ValueOf(w).Elem().FieldByName(field.Name).Set(resourceValue)
		}
	}

	if d.HasChange("labels") {
		w.Labels = models.GetHrefs(d.Get("labels").(*schema.Set).List())
	}

	/* Following code is commented to prevent the race condition
	 * between Workload and Workload Interface Resources. Preserved for future use.
	 * Bug#15
	 */
	// w.Interfaces = expandIllumioWorkloadInterface(d.Get("interfaces").(*schema.Set).List())
}

/* Following code is commented to prevent the race condition
 * between Workload and Workload Interface Resources. Preserved for future use.
 * Bug#15
 */
// func expandIllumioWorkloadInterface(arr []interface{}) []models.WorkloadInterface {
// 	var wi []models.WorkloadInterface
// 	for _, e := range arr {
// 		elem := e.(map[string]interface{})
// 		wi = append(wi, models.WorkloadInterface{
// 			Name:                  elem["name"].(string),
// 			LinkState:             elem["link_state"].(string),
// 			Address:               elem["address"].(string),
// 			CidrBlock:             elem["cidr_block"].(int),
// 			DefaultGatewayAddress: elem["default_gateway_address"].(string),
// 			FriendlyName:          elem["friendly_name"].(string),
// 		})
// 	}
// 	return wi
// }
