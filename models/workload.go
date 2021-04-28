package models

/* Sample
{
  "name": "string",
  "description": "string",
  "external_data_set": null,
  "external_data_reference": null,
  "hostname": "string",
  "service_principal_name": null,
  "public_ip": "string",
  "interfaces": [
    {
      "name": "string",
      "link_state": "up",
      "address": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
      "cidr_block": 0,
      "default_gateway_address": "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
      "friendly_name": "string"
    }
  ],
  "service_provider": "string",
  "data_center": "string",
  "data_center_zone": "string",
  "os_id": "string",
  "os_detail": "string",
  "online": true,
  "labels": [
    {
      "href": "string"
    }
  ],
  "agent": {
    "config": {
      "mode": "idle",
      "log_traffic": true
    }
  },
  "enforcement_mode": "idle"
}
*/

type Workload struct {
	Name                                  string              `json:"name"`
	Description                           string              `json:"description"`
	ExternalDataSet                       string              `json:"external_data_set"`
	ExternalDataReference                 string              `json:"external_data_reference"`
	Hostname                              string              `json:"hostname"`
	ServicePrincipalName                  string              `json:"service_principal_name"`
	PublicIP                              string              `json:"public_ip"`
	Interfaces                            []WorkloadInterface `json:"interfaces"`
	AgentToPceCertificateAuthenticationID string              `json:"agent_to_pce_certificate_authentication_id"`
	DistinguishedName                     string              `json:"distinguished_name"`
	ServiceProvider                       string              `json:"service_provider"`
	DataCenter                            string              `json:"data_center"`
	DataCenterZone                        string              `json:"data_center_zone"`
	OsID                                  string              `json:"os_id"`
	OsDetail                              string              `json:"os_detail"`
	Online                                bool                `json:"online"`
	Labels                                []Href              `json:"labels"`
	EnforcementMode                       string              `json:"enforcement_mode"`
}

// ToMap - Returns map for Workload model
func (w *Workload) ToMap() (map[string]interface{}, error) {
	workloadAttrMap := make(map[string]interface{})

	workloadAttrMap["name"] = w.Name
	workloadAttrMap["hostname"] = w.Hostname
	workloadAttrMap["description"] = w.Description
	workloadAttrMap["distinguished_name"] = w.DistinguishedName
	workloadAttrMap["service_provider"] = w.ServiceProvider
	workloadAttrMap["data_center"] = w.DataCenter
	workloadAttrMap["data_center_zone"] = w.DataCenterZone
	workloadAttrMap["os_id"] = w.OsID
	workloadAttrMap["os_detail"] = w.OsDetail

	workloadAttrMap["online"] = w.Online
	workloadAttrMap["labels"] = GetHrefMaps(w.Labels)
	workloadAttrMap["agent_to_pce_certificate_authentication_id"] = nil
	if w.AgentToPceCertificateAuthenticationID != "" {
		workloadAttrMap["agent_to_pce_certificate_authentication_id"] = w.AgentToPceCertificateAuthenticationID
	}
	workloadAttrMap["service_principal_name"] = nil
	if w.ServicePrincipalName != "" {
		workloadAttrMap["service_principal_name"] = w.ServicePrincipalName
	}
	if w.PublicIP != "" {
		workloadAttrMap["public_ip"] = w.PublicIP
	}
	if w.EnforcementMode != "" {
		workloadAttrMap["enforcement_mode"] = w.EnforcementMode
	}

	workloadAttrMap["external_data_reference"] = nil
	if w.ExternalDataSet != "" {
		workloadAttrMap["external_data_reference"] = w.ExternalDataReference
	}

	workloadAttrMap["external_data_set"] = nil
	if w.ExternalDataReference != "" {
		workloadAttrMap["external_data_set"] = w.ExternalDataSet
	}
	wMapArr := []map[string]interface{}{}
	for _, o := range w.Interfaces {
		m := make(map[string]interface{})
		if o.Name != "" {
			m["name"] = o.Name
		}
		if o.LinkState != "" {
			m["link_state"] = o.LinkState
		}
		if o.Address != "" {
			m["address"] = o.Address
		}
		m["cidr_block"] = o.CidrBlock
		if o.DefaultGatewayAddress != "" {
			m["default_gateway_address"] = o.DefaultGatewayAddress
		}
		if o.FriendlyName != "" {
			m["friendly_name"] = o.FriendlyName

		}
		wMapArr = append(wMapArr, m)
	}
	workloadAttrMap["interfaces"] = wMapArr

	return workloadAttrMap, nil
}
