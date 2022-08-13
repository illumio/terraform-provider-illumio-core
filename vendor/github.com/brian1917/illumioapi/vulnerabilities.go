package illumioapi

// Vulnerability represents a vulnerability in the Illumio PCE
type Vulnerability struct {
	CreatedAt   string     `json:"created_at,omitempty"`
	CreatedBy   *CreatedBy `json:"created_by,omitempty"`
	CveIds      []string   `json:"cve_ids,omitempty"`
	Description string     `json:"description,omitempty"`
	Href        string     `json:"href,omitempty"`
	Name        string     `json:"name,omitempty"`
	Score       int        `json:"score,omitempty"`
	UpdatedAt   string     `json:"updated_at,omitempty"`
	UpdatedBy   *UpdatedBy `json:"updated_by,omitempty"`
}

// VulnerabilityReport represents a vulnerability report in the Illumio PCE
type VulnerabilityReport struct {
	Authoritative      bool       `json:"authoritative,omitempty"`
	CreatedAt          string     `json:"created_at,omitempty"`
	CreatedBy          *CreatedBy `json:"created_by,omitempty"`
	Href               string     `json:"href,omitempty"`
	Name               string     `json:"name,omitempty"`
	NumVulnerabilities int        `json:"num_vulnerabilities,omitempty"`
	ReportType         string     `json:"report_type,omitempty"`
	ScannedIps         []string   `json:"scanned_ips,omitempty"`
	UpdatedAt          string     `json:"updated_at,omitempty"`
	UpdatedBy          *UpdatedBy `json:"updated_by,omitempty"`
}

// GetVulns returns a slice of vulnerabilities from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetVulns(queryParameters map[string]string) (vulns []Vulnerability, api APIResponse, err error) {
	api, err = p.GetCollection("vulnerabilities", false, queryParameters, &vulns)
	if len(vulns) >= 500 {
		vulns = nil
		api, err = p.GetCollection("vulnerabilities", true, queryParameters, &vulns)
	}
	return vulns, api, err
}

// GetVulnReports returns a slice of vulnerabilities from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetVulnReports(queryParameters map[string]string) (vulnReports []VulnerabilityReport, api APIResponse, err error) {
	api, err = p.GetCollection("vulnerability_reports", false, queryParameters, &vulnReports)
	if len(vulnReports) >= 500 {
		vulnReports = nil
		api, err = p.GetCollection("vulnerability_reports", true, queryParameters, &vulnReports)
	}
	return vulnReports, api, err
}
