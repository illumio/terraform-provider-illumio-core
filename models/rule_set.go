// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type RuleSet struct {
	Name                  *string                 `json:"name,omitempty"`
	Description           *string                 `json:"description,omitempty"`
	Enabled               *bool                   `json:"enabled,omitempty"`
	Scopes                [][]*RuleSetScope       `json:"scopes"`
	IPTablesRules         *[]*RuleSetIPTablesRule `json:"ip_tables_rules,omitempty"`
	ExternalDataSet       string                  `json:"external_data_set,omitempty"`
	ExternalDataReference string                  `json:"external_data_reference,omitempty"`
}

type RuleSetScope struct {
	Exclusion  *bool                       `json:"exclusion,omitempty"`
	Label      *LabelOptionalKeyValue      `json:"label,omitempty"`
	LabelGroup *LabelGroupOptionalKeyValue `json:"label_group,omitempty"`
}

type RuleSetIPTablesRulesStatement struct {
	TableName  string `json:"table_name,omitempty"`
	ChainName  string `json:"chain_name,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

type RuleSetIPTablesRulesActor struct {
	Actors     string                 `json:"actors,omitempty"`
	Label      *LabelOptionalKeyValue `json:"label,omitempty"`
	LabelGroup *Href                  `json:"label_group,omitempty"`
	Workload   *Href                  `json:"workload,omitempty"`
}

type RuleSetIPTablesRule struct {
	Enabled     *bool                             `json:"enabled,omitempty"`
	Description *string                           `json:"description,omitempty"`
	Statements  *[]*RuleSetIPTablesRulesStatement `json:"statements,omitempty"`
	Actors      *[]*RuleSetIPTablesRulesActor     `json:"actors,omitempty"`
	IPVersion   string                            `json:"ip_version,omitempty"`
}

// ToMap - Returns map for RuleSet model
func (r *RuleSet) ToMap() (map[string]interface{}, error) {
	return toMap(r)
}
