package illumioapi

import (
	"fmt"
	"strings"
)

// PCE represents an Illumio PCE and the necessary info to authenticate. The policy objects are maps for lookups. ]
// Each map will have multiple look up keys so the length of the map will be larger than the total objects.
// For example, each label will be in the map for an HREF and a key value.
// Policy objects should be called by their corresponding PCE method if you need to iterate or count them (e.g., pce.GetAllLabels)
type PCE struct {
	FriendlyName                   string
	FQDN                           string
	Port                           int
	Org                            int
	User                           string
	Key                            string
	DisableTLSChecking             bool
	LabelsSlice                    []Label               // All labels stored in a slice
	Labels                         map[string]Label      // Labels can be looked up by href or key+value (no character between key and value)
	LabelGroups                    map[string]LabelGroup // Label Groups can be looked up by href or name
	LabelGroupsSlice               []LabelGroup
	IPLists                        map[string]IPList                      // IP Lists can be looked up by href or name
	IPListsSlice                   []IPList                               // All IP Lists stored in a slice
	Workloads                      map[string]Workload                    // Workloads can be looked up by href, hostname, or names
	WorkloadsSlice                 []Workload                             // All Workloads stored in a slice
	VirtualServices                map[string]VirtualService              // VirtualServices can be looked up by href or name
	VirtualServers                 map[string]VirtualServer               // VirtualServers can be looked up by href or name
	Services                       map[string]Service                     // Services can be looked up by href or name
	ServicesSlice                  []Service                              // All services stored in a slice
	ConsumingSecurityPrincipals    map[string]ConsumingSecurityPrincipals // ConsumingSecurityPrincipals can be loooked up by href or name
	RuleSets                       map[string]RuleSet                     // RuleSets can be looked up by href or name
	VENs                           map[string]VEN                         // VENs can be looked up by href or name
	VENsSlice                      []VEN                                  // All VENs stored in a slice
	ContainerClusters              map[string]ContainerCluster
	ContainerClustersSlice         []ContainerCluster
	ContainerWorkloads             map[string]Workload
	ContainerWorkloadsSlice        []Workload
	ContainerWorkloadProfiles      map[string]ContainerWorkloadProfile
	ContainerWorkloadProfilesSlice []ContainerWorkloadProfile
}

// LoadInput tells the p.Load method what objects to load
type LoadInput struct {
	ProvisionStatus             string // Must be draft or active. Blank value is draft
	Labels                      bool
	LabelGroups                 bool
	IPLists                     bool
	Workloads                   bool
	WorkloadsQueryParameters    map[string]string
	VirtualServices             bool
	VirtualServers              bool
	Services                    bool
	ConsumingSecurityPrincipals bool
	RuleSets                    bool
	VENs                        bool
	ContainerClusters           bool
	ContainerWorkloads          bool
}

// Load fills the PCE object maps
func (p *PCE) Load(l LoadInput) (map[string]APIResponse, error) {

	var err error
	var a APIResponse
	apiResps := make(map[string]APIResponse)

	// Check provisionStatus
	provisionStatus := strings.ToLower(l.ProvisionStatus)
	if provisionStatus == "" {
		provisionStatus = "draft"
	}
	if provisionStatus != "draft" && provisionStatus != "active" {
		return apiResps, fmt.Errorf("provisionStatus must be draft or active")
	}

	// Get Label maps
	if l.Labels {
		p.LabelsSlice, a, err = p.GetLabels(nil)
		apiResps["GetAllLabels"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting labels - %s", err)
		}
		p.Labels = make(map[string]Label)
		for _, l := range p.LabelsSlice {
			p.Labels[l.Href] = l
			p.Labels[l.Key+l.Value] = l
		}
	}

	// Get all label groups
	if l.LabelGroups {
		p.LabelGroupsSlice, a, err = p.GetAllLabelGroups(provisionStatus)
		apiResps["GetAllLabelGroups"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting label groups - %s", err)
		}
		p.LabelGroups = make(map[string]LabelGroup)
		for _, lg := range p.LabelGroupsSlice {
			p.LabelGroups[lg.Href] = lg
			p.LabelGroups[lg.Name] = lg
			p.LabelGroups[lg.Key+lg.Name] = lg
		}
	}

	// Get all IPLists
	if l.IPLists {
		p.IPListsSlice, a, err = p.GetIPLists(nil, provisionStatus)
		apiResps["getAllIPLists"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting draft ip lists - %s", err)
		}
		p.IPLists = make(map[string]IPList)
		for _, ipl := range p.IPListsSlice {
			p.IPLists[ipl.Href] = ipl
			p.IPLists[ipl.Name] = ipl
		}
	}

	//  Workloads
	if l.Workloads {
		p.WorkloadsSlice, a, err = p.GetAllWorkloadsQP(l.WorkloadsQueryParameters)
		apiResps["GetAllWorkloadsQP"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting workloads - %s", err)
		}
	}

	// Virtual services
	if l.VirtualServices {
		virtualServices, a, err := p.GetAllVirtualServices(nil, provisionStatus)
		apiResps["GetAllVirtualServices"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting virtual services - %s", err)
		}
		p.VirtualServices = make(map[string]VirtualService)
		for _, vs := range virtualServices {
			p.VirtualServices[vs.Href] = vs
			p.VirtualServices[vs.Name] = vs
		}
	}

	// Services
	if l.Services {
		p.ServicesSlice, a, err = p.GetAllServices(provisionStatus)
		apiResps["GetAllServices"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting all services - %s", err)
		}
		p.Services = make(map[string]Service)
		for _, s := range p.ServicesSlice {
			p.Services[s.Href] = s
			p.Services[s.Name] = s
		}
	}

	// VirtualServers
	if l.VirtualServers {
		virtualServers, a, err := p.GetAllVirtualServers(provisionStatus)
		apiResps["GetAllVirtualServers"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting all virtual servers - %s", err)
		}
		p.VirtualServers = make(map[string]VirtualServer)
		for _, v := range virtualServers {
			p.VirtualServers[v.Href] = v
			p.VirtualServers[v.Name] = v
		}
	}

	// Rulesets
	if l.RuleSets {
		rulesets, a, err := p.GetAllRuleSets(provisionStatus)
		apiResps["GetAllRuleSets"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting all rulesets - %s", err)
		}
		p.RuleSets = make(map[string]RuleSet)
		for _, rs := range rulesets {
			p.RuleSets[rs.Href] = rs
			p.RuleSets[rs.Name] = rs
		}
	}

	// Consuming Security Principals
	if l.ConsumingSecurityPrincipals {
		cps, a, err := p.GetAllADUserGroups()
		apiResps["GetAllADUserGroups"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting all consuming security principals - %s", err)
		}
		p.ConsumingSecurityPrincipals = make(map[string]ConsumingSecurityPrincipals)
		for _, cp := range cps {
			p.ConsumingSecurityPrincipals[cp.Href] = cp
			p.ConsumingSecurityPrincipals[cp.Name] = cp
		}
	}

	// Get VENs
	if l.VENs {
		p.VENsSlice, a, err = p.GetAllVens(nil)
		apiResps["GetAllVens"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting all vens - %s", err)
		}
		p.VENs = make(map[string]VEN)
		for _, v := range p.VENsSlice {
			p.VENs[v.Name] = v
			p.VENs[v.Href] = v
			p.VENs[v.UID] = v
		}
	}

	// Container Clusters
	if l.ContainerClusters {
		p.ContainerClustersSlice, a, err = p.GetAllContainerClusters(nil)
		apiResps["GetAllContainerClusters"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting all container clusters - %s", err)
		}
		p.ContainerClusters = make(map[string]ContainerCluster)
		for _, c := range p.ContainerClustersSlice {
			p.ContainerClusters[c.Href] = c
			p.ContainerClusters[c.Name] = c
		}
	}

	// Container Workloads
	if l.ContainerWorkloads {
		p.ContainerWorkloadsSlice, a, err = p.GetAllContainerWorkloads(nil)
		apiResps["GetAllContainerWorkloads"] = a
		if err != nil {
			return apiResps, fmt.Errorf("getting all container workloads - %s", err)
		}
		p.ContainerWorkloads = map[string]Workload{}
		for _, cw := range p.ContainerWorkloadsSlice {
			p.ContainerWorkloads[cw.Name] = cw
			p.ContainerWorkloads[cw.Href] = cw
		}
	}

	return apiResps, nil
}

// FindObject takes an href and returns what it is and the name
func (p *PCE) FindObject(href string) (key, name string, err error) {

	// IPLists
	if strings.Contains(href, "/ip_lists/") {
		return "iplist", p.IPLists[href].Name, nil
	}
	// Labels
	if strings.Contains(href, "/labels/") {
		return fmt.Sprintf("%s_label", p.Labels[href].Key), p.Labels[href].Value, nil
	}
	// Label Groups
	if strings.Contains(href, "/label_groups/") {
		return fmt.Sprintf("%s_label_group", p.LabelGroups[href].Key), p.LabelGroups[href].Name, nil
	}
	// Virtual Services
	if strings.Contains(href, "/virtual_services/") {
		return "virtual_service", p.VirtualServices[href].Name, nil
	}
	// Workloads
	if strings.Contains(href, "/workloads/") {
		if p.Workloads[href].Hostname != "" {
			return "workload", p.Workloads[href].Hostname, nil
		}
		return "workload", p.Workloads[href].Name, nil
	}

	return "nil", "nil", fmt.Errorf("object not found")
}

// ParseObjectType takes an href and returns one of the following options: iplist, label, label_group, virtual_service, workload, or unknown.
func ParseObjectType(href string) string {
	// IPLists
	if strings.Contains(href, "/ip_lists/") {
		return "iplist"
	}
	// Labels
	if strings.Contains(href, "/labels/") {
		return "label"
	}
	// Label Groups
	if strings.Contains(href, "/label_groups/") {
		return "label_group"
	}
	// Virtual Services
	if strings.Contains(href, "/virtual_services/") {
		return "virtual_service"
	}
	// Workloads
	if strings.Contains(href, "/workloads/") {
		return "workload"
	}
	return "unknown"

}
