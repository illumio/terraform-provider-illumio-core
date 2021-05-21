// Copyright 2021 Illumio, Inc. All Rights Reserved.

package illumiocore

import (
	"context"
	"crypto/aes"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

// getAESGCMKeyFromEnv - returns ILLUMIO_AES_GCM_KEY env variable, err if not set
func getAESGCMKeyFromEnv() (string, error) {
	key := os.Getenv("ILLUMIO_AES_GCM_KEY")
	if key == "" {
		return "", errors.New("[illumio-core_pairing_keys] ILLUMIO_AES_GCM_KEY environment variable is not set")
	}
	return key, nil
}

func pairingKeyPrequisiteValidation() schema.SchemaValidateDiagFunc {
	return func(v interface{}, path cty.Path) diag.Diagnostics {
		var diags diag.Diagnostics
		if d := isPairingProfileHref(v, path); d.HasError() {
			diags = append(diags, d...)
		}
		key, err := getAESGCMKeyFromEnv()
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  err.Error(),
				Detail:   "Please set key to encrypt activation token",
			})
			return diags
		}
		k, err := hex.DecodeString(key)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "[illumio-core_pairing_keys] Could not decode AES GCM key",
				Detail:   "Key should be 128/192/256 bit in hex format",
			})
		}
		_, err = aes.NewCipher(k)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "[illumio-core_pairing_keys] Invalid AES GCM key",
				Detail:   err.Error(),
			})
		}

		return diags
	}
}

func resourceIllumioPairingKeys() *schema.Resource {

	return &schema.Resource{
		CreateContext: resourceIllumioPairingKeysCreate,
		ReadContext:   resourceIllumioPairingKeysRead,
		UpdateContext: resourceIllumioPairingKeysUpdate,
		DeleteContext: resourceIllumioPairingKeysDelete,
		Description:   "Manages Illumio Pairing Keys",
		SchemaVersion: version,

		Schema: map[string]*schema.Schema{
			"pairing_profile_href": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Href of pairing profile",
				ValidateDiagFunc: pairingKeyPrequisiteValidation(),
			},
			"token_count": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "Count of token to generate/maintain. It can be accessed in `activation_tokens` On increasing the count, new activation tokens will be generated. " +
					"On decreasing the count `activation_tokens` list will shrink to that size and extra keys will be discarded. Allowed range is 1 to 5",
				ValidateDiagFunc: validation.ToDiagFunc(validation.IntBetween(1, 5)),
			},
			"activation_tokens": {
				Type:     schema.TypeList,
				Computed: true,
				// Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of activation tokens",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activation_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Encrypted activation token (encrypted by key set in env. variable `ILLUMIO_AES_GCM_KEY`)",
						},
						"nonce": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Nonce used in encrypting activation token",
						},
					},
				},
			},
		},
	}
}

func resourceIllumioPairingKeysCommon(activationTokens []interface{}, addCount int, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Get("pairing_profile_href").(string)

	atLeastOneSuccess := false
	for i := 1; i <= addCount; i++ {
		_, data, err := illumioClient.Create(href+"/pairing_key", &models.PairingKey{})
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  fmt.Sprintf("[illumio-core_pairing_keys] Could not generate activation token - Error: %v", err),
			})
			// As hard limit reached, we should stop calling api
			if strings.Contains(err.Error(), "hard limit reached") {
				log.Printf("[illumio-core_pairing_keys] Hard limit reached for pairing profile - %v", href)
				break
			}
		} else {
			key, _ := getAESGCMKeyFromEnv() // suppressing error as it should hit error in validation phase
			activationCode := data.S("activation_code").Data().(string)
			encryptedToken, nonce, err := aesGcmEncrypt(key, activationCode)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  fmt.Sprintf("[illumio-core_pairing_keys] Could not encrypt activation token - Error: %v", err),
				})
			} else {
				activationTokens = append(activationTokens, map[string]string{
					"activation_token": encryptedToken,
					"nonce":            nonce,
				})
				atLeastOneSuccess = true
			}
		}
	}
	if atLeastOneSuccess {
		// Store generated tokens
		d.SetId(href + "/pairing_keys")
		d.Set("activation_tokens", activationTokens)

		if diags.HasError() {
			newDiags := diag.Diagnostics{}
			newDiags = append(newDiags, diag.Diagnostic{
				Severity: diag.Warning, // was able to generate some tokens, so warning
				Summary:  "[illumio-core_pairing_keys] Could not generate required tokens",
				Detail:   "Generated tokens are saved, please do terraform apply to try again",
			})
			newDiags = append(newDiags, diags...)
			return newDiags
		} else {
			return diags
		}
	}
	newDiags := diag.Diagnostics{}
	newDiags = append(newDiags, diag.Diagnostic{
		Severity: diag.Error, // could not generate single token, so error
		Summary:  "[illumio-core_pairing_keys] Could not generate required tokens",
		Detail:   "Generated tokens are saved, please do terraform apply to try again",
	})
	newDiags = append(newDiags, diags...)
	return newDiags
}

func resourceIllumioPairingKeysCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	tokenCount := d.Get("token_count").(int)
	activationTokens := []interface{}{}
	return resourceIllumioPairingKeysCommon(activationTokens, tokenCount, d, m)

}

func resourceIllumioPairingKeysUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if d.HasChange("pairing_profile_href") {
		return diag.Errorf("[illumio-core_pairing_keys] Can not change pairing_profile_href once set")
	}
	activationTokens := d.Get("activation_tokens").([]interface{})
	old := len(activationTokens) // relying on actual length of activation tokens instead of last user input
	new := d.Get("token_count").(int)
	if new < old {
		activationTokens = activationTokens[:new]
		d.Set("activation_tokens", activationTokens)
		return diag.Diagnostics{}
	}
	return resourceIllumioPairingKeysCommon(activationTokens, new-old, d, m)
}

func resourceIllumioPairingKeysRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return diag.Diagnostics{}
}

func resourceIllumioPairingKeysDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
