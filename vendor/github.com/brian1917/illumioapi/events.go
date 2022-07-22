package illumioapi

import (
	"time"
)

// Event represents an auditable event in the Illumio PCE
type Event struct {
	Href           string          `json:"href"`
	Timestamp      time.Time       `json:"timestamp"`
	PceFqdn        string          `json:"pce_fqdn"`
	EventCreatedBy EventCreatedBy  `json:"created_by"`
	EventType      string          `json:"event_type"`
	Status         string          `json:"status"`
	Severity       string          `json:"severity"`
	Notifications  []Notifications `json:"notifications"`
}

// EventCreatedBy is who created the event
type EventCreatedBy struct {
	Agent            Agent            `json:"agent"`
	User             UserLogin        `json:"user"`
	ContainerCluster ContainerCluster `json:"container_cluster"`
	System           System           `json:"system,omitempty"`
	Name             string
	Href             string
}

// System is an empty struct for system-generated events
type System struct {
}

// Notifications are event notifications
type Notifications struct {
	UUID             string `json:"uuid"`
	NotificationType string `json:"notification_type"`
	Info             Info   `json:"info"`
}

// Info are notification info
type Info struct {
	APIEndpoint string `json:"api_endpoint"`
	APIMethod   string `json:"api_method"`
	SrcIP       string `json:"src_ip"`
}

// GetEvents returns a slice of labels from the PCE.
// queryParameters can be used for filtering in the form of ["parameter"]="value".
// The first API call to the PCE does not use the async option.
// If the slice length is >=500, it re-runs with async.
func (p *PCE) GetEvents(queryParameters map[string]string) (events []Event, api APIResponse, err error) {
	api, err = p.GetCollection("events", false, queryParameters, &events)
	if len(events) >= 500 {
		events = nil
		api, err = p.GetCollection("events", true, queryParameters, &events)
	}
	for i, e := range events {
		e.PopulateCreatedBy()
		events[i] = e
	}
	return events, api, err
}

func (e *Event) PopulateCreatedBy() {
	if e.EventCreatedBy.Agent.Href != "" {
		e.EventCreatedBy.Href = e.EventCreatedBy.Agent.Href
		e.EventCreatedBy.Name = e.EventCreatedBy.Agent.Hostname
	} else if e.EventCreatedBy.User.Href != "" {
		e.EventCreatedBy.Href = e.EventCreatedBy.User.Href
		e.EventCreatedBy.Name = e.EventCreatedBy.User.Username
	} else if e.EventCreatedBy.ContainerCluster.Href != "" {
		e.EventCreatedBy.Href = e.EventCreatedBy.ContainerCluster.Href
		e.EventCreatedBy.Name = e.EventCreatedBy.ContainerCluster.Name
	} else {
		e.EventCreatedBy.Href = "system"
		e.EventCreatedBy.Name = "system"
	}
}
