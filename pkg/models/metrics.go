package models

import "time"

type Metrics struct {
	UniqueVisitorsCount int64
	Timestamp           time.Time
}

type MetricsAggregated struct {
	LastHour          *Metrics
	Last24Hours       *Metrics
	PublicationsCount int64
}
