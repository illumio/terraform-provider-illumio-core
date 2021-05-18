// Copyright 2021 Illumio, Inc. All Rights Reserved. 

package models

/* Sample
{
  "href": "string",
  "name": "string",
  "description": "string",
  "external_data_set": null,
  "external_data_reference": null,
  "ip_ranges": [
    {
      "description": "string",
      "from_ip": "string",
      "to_ip": "string",
      "exclusion": true
    }
  ],
  "fqdns": [
    {
      "fqdn": "string",
      "description": "string"
    }
  ],
  "created_at": "2021-03-02T02:37:59Z",
  "updated_at": "2021-03-02T02:37:59Z",
  "deleted_at": "2021-03-02T02:37:59Z",
  "created_by": {
    "href": "string"
  },
  "updated_by": {
    "href": "string"
  },
  "deleted_by": {
    "href": "string"
  }
}
*/

type IPRange struct {
	Description string `json:"description"`
	FromIP      string `json:"from_ip"`
	ToIP        string `json:"to_ip"`
	Exclusion   bool   `json:"exclusion"`
}

type FQDN struct {
	FQDN        string `json:"fqdn"`
	Description string `json:"description"`
}

type IPList struct {
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	ExternalDataSet       string    `json:"external_data_set"`
	ExternalDataReference string    `json:"external_data_reference"`
	IPRanges              []IPRange `json:"ip_ranges"`
	FQDNs                 []FQDN    `json:"fqdns"`
}

// ToMap - Returns map for IP List model
func (i *IPList) ToMap() (map[string]interface{}, error) {
	ipListAttrMap := make(map[string]interface{})

	if i.Name != "" {
		ipListAttrMap["name"] = i.Name
	}

	ipListAttrMap["description"] = i.Description

	ipListAttrMap["external_data_set"] = nil
	if i.ExternalDataSet != "" {
		ipListAttrMap["external_data_set"] = i.ExternalDataSet
	}

	ipListAttrMap["external_data_reference"] = nil
	if i.ExternalDataReference != "" {
		ipListAttrMap["external_data_reference"] = i.ExternalDataReference
	}

	iprMapArr := []map[string]interface{}{}
	for _, o := range i.IPRanges {
		m := make(map[string]interface{})
		if o.FromIP != "" {
			m["from_ip"] = o.FromIP
		}

		if o.ToIP != "" {
			m["to_ip"] = o.ToIP
		}

		if o.Description != "" {
			m["description"] = o.Description
		}

		if o.Exclusion {
			m["exclusion"] = o.Exclusion
		}

		iprMapArr = append(iprMapArr, m)
	}
	ipListAttrMap["ip_ranges"] = iprMapArr

	fqdnMapArr := []map[string]string{}
	for _, o := range i.FQDNs {
		m := make(map[string]string)

		m["fqdn"] = o.FQDN

		m["description"] = o.Description

		fqdnMapArr = append(fqdnMapArr, m)
	}
	ipListAttrMap["fqdns"] = fqdnMapArr

	return ipListAttrMap, nil
}
