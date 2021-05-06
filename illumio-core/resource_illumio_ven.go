package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

var (
	venStatusValidValues = []string{"active", "suspended"}
)

func resourceIllumioVEN() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioVENCreate,
		ReadContext:   resourceIllumioVENRead,
		UpdateContext: resourceIllumioVENUpdate,
		DeleteContext: resourceIllumioVENDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio VEN",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of VEN",
			},
			"name": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Friendly name for the VEN. The name should be upto 255 characters.",
				ValidateDiagFunc: nameValidation,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
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
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(venStatusValidValues, false)),
				Description:      `Status of the VEN. Allowed values are "active" and "suspended`,
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
				Optional:    true,
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
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface name",
						},
						"loopback": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Loopback for Workload Interface",
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
							Description: "The number of bits in the subnet /24 is 255.255.255.0",
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
				Description: "",
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
							Type:        schema.TypeSet,
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

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioVENCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Error,
			Detail:   "[illumio-core_ven] Cannot use Create Operation.",
			Summary:  "Please use terrform import...",
		},
	}
}

func resourceIllumioVENRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	setIllumioVENState(data, d)

	return diags
}

func resourceIllumioVENUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	ven := &models.VEN{
		Description: d.Get("description").(string),
		Name:        d.Get("name").(string),
	}

	if v, ok := d.GetOk("target_pce_fqdn"); ok {
		v, _ := v.(string)
		ven.TargetPCEFqdn = &v
	}

	if v, ok := d.GetOk("status"); ok {
		ven.Status = v.(string)
	}

	_, err := illumioClient.Update(d.Id(), ven)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioVENRead(ctx, d, m)
}

func resourceIllumioVENDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Warning,
			Detail:   "[illumio-core_ven] Ignoring Delete Operation...",
		},
	}
}
