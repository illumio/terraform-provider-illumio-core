package models

type ServiceBindingPortOverrides struct {
	Port      int `json:"port"`
	Proto     int `json:"proto"`
	NewPort   int `json:"new_port"`
	NewToPort int `json:"new_to_port"`
}

type ServiceBinding struct {
	VirtualService        Href                          `json:"virtual_service"`
	Workload              Href                          `json:"workload"`
	PortOverrides         []ServiceBindingPortOverrides `json:"port_overrides"`
	ExternalDataReference string                        `json:"external_data_reference"`
	ExternalDataSet       string                        `json:"external_data_set"`
	ContainerWorkload     Href                          `json:"container_workload"`
}

func (sb *ServiceBinding) ToMap() (map[string]interface{}, error) {
	sbAttrMap := make(map[string]interface{})

	sbAttrMap["virtual_service"], _ = sb.VirtualService.ToMap()
	if sb.Workload.Href != "" {
		sbAttrMap["workload"], _ = sb.Workload.ToMap()
	}

	if sb.ContainerWorkload.Href != "" {
		sbAttrMap["container_workload"], _ = sb.ContainerWorkload.ToMap()
	}

	sbAttrMap["external_data_set"] = nil
	if sb.ExternalDataSet != "" {
		sbAttrMap["external_data_set"] = sb.ExternalDataSet
	}

	sbAttrMap["external_data_reference"] = nil
	if sb.ExternalDataReference != "" {
		sbAttrMap["external_data_reference"] = sb.ExternalDataReference
	}

	poMap := []map[string]interface{}{}
	for _, o := range sb.PortOverrides {
		m := make(map[string]interface{})
		m["port"] = o.Port
		m["new_port"] = o.NewPort
		m["new_to_port"] = o.NewToPort
		m["proto"] = o.Proto

		poMap = append(poMap, m)
	}
	sbAttrMap["port_overrides"] = poMap

	return map[string]interface{}{"___items___": []interface{}{sbAttrMap}}, nil
}
