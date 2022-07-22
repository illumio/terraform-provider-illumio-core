package illumioapi

// GetContainerWklds returns a slice of container workloads from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetContainerWklds(queryParameters map[string]string) (containerWklds []Workload, api APIResponse, err error) {
	api, err = p.GetCollection("container_workloads", false, queryParameters, &containerWklds)
	if len(containerWklds) >= 500 {
		containerWklds = nil
		api, err = p.GetCollection("container_workloads", true, queryParameters, &containerWklds)
	}
	p.ContainerWorkloads = make(map[string]Workload)
	for _, w := range containerWklds {
		p.ContainerWorkloads[w.Href] = w
		p.ContainerWorkloads[w.Hostname] = w
		p.ContainerWorkloads[w.Name] = w
	}
	p.ContainerWorkloadsSlice = containerWklds
	return containerWklds, api, err
}
