package illumioapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// VEN represents a VEN in the Illumio PCE.
// Not including duplicated fields in a workload - labels, OS information, interfaces, etc.
type VEN struct {
	Href             string            `json:"href,omitempty"`
	Name             string            `json:"name,omitempty"`
	Description      string            `json:"description,omitempty"`
	Hostname         string            `json:"hostname,omitempty"`
	UID              string            `json:"uid,omitempty"`
	Status           string            `json:"status,omitempty"`
	Version          string            `json:"version,omitempty"`
	ActivationType   string            `json:"activation_type,omitempty"`
	ActivePceFqdn    string            `json:"active_pce_fqdn,omitempty"`
	TargetPceFqdn    string            `json:"target_pce_fqdn,omitempty"`
	Workloads        *[]*Workload      `json:"workloads,omitempty"`
	ContainerCluster *ContainerCluster `json:"container_cluster,omitempty"`
}

type VENUpgrade struct {
	VENs    []VEN  `json:"vens"`
	Release string `json:"release"`
	DryRun  bool   `json:"dry_run"`
}

type VENUpgradeResp struct {
	VENUpgradeErrors []VENUpgradeError `json:"errors"`
}

type VENUpgradeError struct {
	Token   string   `json:"token"`
	Message string   `json:"message"`
	Hrefs   []string `json:"hrefs"`
}

// LoadVenMap populates the workload maps based on p.WorkloadSlice
func (p *PCE) LoadVenMap() {
	p.VENs = make(map[string]VEN)
	for _, v := range p.VENsSlice {
		p.VENs[v.Href] = v
		p.VENs[v.Name] = v
		p.VENs[v.Hostname] = v
	}
}

// GetVens returns a slice of VENs from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value"
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetVens(queryParameters map[string]string) (vens []VEN, api APIResponse, err error) {
	api, err = p.GetCollection("vens", false, queryParameters, &vens)
	if len(vens) >= 500 {
		vens = nil
		api, err = p.GetCollection("vens", true, queryParameters, &vens)
	}
	p.VENsSlice = vens
	p.LoadVenMap()
	return vens, api, err
}

// GetVenByHref returns the VEN with a specific href
func (p *PCE) GetVenByHref(href string) (ven VEN, api APIResponse, err error) {
	api, err = p.GetHref(href, &ven)
	return ven, api, err
}

// UpdateVEN updates an existing ven in the Illumio PCE
// The provided ven struct must include an href.
// Properties that cannot be included in the PUT method will be ignored.
func (p *PCE) UpdateVen(ven VEN) (api APIResponse, err error) {

	// Build the new ven with only propertie we can update
	if strings.ToLower(ven.Status) != "active" && strings.ToLower(ven.Status) != "suspended" {
		return api, fmt.Errorf("%s is not a valid status. must be active or suspended", ven.Status)
	}
	venToUpdate := VEN{Href: ven.Href, Name: ven.Name, Description: ven.Description, Status: strings.ToLower(ven.Status)}

	return p.Put(&venToUpdate)
}

func (p *PCE) UpgradeVENs(vens []VEN, release string) (resp VENUpgradeResp, api APIResponse, err error) {
	// Build the API URL
	apiURL, err := url.Parse("https://" + p.cleanFQDN() + ":" + strconv.Itoa(p.Port) + "/api/v2/orgs/" + strconv.Itoa(p.Org) + "/vens/upgrade")
	if err != nil {
		return VENUpgradeResp{}, api, fmt.Errorf("upgrade ven - %s", err)
	}

	// Build the venUpgrade
	venHrefs := []VEN{}
	for _, v := range vens {
		venHrefs = append(venHrefs, VEN{Href: v.Href})
	}
	venUpgrade := VENUpgrade{Release: release, DryRun: false, VENs: venHrefs}

	// Call the API
	venUpgradeJSON, err := json.Marshal(venUpgrade)
	if err != nil {
		return VENUpgradeResp{}, api, err
	}
	api, err = p.httpReq("PUT", apiURL.String(), venUpgradeJSON, false, true)
	if err != nil {
		return VENUpgradeResp{}, api, fmt.Errorf("upgrade ven - %s", err)
	}
	api.ReqBody = string(venUpgradeJSON)
	json.Unmarshal([]byte(api.RespBody), &resp)

	return resp, api, nil
}

// GetVenByHostname gets a VEN by the hostname
func (p *PCE) GetVenByHostname(hostname string) (VEN, APIResponse, error) {
	vens, a, err := p.GetVens(map[string]string{"hostname": hostname})
	if err != nil {
		return VEN{}, a, err
	}
	for _, ven := range vens {
		if ven.Hostname == hostname {
			return ven, a, nil
		}
	}
	return VEN{}, a, nil
}
