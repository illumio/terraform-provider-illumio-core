// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioSyslogDestination() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceIllumioSyslogDestinationRead,

		SchemaVersion: version,
		Description:   "Represents Illumio Syslog Destination",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isSyslogDestinationHref,
				Description:      "URI of the destination",
			},
			"pce_scope": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "pce_scope for syslog destinations",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Destination type",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the destination",
			},
			"audit_event_logger": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "audit_event_logger details for destination. Single element list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"configuration_event_included": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Configuration (Northbound) auditable events",
						},
						"system_event_included": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "System (PCE) auditable events",
						},
						"min_severity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Minimum severity level of audit event messages",
						},
					},
				},
			},
			"traffic_event_logger": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "traffic_event_logger details for destination. Single element list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"traffic_flow_allowed_event_included": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Set to enable traffic flow events",
						},
						"traffic_flow_potentially_blocked_event_included": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Set to enable traffic flow events",
						},
						"traffic_flow_blocked_event_included": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Set to enable traffic flow events",
						},
					},
				},
			},
			"node_status_logger": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "node_status_logger details for destination. Single element list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_status_included": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Syslog messages regarding status of the nodes",
						},
					},
				},
			},
			"remote_syslog": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "remote_syslog details for destination. Single element list",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The remote syslog IP or DNS address",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The remote syslog port",
						},
						"protocol": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The protocol for streaming syslog messages",
						},
						"tls_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "To enable TLS",
						},
						"tls_ca_bundle": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Trustee CA bundle",
						},
						"tls_verify_cert": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Perform TLS verification",
						},
					},
				},
			},
		},
	}
}

func datasourceIllumioSyslogDestinationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Get("href").(string)

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.S("href").Data().(string))

	setIllumioSyslogDestinationState(data, d)

	return diags
}

func setIllumioSyslogDestinationState(data *gabs.Container, d *schema.ResourceData) {
	for _, k := range []string{
		"href",
		"pce_scope",
		"type",
		"description",
	} {
		if data.Exists(k) {
			d.Set(k, data.S(k).Data())

		} else {
			d.Set(k, nil)
		}
	}

	// Read JSON object of key audit_event_logger
	key := "audit_event_logger"
	aelKeys := []string{
		"configuration_event_included",
		"system_event_included",
		"min_severity",
	}
	if data.Exists(key) {
		d.Set(key, []interface{}{extractMap(data.S(key), aelKeys)})
	} else {
		d.Set(key, nil)
	}

	// Read JSON object of key traffic_event_logger
	key = "traffic_event_logger"
	telKey := []string{
		"traffic_flow_allowed_event_included",
		"traffic_flow_potentially_blocked_event_included",
		"traffic_flow_blocked_event_included",
	}
	if data.Exists(key) {
		d.Set(key, []interface{}{extractMap(data.S(key), telKey)})
	} else {
		d.Set(key, nil)
	}

	// Read JSON object of node_status_logger
	key = "node_status_logger"
	nslKeys := []string{
		"node_status_included",
	}
	if data.Exists(key) {
		d.Set(key, []interface{}{extractMap(data.S(key), nslKeys)})
	} else {
		d.Set(key, nil)
	}

	// Read JSON object of remote_syslog
	key = "remote_syslog"
	rsKeys := []string{
		"address",
		"port",
		"protocol",
		"tls_enabled",
		"tls_ca_bundle",
		"tls_verify_cert",
	}
	if data.Exists(key) {
		d.Set(key, []interface{}{extractMap(data.S(key), rsKeys)})
	} else {
		d.Set(key, nil)
	}
}
