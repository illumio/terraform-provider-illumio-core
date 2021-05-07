package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioSyslogDestinations() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioSyslogDestinationsRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Syslog Destinations",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of Syslog Destinations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Syslog Destination",
						},
						"pce_scope": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
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
				},
			},
		},
	}
}

func dataSourceIllumioSyslogDestinationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%v/settings/syslog/destinations", orgID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(fmt.Sprintf("/orgs/%v/settings/syslog/destinations", orgID))))

	dataMap := []map[string]interface{}{}

	for _, child := range data.Children() {
		m := map[string]interface{}{}

		for _, key := range []string{
			"href",
			"pce_scope",
			"type",
			"description",
		} {
			if child.Exists(key) {
				m[key] = child.S(key).Data()
			} else {
				m[key] = nil
			}
		}

		key := "audit_event_logger"
		aelKeys := []string{
			"configuration_event_included",
			"system_event_included",
			"min_severity",
		}
		if child.Exists(key) {
			m[key] = []interface{}{extractMap(child.S(key), aelKeys)}
		} else {
			m[key] = nil
		}

		key = "traffic_event_logger"
		telKey := []string{
			"traffic_flow_allowed_event_included",
			"traffic_flow_potentially_blocked_event_included",
			"traffic_flow_blocked_event_included",
		}
		if child.Exists(key) {
			m[key] = []interface{}{extractMap(child.S(key), telKey)}
		} else {
			m[key] = nil
		}

		key = "node_status_logger"
		nslKeys := []string{
			"node_status_included",
		}
		if child.Exists(key) {
			m[key] = []interface{}{extractMap(child.S(key), nslKeys)}
		} else {
			m[key] = nil
		}

		key = "remote_syslog"
		rsKeys := []string{
			"address",
			"port",
			"protocol",
			"tls_enabled",
			"tls_ca_bundle",
			"tls_verify_cert",
		}
		if child.Exists(key) {
			m[key] = []interface{}{extractMap(child.S(key), rsKeys)}
		} else {
			m[key] = nil
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
