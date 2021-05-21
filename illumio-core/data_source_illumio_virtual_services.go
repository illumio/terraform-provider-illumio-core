// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

/* Sample of API response
[
	{
		"name": "string",
		"description": "string",
		"external_data_set": null,
		"external_data_reference": null,
		"labels": [
			{
			"href": "string"
			}
		],
		"service_ports": [
			{
			"port": 0,
			"to_port": 0,
			"proto": 0
			}
		],
		"service": {},
		"apply_to": "host_only",
		"ip_overrides": [
			"string"
		],
		"service_addresses": [
			{}
		]
	}
]
*/

func datasourceIllumioVirtualServices() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioVirtualServicesRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Virtual Services",

		Schema: map[string]*schema.Schema{
			"pversion": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "draft",
				ValidateDiagFunc: isValidPversion(),
				Description:      `pversion of the security policy. Allowed values are "draft", "active", and numbers greater than 0. Default value: "draft"`,
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of virtual services",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of virtual service",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the virtual service",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The long description of this virtual service",
						},
						"pce_fqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PCE FQDN for this container cluster. Used in Supercluster only",
						},
						"external_data_set": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The data source from which a resource originates",
						},
						"external_data_reference": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique identifier within the external data source",
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Assigned labels",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
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
						"service_ports": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Service Ports",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Port Number. Also starting port when specifying port range",
									},
									"to_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specify port or port range to filter results.",
									},
									"proto": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Transport Protocol",
									},
								},
							},
						},
						"service": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "URI of associated service",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of service",
									},
								},
							},
						},
						"apply_to": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Firewall rule target for workloads bound to this virtual service: host_only or internal_bridge_network",
						},
						"ip_overrides": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Array of IPs or CIDRs as IP overrides",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"service_addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Service Addresses",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fqdn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "FQDN to assign to the virtual service",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Description for given fqdn",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Port Number. Also, the starting port when specifying a range",
									},
									"ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP address to assign to the virtual service",
									},
									"network": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Network URI for this IP address",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"caps": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "CAPS",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"update_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update Type",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Virtual Service was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Virtual Service was last updated",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Virtual Service was deleted",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who created this Virtual Service",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who last updated this Virtual Service",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"deleted_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who deleted this Virtual Service",
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description on which to filter. Supports partial matches",
			},
			"external_data_reference": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A unique identifier within the external data source",
			},
			"external_data_set": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The data source from which a resource originates",
			},
			"labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of lists of label URIs, encoded as a JSON string",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of Virtual Services to return. The integer should be a non-zero positive integer.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name on which to filter. Supports partial matches",
			},
			"service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service URI",
			},
			"service_address_fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "FQDN configured under service_address property, supports partial matches",
			},
			"service_address_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IP address configured under service_address property, supports partial matches",
			},
			"service_ports_port": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringInRange(-1, 65535),
				Description:      "Specify port or port range to filter results. The range is from -1 to 65535.",
			},
			"service_ports_proto": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringANumber(),
				Description:      "Protocol to filter on",
			},
			"usage": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringABoolean(),
				Description:      "Include Virtual Service usage flags",
			},
		},
	}
}

func dataSourceIllumioVirtualServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	pversion := d.Get("pversion").(string)

	paramKeys := []string{
		"description",
		"external_data_reference",
		"external_data_set",
		"labels",
		"max_results",
		"name",
		"service",
		"usage",
	}
	params := resourceDataToMap(d, paramKeys)

	if value, ok := d.GetOk("service_address_fqdn"); ok {
		params["service_address.fqdn"] = value.(string)
	}
	if value, ok := d.GetOk("service_address_ip"); ok {
		params["service_address.ip"] = value.(string)
	}
	if value, ok := d.GetOk("service_ports_port"); ok {
		params["service_ports.port"] = value.(string)
	}
	if value, ok := d.GetOk("service_ports_proto"); ok {
		params["service_ports.proto"] = value.(string)
	}

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%v/sec_policy/%v/virtual_services", pConfig.OrgID, pversion), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}
	keys := []string{
		"href",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"update_type",
		"name",
		"description",
		"external_data_set",
		"external_data_reference",
		"pce_fqdn",
		"labels",
		"ip_overrides",
		"apply_to",
		"caps",
	}

	for _, child := range data.Children() {
		m := extractMap(child, keys)

		key := "service"
		if child.Exists(key) {
			l := []map[string]string{}
			l = append(l, map[string]string{"href": child.S(key, "href").Data().(string)})
			m[key] = l
		} else {
			m[key] = nil
		}

		key = "service_addresses"
		if child.Exists(key) {
			l := []map[string]interface{}{}
			for _, c := range child.S(key).Children() {
				val := map[string]interface{}{}

				if v := c.S("fqdn").Data(); v != nil {
					val["fqdn"] = v.(string)
				}
				if v := c.S("description").Data(); v != nil {
					val["description"] = v.(string)
				}
				if v := c.S("port").Data(); v != nil {
					val["port"] = v
				}
				if v := c.S("ip").Data(); v != nil {
					val["ip"] = v.(string)
				}
				if v := c.S("network").Data(); v != nil {
					val["network"] = v
				}
				l = append(l, val)
			}
			m[key] = l
		} else {
			m[key] = nil
		}

		key = "service_ports"
		if child.Exists(key) {
			sps := []map[string]interface{}{}

			for _, serPort := range child.S(key).Children() {
				sp := map[string]interface{}{}

				for k, v := range serPort.ChildrenMap() {
					if k == "proto" {
						sp[k] = v.Data()
					} else if k == "port" || k == "to_port" {
						sp[k] = v.Data()
					}

					sps = append(sps, sp)
				}
			}

			m[key] = sps
		} else {
			m[key] = nil
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}
