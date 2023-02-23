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

func resourceIllumioLabelType() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioLabelTypeCreate,
		ReadContext:   resourceIllumioLabelTypeRead,
		UpdateContext: resourceIllumioLabelTypeUpdate,
		DeleteContext: resourceIllumioLabelTypeDelete,

		SchemaVersion: 1,
		Description:   "Manages Illumio Label Type. Requires PCE version 22.5.0 or higher",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of this label type",
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringLenBetween(1, LABEL_KEY_LENGTH_MAX),
				),
				Description: `Key in key-value pair. The value must be a string between 1 and 64 characters long`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The display name of the label type",
			},
			"display_info": {
				// XXX: surely there must be a better way to represent
				//      subobjects in the schema
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Object containing UI display information for the label type",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"initial": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "1-2 initial characters for use in the UI display. Defaults to the first letter of the label type's display_name",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.StringLenBetween(1, 2),
							),
						},
						"icon": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Icon for use in the UI display",
						},
						"background_color": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Background color in hexadecimal for UI display",
							ValidateDiagFunc: isValidColorCode,
						},
						"foreground_color": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Foreground color in hexadecimal for UI display",
							ValidateDiagFunc: isValidColorCode,
						},
						"sort_ordinal": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Optional user provided sort order for label type",
							ValidateDiagFunc: isStringANumber(),
						},
						"display_name_plural": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Optional pluralized form of the display name for the label type",
						},
					},
				},
			},
			"usage": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Usage of the label type",
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
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
			"deleted": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Flag to indicate whether the label type has been deleted or not",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this label type was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this label type was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this label type was deleted",
			},
			"created_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who created this label type",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who last updated this label type",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deleted_by": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "User who deleted this label type",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"caps": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User permissions for the object",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioLabelTypeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := illumioClient.OrgID

	labelType := &models.LabelType{
		Key:                   d.Get("key").(string),
		DisplayName:           d.Get("display_name").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
	}

	// set the display info if provided
	if di, ok := d.GetOk("display_info"); ok {
		displayInfo, diags := expandLabelTypeDisplayInfo(di.([]interface{})[0])
		if diags.HasError() {
			return diags
		}

		labelType.LabelTypeDisplayInfo = displayInfo
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/label_dimensions", orgID), labelType)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(data.S("href").Data().(string))

	return resourceIllumioLabelTypeRead(ctx, d, m)
}

func resourceIllumioLabelTypeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, key := range []string{
		"href",
		"key",
		"display_name",
		"usage",
		"external_data_set",
		"external_data_reference",
		"deleted",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"deleted_at",
		"deleted_by",
		"caps",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("display_info") {
		di := data.S("display_info")
		if di.Data() != nil {
			d.Set("display_info", []interface{}{
				extractMap(di, []string{
					"initial",
					"icon",
					"background_color",
					"foreground_color",
					"sort_ordinal",
					"display_name_plural",
				}),
			})
		}
	} else {
		d.Set("display_info", nil)
	}

	return diagnostics
}

func resourceIllumioLabelTypeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	label := &models.LabelType{
		DisplayName:           d.Get("display_name").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
	}

	_, err := illumioClient.Update(d.Id(), label)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioLabelTypeRead(ctx, d, m)
}

func resourceIllumioLabelTypeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func expandLabelTypeDisplayInfo(displayInfo interface{}) (*models.LabelTypeDisplayInfo, diag.Diagnostics) {
	di := displayInfo.(map[string]interface{})

	var sortOrdinal *int
	soStr := di["sort_ordinal"].(string)
	if soStr != "" {
		so, err := strconv.Atoi(di["sort_ordinal"].(string))
		if err != nil {
			return nil, diag.Errorf("[illumio-core_label_type] sort_ordinal must be an integer value")
		}
		sortOrdinal = &so
	}

	prov := &models.LabelTypeDisplayInfo{
		Initial:           di["initial"].(string),
		Icon:              di["icon"].(string),
		BackgroundColor:   di["background_color"].(string),
		ForegroundColor:   di["foreground_color"].(string),
		SortOrdinal:       sortOrdinal,
		DisplayNamePlural: di["display_name_plural"].(string),
	}

	return prov, diag.Diagnostics{}
}
