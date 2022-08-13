package illumioapi

// ConsumingSecurityPrincipals are AD user groups
type ConsumingSecurityPrincipals struct {
	Deleted       bool   `json:"deleted,omitempty"`
	Href          string `json:"href,omitempty"`
	Name          string `json:"name,omitempty"`
	SID           string `json:"sid,omitempty"`
	UsedByRuleSet bool   `json:"used_by_ruleset,omitempty"`
}

// GetADUserGroups returns a slice of AD user groups from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetADUserGroups(queryParameters map[string]string) (groups []ConsumingSecurityPrincipals, api APIResponse, err error) {
	api, err = p.GetCollection("security_principals", false, queryParameters, &groups)
	if len(groups) >= 500 {
		groups = nil
		api, err = p.GetCollection("security_principals", true, queryParameters, &groups)
	}
	return groups, api, err
}

// CreateADUserGroup creates a user group policy object in the PCE
func (p *PCE) CreateADUserGroup(group ConsumingSecurityPrincipals) (createdGroup ConsumingSecurityPrincipals, api APIResponse, err error) {
	api, err = p.Post("security_principals", &group, &createdGroup)
	return createdGroup, api, err
}
