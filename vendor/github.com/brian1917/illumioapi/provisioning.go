package illumioapi

import (
	"strings"
)

// ChangeSubset Hash of pending hrefs, organized by model
type ChangeSubset struct {
	FirewallSettings      []*FirewallSettings      `json:"firewall_settings,omitempty"`
	IPLists               []*IPList                `json:"ip_lists,omitempty"`
	LabelGroups           []*LabelGroup            `json:"label_groups,omitempty"`
	RuleSets              []*RuleSet               `json:"rule_sets,omitempty"`
	SecureConnectGateways []*SecureConnectGateways `json:"secure_connect_gateways,omitempty"`
	Services              []*Service               `json:"services,omitempty"`
	VirtualServers        []*VirtualServer         `json:"virtual_servers,omitempty"`
	VirtualServices       []*VirtualService        `json:"virtual_services,omitempty"`
	EnforcementBoundaries []*EnforcementBoundary   `json:"enforcement_boundaries,omitempty"`
}

// FirewallSettings are a provisionable object
type FirewallSettings struct {
	Href string `json:"href"`
}

// Provision is sent to the PCE to provision policy objects
type Provision struct {
	ChangeSubset      *ChangeSubset `json:"change_subset,omitempty"`
	UpdateDescription string        `json:"update_description,omitempty"`
}

// SecureConnectGateways represent SecureConnectGateways in provisioning
type SecureConnectGateways struct {
	Href string `json:"href"`
}

// VirtualServers reresent virtual servers in provisioning
type VirtualServers struct {
	Href string `json:"href"`
}

// ProvisionCS provisions a ChangeSubset
func (p *PCE) ProvisionCS(cs ChangeSubset, comment string) (api APIResponse, err error) {
	provision := Provision{ChangeSubset: &cs, UpdateDescription: comment}
	if err != nil {
		return APIResponse{}, err
	}
	api, err = p.Post("/sec_policy", &provision, &struct{}{})
	return api, err
}

// ProvisionHref provisions a slice of HREFs
func (p *PCE) ProvisionHref(hrefs []string, comment string) (APIResponse, error) {

	// Build our variables
	var ipl []*IPList
	var services []*Service
	var ruleSets []*RuleSet
	var labelGroups []*LabelGroup
	var virtualServices []*VirtualService
	var virtualServers []*VirtualServer
	var fs []*FirewallSettings
	var secureConnectGateways []*SecureConnectGateways
	var enforcementBoundaries []*EnforcementBoundary

	// Process our list of HREFs
	for _, h := range hrefs {

		if strings.Contains(h, "/ip_lists/") {
			ipl = append(ipl, &IPList{Href: h})
		}
		// Services
		if strings.Contains(h, "/services/") {
			services = append(services, &Service{Href: h})
		}
		// Rule Sets
		if strings.Contains(h, "/rule_sets/") {
			ruleSets = append(ruleSets, &RuleSet{Href: h})
		}
		// Label Groups
		if strings.Contains(h, "/label_groups/") {
			labelGroups = append(labelGroups, &LabelGroup{Href: h})
		}
		// Virtual Services
		if strings.Contains(h, "/virtual_services/") {
			virtualServices = append(virtualServices, &VirtualService{Href: h})
		}
		// Virtual Servers
		if strings.Contains(h, "/virtual_servers/") {
			virtualServers = append(virtualServers, &VirtualServer{Href: h})
		}
		// Firewall Settings
		if strings.Contains(h, "/firewall_settings/") {
			fs = append(fs, &FirewallSettings{Href: h})
		}
		// SecureConnect Gateway
		if strings.Contains(h, "/secure_connect_gateways/") {
			secureConnectGateways = append(secureConnectGateways, &SecureConnectGateways{Href: h})
		}
		// Enforcement Boundaries
		if strings.Contains(h, "/enforcement_boundaries/") {
			enforcementBoundaries = append(enforcementBoundaries, &EnforcementBoundary{Href: h})
		}

	}
	// Build the Provision
	api, err := p.ProvisionCS(ChangeSubset{
		FirewallSettings:      fs,
		IPLists:               ipl,
		LabelGroups:           labelGroups,
		RuleSets:              ruleSets,
		SecureConnectGateways: secureConnectGateways,
		Services:              services,
		VirtualServers:        virtualServers,
		VirtualServices:       virtualServices,
		EnforcementBoundaries: enforcementBoundaries,
	}, comment)
	if err != nil {
		return api, err
	}

	return api, nil
}

// GetPending returns a slice of pending changes from the PCE.
func (p *PCE) GetPendingChanges() (cs ChangeSubset, api APIResponse, err error) {
	api, err = p.GetCollection("sec_policy/pending", false, nil, &cs)
	return cs, api, err
}
