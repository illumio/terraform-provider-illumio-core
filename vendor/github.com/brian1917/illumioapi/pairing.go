package illumioapi

import (
	"fmt"
	"strings"
)

// PairingProfile represents a pairing profile in the Illumio PCE
type PairingProfile struct {
	AllowedUsesPerKey     string     `json:"allowed_uses_per_key,omitempty"`
	AppLabelLock          bool       `json:"app_label_lock"`
	CreatedAt             string     `json:"created_at,omitempty"`
	CreatedBy             *CreatedBy `json:"created_by,omitempty"`
	Description           string     `json:"description,omitempty"`
	Enabled               bool       `json:"enabled"`
	EnvLabelLock          bool       `json:"env_label_lock"`
	ExternalDataReference string     `json:"external_data_reference,omitempty"`
	ExternalDataSet       string     `json:"external_data_set,omitempty"`
	Href                  string     `json:"href,omitempty"`
	IsDefault             bool       `json:"is_default,omitempty"`
	KeyLifespan           string     `json:"key_lifespan,omitempty"`
	Labels                []*Label   `json:"labels,omitempty"`
	LastPairingAt         string     `json:"last_pairing_at,omitempty"`
	LocLabelLock          bool       `json:"loc_label_lock"`
	LogTraffic            bool       `json:"log_traffic"`
	LogTrafficLock        bool       `json:"log_traffic_lock"`
	Mode                  string     `json:"mode,omitempty"`
	ModeLock              bool       `json:"mode_lock"`
	Name                  string     `json:"name,omitempty"`
	RoleLabelLock         bool       `json:"role_label_lock"`
	TotalUseCount         int        `json:"total_use_count,omitempty"`
	UpdatedAt             string     `json:"updated_at,omitempty"`
	UpdatedBy             *UpdatedBy `json:"updated_by,omitempty"`
	VisibilityLevel       string     `json:"visibility_level,omitempty"`
	VisibilityLevelLock   bool       `json:"visibility_level_lock"`
}

// PairingKey represents a VEN pairing key
type PairingKey struct {
	ActivationCode string `json:"activation_code,omitempty"`
}

// GetPairingProfiles returns a slice of labels from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetPairingProfiles(queryParameters map[string]string) (pairingProfiles []PairingProfile, api APIResponse, err error) {
	api, err = p.GetCollection("pairing_profiles", false, queryParameters, &pairingProfiles)
	if len(pairingProfiles) >= 500 {
		pairingProfiles = nil
		api, err = p.GetCollection("pairing_profiles", true, queryParameters, &pairingProfiles)
	}
	return pairingProfiles, api, err
}

// CreatePairingProfile creates a new pairing profile in the PCE.
func (p *PCE) CreatePairingProfile(pairingProfile PairingProfile) (createdPairingProfile PairingProfile, api APIResponse, err error) {
	api, err = p.Post("pairing_profiles", &pairingProfile, &createdPairingProfile)
	return createdPairingProfile, api, err
}

// CreatePairingKey creates a pairing key from a pairing profile.
func (p *PCE) CreatePairingKey(pairingProfile PairingProfile) (pairingKey PairingKey, api APIResponse, err error) {
	api, err = p.Post(strings.TrimPrefix(pairingProfile.Href, fmt.Sprintf("/orgs/%d/", p.Org))+"/pairing_key", &struct{}{}, &pairingKey)
	return pairingKey, api, err
}
