package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceIllumioVirtualServices() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceIllumioVirtualServicesRead,
		SchemaVersion: version,
		Description:   "Represents Illumio Virtual Services",

		Schema: map[string]*schema.Schema{
			"items": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "list of virtual service hrefs",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "URI of virtual service",
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
				ValidateDiagFunc: isStringANumber(),
				Description:      "Maximum number of Virtual Services to return.",
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
	if value, ok := d.GetOk("service_address_port"); ok {
		params["service_address.port"] = value.(string)
	}
	if value, ok := d.GetOk("service_address_proto"); ok {
		params["service_address.proto"] = value.(string)
	}

	_, data, err := illumioClient.AsyncGet(fmt.Sprintf("/orgs/%v/vens", pConfig.OrgID), &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%v", hashcode(paramsString(params))))

	d.Set("items", extractHrefs(data))

	return diagnostics
}
