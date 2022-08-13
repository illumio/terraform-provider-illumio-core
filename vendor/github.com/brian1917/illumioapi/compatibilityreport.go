package illumioapi

import (
	"time"
)

// CompatibilityReport is a compatibility report for a VEN in Idle status
type CompatibilityReport struct {
	LastUpdatedAt time.Time `json:"last_updated_at"`
	Results       Results   `json:"results"`
	QualifyStatus string    `json:"qualify_status"`
}

// QualifyTest is part of compatibility report. Using interface types because API format is not guaranteed.
type QualifyTest struct {
	Status                    string      `json:"status"`
	IpsecServiceEnabled       interface{} `json:"ipsec_service_enabled"`
	Ipv4ForwardingEnabled     interface{} `json:"ipv4_forwarding_enabled"`
	Ipv4ForwardingPktCnt      interface{} `json:"ipv4_forwarding_pkt_cnt"`
	IptablesRuleCnt           interface{} `json:"iptables_rule_cnt"`
	Ipv6GlobalScope           interface{} `json:"ipv6_global_scope"`
	Ipv6ActiveConnCnt         interface{} `json:"ipv6_active_conn_cnt"`
	IP6TablesRuleCnt          interface{} `json:"ip6tables_rule_cnt"`
	RoutingTableConflict      interface{} `json:"routing_table_conflict"`
	IPv6Enabled               interface{} `json:"IPv6_enabled"`
	UnwantedNics              interface{} `json:"Unwanted_nics"`
	GroupPolicy               interface{} `json:"Group_policy"`
	RequiredPackagesInstalled interface{} `json:"required_packages_installed"`
	RequiredPackagesMissing   *[]string   `json:"required_packages_missing"`
}

// Results are the list of qualify tests
type Results struct {
	QualifyTests []QualifyTest `json:"qualify_tests"`
}

// GetCompatibilityReport returns the compatibility report for a VEN
func (p *PCE) GetCompatibilityReport(w Workload) (cr CompatibilityReport, api APIResponse, err error) {
	api, err = p.GetHref(w.Agent.Href+"/compatibility_report", &cr)
	return cr, api, err
}
