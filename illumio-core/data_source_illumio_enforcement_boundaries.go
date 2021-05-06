package illumiocore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioEnforcementBoundaries() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioEnforcementBoundariesRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Enforcement Boundaries",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Enforcement Boundary",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of Enforcement Boundary",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Enforcement Boundary",
						},
						"ingress_service": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of Ingress Service. Only one of the {\"href\"} or {\"proto\", \"port\", \"to_port\"} parameter combination is allowed",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"proto": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Protocol number.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Port number used with protocol or starting port when specifying a range.",
									},
									"to_port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Upper end of port range.",
									},
									"href": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "URI of Ingress Service",
									},
								},
							},
						},
						"illumio_provider": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "providers for Enforcement Boundary. Only one actor can be specified in one illumio_provider block",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"actors": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "actors for illumio_provider.",
									},
									"label": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Href of Label",
										Elem:        hrefSchemaComputed("Label", isLabelHref),
									},
									"label_group": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Href of Label Group",
										Elem:        hrefSchemaComputed("Label Group", isLabelGroupHref),
									},
									"ip_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Href of IP List",
										Elem:        hrefSchemaComputed("IP List", isIPListHref),
									},
								},
							},
						},
						"consumer": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Consumers for Enforcement Boundary. Only one actor can be specified in one consumer block",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"actors": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "actors for consumers parameter. Allowed values are \"ams\" and \"container_host\"",
									},
									"label": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Href of Label",
										Elem:        hrefSchemaComputed("Label", isLabelHref),
									},
									"label_group": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Href of Label Group",
										Elem:        hrefSchemaComputed("Label Group", isLabelGroupHref),
									},
									"ip_list": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Href of IP List",
										Elem:        hrefSchemaComputed("IP List", isIPListHref),
									},
								},
							},
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Enforcement Boundary was first created",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Enforcement Boundary was last updated",
						},
						"deleted_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this Enforcement Boundary was last deleted",
						},
						"created_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who originally created this Enforcement Boundary",
						},
						"updated_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who last updated this Enforcement Boundary",
						},
						"deleted_by": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "User who last deleted this Enforcement Boundary",
						},
						"caps": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "CAPS for Enforcement Boundary",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "List of label URIs, encoded as a JSON string",
			},
			"max_results": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum number of enforcement boundaries to return.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by name supports partial matching",
			},
			"service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Service URI",
			},
			"service_ports_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify port or port range to filter results. The range is from -1 to 65535.",
			},
			"service_ports_proto": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Protocol to filter on",
			},
		},
	}
}

func dataSourceIllumioEnforcementBoundariesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	paramKeys := []string{
		"labels",
		"max_results",
		"name",
		"service",
	}
	params := resourceDataToMap(d, paramKeys)

	if value, ok := d.GetOk("service_address_port"); ok {
		params["service_address.port"] = value.(string)
	}
	if value, ok := d.GetOk("service_address_proto"); ok {
		params["service_address.proto"] = value.(string)
	}

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%v/sec_policy/draft/enforcement_boundaries", pConfig.OrgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	dataMap := []map[string]interface{}{}

	for _, child := range data.Children() {
		m := map[string]interface{}{}
		for _, key := range []string{
			"href",
			"name",
			"created_at",
			"updated_at",
			"deleted_at",
			"created_by",
			"updated_by",
			"deleted_by",
			"caps",
		} {
			if child.Exists(key) {
				m[key] = child.S(key).Data()
			} else {
				m[key] = nil
			}
		}
		if child.Exists("ingress_services") {
			ingServs := child.S("ingress_services").Data().([]interface{})
			iss := []map[string]interface{}{}

			for _, ingServ := range ingServs {
				is := ingServ.(map[string]interface{})

				for k, v := range ingServ.(map[string]interface{}) {
					if k == "href" {
						is[k] = v
					} else {
						is[k] = strconv.Itoa(int(v.(float64)))
					}

					iss = append(iss, is)
				}
			}

			m["ingress_service"] = iss
		} else {
			m["ingress_service"] = nil
		}
		m["illumio_provider"] = getEBActors(child.S("providers"))
		m["consumer"] = getEBActors(child.S("consumers"))

		dataMap = append(dataMap, m)
	}
	d.Set("items", dataMap)

	return diagnostics
}