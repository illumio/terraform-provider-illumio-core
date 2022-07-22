package illumioapi

// Deprecated: use GetLabels instead.
func (p *PCE) GetAllLabels() ([]Label, APIResponse, error) {
	return p.GetLabels(nil)
}

// Deprecated: use GetLabels instead.
func (p *PCE) GetAllLabelsQP(queryParameters map[string]string) ([]Label, APIResponse, error) {
	return p.GetLabels(queryParameters)
}

// Deprecated: Use GetLabelByKeyValue instead.
func (p *PCE) GetLabelbyKeyValue(key, value string) (Label, APIResponse, error) {
	return p.GetLabelByKeyValue(key, value)
}

// Deprecated: Use GetLabelByHref instead.
func (p *PCE) GetLabelbyHref(href string) (Label, APIResponse, error) {
	return p.GetLabelByHref(href)
}

// Deprecated: Use GetWklds instead.
func (p *PCE) GetAllWorkloads() ([]Workload, APIResponse, error) {
	return p.GetWklds(nil)
}

// Deprecated: Use GetWklds instead.
func (p *PCE) GetAllWorkloadsQP(queryParameters map[string]string) ([]Workload, APIResponse, error) {
	return p.GetWklds(queryParameters)
}

// Deprecated: Use GetWklds and the populated workloads map instead.
func (p *PCE) GetWkldHrefMap() (map[string]Workload, APIResponse, error) {
	m := make(map[string]Workload)
	wklds, a, err := p.GetWklds(nil)
	if err != nil {
		return nil, a, err
	}
	for _, w := range wklds {
		m[w.Href] = w
	}
	return m, a, nil
}

// Deprecated: Use GetWklds and the populated workloads map instead.
func (p *PCE) GetWkldHostMap() (map[string]Workload, APIResponse, error) {
	m := make(map[string]Workload)
	wklds, a, err := p.GetAllWorkloads()
	if err != nil {
		return nil, a, err
	}
	for _, w := range wklds {
		m[w.Hostname] = w
	}
	return m, a, nil
}

// Deprecated: Use CreateWkld instead.
func (p *PCE) CreateWorkload(wkld Workload) (Workload, APIResponse, error) {
	return p.CreateWkld(wkld)
}

// Deprecated: Use UpdateWkld instead.
func (p *PCE) UpdateWorkload(wkld Workload) (APIResponse, error) {
	return p.UpdateWkld(wkld)
}

// Deprecated: Use GetContainerClusters instead.
func (p *PCE) GetAllContainerClusters(queryParameters map[string]string) (containerClusters []ContainerCluster, api APIResponse, err error) {
	return p.GetContainerClusters(queryParameters)
}

// Deprecated: Use GetIPListByName instead.
func (p *PCE) GetIPList(name string, pStatus string) (IPList, APIResponse, error) {
	return p.GetIPListByName(name, pStatus)
}

// Deprecated: Use GetIPLists instead.
func (p *PCE) GetAllDraftIPLists() ([]IPList, APIResponse, error) {
	return p.GetIPLists(nil, "draft")
}

// Deprecated: Use GetIPLists instead.
func (p *PCE) GetAllActiveIPLists() ([]IPList, APIResponse, error) {
	return p.GetIPLists(nil, "active")
}

// Deprecated: Use two separate calls to GetIPLists instead.
func (p *PCE) GetAllIPLists() ([]IPList, []APIResponse, error) {
	draftIPLists, a1, err := p.GetIPLists(nil, "draft")
	if err != nil {
		return nil, nil, err
	}
	activeIPLists, a2, err := p.GetIPLists(nil, "active")
	if err != nil {
		return nil, nil, err
	}
	return append(draftIPLists, activeIPLists...), []APIResponse{a1, a2}, nil
}

// Deprecated: Use GetContainerWkldProfiles instead.
func (p *PCE) GetAllContainerWorkloadProfiles(queryParameters map[string]string, containerClusterID string) ([]ContainerWorkloadProfile, APIResponse, error) {
	return p.GetContainerWkldProfiles(queryParameters, containerClusterID)
}

// Deprecated: Use GetContainerWklds instead.
func (p *PCE) GetAllContainerWorkloads(queryParameters map[string]string) ([]Workload, APIResponse, error) {
	return p.GetContainerWklds(queryParameters)
}

// Deprecated: Use GetEvents instead.
func (p *PCE) GetAllEvents(queryParameters map[string]string) ([]Event, APIResponse, error) {
	return p.GetEvents(queryParameters)
}

// Deprecated: Use GetLabelGroups instead.
func (p *PCE) GetAllLabelGroups(pStatus string) ([]LabelGroup, APIResponse, error) {
	return p.GetLabelGroups(nil, pStatus)
}

// Deprecated: Use GetVulns instead.
func (p *PCE) GetAllVulns() ([]Vulnerability, APIResponse, error) {
	return p.GetVulns(nil)
}

// Deprecated: Use GetVulnReports instead.
func (p *PCE) GetAllVulnReports() ([]VulnerabilityReport, APIResponse, error) {
	return p.GetVulnReports(nil)
}

// Deprecated: Use GetPairingProfiles instead.
func (p *PCE) GetAllPairingProfiles() ([]PairingProfile, APIResponse, error) {
	return p.GetPairingProfiles(nil)
}

// Deprecated: Use GetPendingChanges instead.
func (p *PCE) GetAllPending() (ChangeSubset, APIResponse, error) {
	return p.GetPendingChanges()
}

// Deprecated: Use GetRulesets instead.
func (p *PCE) GetAllRuleSetsQP(queryParameters map[string]string, pStatus string) ([]RuleSet, APIResponse, error) {
	return p.GetRulesets(queryParameters, pStatus)

}

// Deprecated: Use GetRulesets instead.
func (p *PCE) GetAllRuleSets(pStatus string) ([]RuleSet, APIResponse, error) {
	return p.GetRulesets(nil, pStatus)
}

// Deprecated: Use GetRulesets and the maps attached to PCE instead.
func (p *PCE) GetRuleSetMapName(pStatus string) (map[string]RuleSet, APIResponse, error) {
	ruleSets, api, err := p.GetRulesets(nil, pStatus)
	if err != nil {
		return nil, api, err
	}
	rsMap := make(map[string]RuleSet)
	for _, rs := range ruleSets {
		rsMap[rs.Name] = rs
	}
	return rsMap, api, nil
}

// Deprecated: Use CreateRule instead.
func (p *PCE) CreateRuleSetRule(rulesetHref string, rule Rule) (Rule, APIResponse, error) {
	return p.CreateRule(rulesetHref, rule)
}

// Deprecated: Use UpdateRuleset instead.
func (p *PCE) UpdateRuleSet(ruleset RuleSet) (APIResponse, error) {
	return p.UpdateRuleset(ruleset)
}

// Deprecated: Use UpdateRule instead.
func (p *PCE) UpdateRuleSetRules(rule Rule) (APIResponse, error) {
	return p.UpdateRule(rule)
}

// Deprecated: Use GetRulesetByHref instead.
func (p *PCE) GetRuleSetByHref(href string) (RuleSet, APIResponse, error) {
	return p.GetRulesetByHref(href)
}

// Deprecated: Use GetRuleByHref instead.
func (p *PCE) GetRuleSetRuleByHref(href string) (Rule, APIResponse, error) {
	return p.GetRuleByHref(href)
}

// Deprecated: Use GetRulesetHref instead.
func (r *Rule) GetRuleSetHrefFromRuleHref() string {
	return r.GetRulesetHref()
}

// Deprecated: Use GetServiceBindings instead.
func (p *PCE) GetAllServiceBindings(virtualService VirtualService) ([]ServiceBinding, APIResponse, error) {
	return p.GetServiceBindings(map[string]string{"virtual_service": virtualService.Href})
}

// Deprecated: Use GetServices instead.
func (p *PCE) GetAllServices(pStatus string) ([]Service, APIResponse, error) {
	return p.GetServices(nil, pStatus)
}

// Deprecated: Use GetTrafficAnalysis instead.
func (p *PCE) GetTrafficAnalysisAPI(t TrafficAnalysisRequest) (returnedTraffic []TrafficAnalysis, api APIResponse, err error) {
	return p.CreateTrafficRequest(t)
}

// Deprecated: Use GetADUserGroups instead.
func (p *PCE) GetAllADUserGroups() ([]ConsumingSecurityPrincipals, APIResponse, error) {
	return p.GetADUserGroups(nil)
}

// Deprecated: Use GetAllVens instead.
func (p *PCE) GetAllVens(queryParameters map[string]string) ([]VEN, APIResponse, error) {
	return p.GetVens(queryParameters)
}

// Deprecated: Use GetVirtualServers instead.
func (p *PCE) GetAllVirtualServers(pStatus string) ([]VirtualServer, APIResponse, error) {
	return p.GetVirtualServers(nil, pStatus)
}

// Deprecated: Use GetVirtualServices instead.
func (p *PCE) GetAllVirtualServices(queryParameters map[string]string, pStatus string) ([]VirtualService, APIResponse, error) {
	return p.GetVirtualServices(queryParameters, pStatus)
}

// Deprecated: Use CreateRuleset instead.
func (p *PCE) CreateRuleSet(rs RuleSet) (createdRS RuleSet, api APIResponse, err error) {
	return p.CreateRuleset(rs)
}
