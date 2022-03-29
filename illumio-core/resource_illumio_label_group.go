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

func resourceIllumioLabelGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioLabelGroupCreate,
		ReadContext:   resourceIllumioLabelGroupRead,
		UpdateContext: resourceIllumioLabelGroupUpdate,
		DeleteContext: resourceIllumioLabelGroupDelete,
		SchemaVersion: version,
		Description:   "Manages Illumio Label Group",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of this label group",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the label group",
				ValidateDiagFunc: nameValidation,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The long description of the label group",
			},
			"key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice(validLabelKeys, false),
				),
				Description: `Key in key-value pair of contained labels or label groups. Allowed values are "role", "loc", "app" and "env"`,
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
			"sub_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Contained label groups",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "URI of label group",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Key in key-value pair",
						},
					},
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
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this label group was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this label group was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this label group was last deleted",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who created this label group",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this label group",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who deleted this label group",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioLabelGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	labelGroup := &models.LabelGroup{
		Name:                  d.Get("name").(string),
		Key:                   d.Get("key").(string),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
	}
	if items, ok := d.GetOk("labels"); ok {
		labelGroup.Labels = models.GetHrefs(items.(*schema.Set).List())
	}
	if items, ok := d.GetOk("sub_groups"); ok {
		labelGroup.SubGroups = models.GetHrefs(items.(*schema.Set).List())
	}
	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/label_groups", orgID), labelGroup)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "label_groups", data.S("href").Data().(string))
	d.SetId(data.S("href").Data().(string))
	return resourceIllumioLabelGroupRead(ctx, d, m)
}

func resourceIllumioLabelGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	// set computed/optional values from api response
	for _, key := range []string{
		"href",
		"name",
		"description",
		"key",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"deleted_at",
		"created_by",
		"updated_by",
		"deleted_by",
	} {

		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("labels") {
		labels := data.S("labels")
		labelI := []map[string]interface{}{}

		for _, l := range labels.Children() {
			labelI = append(labelI, extractMap(l, []string{"href", "key", "value"}))
		}

		d.Set("labels", labelI)
	} else {
		d.Set("labels", nil)
	}

	if data.Exists("sub_groups") {
		sub_groups := data.S("sub_groups")
		sub_groupI := []map[string]interface{}{}

		for _, sg := range sub_groups.Children() {
			sub_groupI = append(sub_groupI, extractMap(sg, []string{"href", "name"}))
		}

		d.Set("sub_groups", sub_groupI)
	} else {
		d.Set("sub_groups", nil)
	}

	return diagnostics
}

func resourceIllumioLabelGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	labelGroup := &models.LabelGroup{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		Labels:                models.GetHrefs(d.Get("labels").(*schema.Set).List()),
		SubGroups:             models.GetHrefs(d.Get("sub_groups").(*schema.Set).List()),
	}

	_, err := illumioClient.Update(d.Id(), labelGroup)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "label_groups", d.Id())

	return resourceIllumioLabelGroupRead(ctx, d, m)
}

func resourceIllumioLabelGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, err := illumioClient.Delete(href)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "label_groups", href)
	d.SetId("")
	return diagnostics
}
