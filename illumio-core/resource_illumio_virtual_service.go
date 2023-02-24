// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

var (
	validApplyToKeys = []string{"host_only", "internal_bridge_network"}
)

func resourceIllumioVirtualService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioVirtualServiceCreate,
		ReadContext:   resourceIllumioVirtualServiceRead,
		UpdateContext: resourceIllumioVirtualServiceUpdate,
		DeleteContext: resourceIllumioVirtualServiceDelete,

		SchemaVersion: 2,
		Description:   "Manages Illumio Virtual Service",

		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the virtual service. The name should be between 1 to 255 characters",
				ValidateDiagFunc: nameValidation,
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of Virtual Service",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The long description of the virtual service",
			},
			"apply_to": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the virtual service. Allowed values are "host_only" and "internal_bridge_network"`,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(validApplyToKeys, false)),
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
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Contained labels",
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
							Description: "Key in key-value pair",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value in key-value pair",
						},
					},
				},
			},
			"service": {
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     1,
				Description:  "Associated service",
				ExactlyOneOf: []string{"service", "service_ports"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "URI of associated service",
						},
					},
				},
			},
			"service_ports": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "URI of associated service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      "Transport protocol. Allowed values are 6 (TCP) and 17 (UDP)",
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"6", "17"}, true)),
						},
						"port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Port Number. Also, the starting port when specifying a range. Allowed range is -1 - 65535",
							ValidateDiagFunc: isStringInRange(-1, 65535),
						},
						"to_port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "High end of port range inclusive if specifying a range. Allowed range is 0 - 65535",
							ValidateDiagFunc: isStringInRange(1, 65535),
						},
					},
				},
			},
			"ip_overrides": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Array of IPs or CIDRs as IP overrides",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"service_addresses": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: "List of service address. Specify one of the combination " +
					"{fqdn, description, port}, {ip, network} or {ip, port}",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
							Description:      "FQDN to assign to the virtual service.  Allowed formats: hostname, IP, or URI",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description for given fqdn",
						},
						"port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Port number of the service. Allowed range is -1 - 65535",
							ValidateDiagFunc: isStringInRange(-1, 65535),
						},
						"ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address to assign to the virtual service",
						},
						"network": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Network URI for this IP address",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "URI of associated service",
									},
								},
							},
						},
					},
				},
			},
			"pce_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PCE FQDN for this container cluster. Used in Supercluster only",
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update type",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this virtual service was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this virtual service was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this virtual service was last deleted",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this virtual service",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this virtual service",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this virtual service",
			},
			"caps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User permissions for the object",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		// Define upgrade from v1 to v2 to migrate the network_href field
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceIllumioVirtualServiceV1().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceIllumioVirtualServiceStateUpgradeV1,
				Version: 1,
			},
		},
	}
}

// XXX: v1 virtual_service resource schema
// required for migration from network_href to
// network.href in service_addresses field
func resourceIllumioVirtualServiceV1() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the virtual service. The name should be between 1 to 255 characters",
				ValidateDiagFunc: nameValidation,
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of Virtual Service",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The long description of the virtual service",
			},
			"apply_to": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `Name of the virtual service. Allowed values are "host_only" and "internal_bridge_network"`,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(validApplyToKeys, false)),
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
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Contained labels",
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
							Description: "Key in key-value pair",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value in key-value pair",
						},
					},
				},
			},
			"service": {
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     1,
				Description:  "Associated service",
				ExactlyOneOf: []string{"service", "service_ports"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "URI of associated service",
						},
					},
				},
			},
			"service_ports": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "URI of associated service",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto": {
							Type:             schema.TypeString,
							Required:         true,
							Description:      "Transport protocol. Allowed values are 6 (TCP) and 17 (UDP)",
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"6", "17"}, true)),
						},
						"port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Port Number. Also, the starting port when specifying a range. Allowed range is -1 - 65535",
							ValidateDiagFunc: isStringInRange(-1, 65535),
						},
						"to_port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "High end of port range inclusive if specifying a range. Allowed range is 0 - 65535",
							ValidateDiagFunc: isStringInRange(1, 65535),
						},
					},
				},
			},
			"ip_overrides": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Array of IPs or CIDRs as IP overrides",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"service_addresses": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: "List of service address. Specify one of the combination " +
					"{fqdn, description, port}, {ip, network_href} or {ip, port}",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
							Description:      "FQDN to assign to the virtual service.  Allowed formats: hostname, IP, or URI",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description for given fqdn",
						},
						"port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Port number of the service. Allowed range is -1 - 65535",
							ValidateDiagFunc: isStringInRange(-1, 65535),
						},
						"ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address to assign to the virtual service",
						},
						"network_href": { // Flattened from network.href
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Network URI for this IP address",
						},
					},
				},
			},
			"pce_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "PCE FQDN for this container cluster. Used in Supercluster only",
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update type",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this virtual service was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this virtual service was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this virtual service was last deleted",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this virtual service",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this virtual service",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this virtual service",
			},
			"caps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User permissions for the object",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceIllumioVirtualServiceStateUpgradeV1(ctx context.Context, rawState map[string]any, meta any) (map[string]any, error) {
	serviceAddresses := rawState["service_addresses"].([]any)
	updatedServiceAddresses := make([]map[string]any, 0, len(serviceAddresses))

	for _, sa := range serviceAddresses {
		ua := map[string]any{}
		for k, v := range sa.(map[string]any) {
			if k == "network_href" {
				if v != "" {
					ua["network"] = []map[string]any{{"href": v}}
				}
			} else {
				ua[k] = v
			}
		}
		updatedServiceAddresses = append(updatedServiceAddresses, ua)
	}

	rawState["service_addresses"] = updatedServiceAddresses

	return rawState, nil
}

func validateServiceAddress(v map[string]interface{}) error {
	if v["fqdn"] != "" && v["ip"] != "" {
		return errors.New("[illumio-core_virtual_service] Exactly One of [fqdn, ip] is allowed inside service address")
	}
	if v["fqdn"] == "" && v["ip"] == "" {
		return errors.New(`[illumio-core_virtual_service] Exactly One of [fqdn, ip] is required inside service address`)
	}
	network := extractNetworkFromVirtualServiceMap(v)
	if v["ip"] != "" && v["port"] == "" && (network == nil || network.Href == "") {
		return errors.New(`[illumio-core_virtual_service] Combination of [network, ip] or [ip, port] is required inside service address`)
	}
	return nil
}

func extractNetworkFromVirtualServiceMap(v map[string]interface{}) *models.Href {
	network := v["network"].(*schema.Set).List()
	if len(network) > 0 {
		nw := network[0].(map[string]interface{})
		return &models.Href{
			Href: nw["href"].(string),
		}
	}

	return nil
}

func resourceIllumioVirtualServiceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	diags := diag.Diagnostics{}
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	orgID := illumioClient.OrgID

	vs := &models.VirtualService{
		Name:                  d.Get("name").(string),
		ApplyTo:               d.Get("apply_to").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		Labels:                []models.Href{},
		ServicePorts:          []models.ServicePort{},
		IPOverrides:           []string{},
		ServiceAddresses:      []models.ServiceAddress{},
	}
	if v, ok := d.GetOk("description"); ok {
		vs.Description = v.(string)
	}

	if v, ok := d.GetOk("service"); ok {
		vs.Service = &models.Href{
			Href: v.([]interface{})[0].(map[string]interface{})["href"].(string),
		}
	} else if v, ok := d.GetOk("service_ports"); ok {
		servicePorts := v.(*schema.Set).List()
		vs.ServicePorts = expandSimpleServicePorts(servicePorts)
	}

	if v, ok := d.GetOk("service_addresses"); ok {
		sas, errs := expandServiceAddresses(v)
		if diags.HasError() {
			diags = append(diags, errs...)
		} else {
			vs.ServiceAddresses = sas
		}
	}

	if v, ok := d.GetOk("labels"); ok {
		vs.Labels = models.GetHrefs(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("ip_overrides"); ok {
		ips := []string{}
		for _, i := range v.(*schema.Set).List() {
			ips = append(ips, i.(string))
		}
		vs.IPOverrides = ips
	}

	if diags.HasError() {
		return diags
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/virtual_services", orgID), vs)
	if err != nil {
		return diag.FromErr(err)
	}

	href := data.S("href").Data().(string)

	pConfig.StoreHref("virtual_services", href)
	d.SetId(href)
	return resourceIllumioVirtualServiceRead(ctx, d, m)
}

func expandSimpleServicePorts(servicePorts []interface{}) []models.ServicePort {
	sps := []models.ServicePort{}
	for _, sp := range servicePorts {
		spmap := sp.(map[string]interface{})
		proto, _ := getInt(spmap["proto"])
		spi := models.ServicePort{Proto: proto}
		if v, ok := getInt(spmap["port"]); ok {
			spi.Port = intPointer(v)
		}
		if v, ok := getInt(spmap["to_port"]); ok {
			spi.ToPort = intPointer(v)
		}
		sps = append(sps, spi)
	}
	return sps
}

func expandServiceAddresses(v interface{}) ([]models.ServiceAddress, diag.Diagnostics) {
	var diags diag.Diagnostics

	saddresses := v.(*schema.Set).List()
	sas := []models.ServiceAddress{}
	for _, sa := range saddresses {
		sai := models.ServiceAddress{}
		samap := sa.(map[string]interface{})
		if err := validateServiceAddress(samap); err != nil {
			diags = append(diags, diag.FromErr(err)...)
			continue // Not valid service address object
		}
		if samap["fqdn"] != "" { // set fqdn object
			sai.Fqdn = samap["fqdn"].(string)
			if samap["description"] != "" {
				sai.Description = samap["description"].(string)
			}
		} else { // set {ip, network} or {ip, port}
			sai.IP = samap["ip"].(string)
			if network := samap["network"]; network != nil {
				vals := network.(*schema.Set).List()
				if len(vals) > 0 {
					sai.Network = &models.Href{
						Href: vals[0].(map[string]interface{})["href"].(string),
					}
				}
			}
		}
		if port, ok := getInt(samap["port"]); ok {
			sai.Port = intPointer(port)
		}
		sas = append(sas, sai)
	}

	return sas, diags
}

func resourceIllumioVirtualServiceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	var fields = []string{
		"href",
		"apply_to",
		"update_type",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"caps",
		"name",
		"description",
		"pce_fqdn",
		"external_data_set",
		"external_data_reference",
		"labels",
		"ip_overrides",
	}
	for _, key := range fields {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}
	key := "service"
	if data.Exists(key) {
		l := []map[string]string{}
		l = append(l, map[string]string{"href": data.S(key, "href").Data().(string)})
		d.Set(key, l)
	} else {
		d.Set(key, nil)
	}

	key = "service_addresses"
	if data.Exists(key) {
		l := []map[string]interface{}{}
		for _, child := range data.S(key).Children() {
			val := map[string]interface{}{}

			if v := child.S("fqdn").Data(); v != nil {
				val["fqdn"] = v.(string)
			}
			if v := child.S("description").Data(); v != nil {
				val["description"] = v.(string)
			}
			if v := child.S("port").Data(); v != nil {
				val["port"] = strconv.Itoa(int(v.(float64)))
			}
			if v := child.S("ip").Data(); v != nil {
				val["ip"] = v.(string)
			}
			if v := child.S("network").Data(); v != nil {
				n := []map[string]string{}
				n = append(n, map[string]string{"href": child.S("network", "href").Data().(string)})
				val["network"] = n
			}
			l = append(l, val)
		}
		d.Set(key, l)
	} else {
		d.Set(key, nil)
	}

	key = "service_ports"
	if data.Exists(key) {
		d.Set(key, extractServicePorts(data))
	} else {
		d.Set(key, nil)
	}

	return diag.Diagnostics{}
}

func resourceIllumioVirtualServiceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	diags := diag.Diagnostics{}
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	vs := &models.VirtualService{
		Name:                  d.Get("name").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		Description:           d.Get("description").(string),
		Labels:                []models.Href{},
		ServicePorts:          []models.ServicePort{},
		IPOverrides:           []string{},
		ServiceAddresses:      []models.ServiceAddress{},
	}

	if d.HasChange("") {
		vs.ApplyTo = d.Get("apply_to").(string)
	}

	if v, ok := d.GetOk("service"); ok {
		vs.Service = &models.Href{
			Href: v.([]interface{})[0].(map[string]interface{})["href"].(string),
		}
	} else if v, ok := d.GetOk("service_ports"); ok {
		servicePorts := v.(*schema.Set).List()
		vs.ServicePorts = expandSimpleServicePorts(servicePorts)
	}

	if v, ok := d.GetOk("service_addresses"); ok {
		sas, errs := expandServiceAddresses(v)
		if diags.HasError() {
			diags = append(diags, errs...)
		} else {
			vs.ServiceAddresses = sas
		}
	}

	if v, ok := d.GetOk("labels"); ok {
		vs.Labels = models.GetHrefs(v.(*schema.Set).List())
	}
	if v, ok := d.GetOk("ip_overrides"); ok {
		ips := []string{}
		for _, i := range v.(*schema.Set).List() {
			ips = append(ips, i.(string))
		}
		vs.IPOverrides = ips
	}
	if diags.HasError() {
		return diags
	}
	_, err := illumioClient.Update(href, vs)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref("virtual_services", href)
	return resourceIllumioVirtualServiceRead(ctx, d, m)
}

func resourceIllumioVirtualServiceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, err := illumioClient.Delete(href)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref("virtual_services", href)
	d.SetId("")
	return diagnostics
}
