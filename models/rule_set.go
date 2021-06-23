// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

/* Schema
{
  "name": "string",
  "description": "string",
  "external_data_set": null,
  "external_data_reference": null,
  "enabled": true,
  "scopes": [
    [
      {
        "label": {
          "href": "string"
        },
        "label_group": {
          "href": "string"
        }
      }
    ]
  ],
  "rules": [
    {
      "enabled": true,
      "description": "string",
      "external_data_set": null,
      "external_data_reference": null,
      "ingress_services": [
        {
          "href": "string"
        }
      ],
      "resolve_labels_as": {
        "providers": [
          "workloads"
        ],
        "consumers": [
          "workloads"
        ]
      },
      "sec_connect": true,
      "stateless": true,
      "machine_auth": true,
      "providers": [
        {
          "actors": "ams",
          "label": {
            "href": "string"
          },
          "label_group": {
            "href": "string"
          },
          "workload": {
            "href": "string"
          },
          "virtual_service": {
            "href": "string"
          },
          "virtual_server": {
            "href": "string"
          },
          "ip_list": {
            "href": "string"
          }
        }
      ],
      "consumers": [
        {
          "actors": "ams",
          "label": {
            "href": "string"
          },
          "label_group": {
            "href": "string"
          },
          "workload": {
            "href": "string"
          },
          "virtual_service": {
            "href": "string"
          },
          "ip_list": {
            "href": "string"
          }
        }
      ],
      "consuming_security_principals": [
        {
          "href": "string"
        }
      ],
      "unscoped_consumers": true
    }
  ],
  "ip_tables_rules": [
    {
      "enabled": true,
      "description": "string",
      "statements": [
        {
          "table_name": "nat",
          "chain_name": "PREROUTING",
          "parameters": "string"
        }
      ],
      "actors": [
        {
          "actors": "string",
          "label": {
            "href": "string"
          },
          "label_group": {
            "href": "string"
          },
          "workload": {
            "href": "string"
          }
        }
      ],
      "ip_version": "4"
    }
  ]
}
*/

type RuleSet struct {
	Name                  string            `json:"name"`
	Description           string            `json:"description"`
	ExternalDataSet       string            `json:"external_data_set"`
	ExternalDataReference string            `json:"external_data_reference"`
	Enabled               bool              `json:"enabled"`
	Scopes                [][]*RuleSetScope `json:"scopes"`
	// Rules                 []*SecurityRule        `json:"rules"`
	IPTablesRules []*RuleSetIPTablesRule `json:"ip_tables_rules"`
}

type RuleSetScope struct {
	Label      *Href `json:"label"`
	LabelGroup *Href `json:"label_group"`
}

type RuleSetIPTablesRulesStatement struct {
	TableName  string `json:"table_name"`
	ChainName  string `json:"chain_name"`
	Parameters string `json:"parameters"`
}

type RuleSetIPTablesRulesActor struct {
	Actors     string `json:"actors"`
	Label      *Href  `json:"label"`
	LabelGroup *Href  `json:"label_group"`
	Workload   *Href  `json:"workload"`
}

type RuleSetIPTablesRule struct {
	Enabled     bool                             `json:"enabled"`
	Description string                           `json:"description"`
	Statements  []*RuleSetIPTablesRulesStatement `json:"statements"`
	Actors      []*RuleSetIPTablesRulesActor     `json:"actors"`
	IPVersion   string                           `json:"ip_version"`
}

// Checks if conflicting paramerters in RuleSetScope are set
func (s *RuleSetScope) HasInnerConflicts() bool {
	if s.Label.Href != "" && s.LabelGroup.Href != "" {
		return false
	}

	return true
}

// ToMap - Returns map for RuleSet model
func (r *RuleSet) ToMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})

	m["enabled"] = r.Enabled
	m["description"] = r.Description

	if r.Name != "" {
		m["name"] = r.Name
	}

	m["external_data_set"] = nil
	if r.ExternalDataSet != "" {
		m["external_data_set"] = r.ExternalDataSet
	}

	m["external_data_reference"] = nil
	if r.ExternalDataReference != "" {
		m["external_data_reference"] = r.ExternalDataReference
	}

	sps := [][]map[string]map[string]interface{}{}
	for _, scope := range r.Scopes {
		sp := []map[string]map[string]interface{}{}
		for _, val := range scope {
			o := make(map[string]map[string]interface{})
			if val.Label != nil && val.Label.Href != "" {
				o["label"], _ = val.Label.ToMap()
			}
			if val.LabelGroup != nil && val.LabelGroup.Href != "" {
				o["label_group"], _ = val.LabelGroup.ToMap()
			}
			sp = append(sp, o)
		}
		sps = append(sps, sp)
	}
	m["scopes"] = sps

	// if r.Rules != nil {
	// 	rls := []map[string]interface{}{}
	// 	for _, rule := range r.Rules {
	// 		rl, _ := rule.ToMap()

	// 		if rl["external_data_set"] == nil {
	// 			delete(rl, "external_data_set")
	// 		}

	// 		if rl["external_data_reference"] == nil {
	// 			delete(rl, "external_data_reference")
	// 		}

	// 		rls = append(rls, rl)
	// 	}

	// 	m["rules"] = rls
	// }

	if r.IPTablesRules != nil {
		iptrs := []map[string]interface{}{}

		for _, ipTableRule := range r.IPTablesRules {
			iptr := make(map[string]interface{})

			iptr["enabled"] = ipTableRule.Enabled

			if ipTableRule.Description != "" {
				iptr["description"] = ipTableRule.Description
			}

			ss := []map[string]interface{}{}
			for _, stat := range ipTableRule.Statements {
				s := make(map[string]interface{})
				if stat.ChainName != "" {
					s["chain_name"] = stat.ChainName
				}
				if stat.TableName != "" {
					s["table_name"] = stat.TableName
				}
				if stat.Parameters != "" {
					s["parameters"] = stat.Parameters
				}

				ss = append(ss, s)
			}

			iptr["statements"] = ss

			acs := []map[string]interface{}{}

			for _, actor := range ipTableRule.Actors {
				ac := make(map[string]interface{})
				if actor.Actors != "" {
					ac["actors"] = actor.Actors
				}
				if actor.Label.Href != "" {
					ac["label"], _ = actor.Label.ToMap()
				}
				if actor.LabelGroup.Href != "" {
					ac["label_group"], _ = actor.LabelGroup.ToMap()
				}
				if actor.Workload.Href != "" {
					ac["workload"], _ = actor.Workload.ToMap()
				}

				acs = append(acs, ac)
			}

			iptr["actors"] = acs

			if ipTableRule.IPVersion != "" {
				iptr["ip_version"] = ipTableRule.IPVersion
			}

			iptrs = append(iptrs, iptr)
		}

		m["ip_tables_rules"] = iptrs
	}
	return m, nil
}
