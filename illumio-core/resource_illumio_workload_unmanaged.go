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

var (
	ValidWorkloadInterfaceLinkStateValues = []string{"up", "down", "unknown"}
	ValidWorkloadEnforcementModeValues    = []string{"idle", "visibility_only", "full", "selective"}
)

func resourceIllumioUnmanagedWorkload() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioUnmanagedWorkloadCreate,
		ReadContext:   resourceIllumioUnmanagedWorkloadRead,
		UpdateContext: resourceIllumioUnmanagedWorkloadUpdate,
		DeleteContext: resourceIllumioUnmanagedWorkloadDelete,
		Description:   "Manages Illumio Workload",
		SchemaVersion: 1,
		Schema:        unmanagedWorkloadSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func unmanagedWorkloadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"interfaces": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Workload network interfaces",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"name": {
						Type:             schema.TypeString,
						Required:         true,
						Description:      "Interface name. Can be up to 255 characters",
						ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
					},
					"link_state": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Interface link state. Allowed values are \"up\", \"down\", and \"unknown\"",
						ValidateDiagFunc: validation.ToDiagFunc(
							validation.StringInSlice(ValidWorkloadInterfaceLinkStateValues, false),
						),
					},
					"address": {
						Type:        schema.TypeString,
						Required:    true,
						Description: "Interface IP address. Must be in IPv4 or IPv6 format",
						ValidateDiagFunc: validation.ToDiagFunc(
							validation.IsIPAddress,
						),
					},
					"cidr_block": {
						Type:        schema.TypeInt,
						Optional:    true,
						Description: "Interface CIDR block bits",
					},
					"default_gateway_address": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "Interface Default Gateway IP address. Must be in IPv4 or IPv6 format",
						ValidateDiagFunc: validation.ToDiagFunc(
							validation.IsIPAddress,
						),
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
						Type:             schema.TypeString,
						Optional:         true,
						Description:      "User-friendly interface name. Can be up to 255 characters",
						ValidateDiagFunc: checkStringZerotoTwoHundredAndFiftyFive,
					},
					"loopback": {
						Type:        schema.TypeBool,
						Computed:    true,
						Description: "Whether or not the interface represents a loopback address on the workload",
					},
				},
			},
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
			Description: "Determines if this workload is online. Default value: true",
			Default:     true,
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
			Description: "User permissions for the object",
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
	}
}

func resourceIllumioUnmanagedWorkloadCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID
	workload := populateUnmanagedWorkloadFromResourceData(d)

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/workloads", orgID), workload)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.S("href").Data().(string))
	return resourceIllumioUnmanagedWorkloadRead(ctx, d, m)
}

func resourceIllumioUnmanagedWorkloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	key = "labels"
	if data.Exists(key) {
		labels := data.S(key)
		labelI := []map[string]interface{}{}

		for _, l := range labels.Children() {
			labelI = append(labelI, extractMap(l, []string{
				"href",
				"key",
				"value",
			}))
		}

		d.Set(key, labelI)
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

func resourceIllumioUnmanagedWorkloadUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	var diags diag.Diagnostics
	workload := populateUnmanagedWorkloadFromResourceData(d)

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), workload)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioUnmanagedWorkloadRead(ctx, d, m)
}

func resourceIllumioUnmanagedWorkloadDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func populateUnmanagedWorkloadFromResourceData(d *schema.ResourceData) *models.Workload {
	labels := expandLabelsOptionalKeyValue(d.Get("labels").(*schema.Set).List())
	interfaces := expandIllumioWorkloadInterface(d.Get("interfaces").(*schema.Set).List())
	ignoredInterfaceNames := getStringList(d.Get("ignored_interface_names"))

	return &models.Workload{
		Name:                                  d.Get("name").(string),
		AgentToPceCertificateAuthenticationID: d.Get("agent_to_pce_certificate_authentication_id").(string),
		DataCenter:                            d.Get("data_center").(string),
		DataCenterZone:                        d.Get("data_center_zone").(string),
		Description:                           d.Get("description").(string),
		EnforcementMode:                       d.Get("enforcement_mode").(string),
		ExternalDataReference:                 d.Get("external_data_reference").(string),
		ExternalDataSet:                       d.Get("external_data_set").(string),
		Hostname:                              d.Get("hostname").(string),
		Online:                                PtrTo(d.Get("online").(bool)),
		OsDetail:                              d.Get("os_detail").(string),
		OsID:                                  d.Get("os_id").(string),
		PublicIP:                              d.Get("public_ip").(string),
		ServicePrincipalName:                  d.Get("service_principal_name").(string),
		ServiceProvider:                       d.Get("service_provider").(string),
		Labels:                                &labels,
		Interfaces:                            &interfaces,
		IgnoredInterfaceNames:                 &ignoredInterfaceNames,
	}
}

func expandIllumioWorkloadInterface(arr []interface{}) []*models.WorkloadInterface {
	wi := make([]*models.WorkloadInterface, 0, len(arr))
	for _, e := range arr {
		elem := e.(map[string]interface{})
		wi = append(wi, &models.WorkloadInterface{
			Name:                  elem["name"].(string),
			LinkState:             elem["link_state"].(string),
			Address:               elem["address"].(string),
			CidrBlock:             elem["cidr_block"].(int),
			DefaultGatewayAddress: elem["default_gateway_address"].(string),
			FriendlyName:          elem["friendly_name"].(string),
		})
	}
	return wi
}
