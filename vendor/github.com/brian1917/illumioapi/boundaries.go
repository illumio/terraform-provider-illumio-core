package illumioapi

type EnforcementBoundary struct {
	Href            string            `json:"href,omitempty"`
	Name            string            `json:"name,omitempty"`
	Providers       []Providers       `json:"providers,omitempty"`
	Consumers       []Consumers       `json:"consumers,omitempty"`
	IngressServices []IngressServices `json:"ingress_services,omitempty"`
}

// CreateEnforcementBoundary creates a new enforcement boundary in the Illumio PCE
func (p *PCE) CreateEnforcementBoundary(eb EnforcementBoundary) (createdEB EnforcementBoundary, api APIResponse, err error) {
	api, err = p.Post("/sec_policy/draft/enforcement_boundaries", &eb, &createdEB)
	return createdEB, api, err
}
