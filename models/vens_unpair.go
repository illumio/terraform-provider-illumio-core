package models

type VENsUnpair struct {
	Hrefs           []Href `json:"vens"`
	FirewallRestore string `json:"firewall_restore"`
}

func (o *VENsUnpair) ToMap() (map[string]interface{}, error) {
	return map[string]interface{}{
		"vens":             GetHrefMaps(o.Hrefs),
		"firewall_restore": o.FirewallRestore,
	}, nil
}
