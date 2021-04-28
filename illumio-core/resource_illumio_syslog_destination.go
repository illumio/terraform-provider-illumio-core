package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

var (
	syslogDestAELMinSeverityValidValues = []string{"error", "warning", "informational"}
	syslogDestTypeValidValues           = []string{"local_syslog", "remote_syslog"}
	syslogDestProtocolValidValues       = []int{6, 17}
)

func resourceIllumioSyslogDestination() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioSyslogDestinationCreate,
		ReadContext:   resourceIllumioSyslogDestinationRead,
		UpdateContext: resourceIllumioSyslogDestinationUpdate,
		DeleteContext: resourceIllumioSyslogDestinationDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio SyslogDestination",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the destination",
			},
			"pce_scope": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "pce_scope for destination",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(syslogDestTypeValidValues, false)),
				Description:      `Destination type. Allowed values are "local_syslog" and "remote_syslog"`,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of the destination",
			},
			"audit_event_logger": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "audit_event_logger details for destination. Single element list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration_event_included": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Configuration (Northbound) auditable events",
						},
						"system_event_included": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "System (PCE) auditable events",
						},
						"min_severity": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(syslogDestAELMinSeverityValidValues, false)),
							Description:      `Minimum severity level of audit event messages. Allowed values are "error", "warning" and "informational"`,
						},
					},
				},
			},
			"traffic_event_logger": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "traffic_event_logger details for destination. Single element list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"traffic_flow_allowed_event_included": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Set to enable traffic flow events",
						},
						"traffic_flow_potentially_blocked_event_included": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Set to enable traffic flow events",
						},
						"traffic_flow_blocked_event_included": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Set to enable traffic flow events",
						},
					},
				},
			},
			"node_status_logger": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "node_status_logger details for destination. Single element list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_status_included": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Syslog messages regarding status of the nodes",
						},
					},
				},
			},
			"remote_syslog": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "remote_syslog details for destination. Single element list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
							Description:      "The remote syslog IP or DNS address",
						},
						"port": {
							Type:             schema.TypeInt,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.IsPortNumberOrZero),
							Description:      "The remote syslog port",
						},
						"protocol": {
							Type:             schema.TypeInt,
							Required:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.IntInSlice(syslogDestProtocolValidValues)),
							Description:      "The protocol for streaming syslog messages. Allowed values are 6 and 17",
						},
						"tls_enabled": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "To enable TLS",
						},
						"tls_ca_bundle": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Trustee CA bundle",
						},
						"tls_verify_cert": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Perform TLS verification",
						},
					},
				},
			},
		},
	}
}

func resourceIllumioSyslogDestinationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	orgID := pConfig.OrgID

	syslogDest := expandIllumioSyslogdestination(d, &diags)

	diags = append(diags, syslogValidation(syslogDest.Type, syslogDest.RemoteSyslog != nil)...)

	if diags.HasError() {
		return diags
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%v/settings/syslog/destinations", orgID), syslogDest)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.S("href").Data().(string))

	return resourceIllumioSyslogDestinationRead(ctx, d, m)
}

func expandIllumioSyslogdestination(d *schema.ResourceData, diags *diag.Diagnostics) *models.SyslogDestination {
	return &models.SyslogDestination{
		PceScope:           getStringList(d.Get("pce_scope").(*schema.Set).List()),
		Description:        d.Get("description").(string),
		Type:               d.Get("type").(string),
		AuditEventLogger:   expandIllumioSyslogdestinationAuditEventLogger(d, diags),
		TrafficEventLogger: expandIllumioSyslogDestinationTrafficEventLogger(d, diags),
		NodeStatusLogger:   expandIllumioSyslogDestinationNodeStatusLogger(d, diags),
		RemoteSyslog:       expandIllumioSyslogDestinationRemoteSyslog(d, diags),
	}
}

func expandIllumioSyslogdestinationAuditEventLogger(d *schema.ResourceData, diags *diag.Diagnostics) *models.SyslogDestinationAuditEventLogger {
	v := d.Get("audit_event_logger").([]interface{})[0].(map[string]interface{})

	return &models.SyslogDestinationAuditEventLogger{
		MinSeverity:                v["min_severity"].(string),
		ConfigurationEventIncluded: v["configuration_event_included"].(bool),
		SystemEventIncluded:        v["system_event_included"].(bool),
	}
}

func expandIllumioSyslogDestinationTrafficEventLogger(d *schema.ResourceData, diags *diag.Diagnostics) *models.SyslogDestinationTrafficEventLogger {
	v := d.Get("traffic_event_logger").([]interface{})[0].(map[string]interface{})

	return &models.SyslogDestinationTrafficEventLogger{
		TrafficFlowAllowedEventIncluded:            v["traffic_flow_allowed_event_included"].(bool),
		TrafficFlowPotentiallyBlockedEventIncluded: v["traffic_flow_potentially_blocked_event_included"].(bool),
		TrafficFlowBlockedEventIncluded:            v["traffic_flow_blocked_event_included"].(bool),
	}
}

func expandIllumioSyslogDestinationNodeStatusLogger(d *schema.ResourceData, diags *diag.Diagnostics) *models.SyslogDestinationNodeStatusLogger {
	v := d.Get("node_status_logger").([]interface{})[0].(map[string]interface{})

	return &models.SyslogDestinationNodeStatusLogger{
		NodeStatusIncluded: v["node_status_included"].(bool),
	}
}

func expandIllumioSyslogDestinationRemoteSyslog(d *schema.ResourceData, diags *diag.Diagnostics) *models.SyslogDestinationRemoteSyslog {
	o, ok := d.GetOk("remote_syslog")

	if !ok {
		return nil
	}

	v := o.([]interface{})[0].(map[string]interface{})

	return &models.SyslogDestinationRemoteSyslog{
		Address:       v["address"].(string),
		Port:          v["port"].(int),
		Protocol:      v["protocol"].(int),
		TLSEnabled:    v["tls_enabled"].(bool),
		TLSCaBundle:   v["tls_ca_bundle"].(string),
		TLSVerifyCert: v["tls_verify_cert"].(bool),
	}
}

func resourceIllumioSyslogDestinationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	setIllumioSyslogDestinationState(data, d)

	return diags
}

func resourceIllumioSyslogDestinationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	if diags.HasError() {
		return diags
	}

	syslogDest := &models.SyslogDestination{
		Description: d.Get("description").(string),
	}

	if d.HasChange("pce_scope") {
		syslogDest.PceScope = d.Get("pce_scope").([]string)
	}

	if d.HasChange("type") {
		syslogDest.Type = d.Get("type").(string)
	}

	if d.HasChange("audit_event_logger") {
		syslogDest.AuditEventLogger = expandIllumioSyslogdestinationAuditEventLogger(d, &diags)
	}

	if d.HasChange("traffic_event_logger") {
		syslogDest.TrafficEventLogger = expandIllumioSyslogDestinationTrafficEventLogger(d, &diags)
	}

	if d.HasChange("node_status_logger") {
		syslogDest.NodeStatusLogger = expandIllumioSyslogDestinationNodeStatusLogger(d, &diags)
	}

	if d.HasChange("remote_syslog") {
		syslogDest.RemoteSyslog = expandIllumioSyslogDestinationRemoteSyslog(d, &diags)
		diags = append(diags, syslogValidation(d.Get("type").(string), syslogDest.RemoteSyslog != nil)...)
	}

	if d.HasChange("type") && !d.HasChange("remote_syslog") {
		_, ok := d.GetOk("remote_syslog")
		diags = append(diags, syslogValidation(d.Get("type").(string), ok)...)
	}

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), syslogDest)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioSyslogDestinationRead(ctx, d, m)
}

func syslogValidation(destType string, isRemoteSyslogSet bool) diag.Diagnostics {
	var diags diag.Diagnostics
	if destType == "remote_syslog" {
		if !isRemoteSyslogSet {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "remote_syslog block expected",
				Detail:        `When type is set to "remote_syslog", remote_syslog block is required`,
				AttributePath: cty.Path{cty.GetAttrStep{Name: "type"}},
			})
		}
	} else {
		if isRemoteSyslogSet {
			diags = append(diags, diag.Diagnostic{
				Severity:      diag.Error,
				Summary:       "remote_syslog block not allowed",
				Detail:        `When type is set to "local_syslog", remote_syslog block is not allowed`,
				AttributePath: cty.Path{cty.GetAttrStep{Name: "type"}},
			})
		}
	}

	return diags
}

func resourceIllumioSyslogDestinationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	_, err := illumioClient.Delete(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}
