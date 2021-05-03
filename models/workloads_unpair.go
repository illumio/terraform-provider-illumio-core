package models

type WorkloadsUnpair struct {
	Hrefs          []Href `json:"workloads"`
	IPTableRestore string `json:"ip_table_restore"`
}

func (o *WorkloadsUnpair) ToMap() (map[string]interface{}, error) {
	return map[string]interface{}{
		"workloads":        GetHrefMaps(o.Hrefs),
		"ip_table_restore": o.IPTableRestore,
	}, nil
}
