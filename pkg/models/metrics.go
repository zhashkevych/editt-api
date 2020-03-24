package models

import "time"

type Metrics struct {
	UniqueVisitorsCount int64 `json:"unique_visitors_count"`
	Timestamp time.Time `json:"timestamp"`
}

type MetricsAggregated struct {
	Last24HoursStats  []*Metrics `json:"last24"`
	PublicationsCount int64      `json:"publications_count"`
}
