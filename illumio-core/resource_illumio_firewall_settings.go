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

var (
	validWorkloadMode = []string{"illuminated", "enforced"}
	validIKEAuthType  = []string{"psk", "certificate"}
)

func schemaScopes(desc string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		Computed:    true,
		Description: fmt.Sprintf("%s. Either label or label_group can be specified", desc),
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"label": {
					Type:        schema.TypeSet,
					Optional:    true,
					Description: "Href of Label",
					Elem:        hrefSchemaRequired("Label", isLabelHref),
				},
				"label_group": {
					Type:        schema.TypeSet,
					Optional:    true,
					Description: "Href of Label Group",
					Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
				},
			},
		},
	}
}

func resourceIllumioFirewallSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioFirewallSettingsCreate,
		ReadContext:   resourceIllumioFirewallSettingsRead,
		UpdateContext: resourceIllumioFirewallSettingsUpdate,
		DeleteContext: resourceIllumioFirewallSettingsDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		SchemaVersion: 1,
		Description:   "Manages Illumio Firewall Settings",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of Firewall Settings",
			},
			"update_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Type of Update",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when these firewall settings were first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when these firewall settings were last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when these firewall settings were deleted",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who created this resource",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who last updated this resource",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deleted_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who deleted this resource",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ike_authentication_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Description:      `IKE authentication type to use for IPsec (SecureConnect and Machine Authentication). Allowed values are "psk" and "certificate"`,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validIKEAuthType, false)),
			},
			"firewall_coexistence": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "Firewall coexistence configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"illumio_primary": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Whether Illumio is primary firewall or not",
						},
						"scope": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of Href of label",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"href": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Href of Label",
									},
								},
							},
						},
						"workload_mode": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      `Match criteria to select workload(s). Allowed values are "enforced" and "illuminated"`,
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validWorkloadMode, false)),
						},
					},
				},
			},
			"static_policy_scopes":                  schemaScopes("scopes for static policy"),
			"containers_inherit_host_policy_scopes": schemaScopes("scopes for container inherit host policy"),
			"loopback_interfaces_in_policy_scopes":  schemaScopes("scopes for loopback interface"),
			"blocked_connection_reject_scopes":      schemaScopes("scopes for reject connections"),
		},
	}
}

func resourceIllumioFirewallSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	d.SetId(fmt.Sprintf("/orgs/%d/sec_policy/draft/firewall_settings", illumioClient.OrgID))
	return resourceIllumioFirewallSettingsUpdate(ctx, d, m)
}

func resourceIllumioFirewallSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	_, data, err := illumioClient.Get(fmt.Sprintf("/orgs/%d/sec_policy/draft/firewall_settings", orgID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	firewallSettingsKeys := []string{
		"href",
		"update_type",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"ike_authentication_type",
	}
	d.SetId(data.S("href").Data().(string))
	for _, key := range firewallSettingsKeys {
		if data.Exists(key) && data.S(key).Data() != nil {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}
	for _, key := range []string{
		"static_policy_scopes",
		"containers_inherit_host_policy_scopes",
		"blocked_connection_reject_scopes",
		"loopback_interfaces_in_policy_scopes",
	} {
		if data.Exists(key) && data.S(key).Data() != nil {
			d.Set(key, extractResourceScopes(data.S(key)))
		} else {
			d.Set(key, nil)
		}
	}
	key := "firewall_coexistence"
	if data.Exists(key) {
		d.Set(key, extractFirewallCoexistence(data))
	} else {
		d.Set(key, nil)
	}
	return diagnostics
}

func resourceIllumioFirewallSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	var diags diag.Diagnostics

	firewallSettings := &models.FirewallSettings{
		IKEAuthType:         d.Get("ike_authentication_type").(string),
		FirewallCoexistence: expandFirewallCoexistence(d.Get("firewall_coexistence").([]interface{})),
	}
	staticScopes, errs := expandScopes(d.Get("static_policy_scopes").([]interface{}), 1, 4)
	diags = append(diags, errs...)
	containerScopes, errs := expandScopes(d.Get("containers_inherit_host_policy_scopes").([]interface{}), 1, 4)
	diags = append(diags, errs...)
	loopbackScopes, errs := expandScopes(d.Get("loopback_interfaces_in_policy_scopes").([]interface{}), 1, 4)
	diags = append(diags, errs...)
	rejectConScopes, errs := expandScopes(d.Get("blocked_connection_reject_scopes").([]interface{}), 0, 4)
	diags = append(diags, errs...)
	firewallSettings.StaticPolicyScopes = staticScopes
	firewallSettings.ContainerPolicyScopes = containerScopes
	firewallSettings.RejectConnectionScopes = rejectConScopes
	firewallSettings.LoopbackPolicyScopes = loopbackScopes

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), firewallSettings)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref("firewall_settings", d.Id())

	return resourceIllumioFirewallSettingsRead(ctx, d, m)
}

func resourceIllumioFirewallSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "[illumio-core_firewall_settings] Ignoring Delete Operation.",
	})

	return diags
}

func expandFirewallCoexistence(firewallCoexistence []interface{}) models.FirewallCoexistence {
	fcs := models.FirewallCoexistence{}
	for _, i := range firewallCoexistence {
		im := i.(map[string]interface{})
		fcs = append(fcs, models.FirewallCoexistenceObj{
			IllumioPrimary: im["illumio_primary"].(bool),
			Scope:          models.GetHrefs(im["scope"].([]interface{})),
			WorkloadMode:   im["workload_mode"].(string),
		})
	}
	return fcs
}

func expandScopes(scopes []interface{}, min, max int) (models.Scopes, diag.Diagnostics) {
	scps := models.Scopes{}
	for _, scope := range scopes {
		scp := models.Scope{}

		if scope == nil {
			scps = append(scps, scp)
			continue
		}

		scopeObj := scope.(map[string]interface{})

		labels := scopeObj["label"].(*schema.Set).List()
		labelGroups := scopeObj["label_group"].(*schema.Set).List()

		if len(labels)+len(labelGroups) > max {
			return models.Scopes{}, diag.Errorf("[illumio-core_firewall_settings] at most %d blocks of label/label_group are allowed", max)
		}
		if len(labels)+len(labelGroups) < min {
			return models.Scopes{}, diag.Errorf("[illumio-core_firewall_settings] at least %d block(s) of label/label_group are required", min)
		}
		for _, label := range labels {
			s := models.ScopeObj{
				Label: getHrefObj(label),
			}
			scp = append(scp, s)
		}
		for _, labelGroup := range labelGroups {
			s := models.ScopeObj{
				LabelGroup: getHrefObj(labelGroup),
			}
			scp = append(scp, s)
		}
		scps = append(scps, scp)
	}
	return scps, diag.Diagnostics{}
}
