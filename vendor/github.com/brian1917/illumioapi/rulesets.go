package illumioapi

import (
	"fmt"
	"strings"
)

// Actors - more info to follow
type Actors struct {
	Actors     string      `json:"actors,omitempty"`
	Label      *Label      `json:"label,omitempty"`
	LabelGroup *LabelGroup `json:"label_group,omitempty"`
	Workload   *Workload   `json:"workload,omitempty"`
}

// Consumers - more info to follow
type Consumers struct {
	Actors         string          `json:"actors,omitempty"`
	IPList         *IPList         `json:"ip_list,omitempty"`
	Label          *Label          `json:"label,omitempty"`
	LabelGroup     *LabelGroup     `json:"label_group,omitempty"`
	VirtualService *VirtualService `json:"virtual_service,omitempty"`
	Workload       *Workload       `json:"workload,omitempty"`
}

// IngressServices - more info to follow
type IngressServices struct {
	Port     *int    `json:"port,omitempty"`
	Protocol *int    `json:"proto,omitempty"`
	ToPort   *int    `json:"to_port,omitempty"`
	Href     *string `json:"href,omitempty"`
}

// IPTablesRules - more info to follow
type IPTablesRules struct {
	Actors      []*Actors     `json:"actors"`
	Description string        `json:"description,omitempty"`
	Enabled     bool          `json:"enabled"`
	Href        string        `json:"href"`
	IPVersion   string        `json:"ip_version"`
	Statements  []*Statements `json:"statements"`
}

// Providers - more info to follow
type Providers struct {
	Actors         string          `json:"actors,omitempty"`
	IPList         *IPList         `json:"ip_list,omitempty"`
	Label          *Label          `json:"label,omitempty"`
	LabelGroup     *LabelGroup     `json:"label_group,omitempty"`
	VirtualServer  *VirtualServer  `json:"virtual_server,omitempty"`
	VirtualService *VirtualService `json:"virtual_service,omitempty"`
	Workload       *Workload       `json:"workload,omitempty"`
}

// ResolveLabelsAs - more info to follow
type ResolveLabelsAs struct {
	Consumers []string `json:"consumers"`
	Providers []string `json:"providers"`
}

// RuleSet - more info to follow
type RuleSet struct {
	CreatedAt             string           `json:"created_at,omitempty"`
	CreatedBy             *CreatedBy       `json:"created_by,omitempty"`
	DeletedAt             string           `json:"deleted_at,omitempty"`
	DeletedBy             *DeletedBy       `json:"deleted_by,omitempty"`
	Description           string           `json:"description,omitempty"`
	Enabled               *bool            `json:"enabled,omitempty"`
	ExternalDataReference string           `json:"external_data_reference,omitempty"`
	ExternalDataSet       string           `json:"external_data_set,omitempty"`
	Href                  string           `json:"href,omitempty"`
	IPTablesRules         []*IPTablesRules `json:"ip_tables_rules,omitempty"`
	Name                  string           `json:"name,omitempty"`
	Rules                 []*Rule          `json:"rules,omitempty"`
	Scopes                [][]*Scopes      `json:"scopes,omitempty"`
	UpdateType            string           `json:"update_type,omitempty"`
	UpdatedAt             string           `json:"updated_at,omitempty"`
	UpdatedBy             *UpdatedBy       `json:"updated_by,omitempty"`
}

// Rule - more info to follow
type Rule struct {
	CreatedAt                   string                         `json:"created_at,omitempty"`
	CreatedBy                   *CreatedBy                     `json:"created_by,omitempty"`
	DeletedAt                   string                         `json:"deleted_at,omitempty"`
	DeletedBy                   *DeletedBy                     `json:"deleted_by,omitempty"`
	Consumers                   []*Consumers                   `json:"consumers,omitempty"`
	ConsumingSecurityPrincipals []*ConsumingSecurityPrincipals `json:"consuming_security_principals,omitempty"`
	Description                 string                         `json:"description,omitempty"`
	Enabled                     *bool                          `json:"enabled,omitempty"`
	ExternalDataReference       string                         `json:"external_data_reference,omitempty"`
	ExternalDataSet             string                         `json:"external_data_set,omitempty"`
	Href                        string                         `json:"href,omitempty"`
	IngressServices             *[]*IngressServices            `json:"ingress_services,omitempty"`
	Providers                   []*Providers                   `json:"providers,omitempty"`
	ResolveLabelsAs             *ResolveLabelsAs               `json:"resolve_labels_as,omitempty"`
	SecConnect                  *bool                          `json:"sec_connect,omitempty"`
	Stateless                   *bool                          `json:"stateless,omitempty"`
	MachineAuth                 *bool                          `json:"machine_auth,omitempty"`
	UnscopedConsumers           *bool                          `json:"unscoped_consumers,omitempty"`
	UpdateType                  string                         `json:"update_type,omitempty"`
	UpdatedAt                   string                         `json:"updated_at,omitempty"`
	UpdatedBy                   *UpdatedBy                     `json:"updated_by,omitempty"`
}

// Scopes - more info to follow
type Scopes struct {
	Label      *Label      `json:"label,omitempty"`
	LabelGroup *LabelGroup `json:"label_group,omitempty"`
}

// Statements are part of a custom IPTables rule
type Statements struct {
	ChainName  string `json:"chain_name"`
	Parameters string `json:"parameters"`
	TableName  string `json:"table_name"`
}

// GetRulesets returns a slice of labels from the PCE. pStatus must be "draft" or "active".
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetRulesets(queryParameters map[string]string, pStatus string) (ruleSets []RuleSet, api APIResponse, err error) {
	api, err = p.GetCollection("sec_policy/"+pStatus+"/rule_sets", false, queryParameters, &ruleSets)
	if len(ruleSets) >= 500 {
		ruleSets = nil
		api, err = p.GetCollection("sec_policy/"+pStatus+"/rule_sets", true, queryParameters, &ruleSets)
	}
	p.RuleSets = make(map[string]RuleSet)
	for _, rs := range ruleSets {
		p.RuleSets[rs.Href] = rs
		p.RuleSets[rs.Name] = rs
	}
	return ruleSets, api, err
}

// CreateRuleSet creates a new ruleset in the PCE.
func (p *PCE) CreateRuleset(rs RuleSet) (createdRS RuleSet, api APIResponse, err error) {
	api, err = p.Post("sec_policy/draft/rule_sets", &rs, &createdRS)
	return createdRS, api, err
}

// CreateRule creates a new rule List in the PCE.
func (p *PCE) CreateRule(rulesetHref string, rule Rule) (createdRule Rule, api APIResponse, err error) {
	api, err = p.Post(strings.TrimPrefix(rulesetHref, fmt.Sprintf("/orgs/%d/", p.Org))+"/sec_rules", &rule, &createdRule)
	return createdRule, api, err
}

// UpdateRuleset updates an existing ruleset in the PCE.
// The provided ruleset must include an Href.
// Properties that cannot be included in the PUT method will be ignored.
func (p *PCE) UpdateRuleset(ruleset RuleSet) (APIResponse, error) {
	ruleset.CreatedAt = ""
	ruleset.CreatedBy = nil
	ruleset.UpdateType = ""
	ruleset.UpdatedAt = ""
	ruleset.UpdatedBy = nil
	ruleset.DeletedAt = ""
	ruleset.DeletedBy = nil
	ruleset.Rules = nil

	return p.Put(&ruleset)
}

// UpdateRule updates an existing rule in the PCE.
// The provided rule must include an Href.
// Properties that cannot be included in the PUT method will be ignored.
func (p *PCE) UpdateRule(rule Rule) (APIResponse, error) {
	rule.CreatedAt = ""
	rule.CreatedBy = nil
	rule.DeletedAt = ""
	rule.DeletedBy = nil
	rule.UpdatedAt = ""
	rule.UpdatedBy = nil

	return p.Put(&rule)
}

// GetRuleByHref returns the rule with a specific href
func (p *PCE) GetRuleByHref(href string) (rule Rule, api APIResponse, err error) {
	api, err = p.GetHref(href, &rule)
	return rule, api, err
}

// GetRulesetByHref returns the rule with a specific href
func (p *PCE) GetRulesetByHref(href string) (ruleset RuleSet, api APIResponse, err error) {
	api, err = p.GetHref(href, &ruleset)
	return ruleset, api, err
}

// GetRulesetHref returns the href of a ruleset based on the rule's href
func (r *Rule) GetRulesetHref() string {
	x := strings.Split(r.Href, "/")
	x = x[:len(x)-2]
	return strings.Join(x, "/")
}
