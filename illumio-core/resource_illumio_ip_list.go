// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
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
		SchemaVersion: 1,
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
				Computed:     true,
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
							Description: `IP address or a low end of IP range. Might be specified with CIDR notation. The IP given should be in CIDR format example "0.0.0.0/0"`,
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
		CustomizeDiff: customdiff.Sequence(
			customizeIPRanges(),
		),
	}
}

// XXX: The PCE automatically performs several actions when IP ranges
// are uploaded that we need to manually recreate in the diff:
//
// * single address CIDR /32 (IPv4) and /128 (IPv6) notations are stripped
// * CIDR ranges are updated so that the IP matches the range's network address
// * identical ranges are merged
//
// To accommodate this, we normalize the HCL ip_ranges to avoid always
// presenting a diff when planning. As with container cluster workloads,
// this means the ip_ranges block needs to be Optional + Computed and we
// need to check the raw HCL against the state here and in the update
// function so the change behaves correctly when ip_ranges is cleared.
func customizeIPRanges() schema.CustomizeDiffFunc {
	return func(ctx context.Context, d *schema.ResourceDiff, m any) error {
		if d.HasChange("ip_ranges") {
			_, change := d.GetChange("ip_ranges")
			diff := change.(*schema.Set)
			normalized, _ := normalizeIPRanges(diff.List())
			d.SetNew("ip_ranges", schema.NewSet(diff.F, normalized))
		} else if ipRangesRemoved(d.GetRawConfig(), d.GetRawState()) {
			d.SetNewComputed("ip_ranges")
		}

		return nil
	}
}

func ipRangesRemoved(conf cty.Value, state cty.Value) bool {
	confMap := conf.AsValueMap()

	if ipranges, ok := confMap["ip_ranges"]; ok {
		if len(ipranges.AsValueSlice()) == 0 {
			if keyState, ok := state.AsValueMap()["ip_ranges"]; ok {
				if len(keyState.AsValueSlice()) > 0 {
					return true
				}
			}
		}
	}

	return false
}

func normalizeIPRanges(ipranges []any) ([]any, diag.Diagnostics) {
	var diags diag.Diagnostics
	normalized := make([]any, 0, len(ipranges))
	subnets := map[string][]string{}

	for _, r := range ipranges {
		iprange := r.(map[string]interface{})
		fromip := iprange["from_ip"].(string)
		ip, ipnet, err := net.ParseCIDR(fromip)

		// if there's no parse error, perform consolidation checks
		if err == nil {
			ones, _ := ipnet.Mask.Size()
			singleAddrOnes := net.IPv4len * 8
			if ip.To4() == nil {
				singleAddrOnes = net.IPv6len * 8
			}

			// check if this is a /32 or /128 range
			if ones == singleAddrOnes {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary: fmt.Sprintf("[illumio-core_ip_list] Detected single-address range (CIDR /32 for IPv4 or /128 for IPv6). "+
						"IP range %s will be converted to %s on the PCE.", fromip, ip.String()),
				})
				fromip = ip.String()
			}

			// the ipnet IP is the network address
			if !ip.Equal(ipnet.IP) {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary: fmt.Sprintf("[illumio-core_ip_list] Detected CIDR notation address range with host bits set: "+
						"IP range %s will be converted to %s on the PCE.", fromip, ipnet.String()),
				})
				fromip = ipnet.String()
			}

			// the PCE only merges subnets that are identical after normalization;
			// if the description or exclusion values are different, both are kept
			stringified := fromip + iprange["description"].(string) + strconv.FormatBool(iprange["exclusion"].(bool))
			if ips, ok := subnets[stringified]; ok {
				subnets[stringified] = append(ips, iprange["from_ip"].(string))
				continue
			} else {
				subnets[stringified] = []string{iprange["from_ip"].(string)}
			}
		}

		iprange["from_ip"] = fromip
		normalized = append(normalized, iprange)
	}

	for _, ranges := range subnets {
		if len(ranges) > 1 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary: fmt.Sprintf("[illumio-core_ip_list] Detected multiple identical address ranges: "+
					"IP ranges %v will be merged on the PCE.", ranges),
			})
		}
	}

	return normalized, diags
}

func resourceIllumioIPListCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	ipRanges, diags := expandIllumioIPListIPRanges(d.Get("ip_ranges").(*schema.Set).List())
	fqdns := expandIllumioIPListFQDNs(d.Get("fqdns").(*schema.Set).List())

	ipList := &models.IPList{
		Name:                  PtrTo(d.Get("name").(string)),
		Description:           PtrTo(d.Get("description").(string)),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		IPRanges:              ipRanges,
		FQDNs:                 fqdns,
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/ip_lists", orgID), ipList)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref("ip_lists", data.S("href").Data().(string))
	d.SetId(data.S("href").Data().(string))

	return append(diags, resourceIllumioIPListRead(ctx, d, m)...)
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
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	var ipRanges []models.IPRange

	if ipRangesRemoved(d.GetRawConfig(), d.GetRawState()) {
		ipRanges = []models.IPRange{}
	} else {
		ipRanges, diags = expandIllumioIPListIPRanges(d.Get("ip_ranges").(*schema.Set).List())
	}

	fqdns := expandIllumioIPListFQDNs(d.Get("fqdns").(*schema.Set).List())

	ipList := &models.IPList{
		Name:                  PtrTo(d.Get("name").(string)),
		Description:           PtrTo(d.Get("description").(string)),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		IPRanges:              ipRanges,
		FQDNs:                 fqdns,
	}

	_, err := illumioClient.Update(d.Id(), ipList)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref("ip_lists", d.Id())

	return append(diags, resourceIllumioIPListRead(ctx, d, m)...)
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

func expandIllumioIPListIPRanges(arr []interface{}) ([]models.IPRange, diag.Diagnostics) {
	ipranges := make([]models.IPRange, 0, len(arr))

	normalized, diags := normalizeIPRanges(arr)
	for _, elem := range normalized {
		iprange := elem.(map[string]any)
		ipranges = append(ipranges, models.IPRange{
			Description: PtrTo(iprange["description"].(string)),
			FromIP:      PtrTo(iprange["from_ip"].(string)),
			ToIP:        PtrTo(iprange["to_ip"].(string)),
			Exclusion:   PtrTo(iprange["exclusion"].(bool)),
		})
	}

	return ipranges, diags
}

func expandIllumioIPListFQDNs(arr []interface{}) []models.FQDN {
	fqdns := make([]models.FQDN, 0, len(arr))
	for _, elem := range arr {
		fqdns = append(fqdns, models.FQDN{
			FQDN:        PtrTo(elem.(map[string]interface{})["fqdn"].(string)),
			Description: PtrTo(elem.(map[string]interface{})["description"].(string)),
		})
	}
	return fqdns
}
