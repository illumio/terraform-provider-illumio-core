package illumiocore

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Jeffail/gabs/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

var (
	validENIngSerProtos           = []string{"6", "17"}
	validENProducerConsumerActors = []string{"ams"}
)

func resourceIllumioEnforcementBoundary() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioEnforcementBoundaryCreate,
		ReadContext:   resourceIllumioEnforcementBoundaryRead,
		UpdateContext: resourceIllumioEnforcementBoundaryUpdate,
		DeleteContext: resourceIllumioEnforcementBoundaryDelete,
		SchemaVersion: version,
		Description:   "Manages Illumio Enforcement Boundary",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of this Enforcement Boundary",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the Enforcement Boundary",
				ValidateDiagFunc: nameValidation,
			},
			"ingress_services": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Collection of Ingress Service. Only one of the {\"href\"} or {\"proto\", \"port\", \"to_port\"} parameter combination is allowed",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"proto": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Protocol number. Allowed values are 6 and 17",
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(validENIngSerProtos, true)),
						},
						"port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Port number used with protocol or starting port when specifying a range. Valid range is 0-65535",
							ValidateDiagFunc: isStringAPortNumber(),
						},
						"to_port": {
							Type:             schema.TypeString,
							Optional:         true,
							Description:      "Upper end of port range. Valid range (0-65535)",
							ValidateDiagFunc: isStringAPortNumber(),
						},
						"href": {
							Type:             schema.TypeString,
							Optional:         true,
							ValidateDiagFunc: isServiceHref,
							Description:      "URI of Service",
						},
					},
				},
			},
			"providers": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "providers for Enforcement Boundary. Only one actor can be specified in one illumio_provider block",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actors": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "actors for illumio_provider. Valid value is \"ams\"",
							ValidateFunc: validation.StringInSlice(validENProducerConsumerActors, false),
						},
						"label": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Href of Label",
							Elem:        hrefSchemaRequired("Label", isLabelHref),
						},
						"label_group": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Href of Label Group",
							Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
						},
						"ip_list": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Href of IP List",
							Elem:        hrefSchemaRequired("IP List", isIPListHref),
						},
					},
				},
			},
			"consumers": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Consumers for Enforcement Boundary. Only one actor can be specified in one consumer block",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"actors": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "actors for consumers parameter. Allowed values are \"ams\" and \"container_host\"",
							ValidateFunc: validation.StringInSlice(validENProducerConsumerActors, false),
						},
						"label": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Href of Label",
							Elem:        hrefSchemaRequired("Label", isLabelHref),
						},
						"label_group": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Href of Label Group",
							Elem:        hrefSchemaRequired("Label Group", isLabelGroupHref),
						},
						"ip_list": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Href of IP List",
							Elem:        hrefSchemaRequired("IP List", isIPListHref),
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
	}
}

func resourceIllumioEnforcementBoundaryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	EnfBoun, diags := expandIllumioEnforcementBoundary(d)

	if diags.HasError() {
		return *diags
	}

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/enforcement_boundaries", orgID), EnfBoun)
	if err != nil {
		return diag.Errorf(err.Error())
	}

	pConfig.StoreHref(pConfig.OrgID, "enforcement_boundaries", data.S("href").Data().(string))

	d.SetId(data.S("href").Data().(string))

	return resourceIllumioEnforcementBoundaryRead(ctx, d, m)
}

func expandIllumioEnforcementBoundary(d *schema.ResourceData) (*models.EnforcementBoundary, *diag.Diagnostics) {
	var diags diag.Diagnostics
	enB := &models.EnforcementBoundary{
		Name: d.Get("name").(string),
	}

	ingServs, errs := expandIllumioEnforcementBoundaryIngressService(d.Get("ingress_services").(*schema.Set).List())
	diags = append(diags, errs...)
	enB.IngressServices = ingServs

	povs, errs := expandIllumioEnforcementBoundaryProviders(d.Get("providers").(*schema.Set).List())
	diags = append(diags, errs...)
	enB.Providers = povs

	cons, errs := expandIllumioEnforcementBoundaryConsumers(d.Get("consumers").(*schema.Set).List())
	diags = append(diags, errs...)
	enB.Consumers = cons

	return enB, &diags
}

func expandIllumioEnforcementBoundaryIngressService(inServices []interface{}) ([]map[string]interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	iss := []map[string]interface{}{}

	if len(inServices) == 0 {
		diags = append(diags, diag.Errorf("At least one ingress_service must be specified")...)
	}
	for _, service := range inServices {
		s := service.(map[string]interface{})

		m := make(map[string]interface{})

		if isIngressServiceSchemaValid(s, &diags) {
			if s["href"].(string) != "" {
				m["href"] = s["href"].(string)
			}

			if v, ok := getInt(s["proto"]); ok {
				m["proto"] = v
				if vPort, ok := getInt(s["port"]); ok {
					m["port"] = vPort
					if vToPort, ok := getInt(s["to_port"]); ok {
						if vToPort <= vPort {
							diags = append(diags, diag.Errorf("value of to_port can't be less or equal to value of port inside ingress_services")...)
						} else {
							m["to_port"] = vToPort
						}
					}
				}
			}
		}

		iss = append(iss, m)
	}

	return iss, diags
}

func expandIllumioEnforcementBoundaryConsumers(consumers []interface{}) ([]*models.EnforcementBoundaryProviderConsumer, diag.Diagnostics) {
	cons := []*models.EnforcementBoundaryProviderConsumer{}

	for _, consumer := range consumers {
		p := consumer.(map[string]interface{})

		con := &models.EnforcementBoundaryProviderConsumer{
			Actors:     p["actors"].(string),
			Label:      getHrefObj(p["label"]),
			LabelGroup: getHrefObj(p["label_group"]),
			IPList:     getHrefObj(p["ip_list"]),
		}

		if !con.HasOneActor() {
			return nil, diag.Errorf("consumer block can have only one rule actor")
		}
		cons = append(cons, con)
	}

	return cons, diag.Diagnostics{}
}

func expandIllumioEnforcementBoundaryProviders(providers []interface{}) ([]*models.EnforcementBoundaryProviderConsumer, diag.Diagnostics) {
	provs := []*models.EnforcementBoundaryProviderConsumer{}

	for _, provider := range providers {
		p := provider.(map[string]interface{})
		prov := &models.EnforcementBoundaryProviderConsumer{
			Actors:     p["actors"].(string),
			Label:      getHrefObj(p["label"]),
			LabelGroup: getHrefObj(p["label_group"]),
			IPList:     getHrefObj(p["ip_list"]),
		}
		if !prov.HasOneActor() {
			return nil, diag.Errorf("provider block can only have one rule actor")
		}

		provs = append(provs, prov)
	}
	return provs, diag.Diagnostics{}
}

func resourceIllumioEnforcementBoundaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	_, data, err := illumioClient.Get(d.Id(), nil)
	if err != nil {
		return diag.FromErr(err)
	}

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
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	if data.Exists("ingress_services") {
		ingServs := data.S("ingress_services").Data().([]interface{})
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

		d.Set("ingress_services", iss)
	} else {
		d.Set("ingress_services", nil)
	}

	d.Set("providers", getEBActors(data.S("providers")))
	d.Set("consumers", getEBActors(data.S("consumers")))

	return diagnostics
}

func getEBActors(data *gabs.Container) []map[string]interface{} {
	actors := []map[string]interface{}{}

	validActors := []string{
		"label",
		"label_group",
		"ip_list",
	}

	for _, actorArray := range data.Children() {

		actor := map[string]interface{}{}
		for k, v := range actorArray.ChildrenMap() {
			if k == "actors" {
				actor[k] = v.Data().(string)
			} else if contains(validActors, k) {
				vM := v.Data().(map[string]interface{})

				hrefs := map[string]string{}
				hrefs["href"] = vM["href"].(string)

				actor[k] = []map[string]string{hrefs}
			}
		}
		actors = append(actors, actor)
	}

	return actors
}

func resourceIllumioEnforcementBoundaryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	var diags diag.Diagnostics

	ingServs, errs := expandIllumioEnforcementBoundaryIngressService(d.Get("ingress_services").(*schema.Set).List())
	diags = append(diags, errs...)

	povs, errs := expandIllumioEnforcementBoundaryProviders(d.Get("providers").(*schema.Set).List())
	diags = append(diags, errs...)

	cons, errs := expandIllumioEnforcementBoundaryConsumers(d.Get("consumers").(*schema.Set).List())
	diags = append(diags, errs...)

	EB := &models.EnforcementBoundary{
		Name:            d.Get("name").(string),
		IngressServices: ingServs,
		Providers:       povs,
		Consumers:       cons,
	}

	if diags.HasError() {
		return diags
	}

	_, err := illumioClient.Update(d.Id(), EB)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref(pConfig.OrgID, "enforcement_boundaries", d.Id())

	return resourceIllumioEnforcementBoundaryRead(ctx, d, m)
}

func resourceIllumioEnforcementBoundaryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, err := illumioClient.Delete(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref(pConfig.OrgID, "enforcement_boundaries", href)

	d.SetId("")
	return diagnostics
}
