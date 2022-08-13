package illumioapi

// ContainerCluster represents a container cluster in the Illumio PCE
type ContainerCluster struct {
	Href             string `json:"href,omitempty"`
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	ContainerRuntime string `json:"container_runtime,omitempty"`
	ManagerType      string `json:"manager_type,omitempty"`
	Online           *bool  `json:"online,omitempty"`
	KubelinkVersion  string `json:"kubelink_version,omitempty"`
	PceFqdn          string `json:"pce_fqdn,omitempty"`
}

// GetContainerClusters returns a slice of ContainerCluster in the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetContainerClusters(queryParameters map[string]string) (containerClusters []ContainerCluster, api APIResponse, err error) {
	api, err = p.GetCollection("container_clusters", false, queryParameters, &containerClusters)
	if len(containerClusters) >= 500 {
		containerClusters = nil
		api, err = p.GetCollection("container_clusters", true, queryParameters, &containerClusters)
	}
	// Load the PCE with the returned workloads
	p.ContainerClustersSlice = containerClusters
	p.ContainerClusters = make(map[string]ContainerCluster)
	for _, c := range containerClusters {
		p.ContainerClusters[c.Href] = c
		p.ContainerClusters[c.Name] = c
	}
	return containerClusters, api, err
}
