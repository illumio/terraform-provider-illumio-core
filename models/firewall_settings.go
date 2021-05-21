// Copyright 2021 Illumio, Inc. All Rights Reserved.

package models

type FirewallCoexistenceScope struct {
}

type FirewallCoexistenceObj struct {
	IllumioPrimary bool   `json:"illumio_primary"`
	Scope          []Href `json:"scope"`
	WorkloadMode   string `json:"workload_mode"`
}

type FirewallCoexistence []FirewallCoexistenceObj

func (f FirewallCoexistence) ToList() []interface{} {
	fcs := []interface{}{}
	for _, fc := range f {
		fco := map[string]interface{}{}
		fco["illumio_primary"] = fc.IllumioPrimary
		fco["scope"] = GetHrefMaps(fc.Scope)
		if fc.WorkloadMode != "" {
			fco["workload_mode"] = fc.WorkloadMode
		}
		fcs = append(fcs, fco)
	}
	return fcs
}

// FirewallSettings represents Firewall Settings resource
type FirewallSettings struct {
	StaticPolicyScopes     Scopes              `json:"static_policy_scopes"`
	ContainerPolicyScopes  Scopes              `json:"containers_inherit_host_policy_scopes"`
	LoopbackPolicyScopes   Scopes              `json:"loopback_interfaces_in_policy_scopes"`
	RejectConnectionScopes Scopes              `json:"blocked_connection_reject_scopes"`
	FirewallCoexistence    FirewallCoexistence `json:"firewall_coexistence"`
	IKEAuthType            string              `json:"ike_authentication_type"`
}

// ToMap - Returns map for FirewallSettings model
func (fs *FirewallSettings) ToMap() (map[string]interface{}, error) {
	m := map[string]interface{}{}
	if fs.IKEAuthType != "" {
		m["ike_authentication_type"] = fs.IKEAuthType
	}
	m["static_policy_scopes"] = fs.StaticPolicyScopes.ToList()
	m["containers_inherit_host_policy_scopes"] = fs.ContainerPolicyScopes.ToList()
	m["blocked_connection_reject_scopes"] = fs.RejectConnectionScopes.ToList()
	m["loopback_interfaces_in_policy_scopes"] = fs.LoopbackPolicyScopes.ToList()
	m["firewall_coexistence"] = fs.FirewallCoexistence.ToList()
	return m, nil
}
