package models

//Sample
/*
	{
  "name": "string",
  "providers": [
    {
      "actors": "ams",
      "label": {
        "href": "string"
      },
      "label_group": {
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
      "ip_list": {
        "href": "string"
      }
    }
  ],
  "ingress_services": [
    {
      "href": "string"
    }
  ]
}
*/

type EnforcementBoundaryProviderConsumer struct {
	Actors     string `json:"actors"`
	Label      *Href  `json:"label"`
	LabelGroup *Href  `json:"label_group"`
	IPList     *Href  `json:"ip_list"`
}

type EnforcementBoundary struct {
	Name            string                                 `json:"name"`
	Providers       []*EnforcementBoundaryProviderConsumer `json:"providers"`
	Consumers       []*EnforcementBoundaryProviderConsumer `json:"consumers"`
	IngressServices []map[string]interface{}               `json:"ingress_services"`
}

func (o *EnforcementBoundaryProviderConsumer) HasOneActor() bool {
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

	return count == 1
}

func (en *EnforcementBoundary) ToMap() (map[string]interface{}, error) {
	enMap := make(map[string]interface{})

	enMap["name"] = en.Name

	cons := []map[string]interface{}{}
	for _, consumer := range en.Consumers {
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

		cons = append(cons, con)
	}
	enMap["consumers"] = cons

	provs := []map[string]interface{}{}
	for _, provider := range en.Providers {
		con := make(map[string]interface{})

		if provider.Actors != "" {
			con["actors"] = provider.Actors
		}

		if provider.IPList.Href != "" {
			con["ip_list"], _ = provider.IPList.ToMap()
		}

		if provider.Label.Href != "" {
			con["label"], _ = provider.Label.ToMap()
		}

		if provider.LabelGroup.Href != "" {
			con["label_group"], _ = provider.LabelGroup.ToMap()
		}

		provs = append(provs, con)
	}
	enMap["providers"] = provs

	enMap["ingress_services"] = en.IngressServices

	return enMap, nil
}
