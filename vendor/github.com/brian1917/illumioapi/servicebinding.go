package illumioapi

// A ServiceBinding binds a worklad to a Virtual Service
type ServiceBinding struct {
	Href           string          `json:"href,omitempty"`
	VirtualService VirtualService  `json:"virtual_service"`
	Workload       Workload        `json:"workload"`
	PortOverrides  []PortOverrides `json:"port_overrides,omitempty"`
}

// PortOverrides override a port on a virtual service binding.
type PortOverrides struct {
	Port    int `json:"port"`
	Proto   int `json:"proto"`
	NewPort int `json:"new_port"`
}

// GetServiceBindings returns a slice of labels from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetServiceBindings(queryParameters map[string]string) (serviceBindings []ServiceBinding, api APIResponse, err error) {
	api, err = p.GetCollection("service_bindings", false, queryParameters, &serviceBindings)
	if len(serviceBindings) >= 500 {
		serviceBindings = nil
		api, err = p.GetCollection("service_bindings", true, queryParameters, &serviceBindings)
	}
	return serviceBindings, api, err
}

// CreateServiceBinding binds new workloads to a virtual service
func (p *PCE) CreateServiceBinding(serviceBindings []ServiceBinding) (createdServiceBindings []ServiceBinding, api APIResponse, err error) {
	// Sanitize Bindings
	sanSBs := []ServiceBinding{}
	for _, sb := range serviceBindings {
		sb.Href = ""
		sb.VirtualService = VirtualService{Href: sb.VirtualService.SetActive().Href}
		sb.Workload = Workload{Href: sb.Workload.Href}
		sanSBs = append(sanSBs, sb)
	}
	serviceBindings = sanSBs

	api, err = p.Post("service_bindings", &serviceBindings, &createdServiceBindings)
	return createdServiceBindings, api, err
}
