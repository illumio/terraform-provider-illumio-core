package models

/* Sample
{
  "href": "string",
  "created_at": "2019-08-24T14:15:22Z",
  "updated_at": "2019-08-24T14:15:22Z",
  "deleted_at": null,
  "created_by": {
    "href": "string"
  },
  "updated_by": {
    "href": "string"
  },
  "deleted_by": {
    "href": "string"
  },
  "update_type": "string",
  "name": "string",
  "description": null,
  "pce_fqdn": "string",
  "external_data_set": null,
  "external_data_reference": null,
  "labels": [
    {
      "href": "string"
    }
  ],
  "service_ports": [
    {
      "port": 0,
      "to_port": 0,
      "proto": 0
    }
  ],
  "service": {},
  "ip_overrides": [
    "string"
  ],
  "service_addresses": [
    {}
  ]
}
*/

type ServicePort struct {
	Port   *int `json:"port"`
	ToPort *int `json:"to_port"`
	Proto  int  `json:"proto"`
}

type ServiceAdd struct {
	IP          string `json:"ip"`
	Network     *Href  `json:"network"`
	Port        *int   `json:"port"`
	Fqdn        string `json:"fqdn"`
	Description string `json:"description"`
}

type VirtualService struct {
	Name                  string        `json:"name"`
	Description           string        `json:"description"`
	Labels                []Href        `json:"labels"`
	ServicePorts          []ServicePort `json:"service_ports"`
	Service               Href          `json:"service"`
	IPOverrides           []string      `json:"ip_overrides"`
	ApplyTo               string        `json:"apply_to"`
	ServiceAddresses      []ServiceAdd  `json:"service_addresses"`
	ExternalDataSet       string        `json:"external_data_set"`
	ExternalDataReference string        `json:"external_data_reference"`
}

// ToMap - Returns map for VirtualService model
func (vs *VirtualService) ToMap() (map[string]interface{}, error) {
	vsAttrmap := make(map[string]interface{})

	vsAttrmap["description"] = vs.Description

	if vs.Name != "" {
		vsAttrmap["name"] = vs.Name
	}
	if vs.ApplyTo != "" {
		vsAttrmap["apply_to"] = vs.ApplyTo
	}

	vsAttrmap["external_data_set"] = nil
	if vs.ExternalDataSet != "" {
		vsAttrmap["external_data_set"] = vs.ExternalDataSet
	}

	vsAttrmap["external_data_reference"] = nil
	if vs.ExternalDataReference != "" {
		vsAttrmap["external_data_reference"] = vs.ExternalDataReference
	}

	vsAttrmap["labels"] = GetHrefMaps(vs.Labels)

	// One of service or service_ports is required
	if vs.Service.Href != "" {
		vsAttrmap["service"], _ = vs.Service.ToMap()
	} else {
		servicePorts := []map[string]int{}
		for _, sp := range vs.ServicePorts {
			spi := map[string]int{"proto": sp.Proto}
			if sp.Port != nil {
				spi["port"] = *sp.Port
			}
			if sp.ToPort != nil {
				spi["to_port"] = *sp.ToPort
			}
			servicePorts = append(servicePorts, spi)
		}
		vsAttrmap["service_ports"] = servicePorts
	}
	if vs.IPOverrides != nil {
		vsAttrmap["ip_overrides"] = vs.IPOverrides
	} else {
		vsAttrmap["ip_overrides"] = []string{}
	}
	sas := []map[string]interface{}{}
	for _, i := range vs.ServiceAddresses {
		sa := map[string]interface{}{}
		if i.IP != "" {
			sa["ip"] = i.IP
		}
		if i.Network != nil {
			sa["network"], _ = i.Network.ToMap()
		}
		if i.Port != nil {
			sa["port"] = i.Port
		}
		if i.Fqdn != "" {
			sa["fqdn"] = i.Fqdn
		}
		if i.Description != "" {
			sa["description"] = i.Description
		}
		sas = append(sas, sa)
	}
	vsAttrmap["service_addresses"] = sas

	return vsAttrmap, nil
}
