// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioManagedWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioManagedWorkloadCreate,
		ReadContext:   resourceIllumioUnmanagedWorkloadRead, // reuse unmanaged workload read
		UpdateContext: resourceIllumioManagedWorkloadUpdate,
		DeleteContext: resourceIllumioManagedWorkloadDelete,
		Description:   "Manages Illumio Managed Workload",

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
				Description:      "Name of the Workload. The name should be up to 255 characters",
				ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
			},
			"hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The hostname of this workload. Set by the VEN.",
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
				Computed:    true,
				Description: "The public IP address of the server. Set by the VEN.",
			},
			"interfaces": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Workload network interfaces",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface name",
						},
						"link_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface link state. Can be \"up\", \"down\", or \"unknown\" ",
						},
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface IP address",
						},
						"cidr_block": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Interface CIDR block bits",
						},
						"default_gateway_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Default Gateway IP address",
						},
						"network": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Interface Network HREFs",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"network_detection_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface Network Detection Mode",
						},
						"friendly_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User-friendly interface name",
						},
						"loopback": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether or not the interface represents a loopback address on the workload",
						},
					},
				},
			},
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OS identifier for Workload. Set by the VEN.",
			},
			"os_detail": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Additional OS details - just displayed to end-user. Set by the VEN.",
			},
			"online": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determines if this workload is online. Set by the VEN.",
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
				Type:        schema.TypeString,
				Computed:    true,
				Description: "X.509 Subject distinguished name. The name should be up to 255 characters",
			},
			"visibility_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Visibility Level of workload(s) to return",
			},
			"ignored_interface_names": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Workload interface names to ignore (e.g. `eth0`). Ignored interfaces will not be included in policy configuration provided by the PCE.",
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
				Description: "VEN for Workload",
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
				Description: "Timestamp when this workload was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this workload was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this workload was last deleted",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this workload",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this workload",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this workload",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioManagedWorkloadCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Detail:   "[illumio-core_managed_workload] Managed workloads cannot be created through Terraform.",
		Summary:  "Please import managed workload objects after the VEN is paired.",
	})

	return diags
}

func resourceIllumioManagedWorkloadUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	var diags diag.Diagnostics
	labels := models.GetHrefs(d.Get("labels").(*schema.Set).List())
	ignoredInterfaceNames := getStringList(d.Get("ignored_interface_names"))

	var workload = &models.Workload{
		Name:                  d.Get("name").(string),
		DataCenter:            d.Get("data_center").(string),
		DataCenterZone:        d.Get("data_center_zone").(string),
		Description:           d.Get("description").(string),
		EnforcementMode:       d.Get("enforcement_mode").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ServicePrincipalName:  d.Get("service_principal_name").(string),
		ServiceProvider:       d.Get("service_provider").(string),
		Labels:                labels,
		IgnoredInterfaceNames: ignoredInterfaceNames,
	}

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), workload)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioUnmanagedWorkloadRead(ctx, d, m)
}

func resourceIllumioManagedWorkloadDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	hrefs := []models.Href{}
	hrefs = append(hrefs, *getHrefObj(d.Get("ven")))

	// unpair the VEN - this will also remove the workload record
	unpairUri := fmt.Sprintf("/orgs/%v/vens/unpair", illumioClient.OrgID)
	venUnpair := &models.VENsUnpair{
		FirewallRestore: "default",
		Hrefs:           hrefs,
	}

	response, err := illumioClient.Update(unpairUri, venUnpair)
	handleUnpairAndUpgradeOperationErrors(err, response, "workload", "managed")

	d.SetId("")
	return diagnostics
}
