// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

var (
	validServiceProtos = []string{"-1", "1", "2", "4", "6", "17", "47", "58", "94"}
)

func resourceIllumioService() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceIllumioServiceRead,
		CreateContext: resourceIllumioServiceCreate,
		UpdateContext: resourceIllumioServiceUpdate,
		DeleteContext: resourceIllumioServiceDelete,
		Description:   "Manages Illumio Security Service",
		SchemaVersion: version,
		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of the service",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the Service (does not need to be unique). The name should be between 1 to 255 characters",
				ValidateDiagFunc: nameValidation,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Long description of the Service",
			},
			"description_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description URL Read-only to prevent XSS attacks",
			},
			"process_name": {
				Type:             schema.TypeString,
				ValidateDiagFunc: nameValidation,
				Optional:         true,
				Description:      "The process name. The name should be between 1 to 255 characters",
			},
			"service_ports": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Service ports",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Port Number. Also, the starting port when specifying a range. Allowed when value of proto is 6 or 17. Allowed range is 0 - 65535",
							ValidateDiagFunc: isStringAPortNumber(),
						},
						"to_port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "High end of port range if specifying a range. Allowed range is 0 - 65535",
							ValidateDiagFunc: isStringAPortNumber(),
						},
						"proto": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      `Transport protocol. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94`,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validServiceProtos, false)),
						},
						"icmp_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "ICMP Type. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 255",
							ValidateDiagFunc: isStringInRange(0, 255),
						},
						"icmp_code": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "ICMP Code. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 15",
							ValidateDiagFunc: isStringInRange(0, 15),
						},
					},
				},
			},
			"windows_services": {
				Type:          schema.TypeSet,
				Optional:      true,
				Description:   "Windows services",
				ConflictsWith: []string{"service_ports"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_name": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: nameValidation,
							Description:      "Name of Windows Service",
						},
						"process_name": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: nameValidation,
							Description:      "Name of running process",
						},
						"port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Port Number. Also, the starting port when specifying a range. Allowed when value of proto is 6 or 17. Allowed range is 0 - 65535",
							ValidateDiagFunc: isStringAPortNumber(),
						},
						"to_port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "High end of port range if specifying a range. Allowed range is 0 - 65535",
							ValidateDiagFunc: isStringAPortNumber(),
						},
						"proto": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      `Transport protocol. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94`,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validServiceProtos, false)),
						},
						"icmp_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "ICMP Type. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range is 0 - 255",
							ValidateDiagFunc: isStringInRange(0, 255),
						},
						"icmp_code": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "ICMP Code. Allowed when proto is 1 (ICMP) or 58 (ICMPv6). Allowed range 0 - 15",
							ValidateDiagFunc: isStringInRange(0, 15),
						},
					},
				},
			},
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "External data set identifier",
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "External data reference identifier",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Service was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Service was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this Service was deleted",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who created this Service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who last updated this Service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deleted_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who deleted this Service",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of update",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	var diags diag.Diagnostics

	orgID := pConfig.OrgID

	service := &models.Service{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		ProcessName:           d.Get("process_name").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
	}

	if val, exists := d.GetOk("service_ports"); exists {
		sps, errs := expandIllumioServiceServicePorts(val.(*schema.Set).List())
		if errs.HasError() {
			diags = append(diags, errs...)
			return diags
		} else {
			service.ServicePorts = sps
		}
	}

	if val, exists := d.GetOk("windows_services"); exists {
		wss, errs := expandIllumioWindowServices(val.(*schema.Set).List())

		if errs.HasError() {
			diags = append(diags, errs...)
			return diags
		} else {
			service.WindowsServices = wss
		}
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/services", orgID), service)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "services", data.S("href").Data().(string))
	d.SetId(data.S("href").Data().(string))
	return resourceIllumioServiceRead(ctx, d, m)
}

func expandIllumioServiceServicePorts(serPorts []interface{}) ([]map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	ports := []map[string]interface{}{}
	for _, serPort := range serPorts {
		s := serPort.(map[string]interface{})
		m := make(map[string]interface{})
		if isPortServiceSchemaValid(s, &diags) {
			if v, ok := getInt(s["proto"]); ok {
				m["proto"] = v
				if vPort, ok := getInt(s["port"]); ok {
					m["port"] = vPort
					if vToPort, ok := getInt(s["to_port"]); ok {
						if vToPort <= vPort {
							diags = append(diags, diag.Errorf("[illumio-core_service] Value of to_port can't be less or equal to value of port inside service_ports")...)
						} else {
							m["to_port"] = vToPort
						}
					}
				}
				if icmpcode, ok := getInt(s["icmp_code"]); ok {
					m["icmp_code"] = icmpcode
				}
				if icmptype, ok := getInt(s["icmp_type"]); ok {
					m["icmp_type"] = icmptype
				}
			}

		}

		ports = append(ports, m)
	}
	return ports, diags
}

func isPortServiceSchemaValid(p map[string]interface{}, diags *diag.Diagnostics) bool {
	portOk := p["port"].(string) != ""
	toPortOk := p["to_port"].(string) != ""
	icmpTypeOk := p["icmp_type"].(string) != ""
	icmpCodeOk := p["icmp_code"].(string) != ""

	vProto := p["proto"].(string)

	if vProto == "6" || vProto == "17" {
		if icmpCodeOk || icmpTypeOk {
			*diags = append(*diags, diag.Errorf("[illumio-core_service] icmp_code and icmp_type are not allowed when TCP or UDP protocol is specified, inside service ports")...)
			return false
		}
		if !portOk && toPortOk {
			*diags = append(*diags, diag.Errorf("[illumio-core_service] to_port parameter should be defined if port is specified, inside service ports")...)
			return false
		}
	} else if vProto == "1" || vProto == "58" {
		if portOk || toPortOk {
			*diags = append(*diags, diag.Errorf("[illumio-core_service] port and to_port parameter are not allowed if ICMP or ICMPv6 protocol is specified, inside service ports")...)
			return false
		}
		if icmpCodeOk && !icmpTypeOk {
			*diags = append(*diags, diag.Errorf("[illumio-core_service] icmp_type is required if icmp_code is specifiedn inside service ports")...)
			return false
		}
	} else {
		if icmpCodeOk || icmpTypeOk {
			*diags = append(*diags, diag.Errorf("[illumio-core_service] icmp_code and icmp_type are not allowed if protocols except ICMP and ICMPv6 are specified")...)
			return false
		}
		if portOk || toPortOk {
			*diags = append(*diags, diag.Errorf("[illumio-core_service] port and to_port parameter are not allowed if protocols except TCP and UDP are specified")...)
			return false
		}
	}
	return true
}

func expandIllumioWindowServices(winServs []interface{}) ([]map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	winServ := []map[string]interface{}{}
	for _, ws := range winServs {
		s := ws.(map[string]interface{})
		m := make(map[string]interface{})
		serviceNameOk := s["service_name"] != ""
		if serviceNameOk {
			m["service_name"] = s["service_name"].(string)
		}
		processNameOk := s["process_name"] != ""
		if processNameOk {
			m["process_name"] = s["process_name"].(string)
		}
		if isPortServiceSchemaValid(s, &diags) {
			if v, ok := getInt(s["proto"]); ok {
				m["proto"] = v
				if vPort, ok := getInt(s["port"]); ok {
					m["port"] = vPort
					if vToPort, ok := getInt(s["to_port"]); ok {
						if vToPort <= vPort {
							diags = append(diags, diag.Errorf("[illumio-core_service] Value of to_port can't be less or equal to value of port inside windows_services")...)
						} else {
							m["to_port"] = vToPort
						}
					}
				}
				if icmpcode, ok := getInt(s["icmp_code"]); ok {
					m["icmp_code"] = icmpcode
				}
				if icmptype, ok := getInt(s["icmp_type"]); ok {
					m["icmp_type"] = icmptype
				}
			}
		}
		winServ = append(winServ, m)
	}

	return winServ, diags
}

func resourceIllumioServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	_, data, err := illumioClient.Get(d.Id(), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, key := range []string{
		"href",
		"name",
		"description",
		"description_url",
		"process_name",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"update_type",
	} {

		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("service_ports") {
		serPorts := data.S("service_ports").Data().([]interface{})

		sps := []map[string]interface{}{}

		for _, serPort := range serPorts {

			sp := serPort.(map[string]interface{})

			for k, v := range serPort.(map[string]interface{}) {
				if v != nil {
					sp[k] = strconv.Itoa(int(v.(float64)))
				}
				sps = append(sps, sp)
			}
		}

		d.Set("service_ports", sps)
	} else {
		d.Set("service_ports", nil)
	}

	if data.Exists("windows_services") {
		winSers := data.S("windows_services").Data().([]interface{})

		wss := []map[string]interface{}{}

		for _, winSer := range winSers {
			ws := winSer.(map[string]interface{})

			for k, v := range winSer.(map[string]interface{}) {

				if v != nil {
					if k == "service_name" || k == "process_name" {
						ws[k] = v.(string)
					} else {
						ws[k] = strconv.Itoa(int(v.(float64)))
					}

				}
				wss = append(wss, ws)
			}
		}

		d.Set("windows_services", wss)
	} else {
		d.Set("windows_services", nil)
	}

	return diags
}

func resourceIllumioServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	service := &models.Service{}

	if d.HasChange("name") {
		service.Name = d.Get("name").(string)
	}

	if d.HasChange("process_name") {
		service.ProcessName = d.Get("process_name").(string)
		if isUpdatedToEmptyString(d.GetChange("process_name")) {
			diags = append(diags, diag.Errorf("[illumio-core_service] Once set, process_name cannot be updated to an empty string")...)
		}
	}

	service.Description = d.Get("description").(string)
	service.ExternalDataSet = d.Get("external_data_set").(string)
	service.ExternalDataReference = d.Get("external_data_reference").(string)

	if d.HasChange("service_ports") {
		sps, errs := expandIllumioServiceServicePorts(d.Get("service_ports").(*schema.Set).List())
		if errs.HasError() {
			diags = append(diags, errs...)
		} else {
			service.ServicePorts = sps
		}
	} else {
		service.ServicePorts = nil
	}

	if d.HasChange("windows_services") {
		wss, errs := expandIllumioWindowServices(d.Get("windows_services").(*schema.Set).List())
		if errs.HasError() {
			diags = append(diags, errs...)
		} else {
			service.WindowsServices = wss
		}
	} else {
		service.WindowsServices = nil
	}

	if service.ServicePorts != nil && service.WindowsServices != nil {
		diags = append(diags, diag.Errorf("[illumio-core_service] Cannot change OS type form windows to service or vice versa")...)
	}

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), service)

	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "services", d.Id())
	return resourceIllumioServiceRead(ctx, d, m)
}

func resourceIllumioServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()
	_, err := illumioClient.Delete(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "services", href)
	d.SetId("")
	return diags
}
