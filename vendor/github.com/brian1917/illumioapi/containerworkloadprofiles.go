package illumioapi

type ContainerWorkloadProfileAssignLabel struct {
	Href string `json:"href,omitempty"`
}

type ContainerWorkloadProfileLabel struct {
	Key        string                                  `json:"key,omitempty"`
	Assignment ContainerWorkloadProfileLabelAssignment `json:"assignment,omitempty"`
}
type ContainerWorkloadProfileLabelAssignment struct {
	Href  string `json:"href,omitempty"`
	Value string `json:"value,omitempty"`
}

// ContainerWorkloadProfile represents a container workload profile in the Illumio PCE
type ContainerWorkloadProfile struct {
	Href            string                                `json:"href,omitempty"`
	Name            string                                `json:"name,omitempty"`
	Namespace       string                                `json:"namespace,omitempty"`
	Description     string                                `json:"description,omitempty"`
	EnforcementMode string                                `json:"enforcement_mode,omitempty"`
	VisibilityLevel string                                `json:"visibility_level,omitempty"`
	Managed         *bool                                 `json:"managed,omitempty"`
	Linked          *bool                                 `json:"linked,omitempty"`
	AssignLabels    []ContainerWorkloadProfileAssignLabel `json:"assign_labels,omitempty"`
	Labels          []ContainerWorkloadProfileLabel       `json:"labels,omitempty"`
	CreatedAt       string                                `json:"created_at,omitempty"`
	CreatedBy       *CreatedBy                            `json:"created_by,omitempty"`
	UpdatedAt       string                                `json:"updated_at,omitempty"`
	UpdatedBy       *UpdatedBy                            `json:"updated_by,omitempty"`
}

// GetContainerWkldProfiles returns a slice of container workload profiles from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetContainerWkldProfiles(queryParameters map[string]string, containerClusterID string) (containerWkldProfiles []ContainerWorkloadProfile, api APIResponse, err error) {
	api, err = p.GetCollection("container_clusters/"+containerClusterID+"/container_workload_profiles", false, queryParameters, &containerWkldProfiles)
	if len(containerWkldProfiles) >= 500 {
		containerWkldProfiles = nil
		api, err = p.GetCollection("container_clusters/"+containerClusterID+"/container_workload_profiles", true, queryParameters, &containerWkldProfiles)
	}
	p.ContainerWorkloadProfilesSlice = containerWkldProfiles
	p.ContainerWorkloadProfiles = make(map[string]ContainerWorkloadProfile)
	for _, c := range containerWkldProfiles {
		p.ContainerWorkloadProfiles[c.Href] = c
		p.ContainerWorkloadProfiles[c.Name] = c
	}
	return containerWkldProfiles, api, err
}
