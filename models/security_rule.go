// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

/* Sample
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
  "unscoped_consumers": true
}
*/

type SecurityRuleResolveLabelAs struct {
	Providers []string `json:"providers"`
	Consumers []string `json:"consumers"`
}

type SecurityRuleProvider struct {
	Actors         string `json:"actors"`
	Label          *Href  `json:"label"`
	LabelGroup     *Href  `json:"label_group"`
	Workload       *Href  `json:"workload"`
	VirtualService *Href  `json:"virtual_service"`
	VirtualServer  *Href  `json:"virtual_server"`
	IPList         *Href  `json:"ip_list"`
}

type SecurityRuleConsumer struct {
	Actors         string `json:"actors"`
	Label          *Href  `json:"label"`
	LabelGroup     *Href  `json:"label_group"`
	Workload       *Href  `json:"workload"`
	VirtualService *Href  `json:"virtual_service"`
	IPList         *Href  `json:"ip_list"`
}

type SecurityRule struct {
	Enabled               bool                        `json:"enabled"`
	Description           string                      `json:"description"`
	IngressServices       []map[string]interface{}    `json:"ingress_services"`
	ResolveLabelsAs       *SecurityRuleResolveLabelAs `json:"resolve_labels_as"`
	SecConnect            bool                        `json:"sec_connect"`
	Stateless             bool                        `json:"stateless"`
	MachineAuth           bool                        `json:"machine_auth"`
	Providers             []*SecurityRuleProvider     `json:"providers"`
	Consumers             []*SecurityRuleConsumer     `json:"consumers"`
	UnscopedConsumers     bool                        `json:"unscoped_consumers"`
	ExternalDataSet       string                      `json:"external_data_set"`
	ExternalDataReference string                      `json:"external_data_reference"`
}

// Checks if Security Rule ResolveLabelAs.providers list containes only "virtual_services"
func (r *SecurityRuleResolveLabelAs) ProviderIsVirtualService() bool {
	if r == nil {
		return false
	}

	if len(r.Providers) == 1 && r.Providers[0] == "virtual_services" {
		return true
	}

	return false
}

// Checks if conflicting values are set in SecurityRule
func (s *SecurityRule) HasConflicts() bool {
	if s == nil {
		return false
	}

	if (s.SecConnect && s.Stateless) || (s.SecConnect && s.MachineAuth) || (s.Stateless && s.MachineAuth) {
		return true
	}
	return false
}

// Validates SecurityRuleConsumer
//
// Checks if SecurityRuleConsumer has only one actor set
func (o *SecurityRuleConsumer) HasOneActor() bool {
	if o == nil {
		return false
	}

	count := 0
	if o.Actors != "" {
		count++
	}
	if o.IPList.Href != "" {
		count++
	}
	if o.Label.Href != "" {
		count++
	}
	if o.LabelGroup.Href != "" {
		count++
	}
	if o.VirtualService.Href != "" {
		count++
	}
	if o.Workload.Href != "" {
		count++
	}

	return count == 1
}

// Validates SecurityRuleProvider
//
// Checks if SecurityRuleProvider has only one actor set
func (o *SecurityRuleProvider) HasOneActor() bool {
	if o == nil {
		return false
	}

	count := 0
	if o.Actors != "" {
		count++
	}
	if o.IPList.Href != "" {
		count++
	}
	if o.Label.Href != "" {
		count++
	}
	if o.LabelGroup.Href != "" {
		count++
	}
	if o.VirtualService.Href != "" {
		count++
	}
	if o.VirtualServer.Href != "" {
		count++
	}
	if o.Workload.Href != "" {
		count++
	}

	return count == 1
}

// ToMap - Returns map for SecurityRule model
func (sr *SecurityRule) ToMap() (map[string]interface{}, error) {
	srMap := make(map[string]interface{})

	srMap["description"] = sr.Description
	srMap["enabled"] = sr.Enabled
	srMap["sec_connect"] = sr.SecConnect
	srMap["stateless"] = sr.Stateless
	srMap["machine_auth"] = sr.MachineAuth
	srMap["unscoped_consumers"] = sr.UnscopedConsumers

	srMap["external_data_set"] = nil
	if sr.ExternalDataSet != "" {
		srMap["external_data_set"] = sr.ExternalDataSet
	}

	srMap["external_data_reference"] = nil
	if sr.ExternalDataReference != "" {
		srMap["external_data_reference"] = sr.ExternalDataReference
	}

	cons := []map[string]interface{}{}
	for _, consumer := range sr.Consumers {
		con := make(map[string]interface{})

		if consumer.Actors != "" {
			con["actors"] = consumer.Actors
		}

		if consumer.IPList.Href != "" {
			con["ip_list"], _ = consumer.IPList.ToMap()
		}

		if consumer.Label.Href != "" {
			con["label"], _ = consumer.Label.ToMap()
		}

		if consumer.LabelGroup.Href != "" {
			con["label_group"], _ = consumer.LabelGroup.ToMap()
		}

		if consumer.VirtualService.Href != "" {
			con["virtual_service"], _ = consumer.VirtualService.ToMap()
		}

		if consumer.Workload.Href != "" {
			con["workload"], _ = consumer.Workload.ToMap()
		}

		cons = append(cons, con)
	}
	srMap["consumers"] = cons

	provs := []map[string]interface{}{}
	for _, provider := range sr.Providers {
		prov := make(map[string]interface{})

		if provider.Actors != "" {
			prov["actors"] = provider.Actors
		}

		if provider.IPList.Href != "" {
			prov["ip_list"], _ = provider.IPList.ToMap()
		}

		if provider.Label.Href != "" {
			prov["label"], _ = provider.Label.ToMap()
		}

		if provider.LabelGroup.Href != "" {
			prov["label_group"], _ = provider.LabelGroup.ToMap()
		}

		if provider.VirtualService.Href != "" {
			prov["virtual_service"], _ = provider.VirtualService.ToMap()
		}

		if provider.VirtualServer.Href != "" {
			prov["virtual_server"], _ = provider.VirtualServer.ToMap()
		}

		if provider.Workload.Href != "" {
			prov["workload"], _ = provider.Workload.ToMap()
		}
		provs = append(provs, prov)
	}
	srMap["providers"] = provs

	if sr.ResolveLabelsAs != nil {
		rsm := make(map[string][]string)
		rsm["providers"] = sr.ResolveLabelsAs.Providers
		rsm["consumers"] = sr.ResolveLabelsAs.Consumers
		srMap["resolve_labels_as"] = rsm
	}

	srMap["ingress_services"] = sr.IngressServices

	return srMap, nil
}
