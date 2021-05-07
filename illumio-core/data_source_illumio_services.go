package illumiocore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/* Sample of API response
[
	{
		href": "string",
		"name": "string",
		"description": "string",
		"description_url": "string",
		"process_name": "string",
		"service_ports": [
			{
			"port": 0,
			"to_port": 0,
			"proto": 0,
			"icmp_type": 0,
			"icmp_code": 0
			}
		],
		"windows_services": [
			{
			"service_name": "string",
			"process_name": "string",
			"port": 0,
			"to_port": 0,
			"proto": 0,
			"icmp_type": 0,
			"icmp_code": 0
			}
		],
		"external_data_set": "string",
		"external_data_reference": "string",
		"created_at": "1970-01-01T00:00:00.000Z",
		"updated_at": "1970-01-01T00:00:00.000Z",
		"deleted_at": "1970-01-01T00:00:00.000Z",
		"created_by": {
			"href": "string"
		},
		"updated_by": {
			"href": "string"
		},
		"deleted_by": {
			"href": "string"
		},
		"update_type": "string"
	}
]
*/

func datasourceIllumioServices() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioServicesRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Services",
		Schema: map[string]*schema.Schema{
			"pversion": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "draft",
				ValidateDiagFunc: isValidPversion(),
				Description:      `pversion of the security policy. Allowed values are "draft", "active" and numbers greater than 0. Default value: "draft"`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Long description of the servcie",
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "External data reference identifier",
			},
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "External data set identifier",
			},
			"max_results": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: isStringGreaterThanZero(),
				Description:      "Maximum number of Services to return. The integer should be a non-zero positive integer.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the servcie (does not need to be unique)",
			},
			"port": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Specify port or port range to filter results. The range is from -1 to 65535 (0 is not supported)",
				ValidateDiagFunc: servicePortValidation(),
			},
			"proto": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      "Protocol to filter on. Allowed values are -1, 1, 2, 4, 6, 17, 47, 58 and 94",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validServiceProtos, false)),
			},
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of services",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of service",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The short friendly name of the service",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Long Description of Service",
						},
						"description_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description URL Read-only to prevent XSS attacks",
						},
						"process_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The process name",
						},
						"service_ports": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Port Number ( the starting port when specifying a range)",
									},
									"to_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "High end of port range",
									},
									"proto": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Transport protocol",
									},
									"icmp_type": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "ICMP Type",
									},
									"icmp_code": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "ICMP Code",
									},
								},
							},
							Description: "Service ports of Illumio Service",
						},
						"windows_services": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of Windows Service",
									},
									"process_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of running process",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Port Number, also the starting port when specifying a range",
									},
									"to_port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "High end of port range",
									},
									"proto": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Transport protocol",
									},
									"icmp_type": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "ICMP Type",
									},
									"icmp_code": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "ICMP Code",
									},
								},
							},
						},
						"external_data_set": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External data set identifier",
						},
						"external_data_reference": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External data reference identifier",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time stamp when this Service was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time stamp when this Service was last updated",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time stamp when this Service was deleted",
						},
						"created_by": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "User who originally created this Service",
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
				},
			},
		},
	}
}

func dataSourceIllumioServicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	pversion := d.Get("pversion").(string)

	paramKeys := []string{
		"description",
		"external_data_reference",
		"external_data_set",
		"max_results",
		"name",
		"port",
		"proto",
	}

	params := resourceDataToMap(d, paramKeys)

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%d/sec_policy/%v/services", pConfig.OrgID, pversion), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}
	keys := []string{
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
	}

	for _, child := range data.Children() {
		m := extractMap(child, keys)

		if child.Exists("service_ports") {
			sps := child.S("service_ports")
			spI := []map[string]interface{}{}

			for _, sp := range sps.Children() {
				spI = append(spI, extractMap(sp, []string{"port", "to_port", "proto", "icmp_type", "icmp_code"}))
			}
			m["service_ports"] = spI
		} else {
			m["service_ports"] = nil
		}

		if child.Exists("windows_services") {
			wss := child.S("windows_services")
			wsI := []map[string]interface{}{}

			for _, ws := range wss.Children() {
				wsI = append(wsI, extractMap(ws, []string{"port", "to_port", "proto", "icmp_type", "icmp_code", "service_name", "process_name"}))
			}
			m["windows_services"] = wsI
		} else {
			m["windows_services"] = nil
		}

		dataMap = append(dataMap, m)
	}

	d.Set("items", dataMap)

	return diagnostics
}

func servicePortValidation() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		k, err := strconv.Atoi(v.(string))

		if err != nil {
			diags = append(diags, diag.Errorf("expected integer value, got: %v", v)...)
			return diags
		}

		if (1 > k || k > 65535) && k != -1 {
			diags = append(diags, diag.Errorf("expected to be in range 1-65535 or -1, got: %v", v)...)
		}

		return diags
	}
}
