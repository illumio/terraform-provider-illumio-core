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
	validEnforcementModes = []string{"idle", "visibility_only", "full", "selective"}
	validVisibilityLevels = []string{"flow_full_detail", "flow_summary", "flow_drops", "flow_off", "enhanced_data_collection"}
)

func resourceIllumioPairingProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioPairingProfileCreate,
		ReadContext:   resourceIllumioPairingProfileRead,
		UpdateContext: resourceIllumioPairingProfileUpdate,
		DeleteContext: resourceIllumioPairingProfileDelete,
		Description:   "Manages Illumio Pairing Profile",
		SchemaVersion: version,

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of this pairing profile",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The short friendly name of the pairing profile",
				ValidateDiagFunc: nameValidation,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The long description of the pairing profile",
			},
			"enforcement_mode": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "visibility_only",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validEnforcementModes, false)),
				Description:      `Flag that controls whether mode can be overridden from pairing script. Allowed values are "idle", "visibility_only", "full" and "selective". Default value: "visibility_only"`,
			},
			"enforcement_mode_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Flag that controls whether enforcement mode can be overridden from pairing script, Default value: "true"`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "The enabled flag of the pairing profile",
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
			"allowed_uses_per_key": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "unlimited",
				ValidateDiagFunc: isUnlimitedOrValidRange(1, 2147483647),
				Description:      `The number of times pairing profile keys can be used. Allowed values are range(1-2147483647) and "unlimited". Default value: "unlimited"`,
			},
			"key_lifespan": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      `Number of seconds pairing profile keys will be valid for. Allowed values are range(1-2147483647) and "unlimited". Default value: "unlimited"`,
				ValidateDiagFunc: isUnlimitedOrValidRange(1, 2147483647),
				Default:          "unlimited",
			},
			"label": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Assigned labels",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Label URI",
						},
					},
				},
			},
			"env_label_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Flag that controls whether env label can be overridden from pairing script. Default value: "true"`,
			},
			"loc_label_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Flag that controls whether loc label can be overridden from pairing script. Default value: "true"`,
			},
			"role_label_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Flag that controls whether role label can be overridden from pairing script. Default value: "true"`,
			},
			"app_label_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Flag that controls whether app label can be overridden from pairing script. Default value: "true"`,
			},
			"log_traffic": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `Status of VEN(alternative of status). Default value: false`,
				Default:     false,
			},
			"log_traffic_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Flag that controls whether log_traffic can be overridden from pairing script. Default value: true`,
			},
			"visibility_level": {
				Type:             schema.TypeString,
				Optional:         true,
				Description:      `Visibility level of the agent. Allowed values are "flow_full_detail", "flow_summary", "flow_drops", "flow_off" and "enhanced_data_collection". Default value: "flow_summary"`,
				Default:          "flow_summary",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validVisibilityLevels, false)),
			},
			"visibility_level_lock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: `Flag that controls whether visibility_level can be overridden from pairing script. Default value: "true"`,
			},
			"agent_software_release": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `Agent software release associated with this paring profile. Default value: "Default ()"`,
				Default:     "Default ()",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this pairing profile was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this pairing profile was last updated",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who originally created this pairing profile",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this pairing profile",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of VEN",
			},
			"total_use_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of times the pairing profile has been used",
			},
			"last_pairing_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this pairing profile was last used for pairing a workload",
			},
			"is_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag indicating this is default auto-created pairing profile",
			},
			"status_lock": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag that controls whether status can be overridden from pairing script",
			},
			"caps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "CAPS for Workload",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceIllumioPairingProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	var diags diag.Diagnostics

	pairingProfile := &models.PairingProfile{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		EnforcementMode:       d.Get("enforcement_mode").(string),
		EnforcementModeLock:   d.Get("enforcement_mode_lock").(bool),
		Enabled:               d.Get("enabled").(bool),
		EnvLabelLock:          d.Get("env_label_lock").(bool),
		LocLabelLock:          d.Get("loc_label_lock").(bool),
		RoleLabelLock:         d.Get("role_label_lock").(bool),
		AppLabelLock:          d.Get("app_label_lock").(bool),
		LogTraffic:            d.Get("log_traffic").(bool),
		LogTrafficLock:        d.Get("log_traffic_lock").(bool),
		VisibilityLevel:       d.Get("visibility_level").(string),
		VisibilityLevelLock:   d.Get("visibility_level_lock").(bool),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		AgentSoftwareRelease:  d.Get("agent_software_release").(string),
		AllowedUsesPerKey:     d.Get("allowed_uses_per_key").(string),
		KeyLifespan:           d.Get("key_lifespan").(string),
	}

	if items, ok := d.GetOk("label"); ok {
		pairingProfile.Labels = models.GetHrefs(items.(*schema.Set).List())
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/pairing_profiles", pConfig.OrgID), pairingProfile)
	if diags.HasError() {
		return diags
	}
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.S("href").Data().(string))
	return resourceIllumioPairingProfileRead(ctx, d, m)
}

func resourceIllumioPairingProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		"enforcement_mode",
		"status",
		"enabled",
		"total_use_count",
		"last_pairing_at",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"is_default",
		"env_label_lock",
		"loc_label_lock",
		"role_label_lock",
		"app_label_lock",
		"enforcement_mode_lock",
		"log_traffic",
		"log_traffic_lock",
		"visibility_level",
		"visibility_level_lock",
		"status_lock",
		"agent_software_release",
		"caps",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	key := "allowed_uses_per_key"
	if data.Exists(key) {
		switch v := data.Data().(type) {
		case string:
			d.Set(key, v)
		case float64:
			d.Set(key, strconv.Itoa(int(v)))
		}
	} else {
		d.Set(key, nil)
	}

	key = "key_lifespan"
	if data.Exists(key) {
		switch v := data.Data().(type) {
		case string:
			d.Set(key, v)
		case float64:
			d.Set(key, strconv.Itoa(int(v)))
		}
	} else {
		d.Set(key, nil)
	}

	if data.Exists("labels") {
		d.Set("label", data.S("labels").Data())
	} else {
		d.Set("label", nil)
	}

	return diagnostics
}

func resourceIllumioPairingProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	var diags diag.Diagnostics

	pairingProfile := &models.PairingProfile{
		Enabled:               d.Get("enabled").(bool),
		EnvLabelLock:          d.Get("env_label_lock").(bool),
		LocLabelLock:          d.Get("loc_label_lock").(bool),
		RoleLabelLock:         d.Get("role_label_lock").(bool),
		AppLabelLock:          d.Get("app_label_lock").(bool),
		LogTraffic:            d.Get("log_traffic").(bool),
		LogTrafficLock:        d.Get("log_traffic_lock").(bool),
		VisibilityLevelLock:   d.Get("visibility_level_lock").(bool),
		EnforcementModeLock:   d.Get("enforcement_mode_lock").(bool),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		Labels:                models.GetHrefs(d.Get("label").(*schema.Set).List()),
	}

	if d.HasChange("name") {
		pairingProfile.Name = d.Get("name").(string)
	}

	if d.HasChange("enforcement_mode") {
		pairingProfile.EnforcementMode = d.Get("enforcement_mode").(string)
	}

	if d.HasChange("allowed_uses_per_key") {
		pairingProfile.AllowedUsesPerKey = d.Get("allowed_uses_per_key").(string)
	}

	if d.HasChange("key_lifespan") {
		pairingProfile.KeyLifespan = d.Get("key_lifespan").(string)
	}

	if d.HasChange("visibility_level") {
		pairingProfile.VisibilityLevel = d.Get("visibility_level").(string)
	}

	if d.HasChange("agent_software_release") {
		pairingProfile.AgentSoftwareRelease = d.Get("agent_software_release").(string)
	}

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), pairingProfile)

	if err != nil {
		return diag.FromErr(err)
	}
	return resourceIllumioPairingProfileRead(ctx, d, m)
}

func resourceIllumioPairingProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	_, err := illumioClient.Delete(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diagnostics
}
