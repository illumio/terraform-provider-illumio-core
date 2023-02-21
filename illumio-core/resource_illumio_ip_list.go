// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioIPList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioIPListCreate,
		ReadContext:   resourceIllumioIPListRead,
		UpdateContext: resourceIllumioIPListUpdate,
		DeleteContext: resourceIllumioIPListDelete,
		SchemaVersion: version,
		Description:   "Manages Illumio IP List",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of this IP List",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the IP List. The name should be between 1 to 255 characters",
				ValidateDiagFunc: nameValidation,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the IP List",
			},
			"ip_ranges": {
				Type:         schema.TypeSet,
				Optional:     true,
				AtLeastOneOf: []string{"ip_ranges", "fqdns"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of IP Range",
						},
						"from_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IP address or a low end of IP range. Might be specified with CIDR notation. The IP given should be in CIDR format example \"0.0.0.0/0\"",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.Any(validation.IsIPAddress, validation.IsCIDR),
							),
						},
						"to_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "High end of an IP range",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.IsIPAddress,
							),
						},
						"exclusion": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether this IP address is an exclusion. Exclusions must be a strict subset of inclusive IP addresses",
							Default:     false,
						},
					},
				},
				Description: "IP addresses or ranges",
			},
			"fqdns": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Fully Qualified Domain Name for IP List. Supported formats are hostname, IP, and URI",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of FQDN",
						},
					},
				},
				Description: "Collection of Fully Qualified Domain Names",
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
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IP List was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IP List was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IP List was last deleted",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this IP List",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this IP List",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last deleted this IP List",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioIPListCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	ipList := &models.IPList{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
	}

	ipList.IPRanges = expandIllumioIPListIPRanges(d.Get("ip_ranges").(*schema.Set).List())
	ipList.FQDNs = expandIllumioIPListFQDNs(d.Get("fqdns").(*schema.Set).List())

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/ip_lists", orgID), ipList)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref("ip_lists", data.S("href").Data().(string))
	d.SetId(data.S("href").Data().(string))
	return resourceIllumioIPListRead(ctx, d, m)
}

func resourceIllumioIPListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
	for _, key := range []string{
		"href",
		"name",
		"description",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"deleted_at",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("ip_ranges") {
		ip_ranges := data.S("ip_ranges")
		ip_rangeI := []map[string]interface{}{}

		for _, ip := range ip_ranges.Children() {
			ip_rangeI = append(ip_rangeI, extractMap(ip, []string{"description", "from_ip", "to_ip", "exclusion"}))
		}
		d.Set("ip_ranges", ip_rangeI)
	} else {
		d.Set("ip_ranges", nil)
	}

	if data.Exists("fqdns") {
		fqdns := data.S("fqdns")
		fqdnI := []map[string]interface{}{}

		for _, ip := range fqdns.Children() {
			fqdnI = append(fqdnI, extractMap(ip, []string{"fqdn", "description"}))
		}

		d.Set("fqdns", fqdnI)
	} else {
		d.Set("fqdns", nil)
	}

	return diagnostics
}

func resourceIllumioIPListUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	ipList := &models.IPList{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		IPRanges:              expandIllumioIPListIPRanges(d.Get("ip_ranges").(*schema.Set).List()),
		FQDNs:                 expandIllumioIPListFQDNs(d.Get("fqdns").(*schema.Set).List()),
	}

	_, err := illumioClient.Update(d.Id(), ipList)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref("ip_lists", d.Id())

	return resourceIllumioIPListRead(ctx, d, m)
}

func resourceIllumioIPListDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()
	_, err := illumioClient.Delete(href)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref("ip_lists", href)

	d.SetId("")
	return diagnostics
}

func expandIllumioIPListIPRanges(arr []interface{}) []models.IPRange {
	var ipranges []models.IPRange
	for _, elem := range arr {
		ipranges = append(ipranges, models.IPRange{
			Description: elem.(map[string]interface{})["description"].(string),
			FromIP:      elem.(map[string]interface{})["from_ip"].(string),
			ToIP:        elem.(map[string]interface{})["to_ip"].(string),
			Exclusion:   BoolPtr(elem.(map[string]interface{})["exclusion"].(bool)),
		})
	}
	return ipranges
}

func expandIllumioIPListFQDNs(arr []interface{}) []models.FQDN {
	var fqdns []models.FQDN
	for _, elem := range arr {
		fqdns = append(fqdns, models.FQDN{
			FQDN:        elem.(map[string]interface{})["fqdn"].(string),
			Description: elem.(map[string]interface{})["description"].(string),
		})
	}
	return fqdns
}
